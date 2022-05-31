// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.4
// source: proto/p2p.proto

package p2p

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

type HelloMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Agent   string `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	Port    uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *HelloMessage) Reset() {
	*x = HelloMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMessage) ProtoMessage() {}

func (x *HelloMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMessage.ProtoReflect.Descriptor instead.
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{0}
}

func (x *HelloMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HelloMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HelloMessage) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *HelloMessage) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers  []string `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	Status string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PeersResponse) Reset() {
	*x = PeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersResponse) ProtoMessage() {}

func (x *PeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersResponse.ProtoReflect.Descriptor instead.
func (*PeersResponse) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{1}
}

func (x *PeersResponse) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *PeersResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PeersRequest) Reset() {
	*x = PeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersRequest) ProtoMessage() {}

func (x *PeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersRequest.ProtoReflect.Descriptor instead.
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{2}
}

type ObjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId []byte `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ObjectId) Reset() {
	*x = ObjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectId) ProtoMessage() {}

func (x *ObjectId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectId.ProtoReflect.Descriptor instead.
func (*ObjectId) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectId) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

func (x *ObjectId) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object []byte `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{4}
}

func (x *Object) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *Object) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Object) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{5}
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxIds   [][]byte `protobuf:"bytes,1,rep,name=txIds,proto3" json:"txIds,omitempty"`
	Nonce   []byte   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	PrevId  []byte   `protobuf:"bytes,3,opt,name=prevId,proto3" json:"prevId,omitempty"`
	Created uint64   `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	T       []byte   `protobuf:"bytes,5,opt,name=t,proto3" json:"t,omitempty"`
	Miner   string   `protobuf:"bytes,6,opt,name=miner,proto3" json:"miner,omitempty"`
	Note    string   `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{6}
}

func (x *Block) GetTxIds() [][]byte {
	if x != nil {
		return x.TxIds
	}
	return nil
}

func (x *Block) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Block) GetPrevId() []byte {
	if x != nil {
		return x.PrevId
	}
	return nil
}

func (x *Block) GetCreated() uint64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Block) GetT() []byte {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *Block) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *Block) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs  []*Input  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Height  int64     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{7}
}

func (x *Tx) GetInputs() []*Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Tx) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Tx) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type OutPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId  []byte `protobuf:"bytes,1,opt,name=txId,proto3" json:"txId,omitempty"`
	Index int64  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *OutPoint) Reset() {
	*x = OutPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPoint) ProtoMessage() {}

func (x *OutPoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPoint.ProtoReflect.Descriptor instead.
func (*OutPoint) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{8}
}

func (x *OutPoint) GetTxId() []byte {
	if x != nil {
		return x.TxId
	}
	return nil
}

func (x *OutPoint) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outpoint *OutPoint `protobuf:"bytes,1,opt,name=outpoint,proto3" json:"outpoint,omitempty"`
	Sig      []byte    `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{9}
}

func (x *Input) GetOutpoint() *OutPoint {
	if x != nil {
		return x.Outpoint
	}
	return nil
}

func (x *Input) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Value  int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_proto_p2p_proto_rawDescGZIP(), []int{10}
}

func (x *Output) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *Output) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_p2p_proto protoreflect.FileDescriptor

var file_proto_p2p_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x70, 0x32, 0x70, 0x22, 0x6a, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a,
	0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x78, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x74,
	0x78, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x76, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x72, 0x65, 0x76,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69,
	0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x22, 0x67, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x22, 0x0a, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x32, 0x70,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x25,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x34, 0x0a,
	0x08, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x44, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x08,
	0x6f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x6f,
	0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x36, 0x0a, 0x06, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xc6, 0x01, 0x0a, 0x03, 0x50, 0x32, 0x50, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x32, 0x70, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x32, 0x70,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0d, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a,
	0x0b, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x2a,
	0x0a, 0x0b, 0x49, 0x48, 0x61, 0x76, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0d, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70,
	0x32, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x70, 0x32, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_p2p_proto_rawDescOnce sync.Once
	file_proto_p2p_proto_rawDescData = file_proto_p2p_proto_rawDesc
)

func file_proto_p2p_proto_rawDescGZIP() []byte {
	file_proto_p2p_proto_rawDescOnce.Do(func() {
		file_proto_p2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_p2p_proto_rawDescData)
	})
	return file_proto_p2p_proto_rawDescData
}

var file_proto_p2p_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_p2p_proto_goTypes = []interface{}{
	(*HelloMessage)(nil),  // 0: p2p.HelloMessage
	(*PeersResponse)(nil), // 1: p2p.PeersResponse
	(*PeersRequest)(nil),  // 2: p2p.PeersRequest
	(*ObjectId)(nil),      // 3: p2p.ObjectId
	(*Object)(nil),        // 4: p2p.Object
	(*Empty)(nil),         // 5: p2p.Empty
	(*Block)(nil),         // 6: p2p.Block
	(*Tx)(nil),            // 7: p2p.Tx
	(*OutPoint)(nil),      // 8: p2p.OutPoint
	(*Input)(nil),         // 9: p2p.Input
	(*Output)(nil),        // 10: p2p.Output
}
var file_proto_p2p_proto_depIdxs = []int32{
	9,  // 0: p2p.Tx.inputs:type_name -> p2p.Input
	10, // 1: p2p.Tx.outputs:type_name -> p2p.Output
	8,  // 2: p2p.Input.outpoint:type_name -> p2p.OutPoint
	0,  // 3: p2p.P2P.SendHello:input_type -> p2p.HelloMessage
	2,  // 4: p2p.P2P.GetPeers:input_type -> p2p.PeersRequest
	3,  // 5: p2p.P2P.GetObject:input_type -> p2p.ObjectId
	3,  // 6: p2p.P2P.IHaveObject:input_type -> p2p.ObjectId
	0,  // 7: p2p.P2P.SendHello:output_type -> p2p.HelloMessage
	1,  // 8: p2p.P2P.GetPeers:output_type -> p2p.PeersResponse
	4,  // 9: p2p.P2P.GetObject:output_type -> p2p.Object
	5,  // 10: p2p.P2P.IHaveObject:output_type -> p2p.Empty
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_p2p_proto_init() }
func file_proto_p2p_proto_init() {
	if File_proto_p2p_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_p2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloMessage); i {
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
		file_proto_p2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersResponse); i {
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
		file_proto_p2p_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersRequest); i {
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
		file_proto_p2p_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectId); i {
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
		file_proto_p2p_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_proto_p2p_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_p2p_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_proto_p2p_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
		file_proto_p2p_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPoint); i {
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
		file_proto_p2p_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_proto_p2p_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_proto_p2p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_p2p_proto_goTypes,
		DependencyIndexes: file_proto_p2p_proto_depIdxs,
		MessageInfos:      file_proto_p2p_proto_msgTypes,
	}.Build()
	File_proto_p2p_proto = out.File
	file_proto_p2p_proto_rawDesc = nil
	file_proto_p2p_proto_goTypes = nil
	file_proto_p2p_proto_depIdxs = nil
}
