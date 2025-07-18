// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: user/event/v1/user_event.proto

package usereventv1

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

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED EventType = 0
	EventType_USER_DELETED           EventType = 1
	EventType_USER_ARCHIVED          EventType = 2
	EventType_USER_RESTORED          EventType = 3
	EventType_USER_BLOCKED           EventType = 4
	EventType_USER_UNBLOCKED         EventType = 5
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "USER_DELETED",
		2: "USER_ARCHIVED",
		3: "USER_RESTORED",
		4: "USER_BLOCKED",
		5: "USER_UNBLOCKED",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED": 0,
		"USER_DELETED":           1,
		"USER_ARCHIVED":          2,
		"USER_RESTORED":          3,
		"USER_BLOCKED":           4,
		"USER_UNBLOCKED":         5,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_event_v1_user_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_user_event_v1_user_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_user_event_v1_user_event_proto_rawDescGZIP(), []int{0}
}

type UserEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventType     EventType              `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=user.event.v1.EventType" json:"event_type,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SyncCode      *string                `protobuf:"bytes,3,opt,name=sync_code,json=syncCode,proto3,oneof" json:"sync_code,omitempty"` // Sync code for tracking
	EventTime     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserEvent) Reset() {
	*x = UserEvent{}
	mi := &file_user_event_v1_user_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEvent) ProtoMessage() {}

func (x *UserEvent) ProtoReflect() protoreflect.Message {
	mi := &file_user_event_v1_user_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEvent.ProtoReflect.Descriptor instead.
func (*UserEvent) Descriptor() ([]byte, []int) {
	return file_user_event_v1_user_event_proto_rawDescGZIP(), []int{0}
}

func (x *UserEvent) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (x *UserEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserEvent) GetSyncCode() string {
	if x != nil && x.SyncCode != nil {
		return *x.SyncCode
	}
	return ""
}

func (x *UserEvent) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

var File_user_event_v1_user_event_proto protoreflect.FileDescriptor

const file_user_event_v1_user_event_proto_rawDesc = "" +
	"\n" +
	"\x1euser/event/v1/user_event.proto\x12\ruser.event.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a#third_party/validate/validate.proto\"\xe5\x01\n" +
	"\tUserEvent\x12A\n" +
	"\n" +
	"event_type\x18\x01 \x01(\x0e2\x18.user.event.v1.EventTypeB\b\xfaB\x05\x82\x01\x02\x10\x01R\teventType\x12!\n" +
	"\auser_id\x18\x02 \x01(\tB\b\xfaB\x05r\x03\xb0\x01\x01R\x06userId\x12)\n" +
	"\tsync_code\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x01H\x00R\bsyncCode\x88\x01\x01\x129\n" +
	"\n" +
	"event_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\teventTimeB\f\n" +
	"\n" +
	"_sync_code*\x85\x01\n" +
	"\tEventType\x12\x1a\n" +
	"\x16EVENT_TYPE_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fUSER_DELETED\x10\x01\x12\x11\n" +
	"\rUSER_ARCHIVED\x10\x02\x12\x11\n" +
	"\rUSER_RESTORED\x10\x03\x12\x10\n" +
	"\fUSER_BLOCKED\x10\x04\x12\x12\n" +
	"\x0eUSER_UNBLOCKED\x10\x05BFZDgithub.com/mandacode-com/accounts-proto/go/user/event/v1;usereventv1b\x06proto3"

var (
	file_user_event_v1_user_event_proto_rawDescOnce sync.Once
	file_user_event_v1_user_event_proto_rawDescData []byte
)

func file_user_event_v1_user_event_proto_rawDescGZIP() []byte {
	file_user_event_v1_user_event_proto_rawDescOnce.Do(func() {
		file_user_event_v1_user_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_event_v1_user_event_proto_rawDesc), len(file_user_event_v1_user_event_proto_rawDesc)))
	})
	return file_user_event_v1_user_event_proto_rawDescData
}

var file_user_event_v1_user_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_event_v1_user_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_event_v1_user_event_proto_goTypes = []any{
	(EventType)(0),                // 0: user.event.v1.EventType
	(*UserEvent)(nil),             // 1: user.event.v1.UserEvent
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_user_event_v1_user_event_proto_depIdxs = []int32{
	0, // 0: user.event.v1.UserEvent.event_type:type_name -> user.event.v1.EventType
	2, // 1: user.event.v1.UserEvent.event_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_event_v1_user_event_proto_init() }
func file_user_event_v1_user_event_proto_init() {
	if File_user_event_v1_user_event_proto != nil {
		return
	}
	file_user_event_v1_user_event_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_event_v1_user_event_proto_rawDesc), len(file_user_event_v1_user_event_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_event_v1_user_event_proto_goTypes,
		DependencyIndexes: file_user_event_v1_user_event_proto_depIdxs,
		EnumInfos:         file_user_event_v1_user_event_proto_enumTypes,
		MessageInfos:      file_user_event_v1_user_event_proto_msgTypes,
	}.Build()
	File_user_event_v1_user_event_proto = out.File
	file_user_event_v1_user_event_proto_goTypes = nil
	file_user_event_v1_user_event_proto_depIdxs = nil
}
