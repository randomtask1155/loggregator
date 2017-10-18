// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package metron_test

import (
	"context"

	"code.cloudfoundry.org/loggregator/plumbing"
	v2 "code.cloudfoundry.org/loggregator/plumbing/v2"

	"google.golang.org/grpc/metadata"
)

type mockDopplerIngestorServerV1 struct {
	PusherCalled chan bool
	PusherInput  struct {
		Arg0 chan plumbing.DopplerIngestor_PusherServer
	}
	PusherOutput struct {
		Ret0 chan error
	}
}

func newMockDopplerIngestorServerV1() *mockDopplerIngestorServerV1 {
	m := &mockDopplerIngestorServerV1{}
	m.PusherCalled = make(chan bool, 100)
	m.PusherInput.Arg0 = make(chan plumbing.DopplerIngestor_PusherServer, 100)
	m.PusherOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockDopplerIngestorServerV1) Pusher(arg0 plumbing.DopplerIngestor_PusherServer) error {
	m.PusherCalled <- true
	m.PusherInput.Arg0 <- arg0
	return <-m.PusherOutput.Ret0
}

type mockDopplerIngestor_PusherServerV1 struct {
	SendAndCloseCalled chan bool
	SendAndCloseInput  struct {
		Arg0 chan *plumbing.PushResponse
	}
	SendAndCloseOutput struct {
		Ret0 chan error
	}
	RecvCalled chan bool
	RecvOutput struct {
		Ret0 chan *plumbing.EnvelopeData
		Ret1 chan error
	}
	SendHeaderCalled chan bool
	SendHeaderInput  struct {
		Arg0 chan metadata.MD
	}
	SendHeaderOutput struct {
		Ret0 chan error
	}
	SetTrailerCalled chan bool
	SetTrailerInput  struct {
		Arg0 chan metadata.MD
	}
	ContextCalled chan bool
	ContextOutput struct {
		Ret0 chan context.Context
	}
	SendMsgCalled chan bool
	SendMsgInput  struct {
		M chan interface{}
	}
	SendMsgOutput struct {
		Ret0 chan error
	}
	RecvMsgCalled chan bool
	RecvMsgInput  struct {
		M chan interface{}
	}
	RecvMsgOutput struct {
		Ret0 chan error
	}
}

func newMockDopplerIngestor_PusherServerV1() *mockDopplerIngestor_PusherServerV1 {
	m := &mockDopplerIngestor_PusherServerV1{}
	m.SendAndCloseCalled = make(chan bool, 100)
	m.SendAndCloseInput.Arg0 = make(chan *plumbing.PushResponse, 100)
	m.SendAndCloseOutput.Ret0 = make(chan error, 100)
	m.RecvCalled = make(chan bool, 100)
	m.RecvOutput.Ret0 = make(chan *plumbing.EnvelopeData, 100)
	m.RecvOutput.Ret1 = make(chan error, 100)
	m.SendHeaderCalled = make(chan bool, 100)
	m.SendHeaderInput.Arg0 = make(chan metadata.MD, 100)
	m.SendHeaderOutput.Ret0 = make(chan error, 100)
	m.SetTrailerCalled = make(chan bool, 100)
	m.SetTrailerInput.Arg0 = make(chan metadata.MD, 100)
	m.ContextCalled = make(chan bool, 100)
	m.ContextOutput.Ret0 = make(chan context.Context, 100)
	m.SendMsgCalled = make(chan bool, 100)
	m.SendMsgInput.M = make(chan interface{}, 100)
	m.SendMsgOutput.Ret0 = make(chan error, 100)
	m.RecvMsgCalled = make(chan bool, 100)
	m.RecvMsgInput.M = make(chan interface{}, 100)
	m.RecvMsgOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockDopplerIngestor_PusherServerV1) SendAndClose(arg0 *plumbing.PushResponse) error {
	m.SendAndCloseCalled <- true
	m.SendAndCloseInput.Arg0 <- arg0
	return <-m.SendAndCloseOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV1) Recv() (*plumbing.EnvelopeData, error) {
	m.RecvCalled <- true
	return <-m.RecvOutput.Ret0, <-m.RecvOutput.Ret1
}
func (m *mockDopplerIngestor_PusherServerV1) SendHeader(arg0 metadata.MD) error {
	m.SendHeaderCalled <- true
	m.SendHeaderInput.Arg0 <- arg0
	return <-m.SendHeaderOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV1) SetTrailer(arg0 metadata.MD) {
	m.SetTrailerCalled <- true
	m.SetTrailerInput.Arg0 <- arg0
}
func (m *mockDopplerIngestor_PusherServerV1) Context() context.Context {
	m.ContextCalled <- true
	return <-m.ContextOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV1) SendMsg(m_ interface{}) error {
	m.SendMsgCalled <- true
	m.SendMsgInput.M <- m_
	return <-m.SendMsgOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV1) RecvMsg(m_ interface{}) error {
	m.RecvMsgCalled <- true
	m.RecvMsgInput.M <- m_
	return <-m.RecvMsgOutput.Ret0
}

type mockDopplerIngressServerV2 struct {
	SenderCalled chan bool
	SenderInput  struct {
		Arg0 chan v2.DopplerIngress_SenderServer
	}
	SenderOutput struct {
		Ret0 chan error
	}

	BatchSenderCalled chan bool
	BatchSenderInput  struct {
		Arg0 chan v2.DopplerIngress_BatchSenderServer
	}
	BatchSenderOutput struct {
		Ret0 chan error
	}
}

func newMockDopplerIngressServerV2() *mockDopplerIngressServerV2 {
	m := &mockDopplerIngressServerV2{}
	m.SenderCalled = make(chan bool, 100)
	m.SenderInput.Arg0 = make(chan v2.DopplerIngress_SenderServer, 100)
	m.SenderOutput.Ret0 = make(chan error, 100)

	m.BatchSenderCalled = make(chan bool, 100)
	m.BatchSenderInput.Arg0 = make(chan v2.DopplerIngress_BatchSenderServer, 100)
	m.BatchSenderOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockDopplerIngressServerV2) Sender(arg0 v2.DopplerIngress_SenderServer) error {
	m.SenderCalled <- true
	m.SenderInput.Arg0 <- arg0
	return <-m.SenderOutput.Ret0
}
func (m *mockDopplerIngressServerV2) BatchSender(arg0 v2.DopplerIngress_BatchSenderServer) error {
	m.BatchSenderCalled <- true
	m.BatchSenderInput.Arg0 <- arg0
	return <-m.BatchSenderOutput.Ret0
}

type mockDopplerIngestor_PusherServerV2 struct {
	SendAndCloseCalled chan bool
	SendAndCloseInput  struct {
		Arg0 chan *v2.SenderResponse
	}
	SendAndCloseOutput struct {
		Ret0 chan error
	}
	RecvCalled chan bool
	RecvOutput struct {
		Ret0 chan *v2.Envelope
		Ret1 chan error
	}
	SendHeaderCalled chan bool
	SendHeaderInput  struct {
		Arg0 chan metadata.MD
	}
	SendHeaderOutput struct {
		Ret0 chan error
	}
	SetTrailerCalled chan bool
	SetTrailerInput  struct {
		Arg0 chan metadata.MD
	}
	ContextCalled chan bool
	ContextOutput struct {
		Ret0 chan context.Context
	}
	SendMsgCalled chan bool
	SendMsgInput  struct {
		M chan interface{}
	}
	SendMsgOutput struct {
		Ret0 chan error
	}
	RecvMsgCalled chan bool
	RecvMsgInput  struct {
		M chan interface{}
	}
	RecvMsgOutput struct {
		Ret0 chan error
	}
}

func newMockDopplerIngestor_PusherServerV2() *mockDopplerIngestor_PusherServerV2 {
	m := &mockDopplerIngestor_PusherServerV2{}
	m.SendAndCloseCalled = make(chan bool, 100)
	m.SendAndCloseInput.Arg0 = make(chan *v2.SenderResponse, 100)
	m.SendAndCloseOutput.Ret0 = make(chan error, 100)
	m.RecvCalled = make(chan bool, 100)
	m.RecvOutput.Ret0 = make(chan *v2.Envelope, 100)
	m.RecvOutput.Ret1 = make(chan error, 100)
	m.SendHeaderCalled = make(chan bool, 100)
	m.SendHeaderInput.Arg0 = make(chan metadata.MD, 100)
	m.SendHeaderOutput.Ret0 = make(chan error, 100)
	m.SetTrailerCalled = make(chan bool, 100)
	m.SetTrailerInput.Arg0 = make(chan metadata.MD, 100)
	m.ContextCalled = make(chan bool, 100)
	m.ContextOutput.Ret0 = make(chan context.Context, 100)
	m.SendMsgCalled = make(chan bool, 100)
	m.SendMsgInput.M = make(chan interface{}, 100)
	m.SendMsgOutput.Ret0 = make(chan error, 100)
	m.RecvMsgCalled = make(chan bool, 100)
	m.RecvMsgInput.M = make(chan interface{}, 100)
	m.RecvMsgOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockDopplerIngestor_PusherServerV2) SendAndClose(arg0 *v2.SenderResponse) error {
	m.SendAndCloseCalled <- true
	m.SendAndCloseInput.Arg0 <- arg0
	return <-m.SendAndCloseOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV2) Recv() (*v2.Envelope, error) {
	m.RecvCalled <- true
	return <-m.RecvOutput.Ret0, <-m.RecvOutput.Ret1
}
func (m *mockDopplerIngestor_PusherServerV2) SendHeader(arg0 metadata.MD) error {
	m.SendHeaderCalled <- true
	m.SendHeaderInput.Arg0 <- arg0
	return <-m.SendHeaderOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV2) SetTrailer(arg0 metadata.MD) {
	m.SetTrailerCalled <- true
	m.SetTrailerInput.Arg0 <- arg0
}
func (m *mockDopplerIngestor_PusherServerV2) Context() context.Context {
	m.ContextCalled <- true
	return <-m.ContextOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV2) SendMsg(m_ interface{}) error {
	m.SendMsgCalled <- true
	m.SendMsgInput.M <- m_
	return <-m.SendMsgOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServerV2) RecvMsg(m_ interface{}) error {
	m.RecvMsgCalled <- true
	m.RecvMsgInput.M <- m_
	return <-m.RecvMsgOutput.Ret0
}
