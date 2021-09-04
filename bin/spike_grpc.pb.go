// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SpikeClient is the client API for Spike service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpikeClient interface {
	MonitorEvent(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Spike_MonitorEventClient, error)
	Subscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Spike_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Success, error)
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Success, error)
}

type spikeClient struct {
	cc grpc.ClientConnInterface
}

func NewSpikeClient(cc grpc.ClientConnInterface) SpikeClient {
	return &spikeClient{cc}
}

func (c *spikeClient) MonitorEvent(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Spike_MonitorEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &Spike_ServiceDesc.Streams[0], "/bin.Spike/MonitorEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &spikeMonitorEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Spike_MonitorEventClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type spikeMonitorEventClient struct {
	grpc.ClientStream
}

func (x *spikeMonitorEventClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *spikeClient) Subscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Spike_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Spike_ServiceDesc.Streams[1], "/bin.Spike/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &spikeSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Spike_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type spikeSubscribeClient struct {
	grpc.ClientStream
}

func (x *spikeSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *spikeClient) Unsubscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/bin.Spike/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spikeClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/bin.Spike/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpikeServer is the server API for Spike service.
// All implementations should embed UnimplementedSpikeServer
// for forward compatibility
type SpikeServer interface {
	MonitorEvent(*Topic, Spike_MonitorEventServer) error
	Subscribe(*Topic, Spike_SubscribeServer) error
	Unsubscribe(context.Context, *Topic) (*Success, error)
	Publish(context.Context, *Message) (*Success, error)
}

// UnimplementedSpikeServer should be embedded to have forward compatible implementations.
type UnimplementedSpikeServer struct {
}

func (UnimplementedSpikeServer) MonitorEvent(*Topic, Spike_MonitorEventServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorEvent not implemented")
}
func (UnimplementedSpikeServer) Subscribe(*Topic, Spike_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSpikeServer) Unsubscribe(context.Context, *Topic) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedSpikeServer) Publish(context.Context, *Message) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}

// UnsafeSpikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpikeServer will
// result in compilation errors.
type UnsafeSpikeServer interface {
	mustEmbedUnimplementedSpikeServer()
}

func RegisterSpikeServer(s grpc.ServiceRegistrar, srv SpikeServer) {
	s.RegisterService(&Spike_ServiceDesc, srv)
}

func _Spike_MonitorEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpikeServer).MonitorEvent(m, &spikeMonitorEventServer{stream})
}

type Spike_MonitorEventServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type spikeMonitorEventServer struct {
	grpc.ServerStream
}

func (x *spikeMonitorEventServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Spike_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpikeServer).Subscribe(m, &spikeSubscribeServer{stream})
}

type Spike_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type spikeSubscribeServer struct {
	grpc.ServerStream
}

func (x *spikeSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Spike_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpikeServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bin.Spike/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpikeServer).Unsubscribe(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spike_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpikeServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bin.Spike/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpikeServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Spike_ServiceDesc is the grpc.ServiceDesc for Spike service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Spike_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bin.Spike",
	HandlerType: (*SpikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _Spike_Unsubscribe_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Spike_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorEvent",
			Handler:       _Spike_MonitorEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Spike_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bin/spike.proto",
}