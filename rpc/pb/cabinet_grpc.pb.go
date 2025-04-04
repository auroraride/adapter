// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: cabinet.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Cabinet_Sync_FullMethodName       = "/pb.Cabinet/Sync"
	Cabinet_Deactivate_FullMethodName = "/pb.Cabinet/Deactivate"
	Cabinet_Biz_FullMethodName        = "/pb.Cabinet/Biz"
	Cabinet_Interrupt_FullMethodName  = "/pb.Cabinet/Interrupt"
	Cabinet_Exchange_FullMethodName   = "/pb.Cabinet/Exchange"
	Cabinet_Statistic_FullMethodName  = "/pb.Cabinet/Statistic"
)

// CabinetClient is the client API for Cabinet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CabinetClient interface {
	// rpc Batch (CabinetBatchRequest) returns (CabinetBatchResponse);
	Sync(ctx context.Context, in *CabinetSyncRequest, opts ...grpc.CallOption) (*CabinetSyncResponse, error)
	Deactivate(ctx context.Context, in *CabinetDeactivateRequest, opts ...grpc.CallOption) (*CabinetDeactivateResponse, error)
	Biz(ctx context.Context, in *CabinetBizRequest, opts ...grpc.CallOption) (*CabinetBizResponse, error)
	Interrupt(ctx context.Context, in *CabinetInterruptRequest, opts ...grpc.CallOption) (*CabinetBizResponse, error)
	Exchange(ctx context.Context, in *CabinetExchangeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CabinetExchangeResponse], error)
	Statistic(ctx context.Context, in *CabinetStatisticRequest, opts ...grpc.CallOption) (*CabinetStatisticResponse, error)
}

type cabinetClient struct {
	cc grpc.ClientConnInterface
}

func NewCabinetClient(cc grpc.ClientConnInterface) CabinetClient {
	return &cabinetClient{cc}
}

func (c *cabinetClient) Sync(ctx context.Context, in *CabinetSyncRequest, opts ...grpc.CallOption) (*CabinetSyncResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CabinetSyncResponse)
	err := c.cc.Invoke(ctx, Cabinet_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cabinetClient) Deactivate(ctx context.Context, in *CabinetDeactivateRequest, opts ...grpc.CallOption) (*CabinetDeactivateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CabinetDeactivateResponse)
	err := c.cc.Invoke(ctx, Cabinet_Deactivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cabinetClient) Biz(ctx context.Context, in *CabinetBizRequest, opts ...grpc.CallOption) (*CabinetBizResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CabinetBizResponse)
	err := c.cc.Invoke(ctx, Cabinet_Biz_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cabinetClient) Interrupt(ctx context.Context, in *CabinetInterruptRequest, opts ...grpc.CallOption) (*CabinetBizResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CabinetBizResponse)
	err := c.cc.Invoke(ctx, Cabinet_Interrupt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cabinetClient) Exchange(ctx context.Context, in *CabinetExchangeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CabinetExchangeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Cabinet_ServiceDesc.Streams[0], Cabinet_Exchange_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CabinetExchangeRequest, CabinetExchangeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Cabinet_ExchangeClient = grpc.ServerStreamingClient[CabinetExchangeResponse]

func (c *cabinetClient) Statistic(ctx context.Context, in *CabinetStatisticRequest, opts ...grpc.CallOption) (*CabinetStatisticResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CabinetStatisticResponse)
	err := c.cc.Invoke(ctx, Cabinet_Statistic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CabinetServer is the server API for Cabinet service.
// All implementations must embed UnimplementedCabinetServer
// for forward compatibility.
type CabinetServer interface {
	// rpc Batch (CabinetBatchRequest) returns (CabinetBatchResponse);
	Sync(context.Context, *CabinetSyncRequest) (*CabinetSyncResponse, error)
	Deactivate(context.Context, *CabinetDeactivateRequest) (*CabinetDeactivateResponse, error)
	Biz(context.Context, *CabinetBizRequest) (*CabinetBizResponse, error)
	Interrupt(context.Context, *CabinetInterruptRequest) (*CabinetBizResponse, error)
	Exchange(*CabinetExchangeRequest, grpc.ServerStreamingServer[CabinetExchangeResponse]) error
	Statistic(context.Context, *CabinetStatisticRequest) (*CabinetStatisticResponse, error)
	mustEmbedUnimplementedCabinetServer()
}

// UnimplementedCabinetServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCabinetServer struct{}

func (UnimplementedCabinetServer) Sync(context.Context, *CabinetSyncRequest) (*CabinetSyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedCabinetServer) Deactivate(context.Context, *CabinetDeactivateRequest) (*CabinetDeactivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedCabinetServer) Biz(context.Context, *CabinetBizRequest) (*CabinetBizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Biz not implemented")
}
func (UnimplementedCabinetServer) Interrupt(context.Context, *CabinetInterruptRequest) (*CabinetBizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interrupt not implemented")
}
func (UnimplementedCabinetServer) Exchange(*CabinetExchangeRequest, grpc.ServerStreamingServer[CabinetExchangeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedCabinetServer) Statistic(context.Context, *CabinetStatisticRequest) (*CabinetStatisticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Statistic not implemented")
}
func (UnimplementedCabinetServer) mustEmbedUnimplementedCabinetServer() {}
func (UnimplementedCabinetServer) testEmbeddedByValue()                 {}

// UnsafeCabinetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CabinetServer will
// result in compilation errors.
type UnsafeCabinetServer interface {
	mustEmbedUnimplementedCabinetServer()
}

func RegisterCabinetServer(s grpc.ServiceRegistrar, srv CabinetServer) {
	// If the following call pancis, it indicates UnimplementedCabinetServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Cabinet_ServiceDesc, srv)
}

func _Cabinet_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CabinetSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabinetServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cabinet_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabinetServer).Sync(ctx, req.(*CabinetSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cabinet_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CabinetDeactivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabinetServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cabinet_Deactivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabinetServer).Deactivate(ctx, req.(*CabinetDeactivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cabinet_Biz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CabinetBizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabinetServer).Biz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cabinet_Biz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabinetServer).Biz(ctx, req.(*CabinetBizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cabinet_Interrupt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CabinetInterruptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabinetServer).Interrupt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cabinet_Interrupt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabinetServer).Interrupt(ctx, req.(*CabinetInterruptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cabinet_Exchange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CabinetExchangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CabinetServer).Exchange(m, &grpc.GenericServerStream[CabinetExchangeRequest, CabinetExchangeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Cabinet_ExchangeServer = grpc.ServerStreamingServer[CabinetExchangeResponse]

func _Cabinet_Statistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CabinetStatisticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabinetServer).Statistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cabinet_Statistic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabinetServer).Statistic(ctx, req.(*CabinetStatisticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cabinet_ServiceDesc is the grpc.ServiceDesc for Cabinet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cabinet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Cabinet",
	HandlerType: (*CabinetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _Cabinet_Sync_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _Cabinet_Deactivate_Handler,
		},
		{
			MethodName: "Biz",
			Handler:    _Cabinet_Biz_Handler,
		},
		{
			MethodName: "Interrupt",
			Handler:    _Cabinet_Interrupt_Handler,
		},
		{
			MethodName: "Statistic",
			Handler:    _Cabinet_Statistic_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exchange",
			Handler:       _Cabinet_Exchange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cabinet.proto",
}
