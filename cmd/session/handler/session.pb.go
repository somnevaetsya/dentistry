// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: handler/session.proto

package handler

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

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_handler_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_handler_session_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SessionModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid      uint64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	CookieValue string `protobuf:"bytes,2,opt,name=cookieValue,proto3" json:"cookieValue,omitempty"`
}

func (x *SessionModel) Reset() {
	*x = SessionModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionModel) ProtoMessage() {}

func (x *SessionModel) ProtoReflect() protoreflect.Message {
	mi := &file_handler_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionModel.ProtoReflect.Descriptor instead.
func (*SessionModel) Descriptor() ([]byte, []int) {
	return file_handler_session_proto_rawDescGZIP(), []int{1}
}

func (x *SessionModel) GetUserid() uint64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *SessionModel) GetCookieValue() string {
	if x != nil {
		return x.CookieValue
	}
	return ""
}

type CodeModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid uint64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CodeModel) Reset() {
	*x = CodeModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeModel) ProtoMessage() {}

func (x *CodeModel) ProtoReflect() protoreflect.Message {
	mi := &file_handler_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeModel.ProtoReflect.Descriptor instead.
func (*CodeModel) Descriptor() ([]byte, []int) {
	return file_handler_session_proto_rawDescGZIP(), []int{2}
}

func (x *CodeModel) GetUserid() uint64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *CodeModel) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_handler_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_handler_session_proto_rawDescGZIP(), []int{3}
}

func (x *Code) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SessionValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CookieValue string `protobuf:"bytes,1,opt,name=cookieValue,proto3" json:"cookieValue,omitempty"`
}

func (x *SessionValue) Reset() {
	*x = SessionValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionValue) ProtoMessage() {}

func (x *SessionValue) ProtoReflect() protoreflect.Message {
	mi := &file_handler_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionValue.ProtoReflect.Descriptor instead.
func (*SessionValue) Descriptor() ([]byte, []int) {
	return file_handler_session_proto_rawDescGZIP(), []int{4}
}

func (x *SessionValue) GetCookieValue() string {
	if x != nil {
		return x.CookieValue
	}
	return ""
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_handler_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_handler_session_proto_rawDescGZIP(), []int{5}
}

func (x *Nothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

var File_handler_session_proto protoreflect.FileDescriptor

var file_handler_session_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1a, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x30, 0x0a, 0x0c, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x07, 0x4e,
	0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0xc1, 0x02, 0x0a,
	0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x0f, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x10, 0x2e, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handler_session_proto_rawDescOnce sync.Once
	file_handler_session_proto_rawDescData = file_handler_session_proto_rawDesc
)

func file_handler_session_proto_rawDescGZIP() []byte {
	file_handler_session_proto_rawDescOnce.Do(func() {
		file_handler_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_session_proto_rawDescData)
	})
	return file_handler_session_proto_rawDescData
}

var file_handler_session_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_handler_session_proto_goTypes = []interface{}{
	(*UserId)(nil),       // 0: handler.UserId
	(*SessionModel)(nil), // 1: handler.SessionModel
	(*CodeModel)(nil),    // 2: handler.CodeModel
	(*Code)(nil),         // 3: handler.Code
	(*SessionValue)(nil), // 4: handler.SessionValue
	(*Nothing)(nil),      // 5: handler.Nothing
}
var file_handler_session_proto_depIdxs = []int32{
	1, // 0: handler.SessionChecker.Create:input_type -> handler.SessionModel
	4, // 1: handler.SessionChecker.Get:input_type -> handler.SessionValue
	4, // 2: handler.SessionChecker.Delete:input_type -> handler.SessionValue
	2, // 3: handler.SessionChecker.CreateCode:input_type -> handler.CodeModel
	0, // 4: handler.SessionChecker.GetCode:input_type -> handler.UserId
	0, // 5: handler.SessionChecker.DeleteCode:input_type -> handler.UserId
	5, // 6: handler.SessionChecker.Create:output_type -> handler.Nothing
	0, // 7: handler.SessionChecker.Get:output_type -> handler.UserId
	5, // 8: handler.SessionChecker.Delete:output_type -> handler.Nothing
	5, // 9: handler.SessionChecker.CreateCode:output_type -> handler.Nothing
	3, // 10: handler.SessionChecker.GetCode:output_type -> handler.Code
	5, // 11: handler.SessionChecker.DeleteCode:output_type -> handler.Nothing
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_handler_session_proto_init() }
func file_handler_session_proto_init() {
	if File_handler_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_handler_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionModel); i {
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
		file_handler_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeModel); i {
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
		file_handler_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_handler_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionValue); i {
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
		file_handler_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
			RawDescriptor: file_handler_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_session_proto_goTypes,
		DependencyIndexes: file_handler_session_proto_depIdxs,
		MessageInfos:      file_handler_session_proto_msgTypes,
	}.Build()
	File_handler_session_proto = out.File
	file_handler_session_proto_rawDesc = nil
	file_handler_session_proto_goTypes = nil
	file_handler_session_proto_depIdxs = nil
}