// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: bin/spike.proto

package bin

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

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_spike_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_bin_spike_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_bin_spike_proto_rawDescGZIP(), []int{0}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Success) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Success) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Offset  int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_spike_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_bin_spike_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_bin_spike_proto_rawDescGZIP(), []int{1}
}

func (x *Topic) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Topic) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Topic) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	GroupId string            `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Offset  int64             `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Value   []byte            `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Header  map[string][]byte `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_spike_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_bin_spike_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_bin_spike_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Message) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Message) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Message) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Message) GetHeader() map[string][]byte {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_bin_spike_proto protoreflect.FileDescriptor

var file_bin_spike_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x69, 0x6e, 0x2f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x62, 0x69, 0x6e, 0x22, 0x51, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x05, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0xac, 0x01, 0x0a, 0x05, 0x53, 0x70, 0x69, 0x6b, 0x65, 0x12, 0x2a, 0x0a,
	0x0c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0a, 0x2e,
	0x62, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x0c, 0x2e, 0x62, 0x69, 0x6e, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x27, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0a, 0x2e, 0x62, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x1a, 0x0c, 0x2e, 0x62, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x30, 0x01, 0x12, 0x27, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x0a, 0x2e, 0x62, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x0c, 0x2e,
	0x62, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x0c, 0x2e, 0x62, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x62, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x2f, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x70,
	0x69, 0x6b, 0x65, 0x2e, 0x62, 0x69, 0x6e, 0x42, 0x0a, 0x53, 0x70, 0x69, 0x6b, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0c, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x2e, 0x69, 0x6f, 0x2f,
	0x62, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bin_spike_proto_rawDescOnce sync.Once
	file_bin_spike_proto_rawDescData = file_bin_spike_proto_rawDesc
)

func file_bin_spike_proto_rawDescGZIP() []byte {
	file_bin_spike_proto_rawDescOnce.Do(func() {
		file_bin_spike_proto_rawDescData = protoimpl.X.CompressGZIP(file_bin_spike_proto_rawDescData)
	})
	return file_bin_spike_proto_rawDescData
}

var file_bin_spike_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bin_spike_proto_goTypes = []interface{}{
	(*Success)(nil), // 0: bin.Success
	(*Topic)(nil),   // 1: bin.Topic
	(*Message)(nil), // 2: bin.Message
	nil,             // 3: bin.Message.HeaderEntry
}
var file_bin_spike_proto_depIdxs = []int32{
	3, // 0: bin.Message.header:type_name -> bin.Message.HeaderEntry
	1, // 1: bin.Spike.MonitorEvent:input_type -> bin.Topic
	1, // 2: bin.Spike.Subscribe:input_type -> bin.Topic
	1, // 3: bin.Spike.Unsubscribe:input_type -> bin.Topic
	2, // 4: bin.Spike.Publish:input_type -> bin.Message
	2, // 5: bin.Spike.MonitorEvent:output_type -> bin.Message
	2, // 6: bin.Spike.Subscribe:output_type -> bin.Message
	0, // 7: bin.Spike.Unsubscribe:output_type -> bin.Success
	0, // 8: bin.Spike.Publish:output_type -> bin.Success
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bin_spike_proto_init() }
func file_bin_spike_proto_init() {
	if File_bin_spike_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bin_spike_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_bin_spike_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_bin_spike_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_bin_spike_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bin_spike_proto_goTypes,
		DependencyIndexes: file_bin_spike_proto_depIdxs,
		MessageInfos:      file_bin_spike_proto_msgTypes,
	}.Build()
	File_bin_spike_proto = out.File
	file_bin_spike_proto_rawDesc = nil
	file_bin_spike_proto_goTypes = nil
	file_bin_spike_proto_depIdxs = nil
}