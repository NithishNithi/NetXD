// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: helloworld/helloworld.proto

package helloworld

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

type AddCustomerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  string `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *AddCustomerDetails) Reset() {
	*x = AddCustomerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCustomerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCustomerDetails) ProtoMessage() {}

func (x *AddCustomerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCustomerDetails.ProtoReflect.Descriptor instead.
func (*AddCustomerDetails) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *AddCustomerDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddCustomerDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCustomerDetails) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

type CustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[2]
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
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

type Customerlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks1 []*AddCustomerDetails `protobuf:"bytes,1,rep,name=tasks1,proto3" json:"tasks1,omitempty"`
}

func (x *Customerlist) Reset() {
	*x = Customerlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customerlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customerlist) ProtoMessage() {}

func (x *Customerlist) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customerlist.ProtoReflect.Descriptor instead.
func (*Customerlist) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *Customerlist) GetTasks1() []*AddCustomerDetails {
	if x != nil {
		return x.Tasks1
	}
	return nil
}

var File_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x31,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x31, 0x32, 0x70, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x11, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x72, 0x61, 0x6e, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_helloworld_proto_rawDescData = file_helloworld_helloworld_proto_rawDesc
)

func file_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_helloworld_proto_rawDescData)
	})
	return file_helloworld_helloworld_proto_rawDescData
}

var file_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_helloworld_helloworld_proto_goTypes = []interface{}{
	(*AddCustomerDetails)(nil), // 0: AddCustomerDetails
	(*CustomerResponse)(nil),   // 1: CustomerResponse
	(*Empty)(nil),              // 2: Empty
	(*Customerlist)(nil),       // 3: Customerlist
}
var file_helloworld_helloworld_proto_depIdxs = []int32{
	0, // 0: Customerlist.tasks1:type_name -> AddCustomerDetails
	0, // 1: Greeting_Customer.AddCustomer:input_type -> AddCustomerDetails
	2, // 2: Greeting_Customer.GetCustomer:input_type -> Empty
	1, // 3: Greeting_Customer.AddCustomer:output_type -> CustomerResponse
	3, // 4: Greeting_Customer.GetCustomer:output_type -> Customerlist
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_helloworld_helloworld_proto_init() }
func file_helloworld_helloworld_proto_init() {
	if File_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCustomerDetails); i {
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
		file_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResponse); i {
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
		file_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customerlist); i {
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
			RawDescriptor: file_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_helloworld_proto = out.File
	file_helloworld_helloworld_proto_rawDesc = nil
	file_helloworld_helloworld_proto_goTypes = nil
	file_helloworld_helloworld_proto_depIdxs = nil
}
