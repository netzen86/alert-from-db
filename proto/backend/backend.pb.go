// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: proto/backend/backend.proto

package backend

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

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Director string `protobuf:"bytes,3,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	mi := &file_proto_backend_backend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{0}
}

func (x *Film) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Film) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Film) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

type AddFilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Film *Film `protobuf:"bytes,1,opt,name=film,proto3" json:"film,omitempty"`
}

func (x *AddFilmRequest) Reset() {
	*x = AddFilmRequest{}
	mi := &file_proto_backend_backend_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddFilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFilmRequest) ProtoMessage() {}

func (x *AddFilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFilmRequest.ProtoReflect.Descriptor instead.
func (*AddFilmRequest) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{1}
}

func (x *AddFilmRequest) GetFilm() *Film {
	if x != nil {
		return x.Film
	}
	return nil
}

type AddFilmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films []*Film `protobuf:"bytes,1,rep,name=films,proto3" json:"films,omitempty"`
	Error string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddFilmResponse) Reset() {
	*x = AddFilmResponse{}
	mi := &file_proto_backend_backend_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddFilmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFilmResponse) ProtoMessage() {}

func (x *AddFilmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFilmResponse.ProtoReflect.Descriptor instead.
func (*AddFilmResponse) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{2}
}

func (x *AddFilmResponse) GetFilms() []*Film {
	if x != nil {
		return x.Films
	}
	return nil
}

func (x *AddFilmResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DelFilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelFilmRequest) Reset() {
	*x = DelFilmRequest{}
	mi := &file_proto_backend_backend_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelFilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelFilmRequest) ProtoMessage() {}

func (x *DelFilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelFilmRequest.ProtoReflect.Descriptor instead.
func (*DelFilmRequest) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{3}
}

func (x *DelFilmRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelFilmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films []*Film `protobuf:"bytes,1,rep,name=films,proto3" json:"films,omitempty"`
	Error string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DelFilmResponse) Reset() {
	*x = DelFilmResponse{}
	mi := &file_proto_backend_backend_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelFilmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelFilmResponse) ProtoMessage() {}

func (x *DelFilmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelFilmResponse.ProtoReflect.Descriptor instead.
func (*DelFilmResponse) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{4}
}

func (x *DelFilmResponse) GetFilms() []*Film {
	if x != nil {
		return x.Films
	}
	return nil
}

func (x *DelFilmResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetFilmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFilmsRequest) Reset() {
	*x = GetFilmsRequest{}
	mi := &file_proto_backend_backend_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFilmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilmsRequest) ProtoMessage() {}

func (x *GetFilmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilmsRequest.ProtoReflect.Descriptor instead.
func (*GetFilmsRequest) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{5}
}

type GetFilmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films []*Film `protobuf:"bytes,1,rep,name=films,proto3" json:"films,omitempty"`
	Error string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetFilmsResponse) Reset() {
	*x = GetFilmsResponse{}
	mi := &file_proto_backend_backend_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFilmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilmsResponse) ProtoMessage() {}

func (x *GetFilmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backend_backend_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilmsResponse.ProtoReflect.Descriptor instead.
func (*GetFilmsResponse) Descriptor() ([]byte, []int) {
	return file_proto_backend_backend_proto_rawDescGZIP(), []int{6}
}

func (x *GetFilmsResponse) GetFilms() []*Film {
	if x != nil {
		return x.Films
	}
	return nil
}

func (x *GetFilmsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_backend_backend_proto protoreflect.FileDescriptor

var file_proto_backend_backend_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22, 0x48, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x22, 0x33, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52,
	0x04, 0x66, 0x69, 0x6c, 0x6d, 0x22, 0x4c, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc6, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x17, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x44, 0x65,
	0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d,
	0x5a, 0x1b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2d, 0x66, 0x72, 0x6f, 0x6d, 0x2d, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_backend_backend_proto_rawDescOnce sync.Once
	file_proto_backend_backend_proto_rawDescData = file_proto_backend_backend_proto_rawDesc
)

func file_proto_backend_backend_proto_rawDescGZIP() []byte {
	file_proto_backend_backend_proto_rawDescOnce.Do(func() {
		file_proto_backend_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_backend_backend_proto_rawDescData)
	})
	return file_proto_backend_backend_proto_rawDescData
}

var file_proto_backend_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_backend_backend_proto_goTypes = []any{
	(*Film)(nil),             // 0: backend.Film
	(*AddFilmRequest)(nil),   // 1: backend.AddFilmRequest
	(*AddFilmResponse)(nil),  // 2: backend.AddFilmResponse
	(*DelFilmRequest)(nil),   // 3: backend.DelFilmRequest
	(*DelFilmResponse)(nil),  // 4: backend.DelFilmResponse
	(*GetFilmsRequest)(nil),  // 5: backend.GetFilmsRequest
	(*GetFilmsResponse)(nil), // 6: backend.GetFilmsResponse
}
var file_proto_backend_backend_proto_depIdxs = []int32{
	0, // 0: backend.AddFilmRequest.film:type_name -> backend.Film
	0, // 1: backend.AddFilmResponse.films:type_name -> backend.Film
	0, // 2: backend.DelFilmResponse.films:type_name -> backend.Film
	0, // 3: backend.GetFilmsResponse.films:type_name -> backend.Film
	1, // 4: backend.Backend.AddFilm:input_type -> backend.AddFilmRequest
	3, // 5: backend.Backend.DelFilm:input_type -> backend.DelFilmRequest
	5, // 6: backend.Backend.GetFilms:input_type -> backend.GetFilmsRequest
	2, // 7: backend.Backend.AddFilm:output_type -> backend.AddFilmResponse
	4, // 8: backend.Backend.DelFilm:output_type -> backend.DelFilmResponse
	6, // 9: backend.Backend.GetFilms:output_type -> backend.GetFilmsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_backend_backend_proto_init() }
func file_proto_backend_backend_proto_init() {
	if File_proto_backend_backend_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_backend_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_backend_backend_proto_goTypes,
		DependencyIndexes: file_proto_backend_backend_proto_depIdxs,
		MessageInfos:      file_proto_backend_backend_proto_msgTypes,
	}.Build()
	File_proto_backend_backend_proto = out.File
	file_proto_backend_backend_proto_rawDesc = nil
	file_proto_backend_backend_proto_goTypes = nil
	file_proto_backend_backend_proto_depIdxs = nil
}
