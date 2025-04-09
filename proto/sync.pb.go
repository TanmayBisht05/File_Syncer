// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/sync.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileChange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Action        string                 `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"` // create, update, delete
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content       []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileChange) Reset() {
	*x = FileChange{}
	mi := &file_proto_sync_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChange) ProtoMessage() {}

func (x *FileChange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sync_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChange.ProtoReflect.Descriptor instead.
func (*FileChange) Descriptor() ([]byte, []int) {
	return file_proto_sync_proto_rawDescGZIP(), []int{0}
}

func (x *FileChange) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileChange) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FileChange) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *FileChange) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_proto_sync_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sync_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_sync_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_sync_proto protoreflect.FileDescriptor

const file_proto_sync_proto_rawDesc = "" +
	"\n" +
	"\x10proto/sync.proto\x12\x04sync\"x\n" +
	"\n" +
	"FileChange\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x16\n" +
	"\x06action\x18\x02 \x01(\tR\x06action\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\x12\x18\n" +
	"\acontent\x18\x04 \x01(\fR\acontent\"\x1d\n" +
	"\x03Ack\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status28\n" +
	"\vSyncService\x12)\n" +
	"\n" +
	"SendChange\x12\x10.sync.FileChange\x1a\t.sync.AckB\x0fZ\r./proto;protob\x06proto3"

var (
	file_proto_sync_proto_rawDescOnce sync.Once
	file_proto_sync_proto_rawDescData []byte
)

func file_proto_sync_proto_rawDescGZIP() []byte {
	file_proto_sync_proto_rawDescOnce.Do(func() {
		file_proto_sync_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_sync_proto_rawDesc), len(file_proto_sync_proto_rawDesc)))
	})
	return file_proto_sync_proto_rawDescData
}

var file_proto_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_sync_proto_goTypes = []any{
	(*FileChange)(nil), // 0: sync.FileChange
	(*Ack)(nil),        // 1: sync.Ack
}
var file_proto_sync_proto_depIdxs = []int32{
	0, // 0: sync.SyncService.SendChange:input_type -> sync.FileChange
	1, // 1: sync.SyncService.SendChange:output_type -> sync.Ack
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sync_proto_init() }
func file_proto_sync_proto_init() {
	if File_proto_sync_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_sync_proto_rawDesc), len(file_proto_sync_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sync_proto_goTypes,
		DependencyIndexes: file_proto_sync_proto_depIdxs,
		MessageInfos:      file_proto_sync_proto_msgTypes,
	}.Build()
	File_proto_sync_proto = out.File
	file_proto_sync_proto_goTypes = nil
	file_proto_sync_proto_depIdxs = nil
}
