// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: db.proto

package db

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

type SingleRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Age       int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SingleRow) Reset() {
	*x = SingleRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleRow) ProtoMessage() {}

func (x *SingleRow) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleRow.ProtoReflect.Descriptor instead.
func (*SingleRow) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{0}
}

func (x *SingleRow) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SingleRow) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *SingleRow) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *SingleRow) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *SingleRow) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Rows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*SingleRow `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Rows) Reset() {
	*x = Rows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rows) ProtoMessage() {}

func (x *Rows) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rows.ProtoReflect.Descriptor instead.
func (*Rows) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{1}
}

func (x *Rows) GetRows() []*SingleRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

type LimitOffset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *LimitOffset) Reset() {
	*x = LimitOffset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitOffset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitOffset) ProtoMessage() {}

func (x *LimitOffset) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitOffset.ProtoReflect.Descriptor instead.
func (*LimitOffset) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{2}
}

func (x *LimitOffset) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LimitOffset) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{4}
}

var File_db_proto protoreflect.FileDescriptor

var file_db_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64, 0x62, 0x22, 0x7d,
	0x0a, 0x09, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x29, 0x0a,
	0x04, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52,
	0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0x61, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x44, 0x42,
	0x12, 0x0f, 0x2e, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x1a, 0x08, 0x2e, 0x64, 0x62, 0x2e, 0x52, 0x6f, 0x77, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a,
	0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0d, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x1a, 0x0d, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x52, 0x6f, 0x77, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_proto_rawDescOnce sync.Once
	file_db_proto_rawDescData = file_db_proto_rawDesc
)

func file_db_proto_rawDescGZIP() []byte {
	file_db_proto_rawDescOnce.Do(func() {
		file_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_proto_rawDescData)
	})
	return file_db_proto_rawDescData
}

var file_db_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_db_proto_goTypes = []interface{}{
	(*SingleRow)(nil),   // 0: db.SingleRow
	(*Rows)(nil),        // 1: db.Rows
	(*LimitOffset)(nil), // 2: db.LimitOffset
	(*Id)(nil),          // 3: db.Id
	(*Empty)(nil),       // 4: db.Empty
}
var file_db_proto_depIdxs = []int32{
	0, // 0: db.Rows.rows:type_name -> db.SingleRow
	2, // 1: db.DatabaseService.GetDB:input_type -> db.LimitOffset
	0, // 2: db.DatabaseService.Insert:input_type -> db.SingleRow
	1, // 3: db.DatabaseService.GetDB:output_type -> db.Rows
	0, // 4: db.DatabaseService.Insert:output_type -> db.SingleRow
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_db_proto_init() }
func file_db_proto_init() {
	if File_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleRow); i {
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
		file_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rows); i {
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
		file_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitOffset); i {
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
		file_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_db_proto_goTypes,
		DependencyIndexes: file_db_proto_depIdxs,
		MessageInfos:      file_db_proto_msgTypes,
	}.Build()
	File_db_proto = out.File
	file_db_proto_rawDesc = nil
	file_db_proto_goTypes = nil
	file_db_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	GetDB(ctx context.Context, in *LimitOffset, opts ...grpc.CallOption) (*Rows, error)
	Insert(ctx context.Context, in *SingleRow, opts ...grpc.CallOption) (*SingleRow, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) GetDB(ctx context.Context, in *LimitOffset, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/db.DatabaseService/GetDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Insert(ctx context.Context, in *SingleRow, opts ...grpc.CallOption) (*SingleRow, error) {
	out := new(SingleRow)
	err := c.cc.Invoke(ctx, "/db.DatabaseService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
type DatabaseServiceServer interface {
	GetDB(context.Context, *LimitOffset) (*Rows, error)
	Insert(context.Context, *SingleRow) (*SingleRow, error)
}

// UnimplementedDatabaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (*UnimplementedDatabaseServiceServer) GetDB(context.Context, *LimitOffset) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDB not implemented")
}
func (*UnimplementedDatabaseServiceServer) Insert(context.Context, *SingleRow) (*SingleRow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

func RegisterDatabaseServiceServer(s *grpc.Server, srv DatabaseServiceServer) {
	s.RegisterService(&_DatabaseService_serviceDesc, srv)
}

func _DatabaseService_GetDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitOffset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DatabaseService/GetDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDB(ctx, req.(*LimitOffset))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DatabaseService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Insert(ctx, req.(*SingleRow))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDB",
			Handler:    _DatabaseService_GetDB_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _DatabaseService_Insert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}
