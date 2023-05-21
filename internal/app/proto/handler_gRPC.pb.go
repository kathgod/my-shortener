// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: handler_gRPC.proto

package proto

import (
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

type GetFuncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url *GetFuncRequest_Httprequrl `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetFuncRequest) Reset() {
	*x = GetFuncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFuncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFuncRequest) ProtoMessage() {}

func (x *GetFuncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFuncRequest.ProtoReflect.Descriptor instead.
func (*GetFuncRequest) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{0}
}

func (x *GetFuncRequest) GetUrl() *GetFuncRequest_Httprequrl {
	if x != nil {
		return x.Url
	}
	return nil
}

type GetFuncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64                `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Out    *GetFuncResponse_Out `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *GetFuncResponse) Reset() {
	*x = GetFuncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFuncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFuncResponse) ProtoMessage() {}

func (x *GetFuncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFuncResponse.ProtoReflect.Descriptor instead.
func (*GetFuncResponse) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{1}
}

func (x *GetFuncResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetFuncResponse) GetOut() *GetFuncResponse_Out {
	if x != nil {
		return x.Out
	}
	return nil
}

type PostFuncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresss string                   `protobuf:"bytes,1,opt,name=addresss,proto3" json:"addresss,omitempty"`
	Longurl  *PostFuncRequest_Longurl `protobuf:"bytes,2,opt,name=longurl,proto3" json:"longurl,omitempty"`
}

func (x *PostFuncRequest) Reset() {
	*x = PostFuncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFuncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFuncRequest) ProtoMessage() {}

func (x *PostFuncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFuncRequest.ProtoReflect.Descriptor instead.
func (*PostFuncRequest) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{2}
}

func (x *PostFuncRequest) GetAddresss() string {
	if x != nil {
		return x.Addresss
	}
	return ""
}

func (x *PostFuncRequest) GetLongurl() *PostFuncRequest_Longurl {
	if x != nil {
		return x.Longurl
	}
	return nil
}

type PostFuncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Out    []byte `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *PostFuncResponse) Reset() {
	*x = PostFuncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFuncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFuncResponse) ProtoMessage() {}

func (x *PostFuncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFuncResponse.ProtoReflect.Descriptor instead.
func (*PostFuncResponse) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{3}
}

func (x *PostFuncResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PostFuncResponse) GetOut() []byte {
	if x != nil {
		return x.Out
	}
	return nil
}

type GetFuncPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFuncPingRequest) Reset() {
	*x = GetFuncPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFuncPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFuncPingRequest) ProtoMessage() {}

func (x *GetFuncPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFuncPingRequest.ProtoReflect.Descriptor instead.
func (*GetFuncPingRequest) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{4}
}

type GetFuncPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetFuncPingResponse) Reset() {
	*x = GetFuncPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFuncPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFuncPingResponse) ProtoMessage() {}

func (x *GetFuncPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFuncPingResponse.ProtoReflect.Descriptor instead.
func (*GetFuncPingResponse) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{5}
}

func (x *GetFuncPingResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetFuncRequest_Httprequrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetFuncRequest_Httprequrl) Reset() {
	*x = GetFuncRequest_Httprequrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFuncRequest_Httprequrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFuncRequest_Httprequrl) ProtoMessage() {}

func (x *GetFuncRequest_Httprequrl) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFuncRequest_Httprequrl.ProtoReflect.Descriptor instead.
func (*GetFuncRequest_Httprequrl) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetFuncRequest_Httprequrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetFuncResponse_Out struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Out string `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *GetFuncResponse_Out) Reset() {
	*x = GetFuncResponse_Out{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFuncResponse_Out) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFuncResponse_Out) ProtoMessage() {}

func (x *GetFuncResponse_Out) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFuncResponse_Out.ProtoReflect.Descriptor instead.
func (*GetFuncResponse_Out) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetFuncResponse_Out) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

type PostFuncRequest_Longurl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PostFuncRequest_Longurl) Reset() {
	*x = PostFuncRequest_Longurl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_gRPC_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFuncRequest_Longurl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFuncRequest_Longurl) ProtoMessage() {}

func (x *PostFuncRequest_Longurl) ProtoReflect() protoreflect.Message {
	mi := &file_handler_gRPC_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFuncRequest_Longurl.ProtoReflect.Descriptor instead.
func (*PostFuncRequest_Longurl) Descriptor() ([]byte, []int) {
	return file_handler_gRPC_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PostFuncRequest_Longurl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_handler_gRPC_proto protoreflect.FileDescriptor

var file_handler_gRPC_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x72, 0x65, 0x71, 0x75, 0x72, 0x6c, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x1a, 0x1e, 0x0a, 0x0a, 0x48, 0x74, 0x74, 0x70, 0x72, 0x65, 0x71, 0x75, 0x72, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x70, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x03,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4f, 0x75, 0x74, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x1a, 0x17, 0x0a, 0x03, 0x4f, 0x75,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6f, 0x75, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x6e,
	0x67, 0x75, 0x72, 0x6c, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x75, 0x72, 0x6c, 0x1a, 0x1b, 0x0a,
	0x07, 0x4c, 0x6f, 0x6e, 0x67, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x10, 0x50, 0x6f,
	0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd1, 0x01,
	0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x50, 0x43, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74,
	0x46, 0x75, 0x6e, 0x63, 0x52, 0x50, 0x43, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x50, 0x43, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x21, 0x5a, 0x1f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handler_gRPC_proto_rawDescOnce sync.Once
	file_handler_gRPC_proto_rawDescData = file_handler_gRPC_proto_rawDesc
)

func file_handler_gRPC_proto_rawDescGZIP() []byte {
	file_handler_gRPC_proto_rawDescOnce.Do(func() {
		file_handler_gRPC_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_gRPC_proto_rawDescData)
	})
	return file_handler_gRPC_proto_rawDescData
}

var file_handler_gRPC_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_handler_gRPC_proto_goTypes = []interface{}{
	(*GetFuncRequest)(nil),            // 0: proto.GetFuncRequest
	(*GetFuncResponse)(nil),           // 1: proto.GetFuncResponse
	(*PostFuncRequest)(nil),           // 2: proto.PostFuncRequest
	(*PostFuncResponse)(nil),          // 3: proto.PostFuncResponse
	(*GetFuncPingRequest)(nil),        // 4: proto.GetFuncPingRequest
	(*GetFuncPingResponse)(nil),       // 5: proto.GetFuncPingResponse
	(*GetFuncRequest_Httprequrl)(nil), // 6: proto.GetFuncRequest.Httprequrl
	(*GetFuncResponse_Out)(nil),       // 7: proto.GetFuncResponse.Out
	(*PostFuncRequest_Longurl)(nil),   // 8: proto.PostFuncRequest.Longurl
}
var file_handler_gRPC_proto_depIdxs = []int32{
	6, // 0: proto.GetFuncRequest.url:type_name -> proto.GetFuncRequest.Httprequrl
	7, // 1: proto.GetFuncResponse.out:type_name -> proto.GetFuncResponse.Out
	8, // 2: proto.PostFuncRequest.longurl:type_name -> proto.PostFuncRequest.Longurl
	0, // 3: proto.MyService.GetFuncRPC:input_type -> proto.GetFuncRequest
	2, // 4: proto.MyService.PostFuncRPC:input_type -> proto.PostFuncRequest
	4, // 5: proto.MyService.GetFuncPingRPC:input_type -> proto.GetFuncPingRequest
	1, // 6: proto.MyService.GetFuncRPC:output_type -> proto.GetFuncResponse
	3, // 7: proto.MyService.PostFuncRPC:output_type -> proto.PostFuncResponse
	5, // 8: proto.MyService.GetFuncPingRPC:output_type -> proto.GetFuncPingResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_handler_gRPC_proto_init() }
func file_handler_gRPC_proto_init() {
	if File_handler_gRPC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_gRPC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFuncRequest); i {
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
		file_handler_gRPC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFuncResponse); i {
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
		file_handler_gRPC_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFuncRequest); i {
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
		file_handler_gRPC_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFuncResponse); i {
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
		file_handler_gRPC_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFuncPingRequest); i {
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
		file_handler_gRPC_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFuncPingResponse); i {
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
		file_handler_gRPC_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFuncRequest_Httprequrl); i {
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
		file_handler_gRPC_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFuncResponse_Out); i {
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
		file_handler_gRPC_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFuncRequest_Longurl); i {
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
			RawDescriptor: file_handler_gRPC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_gRPC_proto_goTypes,
		DependencyIndexes: file_handler_gRPC_proto_depIdxs,
		MessageInfos:      file_handler_gRPC_proto_msgTypes,
	}.Build()
	File_handler_gRPC_proto = out.File
	file_handler_gRPC_proto_rawDesc = nil
	file_handler_gRPC_proto_goTypes = nil
	file_handler_gRPC_proto_depIdxs = nil
}