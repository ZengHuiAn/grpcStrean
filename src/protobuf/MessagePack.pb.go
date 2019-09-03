// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MessagePack.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("MessagePack.proto", fileDescriptor_2bfb66b5d26e7d96) }

var fileDescriptor_2bfb66b5d26e7d96 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf4, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x0d, 0x48, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0x53, 0x49, 0xa5, 0x69, 0x52, 0xc2, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x41, 0xa9, 0xe9,
	0x99, 0xc5, 0x25, 0x10, 0x69, 0xa3, 0x54, 0x2e, 0xde, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0xa8,
	0xa4, 0x50, 0x08, 0x17, 0x2f, 0x44, 0x01, 0x4c, 0x40, 0x41, 0x0f, 0x66, 0x82, 0x1e, 0x44, 0x22,
	0xb5, 0x08, 0x6e, 0x50, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94, 0x22, 0x1e, 0x15, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x49, 0x6c, 0x60, 0x15, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x80,
	0x39, 0x30, 0xa1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	RegistService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) RegistService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error) {
	out := new(RegisterServiceResponse)
	err := c.cc.Invoke(ctx, "/protobuf.StreamService/RegistService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	RegistService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error)
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_RegistService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).RegistService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.StreamService/RegistService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).RegistService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistService",
			Handler:    _StreamService_RegistService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MessagePack.proto",
}