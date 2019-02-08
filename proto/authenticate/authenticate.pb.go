// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authenticate.proto

package authenticate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type AuthenticateRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authenticate_b52fdd447b0a5778, []int{0}
}
func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(dst, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AuthenticateReply struct {
	AccessToken          string               `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string               `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	IdToken              string               `protobuf:"bytes,3,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	User                 string               `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Email                string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Expiry               *timestamp.Timestamp `protobuf:"bytes,6,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthenticateReply) Reset()         { *m = AuthenticateReply{} }
func (m *AuthenticateReply) String() string { return proto.CompactTextString(m) }
func (*AuthenticateReply) ProtoMessage()    {}
func (*AuthenticateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_authenticate_b52fdd447b0a5778, []int{1}
}
func (m *AuthenticateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateReply.Unmarshal(m, b)
}
func (m *AuthenticateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateReply.Marshal(b, m, deterministic)
}
func (dst *AuthenticateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateReply.Merge(dst, src)
}
func (m *AuthenticateReply) XXX_Size() int {
	return xxx_messageInfo_AuthenticateReply.Size(m)
}
func (m *AuthenticateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateReply proto.InternalMessageInfo

func (m *AuthenticateReply) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthenticateReply) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *AuthenticateReply) GetIdToken() string {
	if m != nil {
		return m.IdToken
	}
	return ""
}

func (m *AuthenticateReply) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *AuthenticateReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticateReply) GetExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type ValidateRequest struct {
	IdToken              string   `protobuf:"bytes,1,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authenticate_b52fdd447b0a5778, []int{2}
}
func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (dst *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(dst, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetIdToken() string {
	if m != nil {
		return m.IdToken
	}
	return ""
}

type ValidateReply struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateReply) Reset()         { *m = ValidateReply{} }
func (m *ValidateReply) String() string { return proto.CompactTextString(m) }
func (*ValidateReply) ProtoMessage()    {}
func (*ValidateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_authenticate_b52fdd447b0a5778, []int{3}
}
func (m *ValidateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateReply.Unmarshal(m, b)
}
func (m *ValidateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateReply.Marshal(b, m, deterministic)
}
func (dst *ValidateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateReply.Merge(dst, src)
}
func (m *ValidateReply) XXX_Size() int {
	return xxx_messageInfo_ValidateReply.Size(m)
}
func (m *ValidateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateReply.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateReply proto.InternalMessageInfo

func (m *ValidateReply) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type RefreshRequest struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authenticate_b52fdd447b0a5778, []int{4}
}
func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (dst *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(dst, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshReply struct {
	AccessToken          string               `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Expiry               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RefreshReply) Reset()         { *m = RefreshReply{} }
func (m *RefreshReply) String() string { return proto.CompactTextString(m) }
func (*RefreshReply) ProtoMessage()    {}
func (*RefreshReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_authenticate_b52fdd447b0a5778, []int{5}
}
func (m *RefreshReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshReply.Unmarshal(m, b)
}
func (m *RefreshReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshReply.Marshal(b, m, deterministic)
}
func (dst *RefreshReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshReply.Merge(dst, src)
}
func (m *RefreshReply) XXX_Size() int {
	return xxx_messageInfo_RefreshReply.Size(m)
}
func (m *RefreshReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshReply.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshReply proto.InternalMessageInfo

func (m *RefreshReply) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RefreshReply) GetExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticateRequest)(nil), "authenticate.AuthenticateRequest")
	proto.RegisterType((*AuthenticateReply)(nil), "authenticate.AuthenticateReply")
	proto.RegisterType((*ValidateRequest)(nil), "authenticate.ValidateRequest")
	proto.RegisterType((*ValidateReply)(nil), "authenticate.ValidateReply")
	proto.RegisterType((*RefreshRequest)(nil), "authenticate.RefreshRequest")
	proto.RegisterType((*RefreshReply)(nil), "authenticate.RefreshReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticatorClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateReply, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateReply, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshReply, error)
}

type authenticatorClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticatorClient(cc *grpc.ClientConn) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateReply, error) {
	out := new(AuthenticateReply)
	err := c.cc.Invoke(ctx, "/authenticate.Authenticator/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateReply, error) {
	out := new(ValidateReply)
	err := c.cc.Invoke(ctx, "/authenticate.Authenticator/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshReply, error) {
	out := new(RefreshReply)
	err := c.cc.Invoke(ctx, "/authenticate.Authenticator/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
type AuthenticatorServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateReply, error)
	Validate(context.Context, *ValidateRequest) (*ValidateReply, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshReply, error)
}

func RegisterAuthenticatorServer(s *grpc.Server, srv AuthenticatorServer) {
	s.RegisterService(&_Authenticator_serviceDesc, srv)
}

func _Authenticator_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticate.Authenticator/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticate.Authenticator/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticate.Authenticator/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authenticate.Authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Authenticator_Authenticate_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Authenticator_Validate_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Authenticator_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authenticate.proto",
}

func init() { proto.RegisterFile("authenticate.proto", fileDescriptor_authenticate_b52fdd447b0a5778) }

var fileDescriptor_authenticate_b52fdd447b0a5778 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x5f, 0x79, 0xfc, 0x7b, 0x97, 0xf2, 0x5e, 0xde, 0xd5, 0x45, 0xad, 0x1a, 0xa0, 0x6e,
	0xd0, 0x98, 0x92, 0x60, 0xfc, 0x00, 0x2e, 0x4c, 0x5c, 0x37, 0xc4, 0x2d, 0x29, 0xed, 0x05, 0x26,
	0x16, 0xa6, 0x76, 0xa6, 0x46, 0xbe, 0xa7, 0x9f, 0xc5, 0xb5, 0xe9, 0x4c, 0x2b, 0xad, 0x88, 0x61,
	0xd7, 0x9e, 0xf3, 0x9b, 0x9b, 0x7b, 0xce, 0x0c, 0xa0, 0x9f, 0xca, 0x25, 0xad, 0x25, 0x0b, 0x7c,
	0x49, 0x6e, 0x9c, 0x70, 0xc9, 0xd1, 0x2c, 0x6b, 0x76, 0x6f, 0xc1, 0xf9, 0x22, 0xa2, 0x91, 0xf2,
	0x66, 0xe9, 0x7c, 0x24, 0xd9, 0x8a, 0x84, 0xf4, 0x57, 0xb1, 0xc6, 0x9d, 0x4b, 0x38, 0xba, 0x2b,
	0x1d, 0xf0, 0xe8, 0x39, 0x25, 0x21, 0x11, 0xa1, 0x1e, 0xf0, 0x90, 0x2c, 0xa3, 0x6f, 0x0c, 0xff,
	0x78, 0xea, 0xdb, 0x79, 0x33, 0xe0, 0x7f, 0x95, 0x8d, 0xa3, 0x0d, 0x0e, 0xc0, 0xf4, 0x83, 0x80,
	0x84, 0x98, 0x4a, 0xfe, 0x44, 0xeb, 0xfc, 0x44, 0x47, 0x6b, 0x93, 0x4c, 0xc2, 0x0b, 0xe8, 0x26,
	0x34, 0x4f, 0x48, 0x2c, 0x73, 0xa6, 0xa6, 0x18, 0x33, 0x17, 0x35, 0x74, 0x02, 0x6d, 0x16, 0xe6,
	0xfe, 0x6f, 0xe5, 0xb7, 0x58, 0xa8, 0x2d, 0x84, 0x7a, 0x2a, 0x28, 0xb1, 0xea, 0x7a, 0x99, 0xec,
	0x1b, 0x8f, 0xa1, 0x41, 0x2b, 0x9f, 0x45, 0x56, 0x43, 0x89, 0xfa, 0x07, 0xc7, 0xd0, 0xa4, 0xd7,
	0x98, 0x25, 0x1b, 0xab, 0xd9, 0x37, 0x86, 0x9d, 0xb1, 0xed, 0xea, 0xfc, 0x6e, 0x91, 0xdf, 0x9d,
	0x14, 0xf9, 0xbd, 0x9c, 0x74, 0xae, 0xe1, 0xdf, 0xa3, 0x1f, 0xb1, 0xb0, 0x94, 0xbe, 0xbc, 0x8b,
	0x51, 0xd9, 0xc5, 0xb9, 0x82, 0xee, 0x96, 0xce, 0xf2, 0x67, 0xac, 0x98, 0xbe, 0x64, 0x9a, 0x62,
	0xdb, 0x5e, 0x8b, 0x09, 0x85, 0x38, 0xb7, 0xf0, 0xd7, 0xd3, 0x11, 0x8b, 0xc1, 0x3b, 0x4d, 0x18,
	0xbb, 0x4d, 0x38, 0x04, 0xe6, 0xe7, 0xb1, 0x03, 0x1b, 0xde, 0xe6, 0xae, 0x1d, 0x9a, 0x7b, 0xfc,
	0x6e, 0x40, 0xb7, 0x74, 0x9d, 0x3c, 0xc1, 0x09, 0x98, 0xe5, 0xfb, 0xc5, 0x81, 0x5b, 0x79, 0x5f,
	0xdf, 0xbc, 0x13, 0xbb, 0xf7, 0x13, 0x12, 0x47, 0x1b, 0xe7, 0x17, 0x3e, 0x40, 0xbb, 0x68, 0x0c,
	0xcf, 0xab, 0xf8, 0x97, 0xde, 0xed, 0xd3, 0x7d, 0xb6, 0x9e, 0x74, 0x0f, 0xad, 0xbc, 0x18, 0x3c,
	0xab, 0x92, 0xd5, 0x9a, 0x6d, 0x7b, 0x8f, 0xab, 0xc6, 0xcc, 0x9a, 0xaa, 0x94, 0x9b, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x76, 0x32, 0xe7, 0x1e, 0x3e, 0x03, 0x00, 0x00,
}
