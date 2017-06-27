package component_test

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.cloudfoundry.org/loggregator/testservers"

	"code.cloudfoundry.org/loggregator/plumbing"
	v2 "code.cloudfoundry.org/loggregator/plumbing/v2"

	"code.cloudfoundry.org/loggregator/metron/app"

	"github.com/cloudfoundry/dropsonde/emitter"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Metron", func() {
	var (
		consumerServer *Server
		metronCleanup  func()
		metronConfig   app.Config
		metronInfoPath string
	)

	BeforeEach(func() {
		var err error
		consumerServer, err = NewServer()
		Expect(err).ToNot(HaveOccurred())

		var metronReady func()
		metronCleanup, metronConfig, metronInfoPath, metronReady = testservers.StartMetron(
			testservers.BuildMetronConfig("localhost", consumerServer.Port()),
		)
		metronReady()
	})

	AfterEach(func() {
		consumerServer.Stop()
		metronCleanup()
	})

	It("provides a health endpoint", func() {
		healthURL := testservers.InfoPollString(metronInfoPath, "metron", "health_url")
		resp, err := http.Get(healthURL)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})

	It("accepts connections on the v1 API", func() {
		udpAddr := testservers.InfoPollString(metronInfoPath, "metron", "udp_addr")
		udpEmitter, err := emitter.NewUdpEmitter(udpAddr)
		Expect(err).ToNot(HaveOccurred())
		eventEmitter := emitter.NewEventEmitter(udpEmitter, "some-origin")

		emitEnvelope := &events.Envelope{
			Origin:    proto.String("some-origin"),
			EventType: events.Envelope_Error.Enum(),
			Error: &events.Error{
				Source:  proto.String("some-source"),
				Code:    proto.Int32(1),
				Message: proto.String("message"),
			},
		}

		f := func() int {
			eventEmitter.Emit(emitEnvelope)
			return len(consumerServer.V1.PusherCalled)
		}
		Eventually(f, 5).Should(BeNumerically(">", 0))

		var rx plumbing.DopplerIngestor_PusherServer
		Eventually(consumerServer.V1.PusherInput.Arg0).Should(Receive(&rx))

		data, err := rx.Recv()
		Expect(err).ToNot(HaveOccurred())

		envelope := new(events.Envelope)
		Expect(envelope.Unmarshal(data.Payload)).To(Succeed())
	})

	It("accepts connections on the v2 API", func() {
		emitEnvelope := &v2.Envelope{
			Message: &v2.Envelope_Log{
				Log: &v2.Log{
					Payload: []byte("some-message"),
					Type:    v2.Log_OUT,
				},
			},
		}

		grpcAddr := testservers.InfoPollString(metronInfoPath, "metron", "grpc_addr")
		client := metronClient(grpcAddr, metronConfig.GRPC)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		sender, err := client.Sender(ctx)
		Expect(err).ToNot(HaveOccurred())

		go func() {
			for range time.Tick(time.Nanosecond) {
				sender.Send(emitEnvelope)
			}
		}()

		var rx v2.DopplerIngress_BatchSenderServer
		Eventually(consumerServer.V2.BatchSenderInput.Arg0).Should(Receive(&rx))
		fmt.Printf("%#v\n", rx)

		var envBatch *v2.EnvelopeBatch
		var idx int
		f := func() *v2.Envelope {
			// TODO: this thing blocks and keeps test running forever
			batch, err := rx.Recv()
			Expect(err).ToNot(HaveOccurred())

			defer func() {
				envBatch = batch
			}()

			for i, envelope := range batch.Batch {
				if envelope.GetLog() != nil {
					idx = i
					return envelope
				}
			}

			return nil
		}
		Eventually(f).ShouldNot(BeNil())

		Expect(len(envBatch.Batch)).ToNot(BeZero())

		Expect(*envBatch.Batch[idx]).To(MatchFields(IgnoreExtras, Fields{
			"Message": Equal(&v2.Envelope_Log{
				Log: &v2.Log{Payload: []byte("some-message")},
			}),
			"Tags": Equal(map[string]*v2.Value{
				"auto-tag-1": {
					Data: &v2.Value_Text{"auto-tag-value-1"},
				},
				"auto-tag-2": {
					Data: &v2.Value_Text{"auto-tag-value-2"},
				},
			}),
		}))
	})
})

func HomeAddrToPort(addr net.Addr) int {
	port, err := strconv.Atoi(strings.Replace(addr.String(), "127.0.0.1:", "", 1))
	if err != nil {
		panic(err)
	}
	return port
}

func metronClient(addr string, grpcConfig app.GRPC) v2.IngressClient {
	tlsConfig, err := plumbing.NewMutualTLSConfig(
		grpcConfig.CertFile,
		grpcConfig.KeyFile,
		grpcConfig.CAFile,
		"metron",
	)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err != nil {
		panic(err)
	}
	return v2.NewIngressClient(conn)
}
