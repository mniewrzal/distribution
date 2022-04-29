// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.28
// source: contact.proto

package pb

import (
	bytes "bytes"
	context "context"
	errors "errors"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_contact_proto struct{}

func (drpcEncoding_File_contact_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_contact_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_contact_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_contact_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCContactClient interface {
	DRPCConn() drpc.Conn

	PingNode(ctx context.Context, in *ContactPingRequest) (*ContactPingResponse, error)
}

type drpcContactClient struct {
	cc drpc.Conn
}

func NewDRPCContactClient(cc drpc.Conn) DRPCContactClient {
	return &drpcContactClient{cc}
}

func (c *drpcContactClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcContactClient) PingNode(ctx context.Context, in *ContactPingRequest) (*ContactPingResponse, error) {
	out := new(ContactPingResponse)
	err := c.cc.Invoke(ctx, "/contact.Contact/PingNode", drpcEncoding_File_contact_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCContactServer interface {
	PingNode(context.Context, *ContactPingRequest) (*ContactPingResponse, error)
}

type DRPCContactUnimplementedServer struct{}

func (s *DRPCContactUnimplementedServer) PingNode(context.Context, *ContactPingRequest) (*ContactPingResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCContactDescription struct{}

func (DRPCContactDescription) NumMethods() int { return 1 }

func (DRPCContactDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/contact.Contact/PingNode", drpcEncoding_File_contact_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCContactServer).
					PingNode(
						ctx,
						in1.(*ContactPingRequest),
					)
			}, DRPCContactServer.PingNode, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterContact(mux drpc.Mux, impl DRPCContactServer) error {
	return mux.Register(impl, DRPCContactDescription{})
}

type DRPCContact_PingNodeStream interface {
	drpc.Stream
	SendAndClose(*ContactPingResponse) error
}

type drpcContact_PingNodeStream struct {
	drpc.Stream
}

func (x *drpcContact_PingNodeStream) SendAndClose(m *ContactPingResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_contact_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeClient interface {
	DRPCConn() drpc.Conn

	CheckIn(ctx context.Context, in *CheckInRequest) (*CheckInResponse, error)
	PingMe(ctx context.Context, in *PingMeRequest) (*PingMeResponse, error)
	GetTime(ctx context.Context, in *GetTimeRequest) (*GetTimeResponse, error)
}

type drpcNodeClient struct {
	cc drpc.Conn
}

func NewDRPCNodeClient(cc drpc.Conn) DRPCNodeClient {
	return &drpcNodeClient{cc}
}

func (c *drpcNodeClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodeClient) CheckIn(ctx context.Context, in *CheckInRequest) (*CheckInResponse, error) {
	out := new(CheckInResponse)
	err := c.cc.Invoke(ctx, "/contact.Node/CheckIn", drpcEncoding_File_contact_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeClient) PingMe(ctx context.Context, in *PingMeRequest) (*PingMeResponse, error) {
	out := new(PingMeResponse)
	err := c.cc.Invoke(ctx, "/contact.Node/PingMe", drpcEncoding_File_contact_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeClient) GetTime(ctx context.Context, in *GetTimeRequest) (*GetTimeResponse, error) {
	out := new(GetTimeResponse)
	err := c.cc.Invoke(ctx, "/contact.Node/GetTime", drpcEncoding_File_contact_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCNodeServer interface {
	CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error)
	PingMe(context.Context, *PingMeRequest) (*PingMeResponse, error)
	GetTime(context.Context, *GetTimeRequest) (*GetTimeResponse, error)
}

type DRPCNodeUnimplementedServer struct{}

func (s *DRPCNodeUnimplementedServer) CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCNodeUnimplementedServer) PingMe(context.Context, *PingMeRequest) (*PingMeResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCNodeUnimplementedServer) GetTime(context.Context, *GetTimeRequest) (*GetTimeResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCNodeDescription struct{}

func (DRPCNodeDescription) NumMethods() int { return 3 }

func (DRPCNodeDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/contact.Node/CheckIn", drpcEncoding_File_contact_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeServer).
					CheckIn(
						ctx,
						in1.(*CheckInRequest),
					)
			}, DRPCNodeServer.CheckIn, true
	case 1:
		return "/contact.Node/PingMe", drpcEncoding_File_contact_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeServer).
					PingMe(
						ctx,
						in1.(*PingMeRequest),
					)
			}, DRPCNodeServer.PingMe, true
	case 2:
		return "/contact.Node/GetTime", drpcEncoding_File_contact_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeServer).
					GetTime(
						ctx,
						in1.(*GetTimeRequest),
					)
			}, DRPCNodeServer.GetTime, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterNode(mux drpc.Mux, impl DRPCNodeServer) error {
	return mux.Register(impl, DRPCNodeDescription{})
}

type DRPCNode_CheckInStream interface {
	drpc.Stream
	SendAndClose(*CheckInResponse) error
}

type drpcNode_CheckInStream struct {
	drpc.Stream
}

func (x *drpcNode_CheckInStream) SendAndClose(m *CheckInResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_contact_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNode_PingMeStream interface {
	drpc.Stream
	SendAndClose(*PingMeResponse) error
}

type drpcNode_PingMeStream struct {
	drpc.Stream
}

func (x *drpcNode_PingMeStream) SendAndClose(m *PingMeResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_contact_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNode_GetTimeStream interface {
	drpc.Stream
	SendAndClose(*GetTimeResponse) error
}

type drpcNode_GetTimeStream struct {
	drpc.Stream
}

func (x *drpcNode_GetTimeStream) SendAndClose(m *GetTimeResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_contact_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
