// Copyright 2021 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.1
// source: bp2build_metrics.proto

package bp2build_metrics_proto

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

type Bp2BuildMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of Soong modules converted to generated targets
	GeneratedModuleCount uint64 `protobuf:"varint,1,opt,name=generatedModuleCount,proto3" json:"generatedModuleCount,omitempty"`
	// Total number of Soong modules converted to handcrafted targets
	HandCraftedModuleCount uint64 `protobuf:"varint,2,opt,name=handCraftedModuleCount,proto3" json:"handCraftedModuleCount,omitempty"`
	// Total number of unconverted Soong modules
	UnconvertedModuleCount uint64 `protobuf:"varint,3,opt,name=unconvertedModuleCount,proto3" json:"unconvertedModuleCount,omitempty"`
	// Counts of generated Bazel targets per Bazel rule class
	RuleClassCount map[string]uint64 `protobuf:"bytes,4,rep,name=ruleClassCount,proto3" json:"ruleClassCount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of converted modules
	ConvertedModules []string `protobuf:"bytes,5,rep,name=convertedModules,proto3" json:"convertedModules,omitempty"`
}

func (x *Bp2BuildMetrics) Reset() {
	*x = Bp2BuildMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bp2build_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bp2BuildMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bp2BuildMetrics) ProtoMessage() {}

func (x *Bp2BuildMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_bp2build_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bp2BuildMetrics.ProtoReflect.Descriptor instead.
func (*Bp2BuildMetrics) Descriptor() ([]byte, []int) {
	return file_bp2build_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Bp2BuildMetrics) GetGeneratedModuleCount() uint64 {
	if x != nil {
		return x.GeneratedModuleCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetHandCraftedModuleCount() uint64 {
	if x != nil {
		return x.HandCraftedModuleCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetUnconvertedModuleCount() uint64 {
	if x != nil {
		return x.UnconvertedModuleCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetRuleClassCount() map[string]uint64 {
	if x != nil {
		return x.RuleClassCount
	}
	return nil
}

func (x *Bp2BuildMetrics) GetConvertedModules() []string {
	if x != nil {
		return x.ConvertedModules
	}
	return nil
}

var File_bp2build_metrics_proto protoreflect.FileDescriptor

var file_bp2build_metrics_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x8f, 0x03, 0x0a, 0x0f, 0x42, 0x70, 0x32, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36,
	0x0a, 0x16, 0x68, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16,
	0x68, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x69,
	0x0a, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x70, 0x32, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x2f, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x2f, 0x75, 0x69, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bp2build_metrics_proto_rawDescOnce sync.Once
	file_bp2build_metrics_proto_rawDescData = file_bp2build_metrics_proto_rawDesc
)

func file_bp2build_metrics_proto_rawDescGZIP() []byte {
	file_bp2build_metrics_proto_rawDescOnce.Do(func() {
		file_bp2build_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_bp2build_metrics_proto_rawDescData)
	})
	return file_bp2build_metrics_proto_rawDescData
}

var file_bp2build_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bp2build_metrics_proto_goTypes = []interface{}{
	(*Bp2BuildMetrics)(nil), // 0: soong_build_bp2build_metrics.Bp2BuildMetrics
	nil,                     // 1: soong_build_bp2build_metrics.Bp2BuildMetrics.RuleClassCountEntry
}
var file_bp2build_metrics_proto_depIdxs = []int32{
	1, // 0: soong_build_bp2build_metrics.Bp2BuildMetrics.ruleClassCount:type_name -> soong_build_bp2build_metrics.Bp2BuildMetrics.RuleClassCountEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bp2build_metrics_proto_init() }
func file_bp2build_metrics_proto_init() {
	if File_bp2build_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bp2build_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bp2BuildMetrics); i {
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
			RawDescriptor: file_bp2build_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bp2build_metrics_proto_goTypes,
		DependencyIndexes: file_bp2build_metrics_proto_depIdxs,
		MessageInfos:      file_bp2build_metrics_proto_msgTypes,
	}.Build()
	File_bp2build_metrics_proto = out.File
	file_bp2build_metrics_proto_rawDesc = nil
	file_bp2build_metrics_proto_goTypes = nil
	file_bp2build_metrics_proto_depIdxs = nil
}