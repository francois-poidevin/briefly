// Code generated by protoc-gen-go. DO NOT EDIT.
// source: briefly/schema/v1/schema_api.proto

package schemav1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// GetUnShortCodedURLRequest request for a shorcoded URL to be unshortcoded
type GetUnShortCodedURLRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUnShortCodedURLRequest) Reset()         { *m = GetUnShortCodedURLRequest{} }
func (m *GetUnShortCodedURLRequest) String() string { return proto.CompactTextString(m) }
func (*GetUnShortCodedURLRequest) ProtoMessage()    {}
func (*GetUnShortCodedURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe56cef58ac92f6, []int{0}
}

func (m *GetUnShortCodedURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUnShortCodedURLRequest.Unmarshal(m, b)
}

func (m *GetUnShortCodedURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUnShortCodedURLRequest.Marshal(b, m, deterministic)
}

func (m *GetUnShortCodedURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUnShortCodedURLRequest.Merge(m, src)
}

func (m *GetUnShortCodedURLRequest) XXX_Size() int {
	return xxx_messageInfo_GetUnShortCodedURLRequest.Size(m)
}

func (m *GetUnShortCodedURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUnShortCodedURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUnShortCodedURLRequest proto.InternalMessageInfo

func (m *GetUnShortCodedURLRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// GetUnShortCodedURLResponse
type GetUnShortCodedURLResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUnShortCodedURLResponse) Reset()         { *m = GetUnShortCodedURLResponse{} }
func (m *GetUnShortCodedURLResponse) String() string { return proto.CompactTextString(m) }
func (*GetUnShortCodedURLResponse) ProtoMessage()    {}
func (*GetUnShortCodedURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe56cef58ac92f6, []int{1}
}

func (m *GetUnShortCodedURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUnShortCodedURLResponse.Unmarshal(m, b)
}

func (m *GetUnShortCodedURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUnShortCodedURLResponse.Marshal(b, m, deterministic)
}

func (m *GetUnShortCodedURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUnShortCodedURLResponse.Merge(m, src)
}

func (m *GetUnShortCodedURLResponse) XXX_Size() int {
	return xxx_messageInfo_GetUnShortCodedURLResponse.Size(m)
}

func (m *GetUnShortCodedURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUnShortCodedURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUnShortCodedURLResponse proto.InternalMessageInfo

func (m *GetUnShortCodedURLResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// GetShortCodeHashRequest request for a shorcoded URL to be unshortcoded
type GetShortCodeHashRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShortCodeHashRequest) Reset()         { *m = GetShortCodeHashRequest{} }
func (m *GetShortCodeHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetShortCodeHashRequest) ProtoMessage()    {}
func (*GetShortCodeHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe56cef58ac92f6, []int{2}
}

func (m *GetShortCodeHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShortCodeHashRequest.Unmarshal(m, b)
}

func (m *GetShortCodeHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShortCodeHashRequest.Marshal(b, m, deterministic)
}

func (m *GetShortCodeHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShortCodeHashRequest.Merge(m, src)
}

func (m *GetShortCodeHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetShortCodeHashRequest.Size(m)
}

func (m *GetShortCodeHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShortCodeHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShortCodeHashRequest proto.InternalMessageInfo

func (m *GetShortCodeHashRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// GetShortCodeHashResponse
type GetShortCodeHashResponse struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShortCodeHashResponse) Reset()         { *m = GetShortCodeHashResponse{} }
func (m *GetShortCodeHashResponse) String() string { return proto.CompactTextString(m) }
func (*GetShortCodeHashResponse) ProtoMessage()    {}
func (*GetShortCodeHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe56cef58ac92f6, []int{3}
}

func (m *GetShortCodeHashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShortCodeHashResponse.Unmarshal(m, b)
}

func (m *GetShortCodeHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShortCodeHashResponse.Marshal(b, m, deterministic)
}

func (m *GetShortCodeHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShortCodeHashResponse.Merge(m, src)
}

func (m *GetShortCodeHashResponse) XXX_Size() int {
	return xxx_messageInfo_GetShortCodeHashResponse.Size(m)
}

func (m *GetShortCodeHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShortCodeHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShortCodeHashResponse proto.InternalMessageInfo

func (m *GetShortCodeHashResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUnShortCodedURLRequest)(nil), "briefly.schema.v1.GetUnShortCodedURLRequest")
	proto.RegisterType((*GetUnShortCodedURLResponse)(nil), "briefly.schema.v1.GetUnShortCodedURLResponse")
	proto.RegisterType((*GetShortCodeHashRequest)(nil), "briefly.schema.v1.GetShortCodeHashRequest")
	proto.RegisterType((*GetShortCodeHashResponse)(nil), "briefly.schema.v1.GetShortCodeHashResponse")
}

func init() {
	proto.RegisterFile("briefly/schema/v1/schema_api.proto", fileDescriptor_bfe56cef58ac92f6)
}

var fileDescriptor_bfe56cef58ac92f6 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x2a, 0xca, 0x4c,
	0x4d, 0xcb, 0xa9, 0xd4, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f, 0x33, 0x84, 0xb2, 0xe2,
	0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xa1, 0x6a, 0xf4, 0x20, 0x32,
	0x7a, 0x65, 0x86, 0x4a, 0xfa, 0x5c, 0x92, 0xee, 0xa9, 0x25, 0xa1, 0x79, 0xc1, 0x19, 0xf9, 0x45,
	0x25, 0xce, 0xf9, 0x29, 0xa9, 0x29, 0xa1, 0x41, 0x3e, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x42, 0x5c, 0x2c, 0x19, 0x89, 0xc5, 0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60,
	0xb6, 0x92, 0x1e, 0x97, 0x14, 0x36, 0x0d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x02, 0x5c,
	0xcc, 0xa5, 0x45, 0x39, 0x50, 0x0d, 0x20, 0xa6, 0x92, 0x36, 0x97, 0xb8, 0x7b, 0x6a, 0x09, 0x5c,
	0xb5, 0x47, 0x62, 0x71, 0x06, 0xcc, 0x78, 0x4c, 0xc5, 0x7a, 0x5c, 0x12, 0x98, 0x8a, 0xa1, 0x46,
	0x63, 0x71, 0x8c, 0xd1, 0x6b, 0x46, 0x2e, 0xce, 0x60, 0xb0, 0x5f, 0x1c, 0x03, 0x3c, 0x85, 0x0a,
	0xb9, 0x84, 0x30, 0x9d, 0x26, 0xa4, 0xa3, 0x87, 0xe1, 0x6b, 0x3d, 0x9c, 0x5e, 0x96, 0xd2, 0x25,
	0x52, 0x35, 0xd4, 0x51, 0xd9, 0x5c, 0x02, 0xe8, 0x0e, 0x16, 0xd2, 0xc2, 0x6e, 0x04, 0xb6, 0x20,
	0x90, 0xd2, 0x26, 0x4a, 0x2d, 0xc4, 0x32, 0xa7, 0x46, 0x46, 0x2e, 0xe5, 0xb4, 0xa2, 0xc4, 0xbc,
	0xe4, 0xfc, 0xcc, 0x62, 0xbd, 0x82, 0xfc, 0xcc, 0x94, 0xd4, 0xb2, 0xcc, 0x3c, 0x4c, 0x03, 0x9c,
	0xf8, 0xa0, 0x41, 0x52, 0x90, 0x19, 0x00, 0x8a, 0xf6, 0x00, 0xc6, 0x28, 0x29, 0x8c, 0xc4, 0x61,
	0x0d, 0x61, 0x95, 0x19, 0x2e, 0x62, 0x62, 0x76, 0x0a, 0x8e, 0x58, 0xc5, 0x24, 0xe8, 0x04, 0x35,
	0x07, 0xa2, 0x59, 0x2f, 0xcc, 0xf0, 0x14, 0x5c, 0x2c, 0x06, 0x22, 0x16, 0x13, 0x66, 0x98, 0xc4,
	0x06, 0x4e, 0x49, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xa4, 0x69, 0xab, 0x6f, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchemaAPIClient is the client API for SchemaAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchemaAPIClient interface {
	// Get the unshortcoded URL.
	GetUnShortCodedURL(ctx context.Context, in *GetUnShortCodedURLRequest, opts ...grpc.CallOption) (*GetUnShortCodedURLResponse, error)
	// Get the Hash of the URL shortcoded.
	GetShortCodeHash(ctx context.Context, in *GetShortCodeHashRequest, opts ...grpc.CallOption) (*GetShortCodeHashResponse, error)
}

type schemaAPIClient struct {
	cc *grpc.ClientConn
}

func NewSchemaAPIClient(cc *grpc.ClientConn) SchemaAPIClient {
	return &schemaAPIClient{cc}
}

func (c *schemaAPIClient) GetUnShortCodedURL(ctx context.Context, in *GetUnShortCodedURLRequest, opts ...grpc.CallOption) (*GetUnShortCodedURLResponse, error) {
	out := new(GetUnShortCodedURLResponse)
	err := c.cc.Invoke(ctx, "/briefly.schema.v1.SchemaAPI/GetUnShortCodedURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaAPIClient) GetShortCodeHash(ctx context.Context, in *GetShortCodeHashRequest, opts ...grpc.CallOption) (*GetShortCodeHashResponse, error) {
	out := new(GetShortCodeHashResponse)
	err := c.cc.Invoke(ctx, "/briefly.schema.v1.SchemaAPI/GetShortCodeHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaAPIServer is the server API for SchemaAPI service.
type SchemaAPIServer interface {
	// Get the unshortcoded URL.
	GetUnShortCodedURL(context.Context, *GetUnShortCodedURLRequest) (*GetUnShortCodedURLResponse, error)
	// Get the Hash of the URL shortcoded.
	GetShortCodeHash(context.Context, *GetShortCodeHashRequest) (*GetShortCodeHashResponse, error)
}

// UnimplementedSchemaAPIServer can be embedded to have forward compatible implementations.
type UnimplementedSchemaAPIServer struct{}

func (*UnimplementedSchemaAPIServer) GetUnShortCodedURL(ctx context.Context, req *GetUnShortCodedURLRequest) (*GetUnShortCodedURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnShortCodedURL not implemented")
}

func (*UnimplementedSchemaAPIServer) GetShortCodeHash(ctx context.Context, req *GetShortCodeHashRequest) (*GetShortCodeHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortCodeHash not implemented")
}

func RegisterSchemaAPIServer(s *grpc.Server, srv SchemaAPIServer) {
	s.RegisterService(&_SchemaAPI_serviceDesc, srv)
}

func _SchemaAPI_GetUnShortCodedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnShortCodedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaAPIServer).GetUnShortCodedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/briefly.schema.v1.SchemaAPI/GetUnShortCodedURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaAPIServer).GetUnShortCodedURL(ctx, req.(*GetUnShortCodedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaAPI_GetShortCodeHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortCodeHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaAPIServer).GetShortCodeHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/briefly.schema.v1.SchemaAPI/GetShortCodeHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaAPIServer).GetShortCodeHash(ctx, req.(*GetShortCodeHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchemaAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "briefly.schema.v1.SchemaAPI",
	HandlerType: (*SchemaAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnShortCodedURL",
			Handler:    _SchemaAPI_GetUnShortCodedURL_Handler,
		},
		{
			MethodName: "GetShortCodeHash",
			Handler:    _SchemaAPI_GetShortCodeHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "briefly/schema/v1/schema_api.proto",
}
