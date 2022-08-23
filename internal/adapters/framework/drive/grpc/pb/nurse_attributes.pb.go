// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: nurse_attributes.proto

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

type Nurse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NurseId          int64  `protobuf:"varint,1,opt,name=nurse_id,json=nurseId,proto3" json:"nurse_id,omitempty"`
	RegNumber        string `protobuf:"bytes,2,opt,name=reg_number,json=regNumber,proto3" json:"reg_number,omitempty"`
	Dob              string `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	NurseRank        string `protobuf:"bytes,4,opt,name=nurse_rank,json=nurseRank,proto3" json:"nurse_rank,omitempty"`
	ActiveStatus     bool   `protobuf:"varint,5,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"`
	CenterId         int64  `protobuf:"varint,6,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
	UserId           string `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName         string `protobuf:"bytes,8,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName        string `protobuf:"bytes,9,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName         string `protobuf:"bytes,10,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Nic              string `protobuf:"bytes,11,opt,name=nic,proto3" json:"nic,omitempty"`
	TellNo           string `protobuf:"bytes,12,opt,name=tell_no,json=tellNo,proto3" json:"tell_no,omitempty"`
	Address          string `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	UserEmail        string `protobuf:"bytes,14,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	IsEmailVerified  bool   `protobuf:"varint,15,opt,name=is_email_verified,json=isEmailVerified,proto3" json:"is_email_verified,omitempty"`
	IsTellNoVerified bool   `protobuf:"varint,16,opt,name=is_tell_no_verified,json=isTellNoVerified,proto3" json:"is_tell_no_verified,omitempty"`
}

func (x *Nurse) Reset() {
	*x = Nurse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nurse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nurse) ProtoMessage() {}

func (x *Nurse) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nurse.ProtoReflect.Descriptor instead.
func (*Nurse) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{0}
}

func (x *Nurse) GetNurseId() int64 {
	if x != nil {
		return x.NurseId
	}
	return 0
}

func (x *Nurse) GetRegNumber() string {
	if x != nil {
		return x.RegNumber
	}
	return ""
}

func (x *Nurse) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Nurse) GetNurseRank() string {
	if x != nil {
		return x.NurseRank
	}
	return ""
}

func (x *Nurse) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *Nurse) GetCenterId() int64 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

func (x *Nurse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Nurse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Nurse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Nurse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Nurse) GetNic() string {
	if x != nil {
		return x.Nic
	}
	return ""
}

func (x *Nurse) GetTellNo() string {
	if x != nil {
		return x.TellNo
	}
	return ""
}

func (x *Nurse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Nurse) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Nurse) GetIsEmailVerified() bool {
	if x != nil {
		return x.IsEmailVerified
	}
	return false
}

func (x *Nurse) GetIsTellNoVerified() bool {
	if x != nil {
		return x.IsTellNoVerified
	}
	return false
}

type RegisterNurseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegNumber        string `protobuf:"bytes,1,opt,name=reg_number,json=regNumber,proto3" json:"reg_number,omitempty"`
	Dob              string `protobuf:"bytes,2,opt,name=dob,proto3" json:"dob,omitempty"`
	NurseRank        string `protobuf:"bytes,3,opt,name=nurse_rank,json=nurseRank,proto3" json:"nurse_rank,omitempty"`
	ActiveStatus     bool   `protobuf:"varint,4,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"`
	CenterId         int64  `protobuf:"varint,5,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
	UserId           string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName         string `protobuf:"bytes,7,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName        string `protobuf:"bytes,8,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName         string `protobuf:"bytes,9,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Nic              string `protobuf:"bytes,10,opt,name=nic,proto3" json:"nic,omitempty"`
	TellNo           string `protobuf:"bytes,11,opt,name=tell_no,json=tellNo,proto3" json:"tell_no,omitempty"`
	Address          string `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	UserEmail        string `protobuf:"bytes,13,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserPwd          string `protobuf:"bytes,14,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	IsEmailVerified  bool   `protobuf:"varint,15,opt,name=is_email_verified,json=isEmailVerified,proto3" json:"is_email_verified,omitempty"`
	IsTellNoVerified bool   `protobuf:"varint,16,opt,name=is_tell_no_verified,json=isTellNoVerified,proto3" json:"is_tell_no_verified,omitempty"`
}

func (x *RegisterNurseRequest) Reset() {
	*x = RegisterNurseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNurseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNurseRequest) ProtoMessage() {}

func (x *RegisterNurseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNurseRequest.ProtoReflect.Descriptor instead.
func (*RegisterNurseRequest) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterNurseRequest) GetRegNumber() string {
	if x != nil {
		return x.RegNumber
	}
	return ""
}

func (x *RegisterNurseRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *RegisterNurseRequest) GetNurseRank() string {
	if x != nil {
		return x.NurseRank
	}
	return ""
}

func (x *RegisterNurseRequest) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *RegisterNurseRequest) GetCenterId() int64 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

func (x *RegisterNurseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterNurseRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RegisterNurseRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterNurseRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterNurseRequest) GetNic() string {
	if x != nil {
		return x.Nic
	}
	return ""
}

func (x *RegisterNurseRequest) GetTellNo() string {
	if x != nil {
		return x.TellNo
	}
	return ""
}

func (x *RegisterNurseRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterNurseRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *RegisterNurseRequest) GetUserPwd() string {
	if x != nil {
		return x.UserPwd
	}
	return ""
}

func (x *RegisterNurseRequest) GetIsEmailVerified() bool {
	if x != nil {
		return x.IsEmailVerified
	}
	return false
}

func (x *RegisterNurseRequest) GetIsTellNoVerified() bool {
	if x != nil {
		return x.IsTellNoVerified
	}
	return false
}

type UpdateNurseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegNumber        string `protobuf:"bytes,1,opt,name=reg_number,json=regNumber,proto3" json:"reg_number,omitempty"`
	Dob              string `protobuf:"bytes,2,opt,name=dob,proto3" json:"dob,omitempty"`
	NurseRank        string `protobuf:"bytes,3,opt,name=nurse_rank,json=nurseRank,proto3" json:"nurse_rank,omitempty"`
	ActiveStatus     bool   `protobuf:"varint,4,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"`
	CenterId         int64  `protobuf:"varint,5,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
	UserId           string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName         string `protobuf:"bytes,7,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName        string `protobuf:"bytes,8,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName         string `protobuf:"bytes,9,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Nic              string `protobuf:"bytes,10,opt,name=nic,proto3" json:"nic,omitempty"`
	TellNo           string `protobuf:"bytes,11,opt,name=tell_no,json=tellNo,proto3" json:"tell_no,omitempty"`
	Address          string `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	UserEmail        string `protobuf:"bytes,13,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserPwd          string `protobuf:"bytes,14,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	IsEmailVerified  bool   `protobuf:"varint,15,opt,name=is_email_verified,json=isEmailVerified,proto3" json:"is_email_verified,omitempty"`
	IsTellNoVerified bool   `protobuf:"varint,16,opt,name=is_tell_no_verified,json=isTellNoVerified,proto3" json:"is_tell_no_verified,omitempty"`
	NurseId          int64  `protobuf:"varint,17,opt,name=nurseId,proto3" json:"nurseId,omitempty"`
}

func (x *UpdateNurseRequest) Reset() {
	*x = UpdateNurseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNurseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNurseRequest) ProtoMessage() {}

func (x *UpdateNurseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNurseRequest.ProtoReflect.Descriptor instead.
func (*UpdateNurseRequest) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNurseRequest) GetRegNumber() string {
	if x != nil {
		return x.RegNumber
	}
	return ""
}

func (x *UpdateNurseRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *UpdateNurseRequest) GetNurseRank() string {
	if x != nil {
		return x.NurseRank
	}
	return ""
}

func (x *UpdateNurseRequest) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

func (x *UpdateNurseRequest) GetCenterId() int64 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

func (x *UpdateNurseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateNurseRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UpdateNurseRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateNurseRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateNurseRequest) GetNic() string {
	if x != nil {
		return x.Nic
	}
	return ""
}

func (x *UpdateNurseRequest) GetTellNo() string {
	if x != nil {
		return x.TellNo
	}
	return ""
}

func (x *UpdateNurseRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateNurseRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UpdateNurseRequest) GetUserPwd() string {
	if x != nil {
		return x.UserPwd
	}
	return ""
}

func (x *UpdateNurseRequest) GetIsEmailVerified() bool {
	if x != nil {
		return x.IsEmailVerified
	}
	return false
}

func (x *UpdateNurseRequest) GetIsTellNoVerified() bool {
	if x != nil {
		return x.IsTellNoVerified
	}
	return false
}

func (x *UpdateNurseRequest) GetNurseId() int64 {
	if x != nil {
		return x.NurseId
	}
	return 0
}

type UpdateNurseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateNurseResponse) Reset() {
	*x = UpdateNurseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNurseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNurseResponse) ProtoMessage() {}

func (x *UpdateNurseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNurseResponse.ProtoReflect.Descriptor instead.
func (*UpdateNurseResponse) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNurseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateNurseResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type RegisterNurseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NurseId   int64  `protobuf:"varint,1,opt,name=nurse_id,json=nurseId,proto3" json:"nurse_id,omitempty"`
	RegNumber string `protobuf:"bytes,2,opt,name=reg_number,json=regNumber,proto3" json:"reg_number,omitempty"`
	NurseRank string `protobuf:"bytes,3,opt,name=nurse_rank,json=nurseRank,proto3" json:"nurse_rank,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message   string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Status    int64  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RegisterNurseResponse) Reset() {
	*x = RegisterNurseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNurseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNurseResponse) ProtoMessage() {}

func (x *RegisterNurseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNurseResponse.ProtoReflect.Descriptor instead.
func (*RegisterNurseResponse) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterNurseResponse) GetNurseId() int64 {
	if x != nil {
		return x.NurseId
	}
	return 0
}

func (x *RegisterNurseResponse) GetRegNumber() string {
	if x != nil {
		return x.RegNumber
	}
	return ""
}

func (x *RegisterNurseResponse) GetNurseRank() string {
	if x != nil {
		return x.NurseRank
	}
	return ""
}

func (x *RegisterNurseResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterNurseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterNurseResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetAllNursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterId int64 `protobuf:"varint,1,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
}

func (x *GetAllNursesRequest) Reset() {
	*x = GetAllNursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNursesRequest) ProtoMessage() {}

func (x *GetAllNursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNursesRequest.ProtoReflect.Descriptor instead.
func (*GetAllNursesRequest) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllNursesRequest) GetCenterId() int64 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

type GetAllNursesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nurses []*Nurse `protobuf:"bytes,1,rep,name=nurses,proto3" json:"nurses,omitempty"`
}

func (x *GetAllNursesResponse) Reset() {
	*x = GetAllNursesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNursesResponse) ProtoMessage() {}

func (x *GetAllNursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNursesResponse.ProtoReflect.Descriptor instead.
func (*GetAllNursesResponse) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllNursesResponse) GetNurses() []*Nurse {
	if x != nil {
		return x.Nurses
	}
	return nil
}

type GetNurseFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterId int64  `protobuf:"varint,1,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetNurseFilterRequest) Reset() {
	*x = GetNurseFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nurse_attributes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNurseFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNurseFilterRequest) ProtoMessage() {}

func (x *GetNurseFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nurse_attributes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNurseFilterRequest.ProtoReflect.Descriptor instead.
func (*GetNurseFilterRequest) Descriptor() ([]byte, []int) {
	return file_nurse_attributes_proto_rawDescGZIP(), []int{7}
}

func (x *GetNurseFilterRequest) GetCenterId() int64 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

func (x *GetNurseFilterRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_nurse_attributes_proto protoreflect.FileDescriptor

var file_nurse_attributes_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xe5, 0x03, 0x0a,
	0x05, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x6f, 0x62, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x61, 0x6e,
	0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x6c, 0x6c, 0x5f,
	0x6e, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6c, 0x6c, 0x4e, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x6c,
	0x5f, 0x6e, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x54, 0x65, 0x6c, 0x6c, 0x4e, 0x6f, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x22, 0xf4, 0x03, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6e, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6c, 0x6c, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x77,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x13,
	0x69, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x54, 0x65, 0x6c,
	0x6c, 0x4e, 0x6f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x8c, 0x04, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x6f, 0x62, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x61,
	0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x63, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x6c, 0x6c,
	0x5f, 0x6e, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6c, 0x6c, 0x4e,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x77, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x77, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x69, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x2d, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x69, 0x73, 0x54, 0x65, 0x6c, 0x6c, 0x4e, 0x6f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x06, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x6e, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nurse_attributes_proto_rawDescOnce sync.Once
	file_nurse_attributes_proto_rawDescData = file_nurse_attributes_proto_rawDesc
)

func file_nurse_attributes_proto_rawDescGZIP() []byte {
	file_nurse_attributes_proto_rawDescOnce.Do(func() {
		file_nurse_attributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_nurse_attributes_proto_rawDescData)
	})
	return file_nurse_attributes_proto_rawDescData
}

var file_nurse_attributes_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_nurse_attributes_proto_goTypes = []interface{}{
	(*Nurse)(nil),                 // 0: pb.Nurse
	(*RegisterNurseRequest)(nil),  // 1: pb.RegisterNurseRequest
	(*UpdateNurseRequest)(nil),    // 2: pb.UpdateNurseRequest
	(*UpdateNurseResponse)(nil),   // 3: pb.UpdateNurseResponse
	(*RegisterNurseResponse)(nil), // 4: pb.RegisterNurseResponse
	(*GetAllNursesRequest)(nil),   // 5: pb.GetAllNursesRequest
	(*GetAllNursesResponse)(nil),  // 6: pb.GetAllNursesResponse
	(*GetNurseFilterRequest)(nil), // 7: pb.GetNurseFilterRequest
}
var file_nurse_attributes_proto_depIdxs = []int32{
	0, // 0: pb.GetAllNursesResponse.nurses:type_name -> pb.Nurse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nurse_attributes_proto_init() }
func file_nurse_attributes_proto_init() {
	if File_nurse_attributes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nurse_attributes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nurse); i {
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
		file_nurse_attributes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNurseRequest); i {
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
		file_nurse_attributes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNurseRequest); i {
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
		file_nurse_attributes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNurseResponse); i {
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
		file_nurse_attributes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNurseResponse); i {
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
		file_nurse_attributes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNursesRequest); i {
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
		file_nurse_attributes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNursesResponse); i {
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
		file_nurse_attributes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNurseFilterRequest); i {
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
			RawDescriptor: file_nurse_attributes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nurse_attributes_proto_goTypes,
		DependencyIndexes: file_nurse_attributes_proto_depIdxs,
		MessageInfos:      file_nurse_attributes_proto_msgTypes,
	}.Build()
	File_nurse_attributes_proto = out.File
	file_nurse_attributes_proto_rawDesc = nil
	file_nurse_attributes_proto_goTypes = nil
	file_nurse_attributes_proto_depIdxs = nil
}