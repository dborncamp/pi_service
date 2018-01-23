// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eyepi.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	eyepi.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Defines the request type for the `HelloProxy` method.
type HelloRequest struct {
	HelloText string `protobuf:"bytes,1,opt,name=hello_text,json=helloText" json:"hello_text,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetHelloText() string {
	if m != nil {
		return m.HelloText
	}
	return ""
}

// Defines the response type for the `HelloProxy` method.
type HelloResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protobuf.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "protobuf.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EyePi service

type EyePiClient interface {
	// HelloProxy says 'hello' in a form that is handled by the gateway proxy.
	HelloProxy(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type eyePiClient struct {
	cc *grpc.ClientConn
}

func NewEyePiClient(cc *grpc.ClientConn) EyePiClient {
	return &eyePiClient{cc}
}

func (c *eyePiClient) HelloProxy(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/protobuf.EyePi/HelloProxy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EyePi service

type EyePiServer interface {
	// HelloProxy says 'hello' in a form that is handled by the gateway proxy.
	HelloProxy(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterEyePiServer(s *grpc.Server, srv EyePiServer) {
	s.RegisterService(&_EyePi_serviceDesc, srv)
}

func _EyePi_HelloProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EyePiServer).HelloProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.EyePi/HelloProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EyePiServer).HelloProxy(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EyePi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.EyePi",
	HandlerType: (*EyePiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloProxy",
			Handler:    _EyePi_HelloProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eyepi.proto",
}

func init() { proto.RegisterFile("eyepi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xad, 0x4c, 0x2d,
	0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x52, 0x32,
	0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25,
	0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x75, 0x4a, 0xba, 0x5c, 0x3c, 0x1e, 0xa9, 0x39, 0x39,
	0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xb2, 0x5c, 0x5c, 0x19, 0x20, 0x7e, 0x7c,
	0x49, 0x6a, 0x45, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x27, 0x58, 0x24, 0x24, 0xb5,
	0xa2, 0x44, 0x49, 0x99, 0x8b, 0x17, 0xaa, 0xbc, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x88,
	0x8b, 0x05, 0x49, 0x25, 0x98, 0x6d, 0x94, 0xca, 0xc5, 0xea, 0x5a, 0x99, 0x1a, 0x90, 0x29, 0x14,
	0xc3, 0xc5, 0x05, 0x56, 0x1d, 0x50, 0x94, 0x5f, 0x51, 0x29, 0x24, 0xa6, 0x07, 0x73, 0x93, 0x1e,
	0xb2, 0x95, 0x52, 0xe2, 0x18, 0xe2, 0x10, 0xb3, 0x95, 0x64, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0x24,
	0x2e, 0x24, 0xaa, 0x0f, 0x36, 0x4e, 0xbf, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0xb5, 0x58, 0x1f,
	0xec, 0x9e, 0x24, 0x36, 0xb0, 0x36, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0xc4, 0x71,
	0xe9, 0xf8, 0x00, 0x00, 0x00,
}