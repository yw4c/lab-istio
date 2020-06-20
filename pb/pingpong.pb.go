// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/pingpong.proto

package pingpong

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingPong struct {
	Ping                 int64    `protobuf:"varint,1,opt,name=ping,proto3" json:"ping,omitempty"`
	Pong                 int64    `protobuf:"varint,2,opt,name=pong,proto3" json:"pong,omitempty"`
	Next                 int64    `protobuf:"varint,3,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingPong) Reset()         { *m = PingPong{} }
func (m *PingPong) String() string { return proto.CompactTextString(m) }
func (*PingPong) ProtoMessage()    {}
func (*PingPong) Descriptor() ([]byte, []int) {
	return fileDescriptor_493895313ee6c7da, []int{0}
}

func (m *PingPong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingPong.Unmarshal(m, b)
}
func (m *PingPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingPong.Marshal(b, m, deterministic)
}
func (m *PingPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingPong.Merge(m, src)
}
func (m *PingPong) XXX_Size() int {
	return xxx_messageInfo_PingPong.Size(m)
}
func (m *PingPong) XXX_DiscardUnknown() {
	xxx_messageInfo_PingPong.DiscardUnknown(m)
}

var xxx_messageInfo_PingPong proto.InternalMessageInfo

func (m *PingPong) GetPing() int64 {
	if m != nil {
		return m.Ping
	}
	return 0
}

func (m *PingPong) GetPong() int64 {
	if m != nil {
		return m.Pong
	}
	return 0
}

func (m *PingPong) GetNext() int64 {
	if m != nil {
		return m.Next
	}
	return 0
}

func init() {
	proto.RegisterType((*PingPong)(nil), "pingpong.PingPong")
}

func init() {
	proto.RegisterFile("pb/pingpong.proto", fileDescriptor_493895313ee6c7da)
}

var fileDescriptor_493895313ee6c7da = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0xd2, 0x2f,
	0xc8, 0xcc, 0x4b, 0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80,
	0xf1, 0x95, 0xdc, 0xb8, 0x38, 0x02, 0x32, 0xf3, 0xd2, 0x03, 0xf2, 0xf3, 0xd2, 0x85, 0x84, 0xb8,
	0x58, 0x40, 0xe2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0x36, 0x58, 0x2c, 0x3f, 0x2f,
	0x5d, 0x82, 0x09, 0x2a, 0x06, 0x55, 0x97, 0x97, 0x5a, 0x51, 0x22, 0xc1, 0x0c, 0x11, 0x03, 0xb1,
	0x8d, 0x7c, 0xb9, 0xf8, 0x61, 0xe6, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59, 0x71,
	0x09, 0xc0, 0x84, 0x5c, 0xf3, 0x52, 0x0a, 0xf2, 0x33, 0xf3, 0x4a, 0x84, 0x84, 0xf4, 0xe0, 0x2e,
	0x81, 0xc9, 0x49, 0x61, 0x11, 0x4b, 0x62, 0x03, 0xbb, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0x40, 0x68, 0x24, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PingPongServiceClient is the client API for PingPongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongServiceClient interface {
	PingPongEndpoint(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error)
}

type pingPongServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongServiceClient(cc grpc.ClientConnInterface) PingPongServiceClient {
	return &pingPongServiceClient{cc}
}

func (c *pingPongServiceClient) PingPongEndpoint(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error) {
	out := new(PingPong)
	err := c.cc.Invoke(ctx, "/pingpong.PingPongService/PingPongEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServiceServer is the server API for PingPongService service.
type PingPongServiceServer interface {
	PingPongEndpoint(context.Context, *PingPong) (*PingPong, error)
}

// UnimplementedPingPongServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPingPongServiceServer struct {
}

func (*UnimplementedPingPongServiceServer) PingPongEndpoint(ctx context.Context, req *PingPong) (*PingPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPongEndpoint not implemented")
}

func RegisterPingPongServiceServer(s *grpc.Server, srv PingPongServiceServer) {
	s.RegisterService(&_PingPongService_serviceDesc, srv)
}

func _PingPongService_PingPongEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServiceServer).PingPongEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PingPongService/PingPongEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServiceServer).PingPongEndpoint(ctx, req.(*PingPong))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPongService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPongService",
	HandlerType: (*PingPongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingPongEndpoint",
			Handler:    _PingPongService_PingPongEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pingpong.proto",
}
