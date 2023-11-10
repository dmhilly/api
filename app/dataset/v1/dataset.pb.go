// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: app/dataset/v1/dataset.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Dataset stores the metadata of a dataset.
type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId string                 `protobuf:"bytes,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	TimeCreated    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time_created,json=timeCreated,proto3" json:"time_created,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *Dataset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dataset) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Dataset) GetTimeCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCreated
	}
	return nil
}

// CreateDatasetRequest defines the name and organization ID of a dataset.
type CreateDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId string `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *CreateDatasetRequest) Reset() {
	*x = CreateDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatasetRequest) ProtoMessage() {}

func (x *CreateDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatasetRequest.ProtoReflect.Descriptor instead.
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDatasetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDatasetRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

// CreateDatasetResponse returns the dataset ID of the created dataset.
type CreateDatasetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateDatasetResponse) Reset() {
	*x = CreateDatasetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatasetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatasetResponse) ProtoMessage() {}

func (x *CreateDatasetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatasetResponse.ProtoReflect.Descriptor instead.
func (*CreateDatasetResponse) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDatasetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteDatasetRequest deletes the dataset specified by the dataset ID.
type DeleteDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDatasetRequest) Reset() {
	*x = DeleteDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatasetRequest) ProtoMessage() {}

func (x *DeleteDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatasetRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatasetRequest) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDatasetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDatasetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDatasetResponse) Reset() {
	*x = DeleteDatasetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatasetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatasetResponse) ProtoMessage() {}

func (x *DeleteDatasetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatasetResponse.ProtoReflect.Descriptor instead.
func (*DeleteDatasetResponse) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{4}
}

// RenameDatasetRequest applies the new name to the dataset specified by the dataset ID.
type RenameDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RenameDatasetRequest) Reset() {
	*x = RenameDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameDatasetRequest) ProtoMessage() {}

func (x *RenameDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameDatasetRequest.ProtoReflect.Descriptor instead.
func (*RenameDatasetRequest) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{5}
}

func (x *RenameDatasetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RenameDatasetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RenameDatasetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenameDatasetResponse) Reset() {
	*x = RenameDatasetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameDatasetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameDatasetResponse) ProtoMessage() {}

func (x *RenameDatasetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameDatasetResponse.ProtoReflect.Descriptor instead.
func (*RenameDatasetResponse) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{6}
}

// ListDatasetsByOrganizationIDRequest requests all of the datasets for an organization.
type ListDatasetsByOrganizationIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *ListDatasetsByOrganizationIDRequest) Reset() {
	*x = ListDatasetsByOrganizationIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetsByOrganizationIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetsByOrganizationIDRequest) ProtoMessage() {}

func (x *ListDatasetsByOrganizationIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetsByOrganizationIDRequest.ProtoReflect.Descriptor instead.
func (*ListDatasetsByOrganizationIDRequest) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{7}
}

func (x *ListDatasetsByOrganizationIDRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

// ListDatasetsByOrganizationIDResponse returns all the dataset metadata for the organization.
type ListDatasetsByOrganizationIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets,proto3" json:"datasets,omitempty"`
}

func (x *ListDatasetsByOrganizationIDResponse) Reset() {
	*x = ListDatasetsByOrganizationIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetsByOrganizationIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetsByOrganizationIDResponse) ProtoMessage() {}

func (x *ListDatasetsByOrganizationIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetsByOrganizationIDResponse.ProtoReflect.Descriptor instead.
func (*ListDatasetsByOrganizationIDResponse) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{8}
}

func (x *ListDatasetsByOrganizationIDResponse) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

// ListDatasetsByIDsRequest requests all of the datasets by their dataset IDs.
type ListDatasetsByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ListDatasetsByIDsRequest) Reset() {
	*x = ListDatasetsByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetsByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetsByIDsRequest) ProtoMessage() {}

func (x *ListDatasetsByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetsByIDsRequest.ProtoReflect.Descriptor instead.
func (*ListDatasetsByIDsRequest) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{9}
}

func (x *ListDatasetsByIDsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// ListDatasetsByIDsResponse returns all the dataset metadata for the associated dataset IDs.
type ListDatasetsByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets,proto3" json:"datasets,omitempty"`
}

func (x *ListDatasetsByIDsResponse) Reset() {
	*x = ListDatasetsByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dataset_v1_dataset_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetsByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetsByIDsResponse) ProtoMessage() {}

func (x *ListDatasetsByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_dataset_v1_dataset_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetsByIDsResponse.ProtoReflect.Descriptor instead.
func (*ListDatasetsByIDsResponse) Descriptor() ([]byte, []int) {
	return file_app_dataset_v1_dataset_proto_rawDescGZIP(), []int{10}
}

func (x *ListDatasetsByIDsResponse) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

var File_app_dataset_v1_dataset_proto protoreflect.FileDescriptor

var file_app_dataset_v1_dataset_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x53, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4e, 0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x60, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x22, 0x2c, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x55, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x32, 0xd2, 0x04, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x38, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x2d, 0x2e,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76,
	0x69, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e,
	0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_dataset_v1_dataset_proto_rawDescOnce sync.Once
	file_app_dataset_v1_dataset_proto_rawDescData = file_app_dataset_v1_dataset_proto_rawDesc
)

func file_app_dataset_v1_dataset_proto_rawDescGZIP() []byte {
	file_app_dataset_v1_dataset_proto_rawDescOnce.Do(func() {
		file_app_dataset_v1_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_dataset_v1_dataset_proto_rawDescData)
	})
	return file_app_dataset_v1_dataset_proto_rawDescData
}

var file_app_dataset_v1_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_app_dataset_v1_dataset_proto_goTypes = []interface{}{
	(*Dataset)(nil),                              // 0: viam.app.dataset.v1.Dataset
	(*CreateDatasetRequest)(nil),                 // 1: viam.app.dataset.v1.CreateDatasetRequest
	(*CreateDatasetResponse)(nil),                // 2: viam.app.dataset.v1.CreateDatasetResponse
	(*DeleteDatasetRequest)(nil),                 // 3: viam.app.dataset.v1.DeleteDatasetRequest
	(*DeleteDatasetResponse)(nil),                // 4: viam.app.dataset.v1.DeleteDatasetResponse
	(*RenameDatasetRequest)(nil),                 // 5: viam.app.dataset.v1.RenameDatasetRequest
	(*RenameDatasetResponse)(nil),                // 6: viam.app.dataset.v1.RenameDatasetResponse
	(*ListDatasetsByOrganizationIDRequest)(nil),  // 7: viam.app.dataset.v1.ListDatasetsByOrganizationIDRequest
	(*ListDatasetsByOrganizationIDResponse)(nil), // 8: viam.app.dataset.v1.ListDatasetsByOrganizationIDResponse
	(*ListDatasetsByIDsRequest)(nil),             // 9: viam.app.dataset.v1.ListDatasetsByIDsRequest
	(*ListDatasetsByIDsResponse)(nil),            // 10: viam.app.dataset.v1.ListDatasetsByIDsResponse
	(*timestamppb.Timestamp)(nil),                // 11: google.protobuf.Timestamp
}
var file_app_dataset_v1_dataset_proto_depIdxs = []int32{
	11, // 0: viam.app.dataset.v1.Dataset.time_created:type_name -> google.protobuf.Timestamp
	0,  // 1: viam.app.dataset.v1.ListDatasetsByOrganizationIDResponse.datasets:type_name -> viam.app.dataset.v1.Dataset
	0,  // 2: viam.app.dataset.v1.ListDatasetsByIDsResponse.datasets:type_name -> viam.app.dataset.v1.Dataset
	1,  // 3: viam.app.dataset.v1.DatasetService.CreateDataset:input_type -> viam.app.dataset.v1.CreateDatasetRequest
	3,  // 4: viam.app.dataset.v1.DatasetService.DeleteDataset:input_type -> viam.app.dataset.v1.DeleteDatasetRequest
	5,  // 5: viam.app.dataset.v1.DatasetService.RenameDataset:input_type -> viam.app.dataset.v1.RenameDatasetRequest
	7,  // 6: viam.app.dataset.v1.DatasetService.ListDatasetsByOrganizationID:input_type -> viam.app.dataset.v1.ListDatasetsByOrganizationIDRequest
	9,  // 7: viam.app.dataset.v1.DatasetService.ListDatasetsByIDs:input_type -> viam.app.dataset.v1.ListDatasetsByIDsRequest
	2,  // 8: viam.app.dataset.v1.DatasetService.CreateDataset:output_type -> viam.app.dataset.v1.CreateDatasetResponse
	4,  // 9: viam.app.dataset.v1.DatasetService.DeleteDataset:output_type -> viam.app.dataset.v1.DeleteDatasetResponse
	6,  // 10: viam.app.dataset.v1.DatasetService.RenameDataset:output_type -> viam.app.dataset.v1.RenameDatasetResponse
	8,  // 11: viam.app.dataset.v1.DatasetService.ListDatasetsByOrganizationID:output_type -> viam.app.dataset.v1.ListDatasetsByOrganizationIDResponse
	10, // 12: viam.app.dataset.v1.DatasetService.ListDatasetsByIDs:output_type -> viam.app.dataset.v1.ListDatasetsByIDsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_app_dataset_v1_dataset_proto_init() }
func file_app_dataset_v1_dataset_proto_init() {
	if File_app_dataset_v1_dataset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_dataset_v1_dataset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatasetRequest); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatasetResponse); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatasetRequest); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatasetResponse); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameDatasetRequest); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameDatasetResponse); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetsByOrganizationIDRequest); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetsByOrganizationIDResponse); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetsByIDsRequest); i {
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
		file_app_dataset_v1_dataset_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetsByIDsResponse); i {
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
			RawDescriptor: file_app_dataset_v1_dataset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_dataset_v1_dataset_proto_goTypes,
		DependencyIndexes: file_app_dataset_v1_dataset_proto_depIdxs,
		MessageInfos:      file_app_dataset_v1_dataset_proto_msgTypes,
	}.Build()
	File_app_dataset_v1_dataset_proto = out.File
	file_app_dataset_v1_dataset_proto_rawDesc = nil
	file_app_dataset_v1_dataset_proto_goTypes = nil
	file_app_dataset_v1_dataset_proto_depIdxs = nil
}
