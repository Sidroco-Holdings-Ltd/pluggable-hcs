// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ConsensusUpdateTopic.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// All fields left null will not be updated.
// See [ConsensusService.updateTopic()](#proto.ConsensusService)
type ConsensusUpdateTopicTransactionBody struct {
	TopicID *TopicID `protobuf:"bytes,1,opt,name=topicID,proto3" json:"topicID,omitempty"`
	// Short publicly visible memo about the topic. No guarantee of uniqueness. Null for "do not update".
	Memo           *wrappers.StringValue `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	ValidStartTime *Timestamp            `protobuf:"bytes,3,opt,name=validStartTime,proto3" json:"validStartTime,omitempty"` // Deprecated: Do not use.
	// Effective consensus timestamp at (and after) which all consensus transactions and queries will fail.
	// The expirationTime may be no longer than 90 days from the consensus timestamp of this transaction.
	// If unspecified, no change.
	ExpirationTime *Timestamp `protobuf:"bytes,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// Access control for update/delete of the topic.
	// If unspecified, no change.
	// If empty keyList - the adminKey is cleared.
	AdminKey *Key `protobuf:"bytes,6,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	// Access control for ConsensusService.submitMessage.
	// If unspecified, no change.
	// If empty keyList - the submitKey is cleared.
	SubmitKey *Key `protobuf:"bytes,7,opt,name=submitKey,proto3" json:"submitKey,omitempty"`
	// The amount of time to extend the topic's lifetime automatically at expirationTime if the autoRenewAccount is
	// configured and has funds.
	// Limited to a maximum of 90 days (server-side configuration which may change).
	// If unspecified, no change.
	AutoRenewPeriod *Duration `protobuf:"bytes,8,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	// Optional account to be used at the topic's expirationTime to extend the life of the topic.
	// The topic lifetime will be extended up to a maximum of the autoRenewPeriod or however long the topic
	// can be extended using all funds on the account (whichever is the smaller duration/amount).
	// If specified as the default value (0.0.0), the autoRenewAccount will be removed.
	AutoRenewAccount     *AccountID `protobuf:"bytes,9,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConsensusUpdateTopicTransactionBody) Reset()         { *m = ConsensusUpdateTopicTransactionBody{} }
func (m *ConsensusUpdateTopicTransactionBody) String() string { return proto.CompactTextString(m) }
func (*ConsensusUpdateTopicTransactionBody) ProtoMessage()    {}
func (*ConsensusUpdateTopicTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cb1e7e99accb0fa, []int{0}
}

func (m *ConsensusUpdateTopicTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Unmarshal(m, b)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Marshal(b, m, deterministic)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Merge(m, src)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_Size() int {
	return xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Size(m)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusUpdateTopicTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusUpdateTopicTransactionBody proto.InternalMessageInfo

func (m *ConsensusUpdateTopicTransactionBody) GetTopicID() *TopicID {
	if m != nil {
		return m.TopicID
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetMemo() *wrappers.StringValue {
	if m != nil {
		return m.Memo
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConsensusUpdateTopicTransactionBody) GetValidStartTime() *Timestamp {
	if m != nil {
		return m.ValidStartTime
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetExpirationTime() *Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetAdminKey() *Key {
	if m != nil {
		return m.AdminKey
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetSubmitKey() *Key {
	if m != nil {
		return m.SubmitKey
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetAutoRenewAccount() *AccountID {
	if m != nil {
		return m.AutoRenewAccount
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsensusUpdateTopicTransactionBody)(nil), "proto.ConsensusUpdateTopicTransactionBody")
}

func init() {
	proto.RegisterFile("proto/ConsensusUpdateTopic.proto", fileDescriptor_8cb1e7e99accb0fa)
}

var fileDescriptor_8cb1e7e99accb0fa = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0xc9, 0xee, 0xba, 0x2f, 0x23, 0x74, 0x97, 0x41, 0x25, 0x14, 0x91, 0x45, 0x41, 0x72,
	0xd9, 0x89, 0xe8, 0x45, 0xa1, 0x17, 0x63, 0x0f, 0x96, 0x5e, 0xca, 0x34, 0x7a, 0xf0, 0x36, 0x4d,
	0x1e, 0x93, 0xc1, 0xce, 0x0b, 0xf3, 0x62, 0xcd, 0xf7, 0xf3, 0x83, 0x49, 0x66, 0x92, 0xca, 0xb6,
	0x3d, 0x0d, 0xf9, 0xfd, 0x7f, 0xff, 0xc9, 0xf3, 0x90, 0xa0, 0x7b, 0x6d, 0x94, 0x53, 0xf9, 0x17,
	0x25, 0x2d, 0x48, 0xeb, 0xed, 0x37, 0x5d, 0x33, 0x07, 0xa5, 0xd2, 0xbc, 0x22, 0x21, 0xc2, 0x4f,
	0xc2, 0x31, 0x7d, 0xd5, 0x28, 0xd5, 0x6c, 0x21, 0x0f, 0x4f, 0x1b, 0xff, 0x33, 0xdf, 0x19, 0xa6,
	0x35, 0x18, 0x1b, 0xb5, 0xe9, 0x8b, 0x78, 0x51, 0xc1, 0x2c, 0xaf, 0xca, 0x4e, 0xc3, 0xc8, 0x9f,
	0x45, 0x3e, 0xf7, 0x86, 0x39, 0xae, 0xe4, 0x40, 0x9f, 0x47, 0x5a, 0x72, 0x01, 0xd6, 0x31, 0xa1,
	0x23, 0x7e, 0xfd, 0xf7, 0x1c, 0xbd, 0x39, 0x35, 0x4a, 0x69, 0x98, 0xb4, 0xac, 0xea, 0x2f, 0x28,
	0x54, 0xdd, 0xe1, 0x0c, 0x5d, 0xb9, 0x9e, 0x2f, 0xe6, 0x69, 0x72, 0x9f, 0x64, 0x4f, 0xdf, 0x4f,
	0xe2, 0x05, 0xa4, 0x8c, 0x94, 0x8e, 0x31, 0x7e, 0x87, 0x2e, 0x04, 0x08, 0x95, 0x9e, 0x05, 0xed,
	0x25, 0x89, 0x5b, 0x90, 0x71, 0x0b, 0xb2, 0x76, 0x86, 0xcb, 0xe6, 0x3b, 0xdb, 0x7a, 0xa0, 0xc1,
	0xc4, 0x33, 0x34, 0xf9, 0xcd, 0xb6, 0xbc, 0x5e, 0x3b, 0x66, 0x5c, 0x3f, 0x60, 0x7a, 0x1e, 0xba,
	0x77, 0xe3, 0x2b, 0xc6, 0x99, 0x8b, 0xb3, 0x34, 0xa1, 0x07, 0x2e, 0xfe, 0x88, 0x26, 0xf0, 0x47,
	0xf3, 0xb8, 0x6c, 0x68, 0x5f, 0x9c, 0x6e, 0xd3, 0x03, 0x0f, 0xbf, 0x45, 0xd7, 0xac, 0x16, 0x5c,
	0x2e, 0xa1, 0x4b, 0x2f, 0x43, 0x07, 0x0d, 0x9d, 0x25, 0x74, 0x74, 0x9f, 0xe1, 0x0c, 0xdd, 0x58,
	0xbf, 0x11, 0xdc, 0xf5, 0xe2, 0xd5, 0x91, 0xf8, 0x3f, 0xc4, 0x9f, 0xd0, 0x2d, 0xf3, 0x4e, 0x51,
	0x90, 0xb0, 0x5b, 0x81, 0xe1, 0xaa, 0x4e, 0xaf, 0x83, 0x7f, 0x3b, 0xf8, 0xe3, 0x47, 0xa1, 0x87,
	0x1e, 0x9e, 0xa1, 0xbb, 0x3d, 0xfa, 0x5c, 0x55, 0xca, 0x4b, 0x97, 0xde, 0x3c, 0x5a, 0x64, 0xa0,
	0x8b, 0x39, 0x3d, 0x32, 0x8b, 0xaf, 0x68, 0x5a, 0x29, 0x41, 0x5a, 0xa8, 0xc1, 0x30, 0xd2, 0x32,
	0xdb, 0x36, 0x86, 0xe9, 0x36, 0x36, 0x57, 0xc9, 0x8f, 0xac, 0xe1, 0xae, 0xf5, 0x1b, 0x52, 0x29,
	0x91, 0xef, 0xd3, 0x3c, 0xea, 0x0f, 0xb6, 0xfe, 0xf5, 0xd0, 0xa8, 0xe1, 0x3f, 0xbb, 0x0c, 0xc7,
	0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x4a, 0x6c, 0xa9, 0xa7, 0x02, 0x00, 0x00,
}
