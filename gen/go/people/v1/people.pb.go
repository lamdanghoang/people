// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: people/v1/people.proto

package people

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OfficeStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	WorkingDays int32  `protobuf:"varint,4,opt,name=working_days,json=workingDays,proto3" json:"working_days,omitempty"`
}

func (x *OfficeStaff) Reset() {
	*x = OfficeStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_v1_people_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeStaff) ProtoMessage() {}

func (x *OfficeStaff) ProtoReflect() protoreflect.Message {
	mi := &file_people_v1_people_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeStaff.ProtoReflect.Descriptor instead.
func (*OfficeStaff) Descriptor() ([]byte, []int) {
	return file_people_v1_people_proto_rawDescGZIP(), []int{0}
}

func (x *OfficeStaff) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OfficeStaff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfficeStaff) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OfficeStaff) GetWorkingDays() int32 {
	if x != nil {
		return x.WorkingDays
	}
	return 0
}

type SaleStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	ProductQuantity int32  `protobuf:"varint,4,opt,name=product_quantity,json=productQuantity,proto3" json:"product_quantity,omitempty"`
}

func (x *SaleStaff) Reset() {
	*x = SaleStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_v1_people_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleStaff) ProtoMessage() {}

func (x *SaleStaff) ProtoReflect() protoreflect.Message {
	mi := &file_people_v1_people_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleStaff.ProtoReflect.Descriptor instead.
func (*SaleStaff) Descriptor() ([]byte, []int) {
	return file_people_v1_people_proto_rawDescGZIP(), []int{1}
}

func (x *SaleStaff) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaleStaff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaleStaff) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SaleStaff) GetProductQuantity() int32 {
	if x != nil {
		return x.ProductQuantity
	}
	return 0
}

type IsStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staff *anypb.Any `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
	X     int64      `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *IsStaff) Reset() {
	*x = IsStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_v1_people_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsStaff) ProtoMessage() {}

func (x *IsStaff) ProtoReflect() protoreflect.Message {
	mi := &file_people_v1_people_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsStaff.ProtoReflect.Descriptor instead.
func (*IsStaff) Descriptor() ([]byte, []int) {
	return file_people_v1_people_proto_rawDescGZIP(), []int{2}
}

func (x *IsStaff) GetStaff() *anypb.Any {
	if x != nil {
		return x.Staff
	}
	return nil
}

func (x *IsStaff) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

var File_people_v1_people_proto protoreflect.FileDescriptor

var file_people_v1_people_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a,
	0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x22, 0x70, 0x0a, 0x09, 0x53, 0x61,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x43, 0x0a, 0x07,
	0x49, 0x73, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x78, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x61, 0x6d, 0x64, 0x61, 0x6e, 0x67, 0x68, 0x6f, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x6f,
	0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_people_v1_people_proto_rawDescOnce sync.Once
	file_people_v1_people_proto_rawDescData = file_people_v1_people_proto_rawDesc
)

func file_people_v1_people_proto_rawDescGZIP() []byte {
	file_people_v1_people_proto_rawDescOnce.Do(func() {
		file_people_v1_people_proto_rawDescData = protoimpl.X.CompressGZIP(file_people_v1_people_proto_rawDescData)
	})
	return file_people_v1_people_proto_rawDescData
}

var file_people_v1_people_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_people_v1_people_proto_goTypes = []interface{}{
	(*OfficeStaff)(nil), // 0: people.v1.OfficeStaff
	(*SaleStaff)(nil),   // 1: people.v1.SaleStaff
	(*IsStaff)(nil),     // 2: people.v1.IsStaff
	(*anypb.Any)(nil),   // 3: google.protobuf.Any
}
var file_people_v1_people_proto_depIdxs = []int32{
	3, // 0: people.v1.IsStaff.staff:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_people_v1_people_proto_init() }
func file_people_v1_people_proto_init() {
	if File_people_v1_people_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_people_v1_people_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeStaff); i {
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
		file_people_v1_people_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleStaff); i {
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
		file_people_v1_people_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsStaff); i {
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
			RawDescriptor: file_people_v1_people_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_people_v1_people_proto_goTypes,
		DependencyIndexes: file_people_v1_people_proto_depIdxs,
		MessageInfos:      file_people_v1_people_proto_msgTypes,
	}.Build()
	File_people_v1_people_proto = out.File
	file_people_v1_people_proto_rawDesc = nil
	file_people_v1_people_proto_goTypes = nil
	file_people_v1_people_proto_depIdxs = nil
}
