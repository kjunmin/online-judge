// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: common/api/judge.proto

package api

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

type ProblemsListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemID string   `protobuf:"bytes,1,opt,name=problemID,proto3" json:"problemID,omitempty"`
	Title     string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tags      []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ProblemsListItem) Reset() {
	*x = ProblemsListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_judge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemsListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemsListItem) ProtoMessage() {}

func (x *ProblemsListItem) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_judge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemsListItem.ProtoReflect.Descriptor instead.
func (*ProblemsListItem) Descriptor() ([]byte, []int) {
	return file_common_api_judge_proto_rawDescGZIP(), []int{0}
}

func (x *ProblemsListItem) GetProblemID() string {
	if x != nil {
		return x.ProblemID
	}
	return ""
}

func (x *ProblemsListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProblemsListItem) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemID   *string  `protobuf:"bytes,1,opt,name=problemID,proto3,oneof" json:"problemID,omitempty"`
	Tags        []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_judge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_judge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_common_api_judge_proto_rawDescGZIP(), []int{1}
}

func (x *Problem) GetProblemID() string {
	if x != nil && x.ProblemID != nil {
		return *x.ProblemID
	}
	return ""
}

func (x *Problem) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Problem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Problem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetProblemByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problemId,proto3" json:"problemId,omitempty"`
}

func (x *GetProblemByIdRequest) Reset() {
	*x = GetProblemByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_judge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemByIdRequest) ProtoMessage() {}

func (x *GetProblemByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_judge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProblemByIdRequest) Descriptor() ([]byte, []int) {
	return file_common_api_judge_proto_rawDescGZIP(), []int{2}
}

func (x *GetProblemByIdRequest) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

type GetProblemsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProblemsListRequest) Reset() {
	*x = GetProblemsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_judge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemsListRequest) ProtoMessage() {}

func (x *GetProblemsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_judge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemsListRequest.ProtoReflect.Descriptor instead.
func (*GetProblemsListRequest) Descriptor() ([]byte, []int) {
	return file_common_api_judge_proto_rawDescGZIP(), []int{3}
}

type GetProblemsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemsList []*ProblemsListItem `protobuf:"bytes,1,rep,name=problemsList,proto3" json:"problemsList,omitempty"`
}

func (x *GetProblemsListResponse) Reset() {
	*x = GetProblemsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_judge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemsListResponse) ProtoMessage() {}

func (x *GetProblemsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_judge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemsListResponse.ProtoReflect.Descriptor instead.
func (*GetProblemsListResponse) Descriptor() ([]byte, []int) {
	return file_common_api_judge_proto_rawDescGZIP(), []int{4}
}

func (x *GetProblemsListResponse) GetProblemsList() []*ProblemsListItem {
	if x != nil {
		return x.ProblemsList
	}
	return nil
}

type CreateProblemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Problem *Problem `protobuf:"bytes,1,opt,name=Problem,proto3" json:"Problem,omitempty"`
}

func (x *CreateProblemRequest) Reset() {
	*x = CreateProblemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_judge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProblemRequest) ProtoMessage() {}

func (x *CreateProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_judge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProblemRequest.ProtoReflect.Descriptor instead.
func (*CreateProblemRequest) Descriptor() ([]byte, []int) {
	return file_common_api_judge_proto_rawDescGZIP(), []int{5}
}

func (x *CreateProblemRequest) GetProblem() *Problem {
	if x != nil {
		return x.Problem
	}
	return nil
}

var File_common_api_judge_proto protoreflect.FileDescriptor

var file_common_api_judge_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x5a, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x49, 0x44, 0x22, 0x35, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x52, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x32, 0xd5, 0x01, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6a, 0x75, 0x6e, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_api_judge_proto_rawDescOnce sync.Once
	file_common_api_judge_proto_rawDescData = file_common_api_judge_proto_rawDesc
)

func file_common_api_judge_proto_rawDescGZIP() []byte {
	file_common_api_judge_proto_rawDescOnce.Do(func() {
		file_common_api_judge_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_api_judge_proto_rawDescData)
	})
	return file_common_api_judge_proto_rawDescData
}

var file_common_api_judge_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_api_judge_proto_goTypes = []any{
	(*ProblemsListItem)(nil),        // 0: api.ProblemsListItem
	(*Problem)(nil),                 // 1: api.Problem
	(*GetProblemByIdRequest)(nil),   // 2: api.GetProblemByIdRequest
	(*GetProblemsListRequest)(nil),  // 3: api.GetProblemsListRequest
	(*GetProblemsListResponse)(nil), // 4: api.GetProblemsListResponse
	(*CreateProblemRequest)(nil),    // 5: api.CreateProblemRequest
}
var file_common_api_judge_proto_depIdxs = []int32{
	0, // 0: api.GetProblemsListResponse.problemsList:type_name -> api.ProblemsListItem
	1, // 1: api.CreateProblemRequest.Problem:type_name -> api.Problem
	5, // 2: api.ProblemsService.CreateProblem:input_type -> api.CreateProblemRequest
	3, // 3: api.ProblemsService.GetProblemsList:input_type -> api.GetProblemsListRequest
	2, // 4: api.ProblemsService.GetProblemById:input_type -> api.GetProblemByIdRequest
	1, // 5: api.ProblemsService.CreateProblem:output_type -> api.Problem
	4, // 6: api.ProblemsService.GetProblemsList:output_type -> api.GetProblemsListResponse
	1, // 7: api.ProblemsService.GetProblemById:output_type -> api.Problem
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_api_judge_proto_init() }
func file_common_api_judge_proto_init() {
	if File_common_api_judge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_api_judge_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProblemsListItem); i {
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
		file_common_api_judge_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Problem); i {
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
		file_common_api_judge_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetProblemByIdRequest); i {
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
		file_common_api_judge_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetProblemsListRequest); i {
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
		file_common_api_judge_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetProblemsListResponse); i {
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
		file_common_api_judge_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProblemRequest); i {
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
	file_common_api_judge_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_api_judge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_api_judge_proto_goTypes,
		DependencyIndexes: file_common_api_judge_proto_depIdxs,
		MessageInfos:      file_common_api_judge_proto_msgTypes,
	}.Build()
	File_common_api_judge_proto = out.File
	file_common_api_judge_proto_rawDesc = nil
	file_common_api_judge_proto_goTypes = nil
	file_common_api_judge_proto_depIdxs = nil
}
