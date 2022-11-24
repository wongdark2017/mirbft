// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.12.4
// source: request.proto

package protobufs

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

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *RequestID `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Pubkey    []byte     `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequest) GetRequestId() *RequestID {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *ClientRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ClientRequest) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *ClientRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientSn int32 `protobuf:"varint,1,opt,name=client_sn,json=clientSn,proto3" json:"client_sn,omitempty"`
	OrderSn  int32 `protobuf:"varint,2,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *ClientResponse) GetClientSn() int32 {
	if x != nil {
		return x.ClientSn
	}
	return 0
}

func (x *ClientResponse) GetOrderSn() int32 {
	if x != nil {
		return x.OrderSn
	}
	return 0
}

type RequestID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSn int32 `protobuf:"varint,2,opt,name=client_sn,json=clientSn,proto3" json:"client_sn,omitempty"`
}

func (x *RequestID) Reset() {
	*x = RequestID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestID) ProtoMessage() {}

func (x *RequestID) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestID.ProtoReflect.Descriptor instead.
func (*RequestID) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

func (x *RequestID) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *RequestID) GetClientSn() int32 {
	if x != nil {
		return x.ClientSn
	}
	return 0
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*ClientRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{3}
}

func (x *Batch) GetRequests() []*ClientRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type MissingEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn             int32 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	PayloadRequest bool  `protobuf:"varint,2,opt,name=payload_request,json=payloadRequest,proto3" json:"payload_request,omitempty"`
}

func (x *MissingEntryRequest) Reset() {
	*x = MissingEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissingEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissingEntryRequest) ProtoMessage() {}

func (x *MissingEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissingEntryRequest.ProtoReflect.Descriptor instead.
func (*MissingEntryRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{4}
}

func (x *MissingEntryRequest) GetSn() int32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *MissingEntryRequest) GetPayloadRequest() bool {
	if x != nil {
		return x.PayloadRequest
	}
	return false
}

type MissingEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn      int32  `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Batch   *Batch `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
	Digest  []byte `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Aborted bool   `protobuf:"varint,4,opt,name=aborted,proto3" json:"aborted,omitempty"`
	Suspect int32  `protobuf:"varint,5,opt,name=suspect,proto3" json:"suspect,omitempty"`
	Proof   string `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"` // This is a dummy proof. TODO: Implement a proper one, like a Merkle proof wrt. the checkpoint.
}

func (x *MissingEntry) Reset() {
	*x = MissingEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissingEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissingEntry) ProtoMessage() {}

func (x *MissingEntry) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissingEntry.ProtoReflect.Descriptor instead.
func (*MissingEntry) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{5}
}

func (x *MissingEntry) GetSn() int32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *MissingEntry) GetBatch() *Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

func (x *MissingEntry) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *MissingEntry) GetAborted() bool {
	if x != nil {
		return x.Aborted
	}
	return false
}

func (x *MissingEntry) GetSuspect() int32 {
	if x != nil {
		return x.Suspect
	}
	return 0
}

func (x *MissingEntry) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

type BucketSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *BucketSubscription) Reset() {
	*x = BucketSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketSubscription) ProtoMessage() {}

func (x *BucketSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketSubscription.ProtoReflect.Descriptor instead.
func (*BucketSubscription) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{6}
}

func (x *BucketSubscription) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type BucketAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch   int32                  `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Buckets map[int32]*ListOfInt32 `protobuf:"bytes,2,rep,name=buckets,proto3" json:"buckets,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BucketAssignment) Reset() {
	*x = BucketAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketAssignment) ProtoMessage() {}

func (x *BucketAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketAssignment.ProtoReflect.Descriptor instead.
func (*BucketAssignment) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{7}
}

func (x *BucketAssignment) GetEpoch() int32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *BucketAssignment) GetBuckets() map[int32]*ListOfInt32 {
	if x != nil {
		return x.Buckets
	}
	return nil
}

type ListOfInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []int32 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *ListOfInt32) Reset() {
	*x = ListOfInt32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfInt32) ProtoMessage() {}

func (x *ListOfInt32) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfInt32.ProtoReflect.Descriptor instead.
func (*ListOfInt32) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{8}
}

func (x *ListOfInt32) GetVals() []int32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x48, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x22, 0x45, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x6e, 0x22, 0x3d, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x34, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x22, 0x4e, 0x0a, 0x13, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x73, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x73, 0x70, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x75, 0x73, 0x70, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x31, 0x0a, 0x12,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0xc0, 0x01, 0x0a, 0x10, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x42, 0x0a, 0x07, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x52,
	0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x04, 0x76, 0x61, 0x6c, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_request_proto_goTypes = []interface{}{
	(*ClientRequest)(nil),       // 0: protobufs.ClientRequest
	(*ClientResponse)(nil),      // 1: protobufs.ClientResponse
	(*RequestID)(nil),           // 2: protobufs.RequestID
	(*Batch)(nil),               // 3: protobufs.Batch
	(*MissingEntryRequest)(nil), // 4: protobufs.MissingEntryRequest
	(*MissingEntry)(nil),        // 5: protobufs.MissingEntry
	(*BucketSubscription)(nil),  // 6: protobufs.BucketSubscription
	(*BucketAssignment)(nil),    // 7: protobufs.BucketAssignment
	(*ListOfInt32)(nil),         // 8: protobufs.ListOfInt32
	nil,                         // 9: protobufs.BucketAssignment.BucketsEntry
}
var file_request_proto_depIdxs = []int32{
	2, // 0: protobufs.ClientRequest.request_id:type_name -> protobufs.RequestID
	0, // 1: protobufs.Batch.requests:type_name -> protobufs.ClientRequest
	3, // 2: protobufs.MissingEntry.batch:type_name -> protobufs.Batch
	9, // 3: protobufs.BucketAssignment.buckets:type_name -> protobufs.BucketAssignment.BucketsEntry
	8, // 4: protobufs.BucketAssignment.BucketsEntry.value:type_name -> protobufs.ListOfInt32
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestID); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
		file_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissingEntryRequest); i {
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
		file_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissingEntry); i {
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
		file_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketSubscription); i {
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
		file_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketAssignment); i {
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
		file_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfInt32); i {
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
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
