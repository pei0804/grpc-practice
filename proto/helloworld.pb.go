// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	IntegerRequest
	IntegerResponse
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type IntegerRequest struct {
	Val int32 `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (m *IntegerRequest) Reset()                    { *m = IntegerRequest{} }
func (m *IntegerRequest) String() string            { return proto.CompactTextString(m) }
func (*IntegerRequest) ProtoMessage()               {}
func (*IntegerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IntegerRequest) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type IntegerResponse struct {
	Val int32 `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (m *IntegerResponse) Reset()                    { *m = IntegerResponse{} }
func (m *IntegerResponse) String() string            { return proto.CompactTextString(m) }
func (*IntegerResponse) ProtoMessage()               {}
func (*IntegerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IntegerResponse) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*IntegerRequest)(nil), "helloworld.IntegerRequest")
	proto.RegisterType((*IntegerResponse)(nil), "helloworld.IntegerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Integer(ctx context.Context, in *IntegerRequest, opts ...grpc.CallOption) (*IntegerResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Integer(ctx context.Context, in *IntegerRequest, opts ...grpc.CallOption) (*IntegerResponse, error) {
	out := new(IntegerResponse)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/Integer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Integer(context.Context, *IntegerRequest) (*IntegerResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Integer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Integer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Integer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Integer(ctx, req.(*IntegerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "Integer",
			Handler:    _Greeter_Integer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0xab, 0xa4, 0xc4, 0xc5, 0xe7, 0x99, 0x57, 0x92, 0x9a, 0x9e, 0x5a,
	0x04, 0x33, 0x4d, 0x80, 0x8b, 0xb9, 0x2c, 0x31, 0x07, 0xac, 0x8e, 0x35, 0x08, 0xc4, 0x54, 0x52,
	0xe6, 0xe2, 0x87, 0xab, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5, 0x54, 0x64, 0xd4, 0xcf, 0xc8,
	0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0x64, 0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09,
	0xb6, 0x5f, 0x48, 0x42, 0x0f, 0xc9, 0x2f, 0xc8, 0xce, 0x96, 0x12, 0xc3, 0x22, 0x53, 0x90, 0x53,
	0xa9, 0xc4, 0x20, 0xe4, 0xc2, 0xc5, 0x0e, 0xb5, 0x50, 0x48, 0x0a, 0x59, 0x11, 0xaa, 0x4b, 0xa5,
	0xa4, 0xb1, 0xca, 0x41, 0x5c, 0xa8, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x39, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0x39, 0xb0, 0x8c, 0x4d, 0x01, 0x00, 0x00,
}