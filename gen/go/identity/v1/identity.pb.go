// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: identity/v1/identity.proto

package identitypb

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ExtendedError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Errors  []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ExtendedError) Reset() {
	*x = ExtendedError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendedError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedError) ProtoMessage() {}

func (x *ExtendedError) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedError.ProtoReflect.Descriptor instead.
func (*ExtendedError) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{1}
}

func (x *ExtendedError) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ExtendedError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExtendedError) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{2}
}

func (x *Business) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Business) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Business) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Business) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Business) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Roles []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{3}
}

func (x *Auth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Auth) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

// CreateBusiness
type InviteBusiness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *InviteBusiness) Reset() {
	*x = InviteBusiness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteBusiness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteBusiness) ProtoMessage() {}

func (x *InviteBusiness) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteBusiness.ProtoReflect.Descriptor instead.
func (*InviteBusiness) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{4}
}

func (x *InviteBusiness) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InviteBusiness) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InviteBusiness) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type InviteBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *Auth           `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Business *InviteBusiness `protobuf:"bytes,2,opt,name=business,proto3" json:"business,omitempty"`
}

func (x *InviteBusinessRequest) Reset() {
	*x = InviteBusinessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteBusinessRequest) ProtoMessage() {}

func (x *InviteBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteBusinessRequest.ProtoReflect.Descriptor instead.
func (*InviteBusinessRequest) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{5}
}

func (x *InviteBusinessRequest) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *InviteBusinessRequest) GetBusiness() *InviteBusiness {
	if x != nil {
		return x.Business
	}
	return nil
}

type InviteBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Business *Business      `protobuf:"bytes,1,opt,name=business,proto3" json:"business,omitempty"`
	Error    *ExtendedError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Status   uint32         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *InviteBusinessResponse) Reset() {
	*x = InviteBusinessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteBusinessResponse) ProtoMessage() {}

func (x *InviteBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteBusinessResponse.ProtoReflect.Descriptor instead.
func (*InviteBusinessResponse) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{6}
}

func (x *InviteBusinessResponse) GetBusiness() *Business {
	if x != nil {
		return x.Business
	}
	return nil
}

func (x *InviteBusinessResponse) GetError() *ExtendedError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *InviteBusinessResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// AcceptInvitationAsBusiness
type AcceptInvitationAsBusiness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	TokenValue uint32 `protobuf:"varint,3,opt,name=tokenValue,proto3" json:"tokenValue,omitempty"`
}

func (x *AcceptInvitationAsBusiness) Reset() {
	*x = AcceptInvitationAsBusiness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptInvitationAsBusiness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptInvitationAsBusiness) ProtoMessage() {}

func (x *AcceptInvitationAsBusiness) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptInvitationAsBusiness.ProtoReflect.Descriptor instead.
func (*AcceptInvitationAsBusiness) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{7}
}

func (x *AcceptInvitationAsBusiness) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AcceptInvitationAsBusiness) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AcceptInvitationAsBusiness) GetTokenValue() uint32 {
	if x != nil {
		return x.TokenValue
	}
	return 0
}

type AcceptInvitationAsBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcceptInvitationAsBusiness *AcceptInvitationAsBusiness `protobuf:"bytes,1,opt,name=acceptInvitationAsBusiness,proto3" json:"acceptInvitationAsBusiness,omitempty"`
}

func (x *AcceptInvitationAsBusinessRequest) Reset() {
	*x = AcceptInvitationAsBusinessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptInvitationAsBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptInvitationAsBusinessRequest) ProtoMessage() {}

func (x *AcceptInvitationAsBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptInvitationAsBusinessRequest.ProtoReflect.Descriptor instead.
func (*AcceptInvitationAsBusinessRequest) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{8}
}

func (x *AcceptInvitationAsBusinessRequest) GetAcceptInvitationAsBusiness() *AcceptInvitationAsBusiness {
	if x != nil {
		return x.AcceptInvitationAsBusiness
	}
	return nil
}

type AcceptInvitationAsBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAccepted bool           `protobuf:"varint,1,opt,name=isAccepted,proto3" json:"isAccepted,omitempty"`
	Error      *ExtendedError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Status     uint32         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AcceptInvitationAsBusinessResponse) Reset() {
	*x = AcceptInvitationAsBusinessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptInvitationAsBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptInvitationAsBusinessResponse) ProtoMessage() {}

func (x *AcceptInvitationAsBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptInvitationAsBusinessResponse.ProtoReflect.Descriptor instead.
func (*AcceptInvitationAsBusinessResponse) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{9}
}

func (x *AcceptInvitationAsBusinessResponse) GetIsAccepted() bool {
	if x != nil {
		return x.IsAccepted
	}
	return false
}

func (x *AcceptInvitationAsBusinessResponse) GetError() *ExtendedError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AcceptInvitationAsBusinessResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// Get BusinessById
type GetBusinessByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBusinessByIdRequest) Reset() {
	*x = GetBusinessByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessByIdRequest) ProtoMessage() {}

func (x *GetBusinessByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessByIdRequest.ProtoReflect.Descriptor instead.
func (*GetBusinessByIdRequest) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{10}
}

func (x *GetBusinessByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBusinessByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Business *Business      `protobuf:"bytes,1,opt,name=business,proto3" json:"business,omitempty"`
	Error    *ExtendedError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Status   uint32         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetBusinessByIdResponse) Reset() {
	*x = GetBusinessByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_v1_identity_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessByIdResponse) ProtoMessage() {}

func (x *GetBusinessByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_v1_identity_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessByIdResponse.ProtoReflect.Descriptor instead.
func (*GetBusinessByIdResponse) Descriptor() ([]byte, []int) {
	return file_identity_v1_identity_proto_rawDescGZIP(), []int{11}
}

func (x *GetBusinessByIdResponse) GetBusiness() *Business {
	if x != nil {
		return x.Business
	}
	return nil
}

func (x *GetBusinessByIdResponse) GetError() *ExtendedError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetBusinessByIdResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_identity_v1_identity_proto protoreflect.FileDescriptor

var file_identity_v1_identity_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x33, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x67,
	0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x0e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x77, 0x0a, 0x15, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x37, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x16,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x6e, 0x0a, 0x1a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x21, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x67, 0x0a, 0x1a, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x1a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x22, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x96, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xcf, 0x02, 0x0a, 0x0f, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x1a, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x61, 0x72, 0x2d, 0x61, 0x72, 0x61,
	0x6e, 0x64, 0x6a, 0x69, 0x63, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6b, 0x65,
	0x6c, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_v1_identity_proto_rawDescOnce sync.Once
	file_identity_v1_identity_proto_rawDescData = file_identity_v1_identity_proto_rawDesc
)

func file_identity_v1_identity_proto_rawDescGZIP() []byte {
	file_identity_v1_identity_proto_rawDescOnce.Do(func() {
		file_identity_v1_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_v1_identity_proto_rawDescData)
	})
	return file_identity_v1_identity_proto_rawDescData
}

var file_identity_v1_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_identity_v1_identity_proto_goTypes = []interface{}{
	(*Error)(nil),                              // 0: identity.v1.Error
	(*ExtendedError)(nil),                      // 1: identity.v1.ExtendedError
	(*Business)(nil),                           // 2: identity.v1.Business
	(*Auth)(nil),                               // 3: identity.v1.Auth
	(*InviteBusiness)(nil),                     // 4: identity.v1.InviteBusiness
	(*InviteBusinessRequest)(nil),              // 5: identity.v1.InviteBusinessRequest
	(*InviteBusinessResponse)(nil),             // 6: identity.v1.InviteBusinessResponse
	(*AcceptInvitationAsBusiness)(nil),         // 7: identity.v1.AcceptInvitationAsBusiness
	(*AcceptInvitationAsBusinessRequest)(nil),  // 8: identity.v1.AcceptInvitationAsBusinessRequest
	(*AcceptInvitationAsBusinessResponse)(nil), // 9: identity.v1.AcceptInvitationAsBusinessResponse
	(*GetBusinessByIdRequest)(nil),             // 10: identity.v1.GetBusinessByIdRequest
	(*GetBusinessByIdResponse)(nil),            // 11: identity.v1.GetBusinessByIdResponse
}
var file_identity_v1_identity_proto_depIdxs = []int32{
	0,  // 0: identity.v1.ExtendedError.errors:type_name -> identity.v1.Error
	3,  // 1: identity.v1.InviteBusinessRequest.auth:type_name -> identity.v1.Auth
	4,  // 2: identity.v1.InviteBusinessRequest.business:type_name -> identity.v1.InviteBusiness
	2,  // 3: identity.v1.InviteBusinessResponse.business:type_name -> identity.v1.Business
	1,  // 4: identity.v1.InviteBusinessResponse.error:type_name -> identity.v1.ExtendedError
	7,  // 5: identity.v1.AcceptInvitationAsBusinessRequest.acceptInvitationAsBusiness:type_name -> identity.v1.AcceptInvitationAsBusiness
	1,  // 6: identity.v1.AcceptInvitationAsBusinessResponse.error:type_name -> identity.v1.ExtendedError
	2,  // 7: identity.v1.GetBusinessByIdResponse.business:type_name -> identity.v1.Business
	1,  // 8: identity.v1.GetBusinessByIdResponse.error:type_name -> identity.v1.ExtendedError
	5,  // 9: identity.v1.BusinessService.InviteBusiness:input_type -> identity.v1.InviteBusinessRequest
	8,  // 10: identity.v1.BusinessService.AcceptInvitationAsBusiness:input_type -> identity.v1.AcceptInvitationAsBusinessRequest
	10, // 11: identity.v1.BusinessService.GetBusinessById:input_type -> identity.v1.GetBusinessByIdRequest
	6,  // 12: identity.v1.BusinessService.InviteBusiness:output_type -> identity.v1.InviteBusinessResponse
	9,  // 13: identity.v1.BusinessService.AcceptInvitationAsBusiness:output_type -> identity.v1.AcceptInvitationAsBusinessResponse
	11, // 14: identity.v1.BusinessService.GetBusinessById:output_type -> identity.v1.GetBusinessByIdResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_identity_v1_identity_proto_init() }
func file_identity_v1_identity_proto_init() {
	if File_identity_v1_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_v1_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_identity_v1_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendedError); i {
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
		file_identity_v1_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
		file_identity_v1_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_identity_v1_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteBusiness); i {
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
		file_identity_v1_identity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteBusinessRequest); i {
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
		file_identity_v1_identity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteBusinessResponse); i {
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
		file_identity_v1_identity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptInvitationAsBusiness); i {
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
		file_identity_v1_identity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptInvitationAsBusinessRequest); i {
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
		file_identity_v1_identity_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptInvitationAsBusinessResponse); i {
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
		file_identity_v1_identity_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessByIdRequest); i {
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
		file_identity_v1_identity_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessByIdResponse); i {
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
			RawDescriptor: file_identity_v1_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_v1_identity_proto_goTypes,
		DependencyIndexes: file_identity_v1_identity_proto_depIdxs,
		MessageInfos:      file_identity_v1_identity_proto_msgTypes,
	}.Build()
	File_identity_v1_identity_proto = out.File
	file_identity_v1_identity_proto_rawDesc = nil
	file_identity_v1_identity_proto_goTypes = nil
	file_identity_v1_identity_proto_depIdxs = nil
}
