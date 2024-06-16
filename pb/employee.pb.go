// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: proto/employee.proto

package pb

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

// 列挙型
// 列挙した値のいずれかを持つことができる
// 列挙型の値は、大文字で型は不要
// tag番号は0から始まる また0を含めないといけない
// 0はデフォルト値として使われる
// 0は使わないことが推奨されていて、慣例的にUNKNOWNとして使われる
type Occupation int32

const (
	Occupation_OCCUPATION_UNKNOWN Occupation = 0
	Occupation_ENGINEER           Occupation = 1
	Occupation_DESIGNER           Occupation = 2
	Occupation_MANAGER            Occupation = 3
)

// Enum value maps for Occupation.
var (
	Occupation_name = map[int32]string{
		0: "OCCUPATION_UNKNOWN",
		1: "ENGINEER",
		2: "DESIGNER",
		3: "MANAGER",
	}
	Occupation_value = map[string]int32{
		"OCCUPATION_UNKNOWN": 0,
		"ENGINEER":           1,
		"DESIGNER":           2,
		"MANAGER":            3,
	}
)

func (x Occupation) Enum() *Occupation {
	p := new(Occupation)
	*p = x
	return p
}

func (x Occupation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Occupation) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_employee_proto_enumTypes[0].Descriptor()
}

func (Occupation) Type() protoreflect.EnumType {
	return &file_proto_employee_proto_enumTypes[0]
}

func (x Occupation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Occupation.Descriptor instead.
func (Occupation) EnumDescriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0}
}

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email      string     `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Occupation Occupation `protobuf:"varint,4,opt,name=occupation,proto3,enum=employee.Occupation" json:"occupation,omitempty"`
	// repeatedは配列のようなもの
	PhoneNumbers []string `protobuf:"bytes,5,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	// map はrepeatedを使えない
	// map<string, Project> projects = 6;
	Projects map[string]*Company_Project `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// oneofはどれか一つだけを持つことができる
	//
	// Types that are assignable to Profile:
	//
	//	*Employee_Text
	//	*Employee_Video
	Profile  isEmployee_Profile `protobuf_oneof:"profile"`
	Birthday *Date              `protobuf:"bytes,9,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Employee) GetOccupation() Occupation {
	if x != nil {
		return x.Occupation
	}
	return Occupation_OCCUPATION_UNKNOWN
}

func (x *Employee) GetPhoneNumbers() []string {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *Employee) GetProjects() map[string]*Company_Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (m *Employee) GetProfile() isEmployee_Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (x *Employee) GetText() string {
	if x, ok := x.GetProfile().(*Employee_Text); ok {
		return x.Text
	}
	return ""
}

func (x *Employee) GetVideo() *Video {
	if x, ok := x.GetProfile().(*Employee_Video); ok {
		return x.Video
	}
	return nil
}

func (x *Employee) GetBirthday() *Date {
	if x != nil {
		return x.Birthday
	}
	return nil
}

type isEmployee_Profile interface {
	isEmployee_Profile()
}

type Employee_Text struct {
	Text string `protobuf:"bytes,7,opt,name=text,proto3,oneof"`
}

type Employee_Video struct {
	Video *Video `protobuf:"bytes,8,opt,name=video,proto3,oneof"`
}

func (*Employee_Text) isEmployee_Profile() {}

func (*Employee_Video) isEmployee_Profile() {}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{1}
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{2}
}

// messageのネスト
// 名前の衝突を避けることができる
// Company.Projectという名前でアクセスする
// スコープを限定することができる
type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{3}
}

type Company_Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Company_Project) Reset() {
	*x = Company_Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company_Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company_Project) ProtoMessage() {}

func (x *Company_Project) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company_Project.ProtoReflect.Descriptor instead.
func (*Company_Project) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{3, 0}
}

var File_proto_employee_proto protoreflect.FileDescriptor

var file_proto_employee_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa7, 0x03, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x34, 0x0a, 0x0a, 0x6f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x26, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x1a, 0x56, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x09, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x22, 0x14, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x09, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2a, 0x4d, 0x0a, 0x0a, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x43, 0x43, 0x55, 0x50, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45,
	0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x4e, 0x41,
	0x47, 0x45, 0x52, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_employee_proto_rawDescOnce sync.Once
	file_proto_employee_proto_rawDescData = file_proto_employee_proto_rawDesc
)

func file_proto_employee_proto_rawDescGZIP() []byte {
	file_proto_employee_proto_rawDescOnce.Do(func() {
		file_proto_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_employee_proto_rawDescData)
	})
	return file_proto_employee_proto_rawDescData
}

var file_proto_employee_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_employee_proto_goTypes = []interface{}{
	(Occupation)(0),         // 0: employee.Occupation
	(*Employee)(nil),        // 1: employee.Employee
	(*Project)(nil),         // 2: employee.Project
	(*Video)(nil),           // 3: employee.Video
	(*Company)(nil),         // 4: employee.Company
	nil,                     // 5: employee.Employee.ProjectsEntry
	(*Company_Project)(nil), // 6: employee.Company.Project
	(*Date)(nil),            // 7: date.Date
}
var file_proto_employee_proto_depIdxs = []int32{
	0, // 0: employee.Employee.occupation:type_name -> employee.Occupation
	5, // 1: employee.Employee.projects:type_name -> employee.Employee.ProjectsEntry
	3, // 2: employee.Employee.video:type_name -> employee.Video
	7, // 3: employee.Employee.birthday:type_name -> date.Date
	6, // 4: employee.Employee.ProjectsEntry.value:type_name -> employee.Company.Project
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_employee_proto_init() }
func file_proto_employee_proto_init() {
	if File_proto_employee_proto != nil {
		return
	}
	file_proto_date_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_proto_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_proto_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_employee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_proto_employee_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company_Project); i {
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
	file_proto_employee_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Employee_Text)(nil),
		(*Employee_Video)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_employee_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_employee_proto_goTypes,
		DependencyIndexes: file_proto_employee_proto_depIdxs,
		EnumInfos:         file_proto_employee_proto_enumTypes,
		MessageInfos:      file_proto_employee_proto_msgTypes,
	}.Build()
	File_proto_employee_proto = out.File
	file_proto_employee_proto_rawDesc = nil
	file_proto_employee_proto_goTypes = nil
	file_proto_employee_proto_depIdxs = nil
}
