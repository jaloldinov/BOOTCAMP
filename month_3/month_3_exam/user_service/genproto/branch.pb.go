// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: branch.proto

package user_service

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

type CreateBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Photo           string `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
	Phone           string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	DeliveryTarifId int32  `protobuf:"varint,4,opt,name=delivery_tarif_id,json=deliveryTarifId,proto3" json:"delivery_tarif_id,omitempty"`
	WorkHourStart   string `protobuf:"bytes,5,opt,name=work_hour_start,json=workHourStart,proto3" json:"work_hour_start,omitempty"`
	WorkHourEnd     string `protobuf:"bytes,6,opt,name=work_hour_end,json=workHourEnd,proto3" json:"work_hour_end,omitempty"`
	Address         string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Destination     string `protobuf:"bytes,8,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *CreateBranchRequest) Reset() {
	*x = CreateBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBranchRequest) ProtoMessage() {}

func (x *CreateBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBranchRequest.ProtoReflect.Descriptor instead.
func (*CreateBranchRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBranchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBranchRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateBranchRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateBranchRequest) GetDeliveryTarifId() int32 {
	if x != nil {
		return x.DeliveryTarifId
	}
	return 0
}

func (x *CreateBranchRequest) GetWorkHourStart() string {
	if x != nil {
		return x.WorkHourStart
	}
	return ""
}

func (x *CreateBranchRequest) GetWorkHourEnd() string {
	if x != nil {
		return x.WorkHourEnd
	}
	return ""
}

func (x *CreateBranchRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateBranchRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone           string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Photo           string `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	DeliveryTarifId int32  `protobuf:"varint,5,opt,name=delivery_tarif_id,json=deliveryTarifId,proto3" json:"delivery_tarif_id,omitempty"`
	WorkHourStart   string `protobuf:"bytes,6,opt,name=work_hour_start,json=workHourStart,proto3" json:"work_hour_start,omitempty"`
	WorkHourEnd     string `protobuf:"bytes,7,opt,name=work_hour_end,json=workHourEnd,proto3" json:"work_hour_end,omitempty"`
	Address         string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Destination     string `protobuf:"bytes,9,opt,name=destination,proto3" json:"destination,omitempty"`
	Active          bool   `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt       string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{1}
}

func (x *Branch) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Branch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Branch) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Branch) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Branch) GetDeliveryTarifId() int32 {
	if x != nil {
		return x.DeliveryTarifId
	}
	return 0
}

func (x *Branch) GetWorkHourStart() string {
	if x != nil {
		return x.WorkHourStart
	}
	return ""
}

func (x *Branch) GetWorkHourEnd() string {
	if x != nil {
		return x.WorkHourEnd
	}
	return ""
}

func (x *Branch) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Branch) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Branch) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Branch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Branch) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone           string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Photo           string `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	DeliveryTarifId int32  `protobuf:"varint,5,opt,name=delivery_tarif_id,json=deliveryTarifId,proto3" json:"delivery_tarif_id,omitempty"`
	WorkHourStart   string `protobuf:"bytes,6,opt,name=work_hour_start,json=workHourStart,proto3" json:"work_hour_start,omitempty"`
	WorkHourEnd     string `protobuf:"bytes,7,opt,name=work_hour_end,json=workHourEnd,proto3" json:"work_hour_end,omitempty"`
	Address         string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Destination     string `protobuf:"bytes,9,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *UpdateBranchRequest) Reset() {
	*x = UpdateBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranchRequest) ProtoMessage() {}

func (x *UpdateBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranchRequest.ProtoReflect.Descriptor instead.
func (*UpdateBranchRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBranchRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBranchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBranchRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateBranchRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *UpdateBranchRequest) GetDeliveryTarifId() int32 {
	if x != nil {
		return x.DeliveryTarifId
	}
	return 0
}

func (x *UpdateBranchRequest) GetWorkHourStart() string {
	if x != nil {
		return x.WorkHourStart
	}
	return ""
}

func (x *UpdateBranchRequest) GetWorkHourEnd() string {
	if x != nil {
		return x.WorkHourEnd
	}
	return ""
}

func (x *UpdateBranchRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateBranchRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type ListBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit         int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page          int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAtFrom string `protobuf:"bytes,4,opt,name=created_at_from,json=createdAtFrom,proto3" json:"created_at_from,omitempty"`
	CreatedAtTo   string `protobuf:"bytes,5,opt,name=created_at_to,json=createdAtTo,proto3" json:"created_at_to,omitempty"`
}

func (x *ListBranchRequest) Reset() {
	*x = ListBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchRequest) ProtoMessage() {}

func (x *ListBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchRequest.ProtoReflect.Descriptor instead.
func (*ListBranchRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{3}
}

func (x *ListBranchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListBranchRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListBranchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListBranchRequest) GetCreatedAtFrom() string {
	if x != nil {
		return x.CreatedAtFrom
	}
	return ""
}

func (x *ListBranchRequest) GetCreatedAtTo() string {
	if x != nil {
		return x.CreatedAtTo
	}
	return ""
}

type ListBranchActiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Time  string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ListBranchActiveRequest) Reset() {
	*x = ListBranchActiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchActiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchActiveRequest) ProtoMessage() {}

func (x *ListBranchActiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchActiveRequest.ProtoReflect.Descriptor instead.
func (*ListBranchActiveRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{4}
}

func (x *ListBranchActiveRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListBranchActiveRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListBranchActiveRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type ListBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branches []*Branch `protobuf:"bytes,1,rep,name=branches,proto3" json:"branches,omitempty"`
	Count    int32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListBranchResponse) Reset() {
	*x = ListBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchResponse) ProtoMessage() {}

func (x *ListBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchResponse.ProtoReflect.Descriptor instead.
func (*ListBranchResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{5}
}

func (x *ListBranchResponse) GetBranches() []*Branch {
	if x != nil {
		return x.Branches
	}
	return nil
}

func (x *ListBranchResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{7}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_branch_proto protoreflect.FileDescriptor

var file_branch_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x89, 0x02, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x61, 0x72, 0x69, 0x66, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x48,
	0x6f, 0x75, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe2, 0x02, 0x0a, 0x06, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x61, 0x72, 0x69, 0x66, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x6f,
	0x75, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x99, 0x02,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x61, 0x72, 0x69, 0x66, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b,
	0x48, 0x6f, 0x75, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x54, 0x6f, 0x22, 0x57, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xb8, 0x03, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17,
	0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branch_proto_rawDescOnce sync.Once
	file_branch_proto_rawDescData = file_branch_proto_rawDesc
)

func file_branch_proto_rawDescGZIP() []byte {
	file_branch_proto_rawDescOnce.Do(func() {
		file_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_branch_proto_rawDescData)
	})
	return file_branch_proto_rawDescData
}

var file_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_branch_proto_goTypes = []interface{}{
	(*CreateBranchRequest)(nil),     // 0: user_service.CreateBranchRequest
	(*Branch)(nil),                  // 1: user_service.Branch
	(*UpdateBranchRequest)(nil),     // 2: user_service.UpdateBranchRequest
	(*ListBranchRequest)(nil),       // 3: user_service.ListBranchRequest
	(*ListBranchActiveRequest)(nil), // 4: user_service.ListBranchActiveRequest
	(*ListBranchResponse)(nil),      // 5: user_service.ListBranchResponse
	(*Response)(nil),                // 6: user_service.Response
	(*IdRequest)(nil),               // 7: user_service.IdRequest
}
var file_branch_proto_depIdxs = []int32{
	1, // 0: user_service.ListBranchResponse.branches:type_name -> user_service.Branch
	0, // 1: user_service.BranchService.Create:input_type -> user_service.CreateBranchRequest
	7, // 2: user_service.BranchService.Get:input_type -> user_service.IdRequest
	3, // 3: user_service.BranchService.List:input_type -> user_service.ListBranchRequest
	2, // 4: user_service.BranchService.Update:input_type -> user_service.UpdateBranchRequest
	7, // 5: user_service.BranchService.Delete:input_type -> user_service.IdRequest
	4, // 6: user_service.BranchService.ListActive:input_type -> user_service.ListBranchActiveRequest
	6, // 7: user_service.BranchService.Create:output_type -> user_service.Response
	1, // 8: user_service.BranchService.Get:output_type -> user_service.Branch
	5, // 9: user_service.BranchService.List:output_type -> user_service.ListBranchResponse
	6, // 10: user_service.BranchService.Update:output_type -> user_service.Response
	6, // 11: user_service.BranchService.Delete:output_type -> user_service.Response
	5, // 12: user_service.BranchService.ListActive:output_type -> user_service.ListBranchResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_branch_proto_init() }
func file_branch_proto_init() {
	if File_branch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBranchRequest); i {
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
		file_branch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
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
		file_branch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranchRequest); i {
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
		file_branch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchRequest); i {
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
		file_branch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchActiveRequest); i {
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
		file_branch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchResponse); i {
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
		file_branch_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_branch_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
			RawDescriptor: file_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branch_proto_goTypes,
		DependencyIndexes: file_branch_proto_depIdxs,
		MessageInfos:      file_branch_proto_msgTypes,
	}.Build()
	File_branch_proto = out.File
	file_branch_proto_rawDesc = nil
	file_branch_proto_goTypes = nil
	file_branch_proto_depIdxs = nil
}