// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.12.4
// source: contentiontest.proto

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

type ControlCommand_Cmd int32

const (
	ControlCommand_START_STAT ControlCommand_Cmd = 0
	ControlCommand_STOP_STAT  ControlCommand_Cmd = 1
)

// Enum value maps for ControlCommand_Cmd.
var (
	ControlCommand_Cmd_name = map[int32]string{
		0: "START_STAT",
		1: "STOP_STAT",
	}
	ControlCommand_Cmd_value = map[string]int32{
		"START_STAT": 0,
		"STOP_STAT":  1,
	}
)

func (x ControlCommand_Cmd) Enum() *ControlCommand_Cmd {
	p := new(ControlCommand_Cmd)
	*p = x
	return p
}

func (x ControlCommand_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlCommand_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_contentiontest_proto_enumTypes[0].Descriptor()
}

func (ControlCommand_Cmd) Type() protoreflect.EnumType {
	return &file_contentiontest_proto_enumTypes[0]
}

func (x ControlCommand_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlCommand_Cmd.Descriptor instead.
func (ControlCommand_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_contentiontest_proto_rawDescGZIP(), []int{0, 0}
}

type ControlCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd ControlCommand_Cmd `protobuf:"varint,1,opt,name=cmd,proto3,enum=protobufs.ControlCommand_Cmd" json:"cmd,omitempty"`
}

func (x *ControlCommand) Reset() {
	*x = ControlCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contentiontest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlCommand) ProtoMessage() {}

func (x *ControlCommand) ProtoReflect() protoreflect.Message {
	mi := &file_contentiontest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlCommand.ProtoReflect.Descriptor instead.
func (*ControlCommand) Descriptor() ([]byte, []int) {
	return file_contentiontest_proto_rawDescGZIP(), []int{0}
}

func (x *ControlCommand) GetCmd() ControlCommand_Cmd {
	if x != nil {
		return x.Cmd
	}
	return ControlCommand_START_STAT
}

type ControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ControlResponse) Reset() {
	*x = ControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contentiontest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlResponse) ProtoMessage() {}

func (x *ControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contentiontest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlResponse.ProtoReflect.Descriptor instead.
func (*ControlResponse) Descriptor() ([]byte, []int) {
	return file_contentiontest_proto_rawDescGZIP(), []int{1}
}

var File_contentiontest_proto protoreflect.FileDescriptor

var file_contentiontest_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x73, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x67, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x22, 0x24, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54,
	0x4f, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x10, 0x01, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9f, 0x01, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x30,
	0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x30, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9f,
	0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x31, 0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x31, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x40,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x59, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_contentiontest_proto_rawDescOnce sync.Once
	file_contentiontest_proto_rawDescData = file_contentiontest_proto_rawDesc
)

func file_contentiontest_proto_rawDescGZIP() []byte {
	file_contentiontest_proto_rawDescOnce.Do(func() {
		file_contentiontest_proto_rawDescData = protoimpl.X.CompressGZIP(file_contentiontest_proto_rawDescData)
	})
	return file_contentiontest_proto_rawDescData
}

var file_contentiontest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contentiontest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contentiontest_proto_goTypes = []interface{}{
	(ControlCommand_Cmd)(0), // 0: protobufs.ControlCommand.Cmd
	(*ControlCommand)(nil),  // 1: protobufs.ControlCommand
	(*ControlResponse)(nil), // 2: protobufs.ControlResponse
	(*ClientRequest)(nil),   // 3: protobufs.ClientRequest
	(*ClientResponse)(nil),  // 4: protobufs.ClientResponse
}
var file_contentiontest_proto_depIdxs = []int32{
	0, // 0: protobufs.ControlCommand.cmd:type_name -> protobufs.ControlCommand.Cmd
	3, // 1: protobufs.ContentionTest0.TestContention0:input_type -> protobufs.ClientRequest
	1, // 2: protobufs.ContentionTest0.Control:input_type -> protobufs.ControlCommand
	3, // 3: protobufs.ContentionTest1.TestContention1:input_type -> protobufs.ClientRequest
	1, // 4: protobufs.ContentionTest1.Control:input_type -> protobufs.ControlCommand
	1, // 5: protobufs.ContentionTestControl.Control:input_type -> protobufs.ControlCommand
	4, // 6: protobufs.ContentionTest0.TestContention0:output_type -> protobufs.ClientResponse
	2, // 7: protobufs.ContentionTest0.Control:output_type -> protobufs.ControlResponse
	4, // 8: protobufs.ContentionTest1.TestContention1:output_type -> protobufs.ClientResponse
	2, // 9: protobufs.ContentionTest1.Control:output_type -> protobufs.ControlResponse
	2, // 10: protobufs.ContentionTestControl.Control:output_type -> protobufs.ControlResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contentiontest_proto_init() }
func file_contentiontest_proto_init() {
	if File_contentiontest_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contentiontest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlCommand); i {
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
		file_contentiontest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlResponse); i {
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
			RawDescriptor: file_contentiontest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_contentiontest_proto_goTypes,
		DependencyIndexes: file_contentiontest_proto_depIdxs,
		EnumInfos:         file_contentiontest_proto_enumTypes,
		MessageInfos:      file_contentiontest_proto_msgTypes,
	}.Build()
	File_contentiontest_proto = out.File
	file_contentiontest_proto_rawDesc = nil
	file_contentiontest_proto_goTypes = nil
	file_contentiontest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContentionTest0Client is the client API for ContentionTest0 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentionTest0Client interface {
	TestContention0(ctx context.Context, opts ...grpc.CallOption) (ContentionTest0_TestContention0Client, error)
	Control(ctx context.Context, in *ControlCommand, opts ...grpc.CallOption) (*ControlResponse, error)
}

type contentionTest0Client struct {
	cc grpc.ClientConnInterface
}

func NewContentionTest0Client(cc grpc.ClientConnInterface) ContentionTest0Client {
	return &contentionTest0Client{cc}
}

func (c *contentionTest0Client) TestContention0(ctx context.Context, opts ...grpc.CallOption) (ContentionTest0_TestContention0Client, error) {
	stream, err := c.cc.NewStream(ctx, &_ContentionTest0_serviceDesc.Streams[0], "/protobufs.ContentionTest0/TestContention0", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentionTest0TestContention0Client{stream}
	return x, nil
}

type ContentionTest0_TestContention0Client interface {
	Send(*ClientRequest) error
	Recv() (*ClientResponse, error)
	grpc.ClientStream
}

type contentionTest0TestContention0Client struct {
	grpc.ClientStream
}

func (x *contentionTest0TestContention0Client) Send(m *ClientRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentionTest0TestContention0Client) Recv() (*ClientResponse, error) {
	m := new(ClientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentionTest0Client) Control(ctx context.Context, in *ControlCommand, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/protobufs.ContentionTest0/Control", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentionTest0Server is the server API for ContentionTest0 service.
type ContentionTest0Server interface {
	TestContention0(ContentionTest0_TestContention0Server) error
	Control(context.Context, *ControlCommand) (*ControlResponse, error)
}

// UnimplementedContentionTest0Server can be embedded to have forward compatible implementations.
type UnimplementedContentionTest0Server struct {
}

func (*UnimplementedContentionTest0Server) TestContention0(ContentionTest0_TestContention0Server) error {
	return status.Errorf(codes.Unimplemented, "method TestContention0 not implemented")
}
func (*UnimplementedContentionTest0Server) Control(context.Context, *ControlCommand) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}

func RegisterContentionTest0Server(s *grpc.Server, srv ContentionTest0Server) {
	s.RegisterService(&_ContentionTest0_serviceDesc, srv)
}

func _ContentionTest0_TestContention0_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentionTest0Server).TestContention0(&contentionTest0TestContention0Server{stream})
}

type ContentionTest0_TestContention0Server interface {
	Send(*ClientResponse) error
	Recv() (*ClientRequest, error)
	grpc.ServerStream
}

type contentionTest0TestContention0Server struct {
	grpc.ServerStream
}

func (x *contentionTest0TestContention0Server) Send(m *ClientResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentionTest0TestContention0Server) Recv() (*ClientRequest, error) {
	m := new(ClientRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ContentionTest0_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentionTest0Server).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.ContentionTest0/Control",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentionTest0Server).Control(ctx, req.(*ControlCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentionTest0_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.ContentionTest0",
	HandlerType: (*ContentionTest0Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Control",
			Handler:    _ContentionTest0_Control_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestContention0",
			Handler:       _ContentionTest0_TestContention0_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "contentiontest.proto",
}

// ContentionTest1Client is the client API for ContentionTest1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentionTest1Client interface {
	TestContention1(ctx context.Context, opts ...grpc.CallOption) (ContentionTest1_TestContention1Client, error)
	Control(ctx context.Context, in *ControlCommand, opts ...grpc.CallOption) (*ControlResponse, error)
}

type contentionTest1Client struct {
	cc grpc.ClientConnInterface
}

func NewContentionTest1Client(cc grpc.ClientConnInterface) ContentionTest1Client {
	return &contentionTest1Client{cc}
}

func (c *contentionTest1Client) TestContention1(ctx context.Context, opts ...grpc.CallOption) (ContentionTest1_TestContention1Client, error) {
	stream, err := c.cc.NewStream(ctx, &_ContentionTest1_serviceDesc.Streams[0], "/protobufs.ContentionTest1/TestContention1", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentionTest1TestContention1Client{stream}
	return x, nil
}

type ContentionTest1_TestContention1Client interface {
	Send(*ClientRequest) error
	Recv() (*ClientResponse, error)
	grpc.ClientStream
}

type contentionTest1TestContention1Client struct {
	grpc.ClientStream
}

func (x *contentionTest1TestContention1Client) Send(m *ClientRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentionTest1TestContention1Client) Recv() (*ClientResponse, error) {
	m := new(ClientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentionTest1Client) Control(ctx context.Context, in *ControlCommand, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/protobufs.ContentionTest1/Control", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentionTest1Server is the server API for ContentionTest1 service.
type ContentionTest1Server interface {
	TestContention1(ContentionTest1_TestContention1Server) error
	Control(context.Context, *ControlCommand) (*ControlResponse, error)
}

// UnimplementedContentionTest1Server can be embedded to have forward compatible implementations.
type UnimplementedContentionTest1Server struct {
}

func (*UnimplementedContentionTest1Server) TestContention1(ContentionTest1_TestContention1Server) error {
	return status.Errorf(codes.Unimplemented, "method TestContention1 not implemented")
}
func (*UnimplementedContentionTest1Server) Control(context.Context, *ControlCommand) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}

func RegisterContentionTest1Server(s *grpc.Server, srv ContentionTest1Server) {
	s.RegisterService(&_ContentionTest1_serviceDesc, srv)
}

func _ContentionTest1_TestContention1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentionTest1Server).TestContention1(&contentionTest1TestContention1Server{stream})
}

type ContentionTest1_TestContention1Server interface {
	Send(*ClientResponse) error
	Recv() (*ClientRequest, error)
	grpc.ServerStream
}

type contentionTest1TestContention1Server struct {
	grpc.ServerStream
}

func (x *contentionTest1TestContention1Server) Send(m *ClientResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentionTest1TestContention1Server) Recv() (*ClientRequest, error) {
	m := new(ClientRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ContentionTest1_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentionTest1Server).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.ContentionTest1/Control",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentionTest1Server).Control(ctx, req.(*ControlCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentionTest1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.ContentionTest1",
	HandlerType: (*ContentionTest1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Control",
			Handler:    _ContentionTest1_Control_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestContention1",
			Handler:       _ContentionTest1_TestContention1_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "contentiontest.proto",
}

// ContentionTestControlClient is the client API for ContentionTestControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentionTestControlClient interface {
	Control(ctx context.Context, in *ControlCommand, opts ...grpc.CallOption) (*ControlResponse, error)
}

type contentionTestControlClient struct {
	cc grpc.ClientConnInterface
}

func NewContentionTestControlClient(cc grpc.ClientConnInterface) ContentionTestControlClient {
	return &contentionTestControlClient{cc}
}

func (c *contentionTestControlClient) Control(ctx context.Context, in *ControlCommand, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/protobufs.ContentionTestControl/Control", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentionTestControlServer is the server API for ContentionTestControl service.
type ContentionTestControlServer interface {
	Control(context.Context, *ControlCommand) (*ControlResponse, error)
}

// UnimplementedContentionTestControlServer can be embedded to have forward compatible implementations.
type UnimplementedContentionTestControlServer struct {
}

func (*UnimplementedContentionTestControlServer) Control(context.Context, *ControlCommand) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}

func RegisterContentionTestControlServer(s *grpc.Server, srv ContentionTestControlServer) {
	s.RegisterService(&_ContentionTestControl_serviceDesc, srv)
}

func _ContentionTestControl_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentionTestControlServer).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.ContentionTestControl/Control",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentionTestControlServer).Control(ctx, req.(*ControlCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentionTestControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.ContentionTestControl",
	HandlerType: (*ContentionTestControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Control",
			Handler:    _ContentionTestControl_Control_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contentiontest.proto",
}
