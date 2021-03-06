// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type GetArtifactRequest struct {
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56fa9fd5f0eb7ff, []int{0}
}

func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (m *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(m, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

type GetArtifactResponse struct {
	Artifact             string   `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56fa9fd5f0eb7ff, []int{1}
}

func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (m *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(m, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type SetArtifactRequest struct {
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Artifact             string   `protobuf:"bytes,2,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetArtifactRequest) Reset()         { *m = SetArtifactRequest{} }
func (m *SetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*SetArtifactRequest) ProtoMessage()    {}
func (*SetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56fa9fd5f0eb7ff, []int{2}
}

func (m *SetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetArtifactRequest.Unmarshal(m, b)
}
func (m *SetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetArtifactRequest.Marshal(b, m, deterministic)
}
func (m *SetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetArtifactRequest.Merge(m, src)
}
func (m *SetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_SetArtifactRequest.Size(m)
}
func (m *SetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetArtifactRequest proto.InternalMessageInfo

func (m *SetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *SetArtifactRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type APIResponse struct {
	Error                bool     `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIResponse) Reset()         { *m = APIResponse{} }
func (m *APIResponse) String() string { return proto.CompactTextString(m) }
func (*APIResponse) ProtoMessage()    {}
func (*APIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56fa9fd5f0eb7ff, []int{3}
}

func (m *APIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIResponse.Unmarshal(m, b)
}
func (m *APIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIResponse.Marshal(b, m, deterministic)
}
func (m *APIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIResponse.Merge(m, src)
}
func (m *APIResponse) XXX_Size() int {
	return xxx_messageInfo_APIResponse.Size(m)
}
func (m *APIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_APIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_APIResponse proto.InternalMessageInfo

func (m *APIResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *APIResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*GetArtifactRequest)(nil), "proto.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "proto.GetArtifactResponse")
	proto.RegisterType((*SetArtifactRequest)(nil), "proto.SetArtifactRequest")
	proto.RegisterType((*APIResponse)(nil), "proto.APIResponse")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_a56fa9fd5f0eb7ff) }

var fileDescriptor_a56fa9fd5f0eb7ff = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0x30, 0x9d, 0x51, 0x11, 0xe2, 0x45, 0x3c, 0x3d, 0x2a, 0xc8, 0x0e, 0xd2, 0x1e,
	0xbc, 0xc8, 0x4e, 0x6e, 0x08, 0xe2, 0x41, 0x18, 0x75, 0xb7, 0x1d, 0x46, 0x6c, 0x5f, 0xdb, 0x40,
	0x97, 0xd7, 0x25, 0x6f, 0xdd, 0xbe, 0x99, 0x5f, 0xc0, 0x4f, 0xa2, 0x5f, 0xc3, 0x83, 0x2c, 0x9d,
	0x63, 0xbb, 0x7a, 0xf0, 0x94, 0xf0, 0x4f, 0xfe, 0xff, 0xff, 0x2f, 0xe1, 0x89, 0x73, 0x65, 0x59,
	0xe7, 0x2a, 0x65, 0x17, 0xd5, 0x96, 0x98, 0x64, 0xc7, 0x2f, 0x57, 0xfd, 0xe5, 0x72, 0x19, 0x35,
	0x58, 0x51, 0xaa, 0x33, 0x5c, 0x45, 0x29, 0xcd, 0xe2, 0x82, 0x2a, 0x65, 0x8a, 0xb8, 0x15, 0xad,
	0xaa, 0x99, 0x6c, 0xec, 0x2f, 0xc7, 0x0e, 0x67, 0xca, 0xb0, 0x4e, 0xdb, 0x88, 0x70, 0x2e, 0xe4,
	0x13, 0xf2, 0x60, 0x13, 0x9c, 0xe0, 0x7c, 0x81, 0x8e, 0xe5, 0x44, 0x74, 0x9b, 0xdc, 0x4d, 0x6b,
	0xc5, 0xe5, 0x65, 0x00, 0x41, 0xef, 0x78, 0xf8, 0xf0, 0xf9, 0xfd, 0xf5, 0x11, 0xf4, 0xe5, 0xfd,
	0xb8, 0x44, 0x68, 0x72, 0x07, 0xeb, 0x33, 0xb0, 0x58, 0x29, 0xd6, 0x0d, 0x02, 0x13, 0x70, 0x89,
	0xb0, 0x05, 0x84, 0x0c, 0x73, 0x6d, 0x34, 0x6b, 0x32, 0xe0, 0x98, 0x2c, 0x46, 0xc9, 0x51, 0x93,
	0xbb, 0x91, 0xe2, 0x32, 0x9c, 0x88, 0x8b, 0xbd, 0x4a, 0x57, 0x93, 0x71, 0x28, 0x1f, 0x45, 0xf7,
	0xd7, 0xbe, 0xe9, 0xec, 0xf9, 0xce, 0x50, 0xc2, 0x78, 0x27, 0x1a, 0x32, 0xc5, 0xea, 0x16, 0xc8,
	0x82, 0x5a, 0x97, 0xa8, 0x45, 0xc5, 0x51, 0xb2, 0x75, 0x86, 0xef, 0x81, 0x90, 0xaf, 0xff, 0xfb,
	0xa0, 0x3d, 0xf2, 0x83, 0x3f, 0x93, 0xaf, 0xc4, 0xc9, 0x60, 0xf4, 0xbc, 0xf3, 0x1d, 0x1d, 0xb4,
	0x96, 0xac, 0xc7, 0xed, 0x0e, 0x23, 0x9f, 0xd8, 0x93, 0x37, 0x03, 0x03, 0x5e, 0x07, 0x4a, 0xd3,
	0x85, 0xc5, 0x0c, 0x1c, 0x32, 0x6b, 0x53, 0xec, 0xe1, 0x46, 0x49, 0x6b, 0x96, 0xd7, 0xe2, 0xcc,
	0x6f, 0xa6, 0x33, 0x74, 0x4e, 0x15, 0xd8, 0xf2, 0x25, 0xa7, 0x5e, 0x7c, 0x69, 0xb5, 0xb7, 0x43,
	0x3f, 0x0a, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x67, 0x7b, 0x4d, 0x60, 0x02, 0x00,
	0x00,
}
