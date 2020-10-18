// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: search.proto

package turkionary_proto

import (
	proto "github.com/golang/protobuf/proto"
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

type LangCode int32

const (
	LangCode_EN LangCode = 0
	LangCode_TR LangCode = 1
)

// Enum value maps for LangCode.
var (
	LangCode_name = map[int32]string{
		0: "EN",
		1: "TR",
	}
	LangCode_value = map[string]int32{
		"EN": 0,
		"TR": 1,
	}
)

func (x LangCode) Enum() *LangCode {
	p := new(LangCode)
	*p = x
	return p
}

func (x LangCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LangCode) Descriptor() protoreflect.EnumDescriptor {
	return file_search_proto_enumTypes[0].Descriptor()
}

func (LangCode) Type() protoreflect.EnumType {
	return &file_search_proto_enumTypes[0]
}

func (x LangCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LangCode.Descriptor instead.
func (LangCode) EnumDescriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

type WordType int32

const (
	WordType_NOUN      WordType = 0
	WordType_VERB      WordType = 1
	WordType_ADJECTIVE WordType = 2
)

// Enum value maps for WordType.
var (
	WordType_name = map[int32]string{
		0: "NOUN",
		1: "VERB",
		2: "ADJECTIVE",
	}
	WordType_value = map[string]int32{
		"NOUN":      0,
		"VERB":      1,
		"ADJECTIVE": 2,
	}
)

func (x WordType) Enum() *WordType {
	p := new(WordType)
	*p = x
	return p
}

func (x WordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordType) Descriptor() protoreflect.EnumDescriptor {
	return file_search_proto_enumTypes[1].Descriptor()
}

func (WordType) Type() protoreflect.EnumType {
	return &file_search_proto_enumTypes[1]
}

func (x WordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordType.Descriptor instead.
func (WordType) EnumDescriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Source LangCode `protobuf:"varint,2,opt,name=source,proto3,enum=turkionary_proto.LangCode" json:"source,omitempty"`
	Target LangCode `protobuf:"varint,3,opt,name=target,proto3,enum=turkionary_proto.LangCode" json:"target,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetSource() LangCode {
	if x != nil {
		return x.Source
	}
	return LangCode_EN
}

func (x *SearchRequest) GetTarget() LangCode {
	if x != nil {
		return x.Target
	}
	return LangCode_EN
}

type Definition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang    LangCode `protobuf:"varint,1,opt,name=lang,proto3,enum=turkionary_proto.LangCode" json:"lang,omitempty"`
	Text    string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Context string   `protobuf:"bytes,3,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *Definition) Reset() {
	*x = Definition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Definition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Definition) ProtoMessage() {}

func (x *Definition) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Definition.ProtoReflect.Descriptor instead.
func (*Definition) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *Definition) GetLang() LangCode {
	if x != nil {
		return x.Lang
	}
	return LangCode_EN
}

func (x *Definition) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Definition) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

type Phrase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang LangCode `protobuf:"varint,1,opt,name=lang,proto3,enum=turkionary_proto.LangCode" json:"lang,omitempty"`
	Text string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Phrase) Reset() {
	*x = Phrase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phrase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phrase) ProtoMessage() {}

func (x *Phrase) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phrase.ProtoReflect.Descriptor instead.
func (*Phrase) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *Phrase) GetLang() LangCode {
	if x != nil {
		return x.Lang
	}
	return LangCode_EN
}

func (x *Phrase) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Texts []*Phrase `protobuf:"bytes,1,rep,name=texts,proto3" json:"texts,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *Example) GetTexts() []*Phrase {
	if x != nil {
		return x.Texts
	}
	return nil
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        string     `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	Definitions []string   `protobuf:"bytes,2,rep,name=definitions,proto3" json:"definitions,omitempty"`
	Sayings     []*Phrase  `protobuf:"bytes,3,rep,name=sayings,proto3" json:"sayings,omitempty"`
	Examples    []*Example `protobuf:"bytes,4,rep,name=examples,proto3" json:"examples,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResult) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *SearchResult) GetDefinitions() []string {
	if x != nil {
		return x.Definitions
	}
	return nil
}

func (x *SearchResult) GetSayings() []*Phrase {
	if x != nil {
		return x.Sayings
	}
	return nil
}

func (x *SearchResult) GetExamples() []*Example {
	if x != nil {
		return x.Examples
	}
	return nil
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   string          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Source  LangCode        `protobuf:"varint,2,opt,name=source,proto3,enum=turkionary_proto.LangCode" json:"source,omitempty"`
	Target  LangCode        `protobuf:"varint,3,opt,name=target,proto3,enum=turkionary_proto.LangCode" json:"target,omitempty"`
	Results []*SearchResult `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResponse) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchResponse) GetSource() LangCode {
	if x != nil {
		return x.Source
	}
	return LangCode_EN
}

func (x *SearchResponse) GetTarget() LangCode {
	if x != nil {
		return x.Target
	}
	return LangCode_EN
}

func (x *SearchResponse) GetResults() []*SearchResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x75, 0x72, 0x6b, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74,
	0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x22, 0x6a, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74,
	0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4c, 0x0a, 0x06,
	0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x39, 0x0a, 0x07, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x05,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x07,
	0x73, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x07, 0x73, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x08, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x32, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x75, 0x72, 0x6b, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x75, 0x72, 0x6b,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2a, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06,
	0x0a, 0x02, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x52, 0x10, 0x01, 0x2a, 0x2d,
	0x0a, 0x08, 0x57, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x55, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x45, 0x52, 0x42, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x44, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_search_proto_goTypes = []interface{}{
	(LangCode)(0),          // 0: turkionary_proto.LangCode
	(WordType)(0),          // 1: turkionary_proto.WordType
	(*SearchRequest)(nil),  // 2: turkionary_proto.SearchRequest
	(*Definition)(nil),     // 3: turkionary_proto.Definition
	(*Phrase)(nil),         // 4: turkionary_proto.Phrase
	(*Example)(nil),        // 5: turkionary_proto.Example
	(*SearchResult)(nil),   // 6: turkionary_proto.SearchResult
	(*SearchResponse)(nil), // 7: turkionary_proto.SearchResponse
}
var file_search_proto_depIdxs = []int32{
	0,  // 0: turkionary_proto.SearchRequest.source:type_name -> turkionary_proto.LangCode
	0,  // 1: turkionary_proto.SearchRequest.target:type_name -> turkionary_proto.LangCode
	0,  // 2: turkionary_proto.Definition.lang:type_name -> turkionary_proto.LangCode
	0,  // 3: turkionary_proto.Phrase.lang:type_name -> turkionary_proto.LangCode
	4,  // 4: turkionary_proto.Example.texts:type_name -> turkionary_proto.Phrase
	4,  // 5: turkionary_proto.SearchResult.sayings:type_name -> turkionary_proto.Phrase
	5,  // 6: turkionary_proto.SearchResult.examples:type_name -> turkionary_proto.Example
	0,  // 7: turkionary_proto.SearchResponse.source:type_name -> turkionary_proto.LangCode
	0,  // 8: turkionary_proto.SearchResponse.target:type_name -> turkionary_proto.LangCode
	6,  // 9: turkionary_proto.SearchResponse.results:type_name -> turkionary_proto.SearchResult
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Definition); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phrase); i {
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
		file_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
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
		file_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		EnumInfos:         file_search_proto_enumTypes,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}