// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/authregistry/v0alpha/authregistry.proto

package authregistryv0alphapb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc"
	types "github.com/cs3org/go-cs3apis/cs3/types"
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

type GetAuthProviderRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The type of authentication provider.
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthProviderRequest) Reset()         { *m = GetAuthProviderRequest{} }
func (m *GetAuthProviderRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthProviderRequest) ProtoMessage()    {}
func (*GetAuthProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a936fc685c3709ca, []int{0}
}

func (m *GetAuthProviderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthProviderRequest.Unmarshal(m, b)
}
func (m *GetAuthProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthProviderRequest.Marshal(b, m, deterministic)
}
func (m *GetAuthProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthProviderRequest.Merge(m, src)
}
func (m *GetAuthProviderRequest) XXX_Size() int {
	return xxx_messageInfo_GetAuthProviderRequest.Size(m)
}
func (m *GetAuthProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthProviderRequest proto.InternalMessageInfo

func (m *GetAuthProviderRequest) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetAuthProviderRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetAuthProviderResponse struct {
	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The auth provider handling the requested auth call.
	Provider             *ProviderInfo `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAuthProviderResponse) Reset()         { *m = GetAuthProviderResponse{} }
func (m *GetAuthProviderResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthProviderResponse) ProtoMessage()    {}
func (*GetAuthProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a936fc685c3709ca, []int{1}
}

func (m *GetAuthProviderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthProviderResponse.Unmarshal(m, b)
}
func (m *GetAuthProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthProviderResponse.Marshal(b, m, deterministic)
}
func (m *GetAuthProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthProviderResponse.Merge(m, src)
}
func (m *GetAuthProviderResponse) XXX_Size() int {
	return xxx_messageInfo_GetAuthProviderResponse.Size(m)
}
func (m *GetAuthProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthProviderResponse proto.InternalMessageInfo

func (m *GetAuthProviderResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetAuthProviderResponse) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetAuthProviderResponse) GetProvider() *ProviderInfo {
	if m != nil {
		return m.Provider
	}
	return nil
}

type ListAuthProvidersRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque               *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListAuthProvidersRequest) Reset()         { *m = ListAuthProvidersRequest{} }
func (m *ListAuthProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*ListAuthProvidersRequest) ProtoMessage()    {}
func (*ListAuthProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a936fc685c3709ca, []int{2}
}

func (m *ListAuthProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthProvidersRequest.Unmarshal(m, b)
}
func (m *ListAuthProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthProvidersRequest.Marshal(b, m, deterministic)
}
func (m *ListAuthProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthProvidersRequest.Merge(m, src)
}
func (m *ListAuthProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_ListAuthProvidersRequest.Size(m)
}
func (m *ListAuthProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthProvidersRequest proto.InternalMessageInfo

func (m *ListAuthProvidersRequest) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type ListAuthProvidersResponse struct {
	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The list of auth providers this registry knows about.
	Providers            []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListAuthProvidersResponse) Reset()         { *m = ListAuthProvidersResponse{} }
func (m *ListAuthProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*ListAuthProvidersResponse) ProtoMessage()    {}
func (*ListAuthProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a936fc685c3709ca, []int{3}
}

func (m *ListAuthProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthProvidersResponse.Unmarshal(m, b)
}
func (m *ListAuthProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthProvidersResponse.Marshal(b, m, deterministic)
}
func (m *ListAuthProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthProvidersResponse.Merge(m, src)
}
func (m *ListAuthProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_ListAuthProvidersResponse.Size(m)
}
func (m *ListAuthProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthProvidersResponse proto.InternalMessageInfo

func (m *ListAuthProvidersResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListAuthProvidersResponse) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ListAuthProvidersResponse) GetProviders() []*ProviderInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAuthProviderRequest)(nil), "cs3.authregistryv0alpha.GetAuthProviderRequest")
	proto.RegisterType((*GetAuthProviderResponse)(nil), "cs3.authregistryv0alpha.GetAuthProviderResponse")
	proto.RegisterType((*ListAuthProvidersRequest)(nil), "cs3.authregistryv0alpha.ListAuthProvidersRequest")
	proto.RegisterType((*ListAuthProvidersResponse)(nil), "cs3.authregistryv0alpha.ListAuthProvidersResponse")
}

func init() {
	proto.RegisterFile("cs3/authregistry/v0alpha/authregistry.proto", fileDescriptor_a936fc685c3709ca)
}

var fileDescriptor_a936fc685c3709ca = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0x4d, 0xcb, 0x0d, 0xb9, 0xcc, 0x5d, 0x70, 0x19, 0x45, 0x6a, 0xdd, 0x90, 0x26, 0x46, 0x8c,
	0x49, 0x8b, 0xf4, 0x09, 0x0a, 0x21, 0x48, 0x62, 0x02, 0x99, 0xe2, 0x6f, 0xdc, 0x94, 0x3a, 0x4a,
	0x13, 0x65, 0x86, 0x99, 0x29, 0x09, 0xbe, 0x80, 0x5b, 0x9f, 0xc1, 0xa5, 0x2b, 0x17, 0x3e, 0x85,
	0x4f, 0x65, 0xfa, 0x83, 0x16, 0x69, 0x23, 0x2c, 0xdc, 0x10, 0x72, 0xbe, 0x73, 0xce, 0x9c, 0xef,
	0x4c, 0x07, 0x1c, 0xb8, 0xdc, 0x34, 0x1c, 0x5f, 0x8c, 0x18, 0xbe, 0xf5, 0xb8, 0x60, 0x33, 0x63,
	0x5a, 0x77, 0xee, 0xe8, 0xc8, 0x59, 0x00, 0x75, 0xca, 0x88, 0x20, 0xb0, 0xe2, 0x72, 0x53, 0x4f,
	0xe2, 0x31, 0x57, 0xad, 0x65, 0xba, 0x30, 0xcc, 0x89, 0xcf, 0x5c, 0xcc, 0x23, 0x0b, 0x75, 0x33,
	0x60, 0x32, 0xea, 0x1a, 0x5c, 0x38, 0xc2, 0x9f, 0xa3, 0xe5, 0x00, 0x15, 0x33, 0x8a, 0x79, 0xf4,
	0x1b, 0xc1, 0xda, 0x19, 0xd8, 0xea, 0x60, 0x61, 0xf9, 0x62, 0xd4, 0x67, 0x64, 0xea, 0x5d, 0x63,
	0x86, 0xf0, 0xc4, 0xc7, 0x5c, 0xc0, 0x7d, 0x90, 0x27, 0xd4, 0x99, 0xf8, 0x58, 0x91, 0xaa, 0x52,
	0xed, 0x5f, 0xa3, 0xa4, 0x07, 0xd1, 0x22, 0x6d, 0x2f, 0x1c, 0xa0, 0x98, 0x00, 0x21, 0xf8, 0x13,
	0xe0, 0x8a, 0x5c, 0x95, 0x6a, 0x05, 0x14, 0xfe, 0xd7, 0x5e, 0x25, 0x50, 0x59, 0x72, 0xe6, 0x94,
	0x8c, 0x39, 0x86, 0x7b, 0x20, 0x1f, 0x65, 0x8b, 0xad, 0x8b, 0xa1, 0x35, 0xa3, 0xae, 0x6e, 0x87,
	0x30, 0x8a, 0xc7, 0x89, 0x0c, 0xf2, 0x4f, 0x19, 0x2c, 0xf0, 0x97, 0xc6, 0xe7, 0x28, 0xb9, 0x90,
	0xbc, 0xab, 0x67, 0x74, 0xa9, 0xcf, 0x03, 0x75, 0xc7, 0x37, 0x04, 0x7d, 0xca, 0xb4, 0x36, 0x50,
	0x8e, 0x3d, 0xbe, 0x10, 0x99, 0xaf, 0xdf, 0x86, 0xf6, 0x26, 0x81, 0xed, 0x14, 0x9f, 0x5f, 0xdc,
	0xbd, 0x05, 0x0a, 0xf3, 0x25, 0xb8, 0x92, 0xab, 0xe6, 0x56, 0x5f, 0xfe, 0x4b, 0xd7, 0x78, 0x94,
	0xc1, 0x46, 0x10, 0x19, 0xc5, 0x7c, 0x1b, 0xb3, 0xa9, 0xe7, 0x62, 0x28, 0x40, 0xf1, 0xdb, 0x3d,
	0x42, 0x23, 0xd3, 0x3c, 0xfd, 0x5b, 0x52, 0xeb, 0xab, 0x0b, 0xe2, 0x9a, 0x1e, 0x40, 0x69, 0xa9,
	0x43, 0x78, 0x98, 0x69, 0x93, 0x75, 0x6f, 0x6a, 0x63, 0x1d, 0x49, 0x74, 0x76, 0xf3, 0x49, 0x02,
	0x3b, 0x2e, 0xb9, 0xcf, 0x52, 0x36, 0x4b, 0x56, 0x02, 0xec, 0x07, 0xcf, 0xa8, 0x2f, 0x5d, 0x96,
	0x53, 0x98, 0x74, 0xf8, 0x2c, 0xff, 0x6f, 0x35, 0x7b, 0xe7, 0xd6, 0xc9, 0xe0, 0x08, 0xb5, 0x3b,
	0x5d, 0x7b, 0x80, 0x2e, 0x5e, 0xe4, 0x4a, 0xcb, 0x36, 0xf5, 0x64, 0xd5, 0xa7, 0x75, 0x2b, 0xa0,
	0xbf, 0x87, 0x93, 0xab, 0x94, 0xc9, 0x30, 0x1f, 0xbe, 0x56, 0xf3, 0x23, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0x5a, 0x29, 0x4b, 0x4c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthRegistryServiceClient is the client API for AuthRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthRegistryServiceClient interface {
	// Returns the auth provider that is reponsible for the given
	// resource reference.
	// MUST return CODE_NOT_FOUND if the reference does not exist.
	GetAuthProvider(ctx context.Context, in *GetAuthProviderRequest, opts ...grpc.CallOption) (*GetAuthProviderResponse, error)
	// Returns a list of the available auth providers known by this registry.
	ListAuthProviders(ctx context.Context, in *ListAuthProvidersRequest, opts ...grpc.CallOption) (*ListAuthProvidersResponse, error)
}

type authRegistryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthRegistryServiceClient(cc *grpc.ClientConn) AuthRegistryServiceClient {
	return &authRegistryServiceClient{cc}
}

func (c *authRegistryServiceClient) GetAuthProvider(ctx context.Context, in *GetAuthProviderRequest, opts ...grpc.CallOption) (*GetAuthProviderResponse, error) {
	out := new(GetAuthProviderResponse)
	err := c.cc.Invoke(ctx, "/cs3.authregistryv0alpha.AuthRegistryService/GetAuthProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authRegistryServiceClient) ListAuthProviders(ctx context.Context, in *ListAuthProvidersRequest, opts ...grpc.CallOption) (*ListAuthProvidersResponse, error) {
	out := new(ListAuthProvidersResponse)
	err := c.cc.Invoke(ctx, "/cs3.authregistryv0alpha.AuthRegistryService/ListAuthProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthRegistryServiceServer is the server API for AuthRegistryService service.
type AuthRegistryServiceServer interface {
	// Returns the auth provider that is reponsible for the given
	// resource reference.
	// MUST return CODE_NOT_FOUND if the reference does not exist.
	GetAuthProvider(context.Context, *GetAuthProviderRequest) (*GetAuthProviderResponse, error)
	// Returns a list of the available auth providers known by this registry.
	ListAuthProviders(context.Context, *ListAuthProvidersRequest) (*ListAuthProvidersResponse, error)
}

// UnimplementedAuthRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthRegistryServiceServer struct {
}

func (*UnimplementedAuthRegistryServiceServer) GetAuthProvider(ctx context.Context, req *GetAuthProviderRequest) (*GetAuthProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthProvider not implemented")
}
func (*UnimplementedAuthRegistryServiceServer) ListAuthProviders(ctx context.Context, req *ListAuthProvidersRequest) (*ListAuthProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthProviders not implemented")
}

func RegisterAuthRegistryServiceServer(s *grpc.Server, srv AuthRegistryServiceServer) {
	s.RegisterService(&_AuthRegistryService_serviceDesc, srv)
}

func _AuthRegistryService_GetAuthProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRegistryServiceServer).GetAuthProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authregistryv0alpha.AuthRegistryService/GetAuthProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRegistryServiceServer).GetAuthProvider(ctx, req.(*GetAuthProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthRegistryService_ListAuthProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRegistryServiceServer).ListAuthProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authregistryv0alpha.AuthRegistryService/ListAuthProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRegistryServiceServer).ListAuthProviders(ctx, req.(*ListAuthProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthRegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.authregistryv0alpha.AuthRegistryService",
	HandlerType: (*AuthRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthProvider",
			Handler:    _AuthRegistryService_GetAuthProvider_Handler,
		},
		{
			MethodName: "ListAuthProviders",
			Handler:    _AuthRegistryService_ListAuthProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/authregistry/v0alpha/authregistry.proto",
}
