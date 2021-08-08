// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: prime/protoprime/prime.proto

package protoprime

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

type Prime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Prime) Reset() {
	*x = Prime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prime_protoprime_prime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prime) ProtoMessage() {}

func (x *Prime) ProtoReflect() protoreflect.Message {
	mi := &file_prime_protoprime_prime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prime.ProtoReflect.Descriptor instead.
func (*Prime) Descriptor() ([]byte, []int) {
	return file_prime_protoprime_prime_proto_rawDescGZIP(), []int{0}
}

func (x *Prime) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

type PrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prime *Prime `protobuf:"bytes,1,opt,name=prime,proto3" json:"prime,omitempty"`
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prime_protoprime_prime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prime_protoprime_prime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_prime_protoprime_prime_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeRequest) GetPrime() *Prime {
	if x != nil {
		return x.Prime
	}
	return nil
}

type PrimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeResponse) Reset() {
	*x = PrimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prime_protoprime_prime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeResponse) ProtoMessage() {}

func (x *PrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prime_protoprime_prime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeResponse.ProtoReflect.Descriptor instead.
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return file_prime_protoprime_prime_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_prime_protoprime_prime_proto protoreflect.FileDescriptor

var file_prime_protoprime_prime_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x22, 0x32, 0x0a, 0x0c,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x22, 0x27, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x44, 0x0a, 0x0c, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prime_protoprime_prime_proto_rawDescOnce sync.Once
	file_prime_protoprime_prime_proto_rawDescData = file_prime_protoprime_prime_proto_rawDesc
)

func file_prime_protoprime_prime_proto_rawDescGZIP() []byte {
	file_prime_protoprime_prime_proto_rawDescOnce.Do(func() {
		file_prime_protoprime_prime_proto_rawDescData = protoimpl.X.CompressGZIP(file_prime_protoprime_prime_proto_rawDescData)
	})
	return file_prime_protoprime_prime_proto_rawDescData
}

var file_prime_protoprime_prime_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prime_protoprime_prime_proto_goTypes = []interface{}{
	(*Prime)(nil),         // 0: prime.Prime
	(*PrimeRequest)(nil),  // 1: prime.PrimeRequest
	(*PrimeResponse)(nil), // 2: prime.PrimeResponse
}
var file_prime_protoprime_prime_proto_depIdxs = []int32{
	0, // 0: prime.PrimeRequest.prime:type_name -> prime.Prime
	1, // 1: prime.PrimeService.Prime:input_type -> prime.PrimeRequest
	2, // 2: prime.PrimeService.Prime:output_type -> prime.PrimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prime_protoprime_prime_proto_init() }
func file_prime_protoprime_prime_proto_init() {
	if File_prime_protoprime_prime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prime_protoprime_prime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prime); i {
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
		file_prime_protoprime_prime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeRequest); i {
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
		file_prime_protoprime_prime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeResponse); i {
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
			RawDescriptor: file_prime_protoprime_prime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prime_protoprime_prime_proto_goTypes,
		DependencyIndexes: file_prime_protoprime_prime_proto_depIdxs,
		MessageInfos:      file_prime_protoprime_prime_proto_msgTypes,
	}.Build()
	File_prime_protoprime_prime_proto = out.File
	file_prime_protoprime_prime_proto_rawDesc = nil
	file_prime_protoprime_prime_proto_goTypes = nil
	file_prime_protoprime_prime_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PrimeServiceClient is the client API for PrimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimeServiceClient interface {
	Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_PrimeClient, error)
}

type primeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeServiceClient(cc grpc.ClientConnInterface) PrimeServiceClient {
	return &primeServiceClient{cc}
}

func (c *primeServiceClient) Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_PrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeService_serviceDesc.Streams[0], "/prime.PrimeService/Prime", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServicePrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeService_PrimeClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type primeServicePrimeClient struct {
	grpc.ClientStream
}

func (x *primeServicePrimeClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeServiceServer is the server API for PrimeService service.
type PrimeServiceServer interface {
	Prime(*PrimeRequest, PrimeService_PrimeServer) error
}

// UnimplementedPrimeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrimeServiceServer struct {
}

func (*UnimplementedPrimeServiceServer) Prime(*PrimeRequest, PrimeService_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}

func RegisterPrimeServiceServer(s *grpc.Server, srv PrimeServiceServer) {
	s.RegisterService(&_PrimeService_serviceDesc, srv)
}

func _PrimeService_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeServiceServer).Prime(m, &primeServicePrimeServer{stream})
}

type PrimeService_PrimeServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type primeServicePrimeServer struct {
	grpc.ServerStream
}

func (x *primeServicePrimeServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PrimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prime.PrimeService",
	HandlerType: (*PrimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Prime",
			Handler:       _PrimeService_Prime_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "prime/protoprime/prime.proto",
}
