// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/v1beta1/o_auth_client.proto

package authv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// An OAuth client.
type OAuthClient struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RedirectUris         []string `protobuf:"bytes,2,rep,name=redirect_uris,json=redirectUris,proto3" json:"redirect_uris,omitempty"`
	GrantTypes           []string `protobuf:"bytes,3,rep,name=grant_types,json=grantTypes,proto3" json:"grant_types,omitempty"`
	ResponseTypes        []string `protobuf:"bytes,4,rep,name=response_types,json=responseTypes,proto3" json:"response_types,omitempty"`
	Scopes               []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	Audiences            []string `protobuf:"bytes,6,rep,name=audiences,proto3" json:"audiences,omitempty"`
	Public               bool     `protobuf:"varint,7,opt,name=public,proto3" json:"public,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthClient) Reset()         { *m = OAuthClient{} }
func (m *OAuthClient) String() string { return proto.CompactTextString(m) }
func (*OAuthClient) ProtoMessage()    {}
func (*OAuthClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f0bd39ac3c3851, []int{0}
}

func (m *OAuthClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthClient.Unmarshal(m, b)
}
func (m *OAuthClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthClient.Marshal(b, m, deterministic)
}
func (m *OAuthClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthClient.Merge(m, src)
}
func (m *OAuthClient) XXX_Size() int {
	return xxx_messageInfo_OAuthClient.Size(m)
}
func (m *OAuthClient) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthClient.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthClient proto.InternalMessageInfo

func (m *OAuthClient) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OAuthClient) GetRedirectUris() []string {
	if m != nil {
		return m.RedirectUris
	}
	return nil
}

func (m *OAuthClient) GetGrantTypes() []string {
	if m != nil {
		return m.GrantTypes
	}
	return nil
}

func (m *OAuthClient) GetResponseTypes() []string {
	if m != nil {
		return m.ResponseTypes
	}
	return nil
}

func (m *OAuthClient) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *OAuthClient) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *OAuthClient) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func init() {
	proto.RegisterType((*OAuthClient)(nil), "auth.v1beta1.OAuthClient")
}

func init() { proto.RegisterFile("auth/v1beta1/o_auth_client.proto", fileDescriptor_42f0bd39ac3c3851) }

var fileDescriptor_42f0bd39ac3c3851 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0x87, 0x49, 0xa2, 0xd1, 0xbc, 0xa4, 0x45, 0x66, 0xa1, 0x59, 0x08, 0x06, 0x45, 0xc8, 0x2a,
	0x25, 0x78, 0x82, 0xc6, 0x03, 0x58, 0x82, 0x4a, 0x11, 0x31, 0x4c, 0x92, 0x87, 0x1d, 0x28, 0x33,
	0x61, 0xfe, 0x08, 0x5e, 0xc7, 0xa5, 0x47, 0xf1, 0x04, 0x1e, 0x47, 0x66, 0x3a, 0xc5, 0x2e, 0xbf,
	0xef, 0xf7, 0xcd, 0x62, 0x1e, 0x14, 0xd4, 0xe8, 0xcd, 0xe2, 0xa3, 0xee, 0x51, 0xd3, 0x7a, 0x21,
	0x3a, 0x8b, 0xdd, 0xb0, 0x65, 0xc8, 0x75, 0x35, 0x49, 0xa1, 0x05, 0xc9, 0xac, 0xaa, 0x7c, 0x71,
	0xfd, 0x1b, 0x40, 0xfa, 0xb0, 0x34, 0x7a, 0x73, 0xef, 0x1a, 0x32, 0x87, 0x90, 0x8d, 0x79, 0x50,
	0x04, 0x65, 0xd2, 0x86, 0x6c, 0x24, 0x37, 0x30, 0x93, 0x38, 0x32, 0x89, 0x83, 0xee, 0x8c, 0x64,
	0x2a, 0x0f, 0x8b, 0xa8, 0x4c, 0xda, 0x6c, 0x2f, 0x9f, 0x24, 0x53, 0xe4, 0x0a, 0xd2, 0x77, 0x49,
	0xb9, 0xee, 0xf4, 0xe7, 0x84, 0x2a, 0x8f, 0x5c, 0x02, 0x4e, 0x3d, 0x5a, 0x43, 0x6e, 0x61, 0x2e,
	0x51, 0x4d, 0x82, 0x2b, 0xf4, 0xcd, 0x91, 0x6b, 0x66, 0x7b, 0xbb, 0xcb, 0xce, 0x21, 0x56, 0x83,
	0xb0, 0xf3, 0xb1, 0x9b, 0x3d, 0x91, 0x4b, 0x48, 0xa8, 0x19, 0x19, 0xf2, 0x01, 0x55, 0x1e, 0xbb,
	0xe9, 0x5f, 0xd8, 0x57, 0x93, 0xe9, 0xb7, 0x6c, 0xc8, 0x4f, 0x8a, 0xa0, 0x3c, 0x6d, 0x3d, 0x35,
	0x6f, 0x70, 0xc1, 0x44, 0x85, 0x5c, 0x21, 0x1d, 0x69, 0x75, 0xf8, 0xeb, 0xe6, 0xec, 0xe0, 0xcb,
	0x2b, 0x7b, 0x95, 0x55, 0xf0, 0x92, 0xda, 0xc2, 0x07, 0x5f, 0x61, 0xb4, 0x5c, 0xaf, 0xbf, 0xc3,
	0xcc, 0x66, 0xd5, 0x73, 0xdd, 0x58, 0xf9, 0xb3, 0xc3, 0x57, 0x8f, 0x7d, 0xec, 0xee, 0x79, 0xf7,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x85, 0xad, 0x7c, 0x73, 0x01, 0x00, 0x00,
}
