// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/service.proto

package rpc

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

type PingRequest struct {
	Sequence             uint64   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Time                 *Time    `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PingRequest) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *PingRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PingResponse struct {
	Time                 *Time        `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Request              *PingRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *PingResponse) GetRequest() *PingRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type Time struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int64    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{2}
}

func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Time) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "grpcntt.rpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "grpcntt.rpc.PingResponse")
	proto.RegisterType((*Time)(nil), "grpcntt.rpc.Time")
}

func init() { proto.RegisterFile("rpc/service.proto", fileDescriptor_ec64d44e618a02a6) }

var fileDescriptor_ec64d44e618a02a6 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0xed, 0xac, 0xbc, 0xee, 0xb2, 0xe0, 0x21, 0xee, 0x54, 0x0a, 0x42, 0x4f, 0x15,
	0x2a, 0x78, 0xd7, 0xb3, 0x07, 0xc9, 0x76, 0xf2, 0x56, 0xb3, 0xc7, 0x88, 0xac, 0xc9, 0x33, 0x2f,
	0x0a, 0xfe, 0xf7, 0xb2, 0x66, 0xd5, 0x09, 0xba, 0x5b, 0x3e, 0xde, 0xe3, 0xf7, 0xfb, 0xf2, 0x60,
	0x11, 0xc8, 0xdc, 0x30, 0x86, 0x0f, 0x6b, 0xb0, 0xa5, 0xe0, 0xa3, 0x97, 0xe5, 0x36, 0x90, 0x71,
	0x31, 0xb6, 0x81, 0x4c, 0xfd, 0x0a, 0xe5, 0x93, 0x75, 0x5b, 0x8d, 0x6f, 0xef, 0xc8, 0x51, 0x2e,
	0xe1, 0x82, 0xf7, 0x4f, 0x67, 0x50, 0x89, 0x4a, 0x34, 0xb9, 0xfe, 0xce, 0xf2, 0x1a, 0xf2, 0x68,
	0x07, 0x54, 0x67, 0x95, 0x68, 0xca, 0x6e, 0xd1, 0x1e, 0x61, 0xda, 0xb5, 0x1d, 0x50, 0x8f, 0x63,
	0xa9, 0xa0, 0xa0, 0xfe, 0x73, 0xe7, 0xfb, 0x8d, 0xca, 0x2a, 0xd1, 0xcc, 0xf5, 0x14, 0x6b, 0x0b,
	0xf3, 0xe4, 0x62, 0xf2, 0x8e, 0x7f, 0x80, 0xe2, 0x34, 0xb0, 0x83, 0x22, 0xa4, 0x7a, 0x07, 0xb5,
	0xfa, 0xb5, 0x79, 0x54, 0x5f, 0x4f, 0x8b, 0xf5, 0x1d, 0xe4, 0xeb, 0x43, 0x19, 0x46, 0xe3, 0xdd,
	0x86, 0x47, 0x4b, 0xa6, 0xa7, 0x28, 0x2f, 0x61, 0xe6, 0x7a, 0xe7, 0x79, 0x64, 0x66, 0x3a, 0x85,
	0xee, 0x11, 0x8a, 0x55, 0x3a, 0x96, 0xbc, 0x07, 0x58, 0xd9, 0x81, 0x76, 0xb8, 0x17, 0xc8, 0x7f,
	0x9d, 0xcb, 0xab, 0x3f, 0x26, 0xe9, 0x83, 0x0f, 0xb3, 0xe7, 0x2c, 0x90, 0x79, 0x39, 0x1f, 0xef,
	0x7e, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xdd, 0x88, 0x52, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	SimplePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SimplePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/grpcntt.rpc.Service/SimplePing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	SimplePing(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_SimplePing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SimplePing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcntt.rpc.Service/SimplePing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SimplePing(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcntt.rpc.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SimplePing",
			Handler:    _Service_SimplePing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/service.proto",
}
