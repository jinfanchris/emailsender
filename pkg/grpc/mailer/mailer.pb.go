// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: mailer.proto

package mailer

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

type MailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiKey        string                 `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	Receiver      string                 `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Body          string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MailRequest) Reset() {
	*x = MailRequest{}
	mi := &file_mailer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailRequest) ProtoMessage() {}

func (x *MailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailRequest.ProtoReflect.Descriptor instead.
func (*MailRequest) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *MailRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *MailRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *MailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *MailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Uuid struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Uuid) Reset() {
	*x = Uuid{}
	mi := &file_mailer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Uuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uuid) ProtoMessage() {}

func (x *Uuid) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uuid.ProtoReflect.Descriptor instead.
func (*Uuid) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *Uuid) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Info          string                 `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_mailer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Status) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_mailer_proto protoreflect.FileDescriptor

var file_mailer_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1a, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x63, 0x0a, 0x0b, 0x4d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x1a,
	0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x11, 0x5a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_mailer_proto_rawDescOnce sync.Once
	file_mailer_proto_rawDescData []byte
)

func file_mailer_proto_rawDescGZIP() []byte {
	file_mailer_proto_rawDescOnce.Do(func() {
		file_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mailer_proto_rawDesc), len(file_mailer_proto_rawDesc)))
	})
	return file_mailer_proto_rawDescData
}

var file_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mailer_proto_goTypes = []any{
	(*MailRequest)(nil), // 0: mailer.MailRequest
	(*Uuid)(nil),        // 1: mailer.Uuid
	(*Status)(nil),      // 2: mailer.Status
}
var file_mailer_proto_depIdxs = []int32{
	0, // 0: mailer.MailService.SendMail:input_type -> mailer.MailRequest
	1, // 1: mailer.MailService.State:input_type -> mailer.Uuid
	1, // 2: mailer.MailService.SendMail:output_type -> mailer.Uuid
	2, // 3: mailer.MailService.State:output_type -> mailer.Status
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mailer_proto_init() }
func file_mailer_proto_init() {
	if File_mailer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mailer_proto_rawDesc), len(file_mailer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailer_proto_goTypes,
		DependencyIndexes: file_mailer_proto_depIdxs,
		MessageInfos:      file_mailer_proto_msgTypes,
	}.Build()
	File_mailer_proto = out.File
	file_mailer_proto_goTypes = nil
	file_mailer_proto_depIdxs = nil
}
