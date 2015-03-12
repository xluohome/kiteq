// Code generated by protoc-gen-go.
// source: protocol/kite_remoting.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol/kite_remoting.proto

It has these top-level messages:
	HeartBeat
	ConnMeta
	ConnAuthAck
	MessageStoreAck
	DeliverAck
	TxACKPacket
	Entry
	Header
	BytesMessage
	StringMessage
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// 心跳请求包
type HeartBeat struct {
	Version          *int64 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}

func (m *HeartBeat) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

// 连接的Meta数据包
type ConnMeta struct {
	GroupId          *string `protobuf:"bytes,1,req,name=groupId" json:"groupId,omitempty"`
	SecretKey        *string `protobuf:"bytes,2,req,name=secretKey" json:"secretKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConnMeta) Reset()         { *m = ConnMeta{} }
func (m *ConnMeta) String() string { return proto.CompactTextString(m) }
func (*ConnMeta) ProtoMessage()    {}

func (m *ConnMeta) GetGroupId() string {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return ""
}

func (m *ConnMeta) GetSecretKey() string {
	if m != nil && m.SecretKey != nil {
		return *m.SecretKey
	}
	return ""
}

// 握手确认数据包
type ConnAuthAck struct {
	Status           *bool   `protobuf:"varint,1,req,name=status,def=1" json:"status,omitempty"`
	Feedback         *string `protobuf:"bytes,2,req,name=feedback" json:"feedback,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConnAuthAck) Reset()         { *m = ConnAuthAck{} }
func (m *ConnAuthAck) String() string { return proto.CompactTextString(m) }
func (*ConnAuthAck) ProtoMessage()    {}

const Default_ConnAuthAck_Status bool = true

func (m *ConnAuthAck) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_ConnAuthAck_Status
}

func (m *ConnAuthAck) GetFeedback() string {
	if m != nil && m.Feedback != nil {
		return *m.Feedback
	}
	return ""
}

// 消息确认接收数据包
type MessageStoreAck struct {
	MessageId        *string `protobuf:"bytes,1,req,name=messageId" json:"messageId,omitempty"`
	Status           *bool   `protobuf:"varint,2,req,name=status,def=1" json:"status,omitempty"`
	Feedback         *string `protobuf:"bytes,3,req,name=feedback" json:"feedback,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MessageStoreAck) Reset()         { *m = MessageStoreAck{} }
func (m *MessageStoreAck) String() string { return proto.CompactTextString(m) }
func (*MessageStoreAck) ProtoMessage()    {}

const Default_MessageStoreAck_Status bool = true

func (m *MessageStoreAck) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *MessageStoreAck) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_MessageStoreAck_Status
}

func (m *MessageStoreAck) GetFeedback() string {
	if m != nil && m.Feedback != nil {
		return *m.Feedback
	}
	return ""
}

// 消息接收packet
type DeliverAck struct {
	MessageId        *string `protobuf:"bytes,1,req,name=messageId" json:"messageId,omitempty"`
	Topic            *string `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
	MessageType      *string `protobuf:"bytes,3,req,name=messageType" json:"messageType,omitempty"`
	GroupId          *string `protobuf:"bytes,4,req,name=groupId" json:"groupId,omitempty"`
	Status           *bool   `protobuf:"varint,5,req,name=status,def=1" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeliverAck) Reset()         { *m = DeliverAck{} }
func (m *DeliverAck) String() string { return proto.CompactTextString(m) }
func (*DeliverAck) ProtoMessage()    {}

const Default_DeliverAck_Status bool = true

func (m *DeliverAck) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *DeliverAck) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *DeliverAck) GetMessageType() string {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return ""
}

func (m *DeliverAck) GetGroupId() string {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return ""
}

func (m *DeliverAck) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_DeliverAck_Status
}

// 事务确认数据包
type TxACKPacket struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Status           *int32  `protobuf:"varint,2,req,name=status,def=0" json:"status,omitempty"`
	Feedback         *string `protobuf:"bytes,3,req,name=feedback" json:"feedback,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TxACKPacket) Reset()         { *m = TxACKPacket{} }
func (m *TxACKPacket) String() string { return proto.CompactTextString(m) }
func (*TxACKPacket) ProtoMessage()    {}

const Default_TxACKPacket_Status int32 = 0

func (m *TxACKPacket) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TxACKPacket) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_TxACKPacket_Status
}

func (m *TxACKPacket) GetFeedback() string {
	if m != nil && m.Feedback != nil {
		return *m.Feedback
	}
	return ""
}

type Entry struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}

func (m *Entry) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Entry) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Header struct {
	MessageId        *string  `protobuf:"bytes,1,req,name=messageId" json:"messageId,omitempty"`
	Topic            *string  `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
	MessageType      *string  `protobuf:"bytes,3,req,name=messageType" json:"messageType,omitempty"`
	ExpiredTime      *int64   `protobuf:"varint,4,req,name=expiredTime,def=-1" json:"expiredTime,omitempty"`
	DeliverLimit     *int32   `protobuf:"varint,5,req,name=deliverLimit,def=100" json:"deliverLimit,omitempty"`
	GroupId          *string  `protobuf:"bytes,6,req,name=groupId" json:"groupId,omitempty"`
	Commit           *bool    `protobuf:"varint,7,req,name=commit" json:"commit,omitempty"`
	Fly              *bool    `protobuf:"varint,8,req,name=fly,def=0" json:"fly,omitempty"`
	Properties       []*Entry `protobuf:"bytes,9,rep,name=properties" json:"properties,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}

const Default_Header_ExpiredTime int64 = -1
const Default_Header_DeliverLimit int32 = 100
const Default_Header_Fly bool = false

func (m *Header) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Header) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *Header) GetMessageType() string {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return ""
}

func (m *Header) GetExpiredTime() int64 {
	if m != nil && m.ExpiredTime != nil {
		return *m.ExpiredTime
	}
	return Default_Header_ExpiredTime
}

func (m *Header) GetDeliverLimit() int32 {
	if m != nil && m.DeliverLimit != nil {
		return *m.DeliverLimit
	}
	return Default_Header_DeliverLimit
}

func (m *Header) GetGroupId() string {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return ""
}

func (m *Header) GetCommit() bool {
	if m != nil && m.Commit != nil {
		return *m.Commit
	}
	return false
}

func (m *Header) GetFly() bool {
	if m != nil && m.Fly != nil {
		return *m.Fly
	}
	return Default_Header_Fly
}

func (m *Header) GetProperties() []*Entry {
	if m != nil {
		return m.Properties
	}
	return nil
}

// byte类消息
type BytesMessage struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             []byte  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BytesMessage) Reset()         { *m = BytesMessage{} }
func (m *BytesMessage) String() string { return proto.CompactTextString(m) }
func (*BytesMessage) ProtoMessage()    {}

func (m *BytesMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BytesMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// string类型的message
type StringMessage struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             *string `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}

func (m *StringMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *StringMessage) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
}
