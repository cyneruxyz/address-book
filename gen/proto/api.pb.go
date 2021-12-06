// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/api.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddressField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"` // Unique ID number for this person
}

func (x *AddressField) Reset() {
	*x = AddressField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressField) ProtoMessage() {}

func (x *AddressField) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressField.ProtoReflect.Descriptor instead.
func (*AddressField) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{0}
}

func (x *AddressField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddressField) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressField) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type AddressFieldQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param string `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *AddressFieldQuery) Reset() {
	*x = AddressFieldQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressFieldQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressFieldQuery) ProtoMessage() {}

func (x *AddressFieldQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressFieldQuery.ProtoReflect.Descriptor instead.
func (*AddressFieldQuery) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{1}
}

func (x *AddressFieldQuery) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

type AddressFieldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field *AddressField `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *AddressFieldRequest) Reset() {
	*x = AddressFieldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressFieldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressFieldRequest) ProtoMessage() {}

func (x *AddressFieldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressFieldRequest.ProtoReflect.Descriptor instead.
func (*AddressFieldRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{2}
}

func (x *AddressFieldRequest) GetField() *AddressField {
	if x != nil {
		return x.Field
	}
	return nil
}

type AddressFieldUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalField    *AddressField `protobuf:"bytes,1,opt,name=originalField,proto3" json:"originalField,omitempty"`
	ReplacementField *AddressField `protobuf:"bytes,2,opt,name=replacementField,proto3" json:"replacementField,omitempty"`
}

func (x *AddressFieldUpdateRequest) Reset() {
	*x = AddressFieldUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressFieldUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressFieldUpdateRequest) ProtoMessage() {}

func (x *AddressFieldUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressFieldUpdateRequest.ProtoReflect.Descriptor instead.
func (*AddressFieldUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{3}
}

func (x *AddressFieldUpdateRequest) GetOriginalField() *AddressField {
	if x != nil {
		return x.OriginalField
	}
	return nil
}

func (x *AddressFieldUpdateRequest) GetReplacementField() *AddressField {
	if x != nil {
		return x.ReplacementField
	}
	return nil
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{4}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_api_proto protoreflect.FileDescriptor

var file_proto_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x22, 0x4b, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78,
	0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xad, 0x01,
	0x0a, 0x19, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x4a, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x79,
	0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x27, 0x0a,
	0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xac, 0x04, 0x0a,
	0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x1d, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x22, 0x05, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x74, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x79, 0x6e,
	0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b,
	0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75,
	0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1e, 0x2e, 0x63,
	0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f,
	0x6f, 0x6b, 0x30, 0x01, 0x12, 0x76, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2b, 0x2e, 0x63, 0x79, 0x6e,
	0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75,
	0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x1a, 0x0c, 0x2f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x79, 0x6e, 0x65, 0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x79, 0x6e, 0x65,
	0x72, 0x75, 0x78, 0x79, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_proto_rawDescOnce sync.Once
	file_proto_api_proto_rawDescData = file_proto_api_proto_rawDesc
)

func file_proto_api_proto_rawDescGZIP() []byte {
	file_proto_api_proto_rawDescOnce.Do(func() {
		file_proto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_proto_rawDescData)
	})
	return file_proto_api_proto_rawDescData
}

var file_proto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_api_proto_goTypes = []interface{}{
	(*AddressField)(nil),              // 0: cyneruxyz.api.v1.AddressField
	(*AddressFieldQuery)(nil),         // 1: cyneruxyz.api.v1.AddressFieldQuery
	(*AddressFieldRequest)(nil),       // 2: cyneruxyz.api.v1.AddressFieldRequest
	(*AddressFieldUpdateRequest)(nil), // 3: cyneruxyz.api.v1.AddressFieldUpdateRequest
	(*EchoRequest)(nil),               // 4: cyneruxyz.api.v1.EchoRequest
	(*Response)(nil),                  // 5: cyneruxyz.api.v1.Response
}
var file_proto_api_proto_depIdxs = []int32{
	0, // 0: cyneruxyz.api.v1.AddressFieldRequest.field:type_name -> cyneruxyz.api.v1.AddressField
	0, // 1: cyneruxyz.api.v1.AddressFieldUpdateRequest.originalField:type_name -> cyneruxyz.api.v1.AddressField
	0, // 2: cyneruxyz.api.v1.AddressFieldUpdateRequest.replacementField:type_name -> cyneruxyz.api.v1.AddressField
	4, // 3: cyneruxyz.api.v1.APIService.Echo:input_type -> cyneruxyz.api.v1.EchoRequest
	2, // 4: cyneruxyz.api.v1.APIService.CreateAddressField:input_type -> cyneruxyz.api.v1.AddressFieldRequest
	1, // 5: cyneruxyz.api.v1.APIService.ReadAddressField:input_type -> cyneruxyz.api.v1.AddressFieldQuery
	3, // 6: cyneruxyz.api.v1.APIService.UpdateAddressField:input_type -> cyneruxyz.api.v1.AddressFieldUpdateRequest
	2, // 7: cyneruxyz.api.v1.APIService.DeleteAddressField:input_type -> cyneruxyz.api.v1.AddressFieldRequest
	5, // 8: cyneruxyz.api.v1.APIService.Echo:output_type -> cyneruxyz.api.v1.Response
	0, // 9: cyneruxyz.api.v1.APIService.CreateAddressField:output_type -> cyneruxyz.api.v1.AddressField
	0, // 10: cyneruxyz.api.v1.APIService.ReadAddressField:output_type -> cyneruxyz.api.v1.AddressField
	5, // 11: cyneruxyz.api.v1.APIService.UpdateAddressField:output_type -> cyneruxyz.api.v1.Response
	5, // 12: cyneruxyz.api.v1.APIService.DeleteAddressField:output_type -> cyneruxyz.api.v1.Response
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_api_proto_init() }
func file_proto_api_proto_init() {
	if File_proto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressFieldQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressFieldRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressFieldUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_proto_goTypes,
		DependencyIndexes: file_proto_api_proto_depIdxs,
		MessageInfos:      file_proto_api_proto_msgTypes,
	}.Build()
	File_proto_api_proto = out.File
	file_proto_api_proto_rawDesc = nil
	file_proto_api_proto_goTypes = nil
	file_proto_api_proto_depIdxs = nil
}