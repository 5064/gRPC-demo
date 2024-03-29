// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language.proto

package language

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

// The request message containing country's name.
type CountryLanguageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountryLanguageRequest) Reset()         { *m = CountryLanguageRequest{} }
func (m *CountryLanguageRequest) String() string { return proto.CompactTextString(m) }
func (*CountryLanguageRequest) ProtoMessage()    {}
func (*CountryLanguageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{0}
}

func (m *CountryLanguageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountryLanguageRequest.Unmarshal(m, b)
}
func (m *CountryLanguageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountryLanguageRequest.Marshal(b, m, deterministic)
}
func (m *CountryLanguageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountryLanguageRequest.Merge(m, src)
}
func (m *CountryLanguageRequest) XXX_Size() int {
	return xxx_messageInfo_CountryLanguageRequest.Size(m)
}
func (m *CountryLanguageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountryLanguageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountryLanguageRequest proto.InternalMessageInfo

func (m *CountryLanguageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the languages percentage spoken in the country
type CountryLanguageResponse struct {
	Languages            []*Language `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CountryLanguageResponse) Reset()         { *m = CountryLanguageResponse{} }
func (m *CountryLanguageResponse) String() string { return proto.CompactTextString(m) }
func (*CountryLanguageResponse) ProtoMessage()    {}
func (*CountryLanguageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{1}
}

func (m *CountryLanguageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountryLanguageResponse.Unmarshal(m, b)
}
func (m *CountryLanguageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountryLanguageResponse.Marshal(b, m, deterministic)
}
func (m *CountryLanguageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountryLanguageResponse.Merge(m, src)
}
func (m *CountryLanguageResponse) XXX_Size() int {
	return xxx_messageInfo_CountryLanguageResponse.Size(m)
}
func (m *CountryLanguageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountryLanguageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountryLanguageResponse proto.InternalMessageInfo

func (m *CountryLanguageResponse) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

type Language struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Percentage           float32  `protobuf:"fixed32,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Language) Reset()         { *m = Language{} }
func (m *Language) String() string { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()    {}
func (*Language) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{2}
}

func (m *Language) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Language.Unmarshal(m, b)
}
func (m *Language) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Language.Marshal(b, m, deterministic)
}
func (m *Language) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Language.Merge(m, src)
}
func (m *Language) XXX_Size() int {
	return xxx_messageInfo_Language.Size(m)
}
func (m *Language) XXX_DiscardUnknown() {
	xxx_messageInfo_Language.DiscardUnknown(m)
}

var xxx_messageInfo_Language proto.InternalMessageInfo

func (m *Language) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Language) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func init() {
	proto.RegisterType((*CountryLanguageRequest)(nil), "language.CountryLanguageRequest")
	proto.RegisterType((*CountryLanguageResponse)(nil), "language.CountryLanguageResponse")
	proto.RegisterType((*Language)(nil), "language.Language")
}

func init() { proto.RegisterFile("language.proto", fileDescriptor_e123c61d1ddd0892) }

var fileDescriptor_e123c61d1ddd0892 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x49, 0xcc, 0x4b,
	0x2f, 0x4d, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x74,
	0xb8, 0xc4, 0x9c, 0xf3, 0x4b, 0xf3, 0x4a, 0x8a, 0x2a, 0x7d, 0xa0, 0x42, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x60, 0xb6, 0x92, 0x37, 0x97, 0x38, 0x86, 0xea, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x03, 0x2e, 0x4e, 0x98, 0xa1, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x42, 0x7a,
	0x70, 0x6b, 0xe1, 0xca, 0x11, 0x8a, 0x94, 0xec, 0xb8, 0x38, 0x60, 0xc2, 0xd8, 0x2c, 0x13, 0x92,
	0xe3, 0xe2, 0x2a, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0x49, 0x4c, 0x4f, 0x95, 0x60, 0x52, 0x60,
	0xd4, 0x60, 0x0a, 0x42, 0x12, 0x31, 0xca, 0xe1, 0xe2, 0x47, 0x73, 0x8c, 0x50, 0x24, 0x97, 0x90,
	0x7b, 0x6a, 0x09, 0xba, 0xa8, 0x02, 0xc2, 0x1d, 0xd8, 0xfd, 0x2a, 0xa5, 0x88, 0x47, 0x05, 0xc4,
	0x7f, 0x49, 0x6c, 0xe0, 0x90, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x99, 0xf6, 0xa9, 0x07,
	0x4b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CountryLanguageClient is the client API for CountryLanguage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountryLanguageClient interface {
	GetCountryLanguage(ctx context.Context, in *CountryLanguageRequest, opts ...grpc.CallOption) (*CountryLanguageResponse, error)
}

type countryLanguageClient struct {
	cc *grpc.ClientConn
}

func NewCountryLanguageClient(cc *grpc.ClientConn) CountryLanguageClient {
	return &countryLanguageClient{cc}
}

func (c *countryLanguageClient) GetCountryLanguage(ctx context.Context, in *CountryLanguageRequest, opts ...grpc.CallOption) (*CountryLanguageResponse, error) {
	out := new(CountryLanguageResponse)
	err := c.cc.Invoke(ctx, "/language.CountryLanguage/GetCountryLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryLanguageServer is the server API for CountryLanguage service.
type CountryLanguageServer interface {
	GetCountryLanguage(context.Context, *CountryLanguageRequest) (*CountryLanguageResponse, error)
}

// UnimplementedCountryLanguageServer can be embedded to have forward compatible implementations.
type UnimplementedCountryLanguageServer struct {
}

func (*UnimplementedCountryLanguageServer) GetCountryLanguage(ctx context.Context, req *CountryLanguageRequest) (*CountryLanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountryLanguage not implemented")
}

func RegisterCountryLanguageServer(s *grpc.Server, srv CountryLanguageServer) {
	s.RegisterService(&_CountryLanguage_serviceDesc, srv)
}

func _CountryLanguage_GetCountryLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryLanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryLanguageServer).GetCountryLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/language.CountryLanguage/GetCountryLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryLanguageServer).GetCountryLanguage(ctx, req.(*CountryLanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountryLanguage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "language.CountryLanguage",
	HandlerType: (*CountryLanguageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCountryLanguage",
			Handler:    _CountryLanguage_GetCountryLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language.proto",
}
