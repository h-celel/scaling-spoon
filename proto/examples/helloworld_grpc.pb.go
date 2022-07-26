// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: examples/helloworld.proto

package examples

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	StreamingHello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_StreamingHelloClient, error)
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (Service_BidiStreamClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/helloworld.Service/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) StreamingHello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_StreamingHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/helloworld.Service/StreamingHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceStreamingHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_StreamingHelloClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type serviceStreamingHelloClient struct {
	grpc.ClientStream
}

func (x *serviceStreamingHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (Service_BidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[1], "/helloworld.Service/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceBidiStreamClient{stream}
	return x, nil
}

type Service_BidiStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type serviceBidiStreamClient struct {
	grpc.ClientStream
}

func (x *serviceBidiStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceBidiStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	StreamingHello(*Empty, Service_StreamingHelloServer) error
	BidiStream(Service_BidiStreamServer) error
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedServiceServer) StreamingHello(*Empty, Service_StreamingHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingHello not implemented")
}
func (UnimplementedServiceServer) BidiStream(Service_BidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiStream not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Service/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_StreamingHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).StreamingHello(m, &serviceStreamingHelloServer{stream})
}

type Service_StreamingHelloServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type serviceStreamingHelloServer struct {
	grpc.ServerStream
}

func (x *serviceStreamingHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_BidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).BidiStream(&serviceBidiStreamServer{stream})
}

type Service_BidiStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type serviceBidiStreamServer struct {
	grpc.ServerStream
}

func (x *serviceBidiStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceBidiStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Service_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingHello",
			Handler:       _Service_StreamingHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidiStream",
			Handler:       _Service_BidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/helloworld.proto",
}
