// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: both_stream.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//请求
type BothStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqIdx  string `protobuf:"bytes,1,opt,name=reqIdx,proto3" json:"reqIdx,omitempty"`
	ReqData string `protobuf:"bytes,2,opt,name=reqData,proto3" json:"reqData,omitempty"`
}

func (x *BothStreamRequest) Reset() {
	*x = BothStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_both_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BothStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BothStreamRequest) ProtoMessage() {}

func (x *BothStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_both_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BothStreamRequest.ProtoReflect.Descriptor instead.
func (*BothStreamRequest) Descriptor() ([]byte, []int) {
	return file_both_stream_proto_rawDescGZIP(), []int{0}
}

func (x *BothStreamRequest) GetReqIdx() string {
	if x != nil {
		return x.ReqIdx
	}
	return ""
}

func (x *BothStreamRequest) GetReqData() string {
	if x != nil {
		return x.ReqData
	}
	return ""
}

//响应
type BothStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RspIdx  string `protobuf:"bytes,1,opt,name=rspIdx,proto3" json:"rspIdx,omitempty"`
	RspData string `protobuf:"bytes,2,opt,name=rspData,proto3" json:"rspData,omitempty"`
}

func (x *BothStreamResponse) Reset() {
	*x = BothStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_both_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BothStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BothStreamResponse) ProtoMessage() {}

func (x *BothStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_both_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BothStreamResponse.ProtoReflect.Descriptor instead.
func (*BothStreamResponse) Descriptor() ([]byte, []int) {
	return file_both_stream_proto_rawDescGZIP(), []int{1}
}

func (x *BothStreamResponse) GetRspIdx() string {
	if x != nil {
		return x.RspIdx
	}
	return ""
}

func (x *BothStreamResponse) GetRspData() string {
	if x != nil {
		return x.RspData
	}
	return ""
}

var File_both_stream_proto protoreflect.FileDescriptor

var file_both_stream_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x11, 0x42, 0x6f,
	0x74, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x49, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x71, 0x49, 0x64, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x46, 0x0a, 0x12, 0x42, 0x6f, 0x74, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x73, 0x70, 0x49, 0x64,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x73, 0x70, 0x49, 0x64, 0x78, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x32, 0x54, 0x0a, 0x0a, 0x42, 0x6f, 0x74,
	0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x46, 0x0a, 0x09, 0x64, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x74,
	0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x74, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_both_stream_proto_rawDescOnce sync.Once
	file_both_stream_proto_rawDescData = file_both_stream_proto_rawDesc
)

func file_both_stream_proto_rawDescGZIP() []byte {
	file_both_stream_proto_rawDescOnce.Do(func() {
		file_both_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_both_stream_proto_rawDescData)
	})
	return file_both_stream_proto_rawDescData
}

var file_both_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_both_stream_proto_goTypes = []interface{}{
	(*BothStreamRequest)(nil),  // 0: proto.BothStreamRequest
	(*BothStreamResponse)(nil), // 1: proto.BothStreamResponse
}
var file_both_stream_proto_depIdxs = []int32{
	0, // 0: proto.BothStream.doRequest:input_type -> proto.BothStreamRequest
	1, // 1: proto.BothStream.doRequest:output_type -> proto.BothStreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_both_stream_proto_init() }
func file_both_stream_proto_init() {
	if File_both_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_both_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BothStreamRequest); i {
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
		file_both_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BothStreamResponse); i {
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
			RawDescriptor: file_both_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_both_stream_proto_goTypes,
		DependencyIndexes: file_both_stream_proto_depIdxs,
		MessageInfos:      file_both_stream_proto_msgTypes,
	}.Build()
	File_both_stream_proto = out.File
	file_both_stream_proto_rawDesc = nil
	file_both_stream_proto_goTypes = nil
	file_both_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BothStreamClient is the client API for BothStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BothStreamClient interface {
	DoRequest(ctx context.Context, opts ...grpc.CallOption) (BothStream_DoRequestClient, error)
}

type bothStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewBothStreamClient(cc grpc.ClientConnInterface) BothStreamClient {
	return &bothStreamClient{cc}
}

func (c *bothStreamClient) DoRequest(ctx context.Context, opts ...grpc.CallOption) (BothStream_DoRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BothStream_serviceDesc.Streams[0], "/proto.BothStream/doRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &bothStreamDoRequestClient{stream}
	return x, nil
}

type BothStream_DoRequestClient interface {
	Send(*BothStreamRequest) error
	Recv() (*BothStreamResponse, error)
	grpc.ClientStream
}

type bothStreamDoRequestClient struct {
	grpc.ClientStream
}

func (x *bothStreamDoRequestClient) Send(m *BothStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bothStreamDoRequestClient) Recv() (*BothStreamResponse, error) {
	m := new(BothStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BothStreamServer is the server API for BothStream service.
type BothStreamServer interface {
	DoRequest(BothStream_DoRequestServer) error
}

// UnimplementedBothStreamServer can be embedded to have forward compatible implementations.
type UnimplementedBothStreamServer struct {
}

func (*UnimplementedBothStreamServer) DoRequest(BothStream_DoRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method DoRequest not implemented")
}

func RegisterBothStreamServer(s *grpc.Server, srv BothStreamServer) {
	s.RegisterService(&_BothStream_serviceDesc, srv)
}

func _BothStream_DoRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BothStreamServer).DoRequest(&bothStreamDoRequestServer{stream})
}

type BothStream_DoRequestServer interface {
	Send(*BothStreamResponse) error
	Recv() (*BothStreamRequest, error)
	grpc.ServerStream
}

type bothStreamDoRequestServer struct {
	grpc.ServerStream
}

func (x *bothStreamDoRequestServer) Send(m *BothStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bothStreamDoRequestServer) Recv() (*BothStreamRequest, error) {
	m := new(BothStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BothStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BothStream",
	HandlerType: (*BothStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "doRequest",
			Handler:       _BothStream_DoRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "both_stream.proto",
}
