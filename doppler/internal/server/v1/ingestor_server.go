package v1

import (
	"context"
	"io"
	"log"
	"sync/atomic"
	"time"

	"code.cloudfoundry.org/loggregator/diodes"
	"code.cloudfoundry.org/loggregator/plumbing"
	"code.cloudfoundry.org/loggregator/plumbing/conversion"
	"github.com/cloudfoundry/dropsonde/metricbatcher"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gogo/protobuf/proto"
)

type IngestorServer struct {
	v1Buf   *diodes.ManyToOneEnvelope
	v2Buf   *diodes.ManyToOneEnvelopeV2
	batcher Batcher
	health  HealthRegistrar
}

type Batcher interface {
	BatchCounter(name string) metricbatcher.BatchCounterChainer
}

type IngestorGRPCServer interface {
	plumbing.DopplerIngestor_PusherServer
}

func NewIngestorServer(
	v1Buf *diodes.ManyToOneEnvelope,
	v2Buf *diodes.ManyToOneEnvelopeV2,
	batcher Batcher,
	health HealthRegistrar,
) *IngestorServer {

	return &IngestorServer{
		v1Buf:   v1Buf,
		v2Buf:   v2Buf,
		batcher: batcher,
		health:  health,
	}
}

func (i *IngestorServer) Pusher(pusher plumbing.DopplerIngestor_PusherServer) error {
	i.health.Inc("ingressStreamCount")
	defer i.health.Dec("ingressStreamCount")

	var done int64
	context := pusher.Context()
	go i.monitorContext(context, &done)

	for {
		if atomic.LoadInt64(&done) > 0 {
			return context.Err()
		}

		envelopeData, err := pusher.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		env := &events.Envelope{}
		err = proto.Unmarshal(envelopeData.Payload, env)
		if err != nil {
			log.Printf("Received bad envelope: %s", err)
			continue
		}
		// metric-documentation-v1: (listeners.receivedEnvelopes) Number of received
		// envelopes in v1 ingress server
		i.batcher.BatchCounter("listeners.receivedEnvelopes").
			SetTag("protocol", "grpc").
			SetTag("event_type", env.GetEventType().String()).
			Increment()

		v2e := conversion.ToV2(env, true)
		i.v1Buf.Set(env)
		i.v2Buf.Set(v2e)

		// metric-documentation-v1: (listeners.totalReceivedMessageCount) Total
		// number of messages received by doppler.
		i.batcher.BatchCounter("listeners.totalReceivedMessageCount").
			Increment()
	}
	return nil
}

func (i *IngestorServer) monitorContext(ctx context.Context, done *int64) {
	<-ctx.Done()
	atomic.StoreInt64(done, 1)
}
