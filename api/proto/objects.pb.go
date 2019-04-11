// Code generated by protoc-gen-go. DO NOT EDIT.
// source: objects.proto

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

type ObjectReference struct {
	Client               *ClientReference          `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Hunt                 *HuntReference            `protobuf:"bytes,3,opt,name=hunt,proto3" json:"hunt,omitempty"`
	Flow                 *FlowReference            `protobuf:"bytes,4,opt,name=flow,proto3" json:"flow,omitempty"`
	VfsFile              *VfsFileReference         `protobuf:"bytes,6,opt,name=vfs_file,json=vfsFile,proto3" json:"vfs_file,omitempty"`
	ApprovalRequest      *ApprovalRequestReference `protobuf:"bytes,7,opt,name=approval_request,json=approvalRequest,proto3" json:"approval_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ObjectReference) Reset()         { *m = ObjectReference{} }
func (m *ObjectReference) String() string { return proto.CompactTextString(m) }
func (*ObjectReference) ProtoMessage()    {}
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da965bc36916fc1, []int{0}
}

func (m *ObjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectReference.Unmarshal(m, b)
}
func (m *ObjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectReference.Marshal(b, m, deterministic)
}
func (m *ObjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectReference.Merge(m, src)
}
func (m *ObjectReference) XXX_Size() int {
	return xxx_messageInfo_ObjectReference.Size(m)
}
func (m *ObjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectReference proto.InternalMessageInfo

func (m *ObjectReference) GetClient() *ClientReference {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *ObjectReference) GetHunt() *HuntReference {
	if m != nil {
		return m.Hunt
	}
	return nil
}

func (m *ObjectReference) GetFlow() *FlowReference {
	if m != nil {
		return m.Flow
	}
	return nil
}

func (m *ObjectReference) GetVfsFile() *VfsFileReference {
	if m != nil {
		return m.VfsFile
	}
	return nil
}

func (m *ObjectReference) GetApprovalRequest() *ApprovalRequestReference {
	if m != nil {
		return m.ApprovalRequest
	}
	return nil
}

type ClientReference struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientReference) Reset()         { *m = ClientReference{} }
func (m *ClientReference) String() string { return proto.CompactTextString(m) }
func (*ClientReference) ProtoMessage()    {}
func (*ClientReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da965bc36916fc1, []int{1}
}

func (m *ClientReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientReference.Unmarshal(m, b)
}
func (m *ClientReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientReference.Marshal(b, m, deterministic)
}
func (m *ClientReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientReference.Merge(m, src)
}
func (m *ClientReference) XXX_Size() int {
	return xxx_messageInfo_ClientReference.Size(m)
}
func (m *ClientReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientReference.DiscardUnknown(m)
}

var xxx_messageInfo_ClientReference proto.InternalMessageInfo

func (m *ClientReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type HuntReference struct {
	HuntId               string   `protobuf:"bytes,1,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HuntReference) Reset()         { *m = HuntReference{} }
func (m *HuntReference) String() string { return proto.CompactTextString(m) }
func (*HuntReference) ProtoMessage()    {}
func (*HuntReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da965bc36916fc1, []int{2}
}

func (m *HuntReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuntReference.Unmarshal(m, b)
}
func (m *HuntReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuntReference.Marshal(b, m, deterministic)
}
func (m *HuntReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuntReference.Merge(m, src)
}
func (m *HuntReference) XXX_Size() int {
	return xxx_messageInfo_HuntReference.Size(m)
}
func (m *HuntReference) XXX_DiscardUnknown() {
	xxx_messageInfo_HuntReference.DiscardUnknown(m)
}

var xxx_messageInfo_HuntReference proto.InternalMessageInfo

func (m *HuntReference) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

type FlowReference struct {
	FlowId               string   `protobuf:"bytes,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowReference) Reset()         { *m = FlowReference{} }
func (m *FlowReference) String() string { return proto.CompactTextString(m) }
func (*FlowReference) ProtoMessage()    {}
func (*FlowReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da965bc36916fc1, []int{3}
}

func (m *FlowReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowReference.Unmarshal(m, b)
}
func (m *FlowReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowReference.Marshal(b, m, deterministic)
}
func (m *FlowReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowReference.Merge(m, src)
}
func (m *FlowReference) XXX_Size() int {
	return xxx_messageInfo_FlowReference.Size(m)
}
func (m *FlowReference) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowReference.DiscardUnknown(m)
}

var xxx_messageInfo_FlowReference proto.InternalMessageInfo

func (m *FlowReference) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *FlowReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type VfsFileReference struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	VfsPathComponents    []string `protobuf:"bytes,2,rep,name=vfs_path_components,json=vfsPathComponents,proto3" json:"vfs_path_components,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VfsFileReference) Reset()         { *m = VfsFileReference{} }
func (m *VfsFileReference) String() string { return proto.CompactTextString(m) }
func (*VfsFileReference) ProtoMessage()    {}
func (*VfsFileReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da965bc36916fc1, []int{4}
}

func (m *VfsFileReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VfsFileReference.Unmarshal(m, b)
}
func (m *VfsFileReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VfsFileReference.Marshal(b, m, deterministic)
}
func (m *VfsFileReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VfsFileReference.Merge(m, src)
}
func (m *VfsFileReference) XXX_Size() int {
	return xxx_messageInfo_VfsFileReference.Size(m)
}
func (m *VfsFileReference) XXX_DiscardUnknown() {
	xxx_messageInfo_VfsFileReference.DiscardUnknown(m)
}

var xxx_messageInfo_VfsFileReference proto.InternalMessageInfo

func (m *VfsFileReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VfsFileReference) GetVfsPathComponents() []string {
	if m != nil {
		return m.VfsPathComponents
	}
	return nil
}

type ApprovalRequestReference struct {
	//    ApprovalRequest.ApprovalType approval_type = 1;
	ApprovalId           string   `protobuf:"bytes,2,opt,name=approval_id,json=approvalId,proto3" json:"approval_id,omitempty"`
	SubjectId            string   `protobuf:"bytes,3,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApprovalRequestReference) Reset()         { *m = ApprovalRequestReference{} }
func (m *ApprovalRequestReference) String() string { return proto.CompactTextString(m) }
func (*ApprovalRequestReference) ProtoMessage()    {}
func (*ApprovalRequestReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da965bc36916fc1, []int{5}
}

func (m *ApprovalRequestReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApprovalRequestReference.Unmarshal(m, b)
}
func (m *ApprovalRequestReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApprovalRequestReference.Marshal(b, m, deterministic)
}
func (m *ApprovalRequestReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovalRequestReference.Merge(m, src)
}
func (m *ApprovalRequestReference) XXX_Size() int {
	return xxx_messageInfo_ApprovalRequestReference.Size(m)
}
func (m *ApprovalRequestReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovalRequestReference.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovalRequestReference proto.InternalMessageInfo

func (m *ApprovalRequestReference) GetApprovalId() string {
	if m != nil {
		return m.ApprovalId
	}
	return ""
}

func (m *ApprovalRequestReference) GetSubjectId() string {
	if m != nil {
		return m.SubjectId
	}
	return ""
}

func init() {
	proto.RegisterType((*ObjectReference)(nil), "proto.ObjectReference")
	proto.RegisterType((*ClientReference)(nil), "proto.ClientReference")
	proto.RegisterType((*HuntReference)(nil), "proto.HuntReference")
	proto.RegisterType((*FlowReference)(nil), "proto.FlowReference")
	proto.RegisterType((*VfsFileReference)(nil), "proto.VfsFileReference")
	proto.RegisterType((*ApprovalRequestReference)(nil), "proto.ApprovalRequestReference")
}

func init() { proto.RegisterFile("objects.proto", fileDescriptor_7da965bc36916fc1) }

var fileDescriptor_7da965bc36916fc1 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x8a, 0xd4, 0x30,
	0x18, 0xc5, 0xe9, 0xcc, 0xda, 0xd9, 0x7e, 0xcb, 0x30, 0x6b, 0x14, 0xb7, 0x28, 0xb2, 0xa1, 0x17,
	0x52, 0x6f, 0x32, 0xb0, 0x82, 0x17, 0xe2, 0x8d, 0x2e, 0xac, 0xd6, 0x0b, 0x95, 0xa2, 0x82, 0xde,
	0x94, 0x6c, 0x93, 0x6e, 0x2b, 0x9d, 0xa4, 0x36, 0x99, 0xd6, 0x17, 0xf1, 0x49, 0x7c, 0x05, 0x9f,
	0x44, 0x5f, 0xc3, 0x0b, 0x49, 0xd2, 0xf9, 0xd3, 0x45, 0xbd, 0x0a, 0x3d, 0xe7, 0xf7, 0x1d, 0x7a,
	0xf2, 0x05, 0xe6, 0xf2, 0xf2, 0x33, 0xcf, 0xb5, 0x22, 0x4d, 0x2b, 0xb5, 0x44, 0x37, 0xec, 0x71,
	0xf7, 0x49, 0xdf, 0xf7, 0xa4, 0xe3, 0xb5, 0xcc, 0x2b, 0xc6, 0xbf, 0x92, 0x5c, 0xae, 0x96, 0x57,
	0xb2, 0xa6, 0xe2, 0x6a, 0xe9, 0xc4, 0x96, 0x36, 0x5a, 0xb6, 0x4b, 0x0b, 0x2f, 0x15, 0x5f, 0x51,
	0xa1, 0xab, 0xdc, 0x45, 0x44, 0xdf, 0x26, 0xb0, 0x78, 0x63, 0x43, 0x53, 0x5e, 0xf0, 0x96, 0x8b,
	0x9c, 0x23, 0x02, 0x7e, 0x5e, 0x57, 0x5c, 0xe8, 0x70, 0x82, 0xbd, 0xf8, 0xe8, 0xec, 0x8e, 0x63,
	0xc9, 0xb9, 0x15, 0xb7, 0x5c, 0x3a, 0x50, 0x28, 0x86, 0x83, 0x72, 0x2d, 0x74, 0x38, 0xb5, 0xf4,
	0xed, 0x81, 0x7e, 0xb9, 0xde, 0x67, 0x2d, 0x61, 0xc8, 0xa2, 0x96, 0x7d, 0x78, 0x30, 0x22, 0x2f,
	0x6a, 0xd9, 0xef, 0x91, 0x86, 0x40, 0x67, 0x70, 0xd8, 0x15, 0x2a, 0x2b, 0xaa, 0x9a, 0x87, 0xbe,
	0xa5, 0x4f, 0x06, 0xfa, 0x43, 0xa1, 0x2e, 0xaa, 0x9a, 0xef, 0x06, 0x66, 0x9d, 0x53, 0xd0, 0x2b,
	0x38, 0xa6, 0x4d, 0xd3, 0xca, 0x8e, 0xd6, 0x59, 0xcb, 0xbf, 0xac, 0xb9, 0xd2, 0xe1, 0xcc, 0xce,
	0x9e, 0x0e, 0xb3, 0xcf, 0x06, 0x3b, 0x75, 0xee, 0x2e, 0x63, 0x41, 0xc7, 0x4e, 0x44, 0x60, 0x71,
	0xad, 0x2e, 0xba, 0x07, 0x81, 0x2b, 0x9c, 0x55, 0x2c, 0xf4, 0xb0, 0x17, 0x07, 0xe9, 0xa1, 0x13,
	0x12, 0x16, 0xc5, 0x30, 0x1f, 0x15, 0x46, 0x27, 0x30, 0x33, 0x95, 0x77, 0xac, 0x6f, 0x3e, 0x13,
	0x16, 0x7d, 0xf7, 0x60, 0x3e, 0x6a, 0x8c, 0x5e, 0xc0, 0xcc, 0x74, 0xde, 0xa2, 0xcf, 0xc9, 0xcf,
	0xdf, 0xbf, 0x7e, 0x78, 0x31, 0x7a, 0xf0, 0xae, 0xe4, 0x58, 0x71, 0xa5, 0x2a, 0x29, 0x70, 0xc5,
	0xb0, 0x2c, 0xb0, 0x2e, 0x39, 0x6e, 0x37, 0x93, 0x0c, 0x9b, 0x39, 0x92, 0xfa, 0xe6, 0x48, 0x18,
	0xfa, 0xb8, 0xff, 0x87, 0x13, 0x1b, 0xf5, 0xd4, 0x46, 0x3d, 0x86, 0xc0, 0xb5, 0x79, 0x9f, 0xbe,
	0x46, 0x0f, 0x4d, 0xaa, 0xe3, 0x4c, 0xe8, 0x5f, 0x12, 0x71, 0x4b, 0x05, 0x96, 0x82, 0xec, 0xf5,
	0xcb, 0xe0, 0xf8, 0xfa, 0xc5, 0xff, 0xf7, 0x42, 0x10, 0x81, 0x5b, 0x66, 0x81, 0x0d, 0xd5, 0x65,
	0x96, 0xcb, 0x55, 0x23, 0x05, 0x17, 0x5a, 0x85, 0x13, 0x3c, 0x8d, 0x83, 0xf4, 0x66, 0x57, 0xa8,
	0xb7, 0x54, 0x97, 0xe7, 0x5b, 0x23, 0xfa, 0x04, 0xe1, 0xbf, 0xb6, 0x83, 0x4e, 0xe1, 0x68, 0xbb,
	0xd8, 0x4d, 0xb3, 0x14, 0x36, 0x52, 0xc2, 0xd0, 0x7d, 0x00, 0xb5, 0xb6, 0xaf, 0xd8, 0xf8, 0x53,
	0xeb, 0x07, 0x83, 0x92, 0xb0, 0x4b, 0xdf, 0x6e, 0xff, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0xa5, 0x7f, 0x01, 0x3f, 0x03, 0x00, 0x00,
}