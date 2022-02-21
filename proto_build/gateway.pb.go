// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: gateway.proto

package proto_build

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *RegisterRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

type UnRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *UnRegisterRequest) Reset() {
	*x = UnRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnRegisterRequest) ProtoMessage() {}

func (x *UnRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnRegisterRequest.ProtoReflect.Descriptor instead.
func (*UnRegisterRequest) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *UnRegisterRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *UnRegisterRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type UnRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnRegisterResponse) Reset() {
	*x = UnRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnRegisterResponse) ProtoMessage() {}

func (x *UnRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnRegisterResponse.ProtoReflect.Descriptor instead.
func (*UnRegisterResponse) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{3}
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *SendMessageRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *SendMessageResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PushToAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PushToAllRequest) Reset() {
	*x = PushToAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushToAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushToAllRequest) ProtoMessage() {}

func (x *PushToAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushToAllRequest.ProtoReflect.Descriptor instead.
func (*PushToAllRequest) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *PushToAllRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *PushToAllRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PushToAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data uint64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PushToAllResponse) Reset() {
	*x = PushToAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushToAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushToAllResponse) ProtoMessage() {}

func (x *PushToAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushToAllResponse.ProtoReflect.Descriptor instead.
func (*PushToAllResponse) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *PushToAllResponse) GetData() uint64 {
	if x != nil {
		return x.Data
	}
	return 0
}

var File_gateway_proto protoreflect.FileDescriptor

var file_gateway_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x40, 0x76, 0x30, 0x2e, 0x36, 0x2e, 0x33, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x0a, 0x11, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59,
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x13, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a,
	0x10, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc8, 0x03, 0x0a, 0x0e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01,
	0x2a, 0x12, 0x6c, 0x0a, 0x0a, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x75, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12,
	0x6d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x28, 0x01, 0x30, 0x01, 0x12, 0x73,
	0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x54, 0x6f, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_proto_rawDescOnce sync.Once
	file_gateway_proto_rawDescData = file_gateway_proto_rawDesc
)

func file_gateway_proto_rawDescGZIP() []byte {
	file_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_rawDescData)
	})
	return file_gateway_proto_rawDescData
}

var file_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gateway_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),     // 0: message.RegisterRequest
	(*RegisterResponse)(nil),    // 1: message.RegisterResponse
	(*UnRegisterRequest)(nil),   // 2: message.UnRegisterRequest
	(*UnRegisterResponse)(nil),  // 3: message.UnRegisterResponse
	(*SendMessageRequest)(nil),  // 4: message.SendMessageRequest
	(*SendMessageResponse)(nil), // 5: message.SendMessageResponse
	(*PushToAllRequest)(nil),    // 6: message.PushToAllRequest
	(*PushToAllResponse)(nil),   // 7: message.PushToAllResponse
}
var file_gateway_proto_depIdxs = []int32{
	0, // 0: message.GatewayService.Register:input_type -> message.RegisterRequest
	2, // 1: message.GatewayService.UnRegister:input_type -> message.UnRegisterRequest
	4, // 2: message.GatewayService.SendMessage:input_type -> message.SendMessageRequest
	6, // 3: message.GatewayService.PushToAllMessage:input_type -> message.PushToAllRequest
	1, // 4: message.GatewayService.Register:output_type -> message.RegisterResponse
	3, // 5: message.GatewayService.UnRegister:output_type -> message.UnRegisterResponse
	5, // 6: message.GatewayService.SendMessage:output_type -> message.SendMessageResponse
	7, // 7: message.GatewayService.PushToAllMessage:output_type -> message.PushToAllResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_proto_init() }
func file_gateway_proto_init() {
	if File_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnRegisterRequest); i {
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
		file_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnRegisterResponse); i {
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
		file_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushToAllRequest); i {
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
		file_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushToAllResponse); i {
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
			RawDescriptor: file_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_proto_msgTypes,
	}.Build()
	File_gateway_proto = out.File
	file_gateway_proto_rawDesc = nil
	file_gateway_proto_goTypes = nil
	file_gateway_proto_depIdxs = nil
}