// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ditto_v1/ditto.proto

package ditto_v1

import (
	core_v1 "github.com/kutty-kumar/ho_oh_go/core_v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PrinterDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Status        core_v1.Status         `protobuf:"varint,5,opt,name=status,proto3,enum=core_v1.Status" json:"status,omitempty"`
	ExternalId    string                 `protobuf:"bytes,6,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	ProductNumber string                 `protobuf:"bytes,7,opt,name=product_number,json=productNumber,proto3" json:"product_number,omitempty"`
	FromDate      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate        *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	FromIndex     uint64                 `protobuf:"varint,10,opt,name=from_index,json=fromIndex,proto3" json:"from_index,omitempty"`
	ToIndex       uint64                 `protobuf:"varint,11,opt,name=to_index,json=toIndex,proto3" json:"to_index,omitempty"`
}

func (x *PrinterDto) Reset() {
	*x = PrinterDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrinterDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterDto) ProtoMessage() {}

func (x *PrinterDto) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterDto.ProtoReflect.Descriptor instead.
func (*PrinterDto) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{0}
}

func (x *PrinterDto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PrinterDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrinterDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PrinterDto) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *PrinterDto) GetStatus() core_v1.Status {
	if x != nil {
		return x.Status
	}
	return core_v1.Status_unknown_status
}

func (x *PrinterDto) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *PrinterDto) GetProductNumber() string {
	if x != nil {
		return x.ProductNumber
	}
	return ""
}

func (x *PrinterDto) GetFromDate() *timestamppb.Timestamp {
	if x != nil {
		return x.FromDate
	}
	return nil
}

func (x *PrinterDto) GetToDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ToDate
	}
	return nil
}

func (x *PrinterDto) GetFromIndex() uint64 {
	if x != nil {
		return x.FromIndex
	}
	return 0
}

func (x *PrinterDto) GetToIndex() uint64 {
	if x != nil {
		return x.ToIndex
	}
	return 0
}

type CreatePrinterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *PrinterDto `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *CreatePrinterRequest) Reset() {
	*x = CreatePrinterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrinterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrinterRequest) ProtoMessage() {}

func (x *CreatePrinterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrinterRequest.ProtoReflect.Descriptor instead.
func (*CreatePrinterRequest) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePrinterRequest) GetRequest() *PrinterDto {
	if x != nil {
		return x.Request
	}
	return nil
}

type CreatePrinterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *PrinterDto `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreatePrinterResponse) Reset() {
	*x = CreatePrinterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrinterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrinterResponse) ProtoMessage() {}

func (x *CreatePrinterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrinterResponse.ProtoReflect.Descriptor instead.
func (*CreatePrinterResponse) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePrinterResponse) GetResponse() *PrinterDto {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdatePrinterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrinterId string      `protobuf:"bytes,1,opt,name=printer_id,json=printerId,proto3" json:"printer_id,omitempty"`
	Request   *PrinterDto `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *UpdatePrinterRequest) Reset() {
	*x = UpdatePrinterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrinterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrinterRequest) ProtoMessage() {}

func (x *UpdatePrinterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrinterRequest.ProtoReflect.Descriptor instead.
func (*UpdatePrinterRequest) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePrinterRequest) GetPrinterId() string {
	if x != nil {
		return x.PrinterId
	}
	return ""
}

func (x *UpdatePrinterRequest) GetRequest() *PrinterDto {
	if x != nil {
		return x.Request
	}
	return nil
}

type UpdatePrinterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *PrinterDto `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdatePrinterResponse) Reset() {
	*x = UpdatePrinterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrinterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrinterResponse) ProtoMessage() {}

func (x *UpdatePrinterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrinterResponse.ProtoReflect.Descriptor instead.
func (*UpdatePrinterResponse) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePrinterResponse) GetResponse() *PrinterDto {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetPrinterByExternalIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrinterId string `protobuf:"bytes,1,opt,name=printer_id,json=printerId,proto3" json:"printer_id,omitempty"`
}

func (x *GetPrinterByExternalIdRequest) Reset() {
	*x = GetPrinterByExternalIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrinterByExternalIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrinterByExternalIdRequest) ProtoMessage() {}

func (x *GetPrinterByExternalIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrinterByExternalIdRequest.ProtoReflect.Descriptor instead.
func (*GetPrinterByExternalIdRequest) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{5}
}

func (x *GetPrinterByExternalIdRequest) GetPrinterId() string {
	if x != nil {
		return x.PrinterId
	}
	return ""
}

type GetPrinterByExternalIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *PrinterDto `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetPrinterByExternalIdResponse) Reset() {
	*x = GetPrinterByExternalIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrinterByExternalIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrinterByExternalIdResponse) ProtoMessage() {}

func (x *GetPrinterByExternalIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrinterByExternalIdResponse.ProtoReflect.Descriptor instead.
func (*GetPrinterByExternalIdResponse) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{6}
}

func (x *GetPrinterByExternalIdResponse) GetResponse() *PrinterDto {
	if x != nil {
		return x.Response
	}
	return nil
}

type MultiGetPrintersByExternalIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrinterIds []string `protobuf:"bytes,1,rep,name=printer_ids,json=printerIds,proto3" json:"printer_ids,omitempty"`
}

func (x *MultiGetPrintersByExternalIdRequest) Reset() {
	*x = MultiGetPrintersByExternalIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiGetPrintersByExternalIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiGetPrintersByExternalIdRequest) ProtoMessage() {}

func (x *MultiGetPrintersByExternalIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiGetPrintersByExternalIdRequest.ProtoReflect.Descriptor instead.
func (*MultiGetPrintersByExternalIdRequest) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{7}
}

func (x *MultiGetPrintersByExternalIdRequest) GetPrinterIds() []string {
	if x != nil {
		return x.PrinterIds
	}
	return nil
}

type MultiGetPrintersByExternalIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*PrinterDto `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *MultiGetPrintersByExternalIdResponse) Reset() {
	*x = MultiGetPrintersByExternalIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiGetPrintersByExternalIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiGetPrintersByExternalIdResponse) ProtoMessage() {}

func (x *MultiGetPrintersByExternalIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiGetPrintersByExternalIdResponse.ProtoReflect.Descriptor instead.
func (*MultiGetPrintersByExternalIdResponse) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{8}
}

func (x *MultiGetPrintersByExternalIdResponse) GetResult() []*PrinterDto {
	if x != nil {
		return x.Result
	}
	return nil
}

type NoOpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoOpRequest) Reset() {
	*x = NoOpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoOpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoOpRequest) ProtoMessage() {}

func (x *NoOpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoOpRequest.ProtoReflect.Descriptor instead.
func (*NoOpRequest) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{9}
}

type DeletePrinterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrinterId string `protobuf:"bytes,1,opt,name=printer_id,json=printerId,proto3" json:"printer_id,omitempty"`
}

func (x *DeletePrinterRequest) Reset() {
	*x = DeletePrinterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ditto_v1_ditto_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePrinterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePrinterRequest) ProtoMessage() {}

func (x *DeletePrinterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ditto_v1_ditto_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePrinterRequest.ProtoReflect.Descriptor instead.
func (*DeletePrinterRequest) Descriptor() ([]byte, []int) {
	return file_ditto_v1_ditto_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePrinterRequest) GetPrinterId() string {
	if x != nil {
		return x.PrinterId
	}
	return ""
}

var File_ditto_v1_ditto_proto protoreflect.FileDescriptor

var file_ditto_v1_ditto_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x74, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31,
	0x1a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44,
	0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x33, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x74,
	0x6f, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x46, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f,
	0x5f, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x65, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x74,
	0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f,
	0x5f, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x23, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x22, 0x54, 0x0a, 0x24, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f,
	0x5f, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x6f, 0x4f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x32, 0xb3, 0x06, 0x0a,
	0x0e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x7c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x8e,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x27, 0x2e, 0x64, 0x69, 0x74, 0x74,
	0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x42,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0xa0, 0x01, 0x0a, 0x1c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x2d, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x1a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2d, 0x67, 0x65, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x4f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x1a, 0x1b, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2d, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1e,
	0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x75, 0x74, 0x74, 0x79, 0x2d, 0x6b, 0x75, 0x6d, 0x61, 0x72, 0x2f, 0x68, 0x6f, 0x5f,
	0x6f, 0x68, 0x5f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x3b, 0x64,
	0x69, 0x74, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ditto_v1_ditto_proto_rawDescOnce sync.Once
	file_ditto_v1_ditto_proto_rawDescData = file_ditto_v1_ditto_proto_rawDesc
)

func file_ditto_v1_ditto_proto_rawDescGZIP() []byte {
	file_ditto_v1_ditto_proto_rawDescOnce.Do(func() {
		file_ditto_v1_ditto_proto_rawDescData = protoimpl.X.CompressGZIP(file_ditto_v1_ditto_proto_rawDescData)
	})
	return file_ditto_v1_ditto_proto_rawDescData
}

var file_ditto_v1_ditto_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ditto_v1_ditto_proto_goTypes = []interface{}{
	(*PrinterDto)(nil),                           // 0: ditto_v1.PrinterDto
	(*CreatePrinterRequest)(nil),                 // 1: ditto_v1.CreatePrinterRequest
	(*CreatePrinterResponse)(nil),                // 2: ditto_v1.CreatePrinterResponse
	(*UpdatePrinterRequest)(nil),                 // 3: ditto_v1.UpdatePrinterRequest
	(*UpdatePrinterResponse)(nil),                // 4: ditto_v1.UpdatePrinterResponse
	(*GetPrinterByExternalIdRequest)(nil),        // 5: ditto_v1.GetPrinterByExternalIdRequest
	(*GetPrinterByExternalIdResponse)(nil),       // 6: ditto_v1.GetPrinterByExternalIdResponse
	(*MultiGetPrintersByExternalIdRequest)(nil),  // 7: ditto_v1.MultiGetPrintersByExternalIdRequest
	(*MultiGetPrintersByExternalIdResponse)(nil), // 8: ditto_v1.MultiGetPrintersByExternalIdResponse
	(*NoOpRequest)(nil),                          // 9: ditto_v1.NoOpRequest
	(*DeletePrinterRequest)(nil),                 // 10: ditto_v1.DeletePrinterRequest
	(core_v1.Status)(0),                          // 11: core_v1.Status
	(*timestamppb.Timestamp)(nil),                // 12: google.protobuf.Timestamp
}
var file_ditto_v1_ditto_proto_depIdxs = []int32{
	11, // 0: ditto_v1.PrinterDto.status:type_name -> core_v1.Status
	12, // 1: ditto_v1.PrinterDto.from_date:type_name -> google.protobuf.Timestamp
	12, // 2: ditto_v1.PrinterDto.to_date:type_name -> google.protobuf.Timestamp
	0,  // 3: ditto_v1.CreatePrinterRequest.request:type_name -> ditto_v1.PrinterDto
	0,  // 4: ditto_v1.CreatePrinterResponse.response:type_name -> ditto_v1.PrinterDto
	0,  // 5: ditto_v1.UpdatePrinterRequest.request:type_name -> ditto_v1.PrinterDto
	0,  // 6: ditto_v1.UpdatePrinterResponse.response:type_name -> ditto_v1.PrinterDto
	0,  // 7: ditto_v1.GetPrinterByExternalIdResponse.response:type_name -> ditto_v1.PrinterDto
	0,  // 8: ditto_v1.MultiGetPrintersByExternalIdResponse.result:type_name -> ditto_v1.PrinterDto
	1,  // 9: ditto_v1.PrinterService.CreatePrinter:input_type -> ditto_v1.CreatePrinterRequest
	3,  // 10: ditto_v1.PrinterService.UpdatePrinter:input_type -> ditto_v1.UpdatePrinterRequest
	5,  // 11: ditto_v1.PrinterService.GetPrinterByExternalId:input_type -> ditto_v1.GetPrinterByExternalIdRequest
	7,  // 12: ditto_v1.PrinterService.MultiGetPrintersByExternalId:input_type -> ditto_v1.MultiGetPrintersByExternalIdRequest
	9,  // 13: ditto_v1.PrinterService.MultiGetPrintersForUser:input_type -> ditto_v1.NoOpRequest
	10, // 14: ditto_v1.PrinterService.DeletePrinter:input_type -> ditto_v1.DeletePrinterRequest
	2,  // 15: ditto_v1.PrinterService.CreatePrinter:output_type -> ditto_v1.CreatePrinterResponse
	4,  // 16: ditto_v1.PrinterService.UpdatePrinter:output_type -> ditto_v1.UpdatePrinterResponse
	6,  // 17: ditto_v1.PrinterService.GetPrinterByExternalId:output_type -> ditto_v1.GetPrinterByExternalIdResponse
	8,  // 18: ditto_v1.PrinterService.MultiGetPrintersByExternalId:output_type -> ditto_v1.MultiGetPrintersByExternalIdResponse
	8,  // 19: ditto_v1.PrinterService.MultiGetPrintersForUser:output_type -> ditto_v1.MultiGetPrintersByExternalIdResponse
	4,  // 20: ditto_v1.PrinterService.DeletePrinter:output_type -> ditto_v1.UpdatePrinterResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_ditto_v1_ditto_proto_init() }
func file_ditto_v1_ditto_proto_init() {
	if File_ditto_v1_ditto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ditto_v1_ditto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrinterDto); i {
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
		file_ditto_v1_ditto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePrinterRequest); i {
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
		file_ditto_v1_ditto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePrinterResponse); i {
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
		file_ditto_v1_ditto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrinterRequest); i {
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
		file_ditto_v1_ditto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrinterResponse); i {
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
		file_ditto_v1_ditto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrinterByExternalIdRequest); i {
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
		file_ditto_v1_ditto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrinterByExternalIdResponse); i {
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
		file_ditto_v1_ditto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiGetPrintersByExternalIdRequest); i {
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
		file_ditto_v1_ditto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiGetPrintersByExternalIdResponse); i {
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
		file_ditto_v1_ditto_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoOpRequest); i {
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
		file_ditto_v1_ditto_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePrinterRequest); i {
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
			RawDescriptor: file_ditto_v1_ditto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ditto_v1_ditto_proto_goTypes,
		DependencyIndexes: file_ditto_v1_ditto_proto_depIdxs,
		MessageInfos:      file_ditto_v1_ditto_proto_msgTypes,
	}.Build()
	File_ditto_v1_ditto_proto = out.File
	file_ditto_v1_ditto_proto_rawDesc = nil
	file_ditto_v1_ditto_proto_goTypes = nil
	file_ditto_v1_ditto_proto_depIdxs = nil
}
