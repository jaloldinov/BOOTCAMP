// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: staff_transaction.proto

package sale_service

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

type CreateStaffTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleId     string  `protobuf:"bytes,1,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	StaffId    string  `protobuf:"bytes,2,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	TrType     string  `protobuf:"bytes,3,opt,name=tr_type,json=trType,proto3" json:"tr_type,omitempty"`
	SourceType string  `protobuf:"bytes,4,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	Amount     float32 `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount,omitempty"`
	AboutText  string  `protobuf:"bytes,6,opt,name=about_text,json=aboutText,proto3" json:"about_text,omitempty"`
}

func (x *CreateStaffTransactionRequest) Reset() {
	*x = CreateStaffTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffTransactionRequest) ProtoMessage() {}

func (x *CreateStaffTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateStaffTransactionRequest) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStaffTransactionRequest) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *CreateStaffTransactionRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *CreateStaffTransactionRequest) GetTrType() string {
	if x != nil {
		return x.TrType
	}
	return ""
}

func (x *CreateStaffTransactionRequest) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *CreateStaffTransactionRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateStaffTransactionRequest) GetAboutText() string {
	if x != nil {
		return x.AboutText
	}
	return ""
}

type StaffTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SaleId     string  `protobuf:"bytes,2,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	StaffId    string  `protobuf:"bytes,3,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	TrType     string  `protobuf:"bytes,4,opt,name=tr_type,json=trType,proto3" json:"tr_type,omitempty"`
	SourceType string  `protobuf:"bytes,5,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	Amount     float32 `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	AboutText  string  `protobuf:"bytes,7,opt,name=about_text,json=aboutText,proto3" json:"about_text,omitempty"`
	CreatedAt  string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  string  `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *StaffTransaction) Reset() {
	*x = StaffTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffTransaction) ProtoMessage() {}

func (x *StaffTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffTransaction.ProtoReflect.Descriptor instead.
func (*StaffTransaction) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *StaffTransaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StaffTransaction) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *StaffTransaction) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *StaffTransaction) GetTrType() string {
	if x != nil {
		return x.TrType
	}
	return ""
}

func (x *StaffTransaction) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *StaffTransaction) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *StaffTransaction) GetAboutText() string {
	if x != nil {
		return x.AboutText
	}
	return ""
}

func (x *StaffTransaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *StaffTransaction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *StaffTransaction) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CreateStaffTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateStaffTransactionResponse) Reset() {
	*x = CreateStaffTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffTransactionResponse) ProtoMessage() {}

func (x *CreateStaffTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateStaffTransactionResponse) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStaffTransactionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetStaffTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffTransaction *StaffTransaction `protobuf:"bytes,1,opt,name=StaffTransaction,proto3" json:"StaffTransaction,omitempty"`
}

func (x *GetStaffTransactionResponse) Reset() {
	*x = GetStaffTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaffTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaffTransactionResponse) ProtoMessage() {}

func (x *GetStaffTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaffTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetStaffTransactionResponse) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *GetStaffTransactionResponse) GetStaffTransaction() *StaffTransaction {
	if x != nil {
		return x.StaffTransaction
	}
	return nil
}

type UpdateStaffTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SaleId     string  `protobuf:"bytes,2,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	StaffId    string  `protobuf:"bytes,3,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	TrType     string  `protobuf:"bytes,4,opt,name=tr_type,json=trType,proto3" json:"tr_type,omitempty"`
	SourceType string  `protobuf:"bytes,5,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	Amount     float32 `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	AboutText  string  `protobuf:"bytes,7,opt,name=about_text,json=aboutText,proto3" json:"about_text,omitempty"`
}

func (x *UpdateStaffTransactionRequest) Reset() {
	*x = UpdateStaffTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaffTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffTransactionRequest) ProtoMessage() {}

func (x *UpdateStaffTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffTransactionRequest.ProtoReflect.Descriptor instead.
func (*UpdateStaffTransactionRequest) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStaffTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStaffTransactionRequest) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *UpdateStaffTransactionRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *UpdateStaffTransactionRequest) GetTrType() string {
	if x != nil {
		return x.TrType
	}
	return ""
}

func (x *UpdateStaffTransactionRequest) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *UpdateStaffTransactionRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateStaffTransactionRequest) GetAboutText() string {
	if x != nil {
		return x.AboutText
	}
	return ""
}

type ListStaffTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *ListStaffTransactionRequest) Reset() {
	*x = ListStaffTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffTransactionRequest) ProtoMessage() {}

func (x *ListStaffTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffTransactionRequest.ProtoReflect.Descriptor instead.
func (*ListStaffTransactionRequest) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *ListStaffTransactionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListStaffTransactionRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListStaffTransactionRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type ListStaffTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffTransactions []*StaffTransaction `protobuf:"bytes,1,rep,name=StaffTransactions,proto3" json:"StaffTransactions,omitempty"`
	Count             int32               `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListStaffTransactionResponse) Reset() {
	*x = ListStaffTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffTransactionResponse) ProtoMessage() {}

func (x *ListStaffTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffTransactionResponse.ProtoReflect.Descriptor instead.
func (*ListStaffTransactionResponse) Descriptor() ([]byte, []int) {
	return file_staff_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *ListStaffTransactionResponse) GetStaffTransactions() []*StaffTransaction {
	if x != nil {
		return x.StaffTransactions
	}
	return nil
}

func (x *ListStaffTransactionResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_staff_transaction_proto protoreflect.FileDescriptor

var file_staff_transaction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x62, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0xa4, 0x02, 0x0a, 0x10, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x30, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x69, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd4,
	0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0x5f, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x11, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbc, 0x03, 0x0a, 0x17,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staff_transaction_proto_rawDescOnce sync.Once
	file_staff_transaction_proto_rawDescData = file_staff_transaction_proto_rawDesc
)

func file_staff_transaction_proto_rawDescGZIP() []byte {
	file_staff_transaction_proto_rawDescOnce.Do(func() {
		file_staff_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_transaction_proto_rawDescData)
	})
	return file_staff_transaction_proto_rawDescData
}

var file_staff_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_staff_transaction_proto_goTypes = []interface{}{
	(*CreateStaffTransactionRequest)(nil),  // 0: sale_service.CreateStaffTransactionRequest
	(*StaffTransaction)(nil),               // 1: sale_service.StaffTransaction
	(*CreateStaffTransactionResponse)(nil), // 2: sale_service.CreateStaffTransactionResponse
	(*GetStaffTransactionResponse)(nil),    // 3: sale_service.GetStaffTransactionResponse
	(*UpdateStaffTransactionRequest)(nil),  // 4: sale_service.UpdateStaffTransactionRequest
	(*ListStaffTransactionRequest)(nil),    // 5: sale_service.ListStaffTransactionRequest
	(*ListStaffTransactionResponse)(nil),   // 6: sale_service.ListStaffTransactionResponse
	(*IdRequest)(nil),                      // 7: sale_service.IdRequest
	(*Response)(nil),                       // 8: sale_service.Response
}
var file_staff_transaction_proto_depIdxs = []int32{
	1, // 0: sale_service.GetStaffTransactionResponse.StaffTransaction:type_name -> sale_service.StaffTransaction
	1, // 1: sale_service.ListStaffTransactionResponse.StaffTransactions:type_name -> sale_service.StaffTransaction
	0, // 2: sale_service.StaffTransactionService.Create:input_type -> sale_service.CreateStaffTransactionRequest
	7, // 3: sale_service.StaffTransactionService.Get:input_type -> sale_service.IdRequest
	5, // 4: sale_service.StaffTransactionService.List:input_type -> sale_service.ListStaffTransactionRequest
	4, // 5: sale_service.StaffTransactionService.Update:input_type -> sale_service.UpdateStaffTransactionRequest
	7, // 6: sale_service.StaffTransactionService.Delete:input_type -> sale_service.IdRequest
	2, // 7: sale_service.StaffTransactionService.Create:output_type -> sale_service.CreateStaffTransactionResponse
	3, // 8: sale_service.StaffTransactionService.Get:output_type -> sale_service.GetStaffTransactionResponse
	6, // 9: sale_service.StaffTransactionService.List:output_type -> sale_service.ListStaffTransactionResponse
	8, // 10: sale_service.StaffTransactionService.Update:output_type -> sale_service.Response
	8, // 11: sale_service.StaffTransactionService.Delete:output_type -> sale_service.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_staff_transaction_proto_init() }
func file_staff_transaction_proto_init() {
	if File_staff_transaction_proto != nil {
		return
	}
	file_sale_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staff_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffTransactionRequest); i {
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
		file_staff_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffTransaction); i {
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
		file_staff_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffTransactionResponse); i {
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
		file_staff_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaffTransactionResponse); i {
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
		file_staff_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaffTransactionRequest); i {
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
		file_staff_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffTransactionRequest); i {
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
		file_staff_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffTransactionResponse); i {
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
			RawDescriptor: file_staff_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_transaction_proto_goTypes,
		DependencyIndexes: file_staff_transaction_proto_depIdxs,
		MessageInfos:      file_staff_transaction_proto_msgTypes,
	}.Build()
	File_staff_transaction_proto = out.File
	file_staff_transaction_proto_rawDesc = nil
	file_staff_transaction_proto_goTypes = nil
	file_staff_transaction_proto_depIdxs = nil
}