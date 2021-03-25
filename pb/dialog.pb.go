// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: dialog.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RecognizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Audio   []byte   `protobuf:"bytes,2,opt,name=audio,proto3" json:"audio,omitempty"`
	Context *Context `protobuf:"bytes,3,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *RecognizeRequest) Reset() {
	*x = RecognizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeRequest) ProtoMessage() {}

func (x *RecognizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dialog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeRequest.ProtoReflect.Descriptor instead.
func (*RecognizeRequest) Descriptor() ([]byte, []int) {
	return file_dialog_proto_rawDescGZIP(), []int{0}
}

func (x *RecognizeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *RecognizeRequest) GetAudio() []byte {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *RecognizeRequest) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

type RecognizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vendor                   string           `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Skill                    string           `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Command                  string           `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Query                    string           `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Text                     string           `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	Confidence               float32          `protobuf:"fixed32,6,opt,name=confidence,proto3" json:"confidence,omitempty"`
	AllRequiredParamsPresent bool             `protobuf:"varint,7,opt,name=all_required_params_present,json=allRequiredParamsPresent,proto3" json:"all_required_params_present,omitempty"`
	Parameters               *_struct.Struct  `protobuf:"bytes,8,opt,name=parameters,proto3" json:"parameters,omitempty"`
	OutputContexts           []*OutputContext `protobuf:"bytes,9,rep,name=output_contexts,json=outputContexts,proto3" json:"output_contexts,omitempty"`
	Context                  *Context         `protobuf:"bytes,10,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *RecognizeResponse) Reset() {
	*x = RecognizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeResponse) ProtoMessage() {}

func (x *RecognizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dialog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeResponse.ProtoReflect.Descriptor instead.
func (*RecognizeResponse) Descriptor() ([]byte, []int) {
	return file_dialog_proto_rawDescGZIP(), []int{1}
}

func (x *RecognizeResponse) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *RecognizeResponse) GetSkill() string {
	if x != nil {
		return x.Skill
	}
	return ""
}

func (x *RecognizeResponse) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *RecognizeResponse) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *RecognizeResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *RecognizeResponse) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *RecognizeResponse) GetAllRequiredParamsPresent() bool {
	if x != nil {
		return x.AllRequiredParamsPresent
	}
	return false
}

func (x *RecognizeResponse) GetParameters() *_struct.Struct {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *RecognizeResponse) GetOutputContexts() []*OutputContext {
	if x != nil {
		return x.OutputContexts
	}
	return nil
}

func (x *RecognizeResponse) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

type OutputContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lifespan int32  `protobuf:"varint,2,opt,name=lifespan,proto3" json:"lifespan,omitempty"`
}

func (x *OutputContext) Reset() {
	*x = OutputContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputContext) ProtoMessage() {}

func (x *OutputContext) ProtoReflect() protoreflect.Message {
	mi := &file_dialog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputContext.ProtoReflect.Descriptor instead.
func (*OutputContext) Descriptor() ([]byte, []int) {
	return file_dialog_proto_rawDescGZIP(), []int{2}
}

func (x *OutputContext) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OutputContext) GetLifespan() int32 {
	if x != nil {
		return x.Lifespan
	}
	return 0
}

type DetectedIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Confidence string `protobuf:"bytes,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *DetectedIdentity) Reset() {
	*x = DetectedIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectedIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectedIdentity) ProtoMessage() {}

func (x *DetectedIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_dialog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectedIdentity.ProtoReflect.Descriptor instead.
func (*DetectedIdentity) Descriptor() ([]byte, []int) {
	return file_dialog_proto_rawDescGZIP(), []int{3}
}

func (x *DetectedIdentity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DetectedIdentity) GetConfidence() string {
	if x != nil {
		return x.Confidence
	}
	return ""
}

var File_dialog_proto protoreflect.FileDescriptor

var file_dialog_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x89, 0x03,
	0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a,
	0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x2a, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3f, 0x0a, 0x0d, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x32, 0x4c, 0x0a, 0x06, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x42, 0x0a, 0x09,
	0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x07, 0x5a, 0x05, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dialog_proto_rawDescOnce sync.Once
	file_dialog_proto_rawDescData = file_dialog_proto_rawDesc
)

func file_dialog_proto_rawDescGZIP() []byte {
	file_dialog_proto_rawDescOnce.Do(func() {
		file_dialog_proto_rawDescData = protoimpl.X.CompressGZIP(file_dialog_proto_rawDescData)
	})
	return file_dialog_proto_rawDescData
}

var file_dialog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dialog_proto_goTypes = []interface{}{
	(*RecognizeRequest)(nil),  // 0: dialog.RecognizeRequest
	(*RecognizeResponse)(nil), // 1: dialog.RecognizeResponse
	(*OutputContext)(nil),     // 2: dialog.OutputContext
	(*DetectedIdentity)(nil),  // 3: dialog.DetectedIdentity
	(*Context)(nil),           // 4: context.Context
	(*_struct.Struct)(nil),    // 5: google.protobuf.Struct
}
var file_dialog_proto_depIdxs = []int32{
	4, // 0: dialog.RecognizeRequest.context:type_name -> context.Context
	5, // 1: dialog.RecognizeResponse.parameters:type_name -> google.protobuf.Struct
	2, // 2: dialog.RecognizeResponse.output_contexts:type_name -> dialog.OutputContext
	4, // 3: dialog.RecognizeResponse.context:type_name -> context.Context
	0, // 4: dialog.Dialog.Recognize:input_type -> dialog.RecognizeRequest
	1, // 5: dialog.Dialog.Recognize:output_type -> dialog.RecognizeResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dialog_proto_init() }
func file_dialog_proto_init() {
	if File_dialog_proto != nil {
		return
	}
	file_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dialog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeRequest); i {
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
		file_dialog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeResponse); i {
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
		file_dialog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputContext); i {
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
		file_dialog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectedIdentity); i {
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
			RawDescriptor: file_dialog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dialog_proto_goTypes,
		DependencyIndexes: file_dialog_proto_depIdxs,
		MessageInfos:      file_dialog_proto_msgTypes,
	}.Build()
	File_dialog_proto = out.File
	file_dialog_proto_rawDesc = nil
	file_dialog_proto_goTypes = nil
	file_dialog_proto_depIdxs = nil
}
