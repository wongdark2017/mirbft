// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: testservice.proto

package protobufs

import (
	context "context"
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

type TestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn      int32  `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TestReq) Reset() {
	*x = TestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReq) ProtoMessage() {}

func (x *TestReq) ProtoReflect() protoreflect.Message {
	mi := &file_testservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReq.ProtoReflect.Descriptor instead.
func (*TestReq) Descriptor() ([]byte, []int) {
	return file_testservice_proto_rawDescGZIP(), []int{0}
}

func (x *TestReq) GetSn() int32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *TestReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type TestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn      int32  `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TestResp) Reset() {
	*x = TestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResp) ProtoMessage() {}

func (x *TestResp) ProtoReflect() protoreflect.Message {
	mi := &file_testservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResp.ProtoReflect.Descriptor instead.
func (*TestResp) Descriptor() ([]byte, []int) {
	return file_testservice_proto_rawDescGZIP(), []int{1}
}

func (x *TestResp) GetSn() int32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *TestResp) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_testservice_proto protoreflect.FileDescriptor

var file_testservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x22, 0x33,
	0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x34, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x73, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x42, 0x0a, 0x0b, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testservice_proto_rawDescOnce sync.Once
	file_testservice_proto_rawDescData = file_testservice_proto_rawDesc
)

func file_testservice_proto_rawDescGZIP() []byte {
	file_testservice_proto_rawDescOnce.Do(func() {
		file_testservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_testservice_proto_rawDescData)
	})
	return file_testservice_proto_rawDescData
}

var file_testservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testservice_proto_goTypes = []interface{}{
	(*TestReq)(nil),  // 0: protobufs.TestReq
	(*TestResp)(nil), // 1: protobufs.TestResp
}
var file_testservice_proto_depIdxs = []int32{
	0, // 0: protobufs.TestService.Test:input_type -> protobufs.TestReq
	1, // 1: protobufs.TestService.Test:output_type -> protobufs.TestResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testservice_proto_init() }
func file_testservice_proto_init() {
	if File_testservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReq); i {
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
		file_testservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResp); i {
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
			RawDescriptor: file_testservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testservice_proto_goTypes,
		DependencyIndexes: file_testservice_proto_depIdxs,
		MessageInfos:      file_testservice_proto_msgTypes,
	}.Build()
	File_testservice_proto = out.File
	file_testservice_proto_rawDesc = nil
	file_testservice_proto_goTypes = nil
	file_testservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	Test(ctx context.Context, opts ...grpc.CallOption) (TestService_TestClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Test(ctx context.Context, opts ...grpc.CallOption) (TestService_TestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/protobufs.TestService/Test", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestClient{stream}
	return x, nil
}

type TestService_TestClient interface {
	Send(*TestReq) error
	Recv() (*TestResp, error)
	grpc.ClientStream
}

type testServiceTestClient struct {
	grpc.ClientStream
}

func (x *testServiceTestClient) Send(m *TestReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceTestClient) Recv() (*TestResp, error) {
	m := new(TestResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	Test(TestService_TestServer) error
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) Test(TestService_TestServer) error {
	return status.Errorf(codes.Unimplemented, "method Test not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Test_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).Test(&testServiceTestServer{stream})
}

type TestService_TestServer interface {
	Send(*TestResp) error
	Recv() (*TestReq, error)
	grpc.ServerStream
}

type testServiceTestServer struct {
	grpc.ServerStream
}

func (x *testServiceTestServer) Send(m *TestResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceTestServer) Recv() (*TestReq, error) {
	m := new(TestReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Test",
			Handler:       _TestService_Test_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "testservice.proto",
}
