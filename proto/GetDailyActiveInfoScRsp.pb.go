// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetDailyActiveInfoScRsp.proto

package proto

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

type GetDailyActiveInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                uint32               `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DailyActiveLevelList   []*DailyActivityInfo `protobuf:"bytes,10,rep,name=daily_active_level_list,json=dailyActiveLevelList,proto3" json:"daily_active_level_list,omitempty"`
	DailyActiveQuestIdList []uint32             `protobuf:"varint,3,rep,packed,name=daily_active_quest_id_list,json=dailyActiveQuestIdList,proto3" json:"daily_active_quest_id_list,omitempty"`
	DailyActivePoint       uint32               `protobuf:"varint,2,opt,name=daily_active_point,json=dailyActivePoint,proto3" json:"daily_active_point,omitempty"`
}

func (x *GetDailyActiveInfoScRsp) Reset() {
	*x = GetDailyActiveInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDailyActiveInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyActiveInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyActiveInfoScRsp) ProtoMessage() {}

func (x *GetDailyActiveInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetDailyActiveInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyActiveInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetDailyActiveInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetDailyActiveInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetDailyActiveInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetDailyActiveInfoScRsp) GetDailyActiveLevelList() []*DailyActivityInfo {
	if x != nil {
		return x.DailyActiveLevelList
	}
	return nil
}

func (x *GetDailyActiveInfoScRsp) GetDailyActiveQuestIdList() []uint32 {
	if x != nil {
		return x.DailyActiveQuestIdList
	}
	return nil
}

func (x *GetDailyActiveInfoScRsp) GetDailyActivePoint() uint32 {
	if x != nil {
		return x.DailyActivePoint
	}
	return 0
}

var File_GetDailyActiveInfoScRsp_proto protoreflect.FileDescriptor

var file_GetDailyActiveInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x49,
	0x0a, 0x17, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x14, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x1a, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetDailyActiveInfoScRsp_proto_rawDescOnce sync.Once
	file_GetDailyActiveInfoScRsp_proto_rawDescData = file_GetDailyActiveInfoScRsp_proto_rawDesc
)

func file_GetDailyActiveInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetDailyActiveInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetDailyActiveInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetDailyActiveInfoScRsp_proto_rawDescData)
	})
	return file_GetDailyActiveInfoScRsp_proto_rawDescData
}

var file_GetDailyActiveInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetDailyActiveInfoScRsp_proto_goTypes = []interface{}{
	(*GetDailyActiveInfoScRsp)(nil), // 0: GetDailyActiveInfoScRsp
	(*DailyActivityInfo)(nil),       // 1: DailyActivityInfo
}
var file_GetDailyActiveInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetDailyActiveInfoScRsp.daily_active_level_list:type_name -> DailyActivityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetDailyActiveInfoScRsp_proto_init() }
func file_GetDailyActiveInfoScRsp_proto_init() {
	if File_GetDailyActiveInfoScRsp_proto != nil {
		return
	}
	file_DailyActivityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetDailyActiveInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDailyActiveInfoScRsp); i {
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
			RawDescriptor: file_GetDailyActiveInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetDailyActiveInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetDailyActiveInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetDailyActiveInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetDailyActiveInfoScRsp_proto = out.File
	file_GetDailyActiveInfoScRsp_proto_rawDesc = nil
	file_GetDailyActiveInfoScRsp_proto_goTypes = nil
	file_GetDailyActiveInfoScRsp_proto_depIdxs = nil
}
