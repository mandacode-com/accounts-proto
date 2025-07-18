// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: mailer/v1/email_verification.proto

package mailerv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type EmailVerificationEvent struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Email            string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	VerificationLink string                 `protobuf:"bytes,2,opt,name=verification_link,json=verificationLink,proto3" json:"verification_link,omitempty"`
	EventTime        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *EmailVerificationEvent) Reset() {
	*x = EmailVerificationEvent{}
	mi := &file_mailer_v1_email_verification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailVerificationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailVerificationEvent) ProtoMessage() {}

func (x *EmailVerificationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_v1_email_verification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailVerificationEvent.ProtoReflect.Descriptor instead.
func (*EmailVerificationEvent) Descriptor() ([]byte, []int) {
	return file_mailer_v1_email_verification_proto_rawDescGZIP(), []int{0}
}

func (x *EmailVerificationEvent) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailVerificationEvent) GetVerificationLink() string {
	if x != nil {
		return x.VerificationLink
	}
	return ""
}

func (x *EmailVerificationEvent) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

var File_mailer_v1_email_verification_proto protoreflect.FileDescriptor

const file_mailer_v1_email_verification_proto_rawDesc = "" +
	"\n" +
	"\"mailer/v1/email_verification.proto\x12\tmailer.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a#third_party/validate/validate.proto\"\xa9\x01\n" +
	"\x16EmailVerificationEvent\x12\x1d\n" +
	"\x05email\x18\x01 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x125\n" +
	"\x11verification_link\x18\x02 \x01(\tB\b\xfaB\x05r\x03\x88\x01\x01R\x10verificationLink\x129\n" +
	"\n" +
	"event_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\teventTimeB?Z=github.com/mandacode-com/accounts-proto/go/mailer/v1;mailerv1b\x06proto3"

var (
	file_mailer_v1_email_verification_proto_rawDescOnce sync.Once
	file_mailer_v1_email_verification_proto_rawDescData []byte
)

func file_mailer_v1_email_verification_proto_rawDescGZIP() []byte {
	file_mailer_v1_email_verification_proto_rawDescOnce.Do(func() {
		file_mailer_v1_email_verification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mailer_v1_email_verification_proto_rawDesc), len(file_mailer_v1_email_verification_proto_rawDesc)))
	})
	return file_mailer_v1_email_verification_proto_rawDescData
}

var file_mailer_v1_email_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mailer_v1_email_verification_proto_goTypes = []any{
	(*EmailVerificationEvent)(nil), // 0: mailer.v1.EmailVerificationEvent
	(*timestamppb.Timestamp)(nil),  // 1: google.protobuf.Timestamp
}
var file_mailer_v1_email_verification_proto_depIdxs = []int32{
	1, // 0: mailer.v1.EmailVerificationEvent.event_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mailer_v1_email_verification_proto_init() }
func file_mailer_v1_email_verification_proto_init() {
	if File_mailer_v1_email_verification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mailer_v1_email_verification_proto_rawDesc), len(file_mailer_v1_email_verification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mailer_v1_email_verification_proto_goTypes,
		DependencyIndexes: file_mailer_v1_email_verification_proto_depIdxs,
		MessageInfos:      file_mailer_v1_email_verification_proto_msgTypes,
	}.Build()
	File_mailer_v1_email_verification_proto = out.File
	file_mailer_v1_email_verification_proto_goTypes = nil
	file_mailer_v1_email_verification_proto_depIdxs = nil
}
