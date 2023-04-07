// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: chifra.proto

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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parts int64    `protobuf:"varint,1,opt,name=parts,proto3" json:"parts,omitempty"`
	Terms []string `protobuf:"bytes,2,rep,name=terms,proto3" json:"terms,omitempty"`
	Sort  int64    `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chifra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chifra_proto_msgTypes[0]
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
	return file_chifra_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetParts() int64 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *SearchRequest) GetTerms() []string {
	if x != nil {
		return x.Terms
	}
	return nil
}

func (x *SearchRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []*Name `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chifra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chifra_proto_msgTypes[1]
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
	return file_chifra_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetNames() []*Name {
	if x != nil {
		return x.Names
	}
	return nil
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Decimals   *uint64 `protobuf:"varint,2,opt,name=decimals,proto3,oneof" json:"decimals,omitempty"`
	Deleted    *bool   `protobuf:"varint,3,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	IsContract *bool   `protobuf:"varint,4,opt,name=isContract,proto3,oneof" json:"isContract,omitempty"`
	IsCustom   *bool   `protobuf:"varint,5,opt,name=isCustom,proto3,oneof" json:"isCustom,omitempty"`
	IsErc20    *bool   `protobuf:"varint,6,opt,name=isErc20,proto3,oneof" json:"isErc20,omitempty"`
	IsErc721   *bool   `protobuf:"varint,7,opt,name=isErc721,proto3,oneof" json:"isErc721,omitempty"`
	IsPrefund  *bool   `protobuf:"varint,8,opt,name=isPrefund,proto3,oneof" json:"isPrefund,omitempty"`
	Name       string  `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Petname    *string `protobuf:"bytes,10,opt,name=petname,proto3,oneof" json:"petname,omitempty"`
	Source     *string `protobuf:"bytes,11,opt,name=source,proto3,oneof" json:"source,omitempty"`
	Symbol     *string `protobuf:"bytes,12,opt,name=symbol,proto3,oneof" json:"symbol,omitempty"`
	Tags       *string `protobuf:"bytes,13,opt,name=tags,proto3,oneof" json:"tags,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chifra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_chifra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_chifra_proto_rawDescGZIP(), []int{2}
}

func (x *Name) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Name) GetDecimals() uint64 {
	if x != nil && x.Decimals != nil {
		return *x.Decimals
	}
	return 0
}

func (x *Name) GetDeleted() bool {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return false
}

func (x *Name) GetIsContract() bool {
	if x != nil && x.IsContract != nil {
		return *x.IsContract
	}
	return false
}

func (x *Name) GetIsCustom() bool {
	if x != nil && x.IsCustom != nil {
		return *x.IsCustom
	}
	return false
}

func (x *Name) GetIsErc20() bool {
	if x != nil && x.IsErc20 != nil {
		return *x.IsErc20
	}
	return false
}

func (x *Name) GetIsErc721() bool {
	if x != nil && x.IsErc721 != nil {
		return *x.IsErc721
	}
	return false
}

func (x *Name) GetIsPrefund() bool {
	if x != nil && x.IsPrefund != nil {
		return *x.IsPrefund
	}
	return false
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Name) GetPetname() string {
	if x != nil && x.Petname != nil {
		return *x.Petname
	}
	return ""
}

func (x *Name) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *Name) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *Name) GetTags() string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chifra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chifra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_chifra_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_chifra_proto protoreflect.FileDescriptor

var file_chifra_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x68, 0x69, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22,
	0x2d, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x96,
	0x04, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x08, 0x69, 0x73, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x72, 0x63,
	0x32, 0x30, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x07, 0x69, 0x73, 0x45, 0x72,
	0x63, 0x32, 0x30, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x69, 0x73, 0x45, 0x72, 0x63, 0x37,
	0x32, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x08, 0x69, 0x73, 0x45, 0x72,
	0x63, 0x37, 0x32, 0x31, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x09, 0x69, 0x73,
	0x50, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x07, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x07, 0x52, 0x07, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x45, 0x72, 0x63, 0x32,
	0x30, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x45, 0x72, 0x63, 0x37, 0x32, 0x31, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x50, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x22, 0x4f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x83, 0x01, 0x0a, 0x05, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x29, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x22, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x72, 0x75,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x74, 0x72, 0x75, 0x65, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x63, 0x68, 0x69, 0x66, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chifra_proto_rawDescOnce sync.Once
	file_chifra_proto_rawDescData = file_chifra_proto_rawDesc
)

func file_chifra_proto_rawDescGZIP() []byte {
	file_chifra_proto_rawDescOnce.Do(func() {
		file_chifra_proto_rawDescData = protoimpl.X.CompressGZIP(file_chifra_proto_rawDescData)
	})
	return file_chifra_proto_rawDescData
}

var file_chifra_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chifra_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),  // 0: SearchRequest
	(*SearchResponse)(nil), // 1: SearchResponse
	(*Name)(nil),           // 2: Name
	(*CreateResponse)(nil), // 3: CreateResponse
}
var file_chifra_proto_depIdxs = []int32{
	2, // 0: SearchResponse.names:type_name -> Name
	0, // 1: Names.Search:input_type -> SearchRequest
	0, // 2: Names.SearchStream:input_type -> SearchRequest
	2, // 3: Names.Create:input_type -> Name
	1, // 4: Names.Search:output_type -> SearchResponse
	2, // 5: Names.SearchStream:output_type -> Name
	3, // 6: Names.Create:output_type -> CreateResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chifra_proto_init() }
func file_chifra_proto_init() {
	if File_chifra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chifra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_chifra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_chifra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_chifra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
	file_chifra_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_chifra_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chifra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chifra_proto_goTypes,
		DependencyIndexes: file_chifra_proto_depIdxs,
		MessageInfos:      file_chifra_proto_msgTypes,
	}.Build()
	File_chifra_proto = out.File
	file_chifra_proto_rawDesc = nil
	file_chifra_proto_goTypes = nil
	file_chifra_proto_depIdxs = nil
}
