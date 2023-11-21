// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/index.proto

package authzv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartIndexJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartIndexJobRequest) Reset() {
	*x = StartIndexJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartIndexJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartIndexJobRequest) ProtoMessage() {}

func (x *StartIndexJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartIndexJobRequest.ProtoReflect.Descriptor instead.
func (*StartIndexJobRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP(), []int{0}
}

type StartIndexJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *StartIndexJobResponse) Reset() {
	*x = StartIndexJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartIndexJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartIndexJobResponse) ProtoMessage() {}

func (x *StartIndexJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartIndexJobResponse.ProtoReflect.Descriptor instead.
func (*StartIndexJobResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP(), []int{1}
}

func (x *StartIndexJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type LookupResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Universe    string `protobuf:"bytes,1,opt,name=universe,proto3" json:"universe,omitempty"`
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// The principal to look up accessible resources for.
	Principal *UID `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`
	// Filter entities for a particular type
	ResourceType string `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The token for the next page.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *LookupResourcesRequest) Reset() {
	*x = LookupResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResourcesRequest) ProtoMessage() {}

func (x *LookupResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResourcesRequest.ProtoReflect.Descriptor instead.
func (*LookupResourcesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP(), []int{2}
}

func (x *LookupResourcesRequest) GetUniverse() string {
	if x != nil {
		return x.Universe
	}
	return ""
}

func (x *LookupResourcesRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *LookupResourcesRequest) GetPrincipal() *UID {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *LookupResourcesRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *LookupResourcesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ResourcePermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Entity `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// the actions that are allowed on the entity.
	Actions []*AllowedAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *ResourcePermission) Reset() {
	*x = ResourcePermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcePermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePermission) ProtoMessage() {}

func (x *ResourcePermission) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePermission.ProtoReflect.Descriptor instead.
func (*ResourcePermission) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP(), []int{3}
}

func (x *ResourcePermission) GetResource() *Entity {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourcePermission) GetActions() []*AllowedAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type AllowedAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *UID `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *AllowedAction) Reset() {
	*x = AllowedAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedAction) ProtoMessage() {}

func (x *AllowedAction) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedAction.ProtoReflect.Descriptor instead.
func (*AllowedAction) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP(), []int{4}
}

func (x *AllowedAction) GetAction() *UID {
	if x != nil {
		return x.Action
	}
	return nil
}

type LookupResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources     []*ResourcePermission `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	NextPageToken string                `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *LookupResourcesResponse) Reset() {
	*x = LookupResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResourcesResponse) ProtoMessage() {}

func (x *LookupResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_index_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResourcesResponse.ProtoReflect.Descriptor instead.
func (*LookupResourcesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP(), []int{5}
}

func (x *LookupResourcesResponse) GetResources() []*ResourcePermission {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *LookupResourcesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_commonfate_authz_v1alpha1_index_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_index_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2e, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x22, 0xd8, 0x01, 0x0a, 0x16, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x09, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x42, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8e,
	0x01, 0x0a, 0x17, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32,
	0x80, 0x02, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x74, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4a, 0x6f,
	0x62, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0f, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0xfa, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x69, 0x65, 0x6d, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41,
	0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_authz_v1alpha1_index_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_index_proto_rawDescData = file_commonfate_authz_v1alpha1_index_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_index_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_index_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_index_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_index_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commonfate_authz_v1alpha1_index_proto_goTypes = []interface{}{
	(*StartIndexJobRequest)(nil),    // 0: commonfate.authz.v1alpha1.StartIndexJobRequest
	(*StartIndexJobResponse)(nil),   // 1: commonfate.authz.v1alpha1.StartIndexJobResponse
	(*LookupResourcesRequest)(nil),  // 2: commonfate.authz.v1alpha1.LookupResourcesRequest
	(*ResourcePermission)(nil),      // 3: commonfate.authz.v1alpha1.ResourcePermission
	(*AllowedAction)(nil),           // 4: commonfate.authz.v1alpha1.AllowedAction
	(*LookupResourcesResponse)(nil), // 5: commonfate.authz.v1alpha1.LookupResourcesResponse
	(*UID)(nil),                     // 6: commonfate.authz.v1alpha1.UID
	(*Entity)(nil),                  // 7: commonfate.authz.v1alpha1.Entity
}
var file_commonfate_authz_v1alpha1_index_proto_depIdxs = []int32{
	6, // 0: commonfate.authz.v1alpha1.LookupResourcesRequest.principal:type_name -> commonfate.authz.v1alpha1.UID
	7, // 1: commonfate.authz.v1alpha1.ResourcePermission.resource:type_name -> commonfate.authz.v1alpha1.Entity
	4, // 2: commonfate.authz.v1alpha1.ResourcePermission.actions:type_name -> commonfate.authz.v1alpha1.AllowedAction
	6, // 3: commonfate.authz.v1alpha1.AllowedAction.action:type_name -> commonfate.authz.v1alpha1.UID
	3, // 4: commonfate.authz.v1alpha1.LookupResourcesResponse.resources:type_name -> commonfate.authz.v1alpha1.ResourcePermission
	0, // 5: commonfate.authz.v1alpha1.IndexService.StartIndexJob:input_type -> commonfate.authz.v1alpha1.StartIndexJobRequest
	2, // 6: commonfate.authz.v1alpha1.IndexService.LookupResources:input_type -> commonfate.authz.v1alpha1.LookupResourcesRequest
	1, // 7: commonfate.authz.v1alpha1.IndexService.StartIndexJob:output_type -> commonfate.authz.v1alpha1.StartIndexJobResponse
	5, // 8: commonfate.authz.v1alpha1.IndexService.LookupResources:output_type -> commonfate.authz.v1alpha1.LookupResourcesResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_index_proto_init() }
func file_commonfate_authz_v1alpha1_index_proto_init() {
	if File_commonfate_authz_v1alpha1_index_proto != nil {
		return
	}
	file_commonfate_authz_v1alpha1_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartIndexJobRequest); i {
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
		file_commonfate_authz_v1alpha1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartIndexJobResponse); i {
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
		file_commonfate_authz_v1alpha1_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupResourcesRequest); i {
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
		file_commonfate_authz_v1alpha1_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcePermission); i {
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
		file_commonfate_authz_v1alpha1_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedAction); i {
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
		file_commonfate_authz_v1alpha1_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupResourcesResponse); i {
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
			RawDescriptor: file_commonfate_authz_v1alpha1_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_index_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_index_proto_depIdxs,
		MessageInfos:      file_commonfate_authz_v1alpha1_index_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_index_proto = out.File
	file_commonfate_authz_v1alpha1_index_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_index_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_index_proto_depIdxs = nil
}
