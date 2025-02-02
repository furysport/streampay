// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya/streampay/v1/streampay.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StreamType int32

const (
	TypeDelayed    StreamType = 0
	TypeContinuous StreamType = 1
	TypePeriodic   StreamType = 2
)

var StreamType_name = map[int32]string{
	0: "STREAM_PAYMENT_TYPE_DELAYED",
	1: "STREAM_PAYMENT_TYPE_CONTINUOUS",
	2: "STREAM_PAYMENT_TYPE_PERIODIC",
}

var StreamType_value = map[string]int32{
	"STREAM_PAYMENT_TYPE_DELAYED":    0,
	"STREAM_PAYMENT_TYPE_CONTINUOUS": 1,
	"STREAM_PAYMENT_TYPE_PERIODIC":   2,
}

func (x StreamType) String() string {
	return proto.EnumName(StreamType_name, int32(x))
}

func (StreamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc977779d6db321d, []int{0}
}

type StreamStatus int32

const (
	StatusActive    StreamStatus = 0
	StatusCompleted StreamStatus = 1
	StatusCancelled StreamStatus = 2
)

var StreamStatus_name = map[int32]string{
	0: "STREAM_PAYMENT_STATUS_ACTIVE",
	1: "STREAM_PAYMENT_STATUS_COMPLETED",
	2: "STREAM_PAYMENT_STATUS_CANCELLED",
}

var StreamStatus_value = map[string]int32{
	"STREAM_PAYMENT_STATUS_ACTIVE":    0,
	"STREAM_PAYMENT_STATUS_COMPLETED": 1,
	"STREAM_PAYMENT_STATUS_CANCELLED": 2,
}

func (x StreamStatus) String() string {
	return proto.EnumName(StreamStatus_name, int32(x))
}

func (StreamStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc977779d6db321d, []int{1}
}

type StreamPayment struct {
	Id             string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender         string       `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient      string       `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	TotalAmount    types.Coin   `protobuf:"bytes,4,opt,name=total_amount,json=totalAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_amount" yaml:"total_amount"`
	StreamType     StreamType   `protobuf:"varint,5,opt,name=stream_type,json=streamType,proto3,enum=furya.streampay.v1.StreamType" json:"stream_type,omitempty" yaml:"stream_type"`
	Periods        []*Period    `protobuf:"bytes,6,rep,name=periods,proto3" json:"periods,omitempty"`
	Cancellable    bool         `protobuf:"varint,7,opt,name=cancellable,proto3" json:"cancellable,omitempty"`
	StartTime      time.Time    `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	EndTime        time.Time    `protobuf:"bytes,9,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	StreamedAmount types.Coin   `protobuf:"bytes,10,opt,name=streamed_amount,json=streamedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"streamed_amount" yaml:"streamed_amount"`
	LastClaimedAt  time.Time    `protobuf:"bytes,11,opt,name=last_claimed_at,json=lastClaimedAt,proto3,stdtime" json:"last_claimed_at" yaml:"last_claimed_at"`
	Status         StreamStatus `protobuf:"varint,12,opt,name=status,proto3,enum=furya.streampay.v1.StreamStatus" json:"status,omitempty"`
}

func (m *StreamPayment) Reset()         { *m = StreamPayment{} }
func (m *StreamPayment) String() string { return proto.CompactTextString(m) }
func (*StreamPayment) ProtoMessage()    {}
func (*StreamPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc977779d6db321d, []int{0}
}
func (m *StreamPayment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamPayment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPayment.Merge(m, src)
}
func (m *StreamPayment) XXX_Size() int {
	return m.Size()
}
func (m *StreamPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPayment.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPayment proto.InternalMessageInfo

func (m *StreamPayment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamPayment) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *StreamPayment) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *StreamPayment) GetTotalAmount() types.Coin {
	if m != nil {
		return m.TotalAmount
	}
	return types.Coin{}
}

func (m *StreamPayment) GetStreamType() StreamType {
	if m != nil {
		return m.StreamType
	}
	return TypeDelayed
}

func (m *StreamPayment) GetPeriods() []*Period {
	if m != nil {
		return m.Periods
	}
	return nil
}

func (m *StreamPayment) GetCancellable() bool {
	if m != nil {
		return m.Cancellable
	}
	return false
}

func (m *StreamPayment) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *StreamPayment) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *StreamPayment) GetStreamedAmount() types.Coin {
	if m != nil {
		return m.StreamedAmount
	}
	return types.Coin{}
}

func (m *StreamPayment) GetLastClaimedAt() time.Time {
	if m != nil {
		return m.LastClaimedAt
	}
	return time.Time{}
}

func (m *StreamPayment) GetStatus() StreamStatus {
	if m != nil {
		return m.Status
	}
	return StatusActive
}

type Period struct {
	Amount   int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty" yaml:"amount"`
	Duration int64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty" yaml:"duration"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc977779d6db321d, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Period) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterEnum("furya.streampay.v1.StreamType", StreamType_name, StreamType_value)
	proto.RegisterEnum("furya.streampay.v1.StreamStatus", StreamStatus_name, StreamStatus_value)
	proto.RegisterType((*StreamPayment)(nil), "furya.streampay.v1.StreamPayment")
	proto.RegisterType((*Period)(nil), "furya.streampay.v1.Period")
}

func init() {
	proto.RegisterFile("furya/streampay/v1/streampay.proto", fileDescriptor_fc977779d6db321d)
}

var fileDescriptor_fc977779d6db321d = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x41, 0x6f, 0xe3, 0x44,
	0x18, 0xcd, 0xa4, 0x25, 0x6d, 0x27, 0x6d, 0x93, 0x9d, 0x42, 0x65, 0xb2, 0xbb, 0xb6, 0x09, 0x42,
	0x84, 0x95, 0xb0, 0x69, 0x91, 0x10, 0x02, 0x71, 0x70, 0x1c, 0xaf, 0x14, 0xa9, 0x6d, 0x22, 0xc7,
	0x45, 0x74, 0x2f, 0xd1, 0xc4, 0x9e, 0x0d, 0x23, 0x6c, 0x8f, 0x15, 0x4f, 0x2a, 0x72, 0xe7, 0x80,
	0x2a, 0x21, 0xed, 0x1f, 0xe8, 0x89, 0x1b, 0x27, 0x7e, 0x00, 0x3f, 0xa0, 0xc7, 0x3d, 0x72, 0xca,
	0xa2, 0xf6, 0x1f, 0xf4, 0xcc, 0x01, 0x79, 0xc6, 0x6e, 0xb2, 0x28, 0xb4, 0xe2, 0x94, 0x99, 0x6f,
	0xde, 0x7b, 0x99, 0xef, 0xbd, 0x6f, 0x64, 0xf8, 0x51, 0x2f, 0x8a, 0xe9, 0xf3, 0x90, 0xfe, 0x68,
	0xa6, 0x7c, 0x42, 0x70, 0x94, 0xe0, 0x99, 0x79, 0x7e, 0xb0, 0xd8, 0x18, 0xc9, 0x84, 0x71, 0x86,
	0xde, 0x2b, 0x60, 0xc6, 0xe2, 0xe4, 0xfc, 0xa0, 0xa1, 0xfa, 0x2c, 0x8d, 0x58, 0x6a, 0x8e, 0x70,
	0x4a, 0xcc, 0xf3, 0x83, 0x11, 0xe1, 0xf8, 0xc0, 0xf4, 0x19, 0x8d, 0x25, 0xad, 0xf1, 0xee, 0x98,
	0x8d, 0x99, 0x58, 0x9a, 0xd9, 0x2a, 0xaf, 0x6a, 0x63, 0xc6, 0xc6, 0x21, 0x31, 0xc5, 0x6e, 0x34,
	0x7d, 0x69, 0x72, 0x1a, 0x91, 0x94, 0xe3, 0x28, 0x91, 0x80, 0xe6, 0xdf, 0x15, 0xb8, 0x33, 0x10,
	0xff, 0xd3, 0xc7, 0xb3, 0x88, 0xc4, 0x1c, 0xed, 0xc2, 0x32, 0x0d, 0x14, 0xa0, 0x83, 0xd6, 0x96,
	0x5b, 0xa6, 0x01, 0xda, 0x87, 0x95, 0x94, 0xc4, 0x01, 0x99, 0x28, 0x65, 0x51, 0xcb, 0x77, 0xe8,
	0x09, 0xdc, 0x9a, 0x10, 0x9f, 0x26, 0x94, 0xc4, 0x5c, 0x59, 0x13, 0x47, 0x8b, 0x02, 0xfa, 0x09,
	0xc0, 0x6d, 0xce, 0x38, 0x0e, 0x87, 0x38, 0x62, 0xd3, 0x98, 0x2b, 0xeb, 0x3a, 0x68, 0x55, 0x0f,
	0xdf, 0x37, 0x64, 0x1b, 0x46, 0xd6, 0x86, 0x91, 0xb7, 0x61, 0xd8, 0x8c, 0xc6, 0xed, 0xe7, 0x57,
	0x73, 0xad, 0x74, 0x3b, 0xd7, 0xf6, 0x66, 0x38, 0x0a, 0xbf, 0x6a, 0x2e, 0x93, 0x9b, 0xbf, 0xbd,
	0xd1, 0x3e, 0x1e, 0x53, 0xfe, 0xfd, 0x74, 0x64, 0xf8, 0x2c, 0x32, 0x73, 0x2b, 0xe4, 0xcf, 0xa7,
	0x69, 0xf0, 0x83, 0xc9, 0x67, 0x09, 0x49, 0x85, 0x8e, 0x5b, 0x15, 0x4c, 0x4b, 0x10, 0xd1, 0x0b,
	0x58, 0x95, 0x2e, 0x0e, 0x33, 0x84, 0xf2, 0x8e, 0x0e, 0x5a, 0xbb, 0x87, 0x1f, 0x18, 0x2b, 0x2d,
	0x36, 0xa4, 0x0f, 0xde, 0x2c, 0x21, 0xed, 0xfd, 0xdb, 0xb9, 0x86, 0xe4, 0x45, 0x96, 0xf8, 0x4d,
	0x17, 0xa6, 0x77, 0x18, 0xf4, 0x0d, 0xdc, 0x48, 0xc8, 0x84, 0xb2, 0x20, 0x55, 0x2a, 0xfa, 0x5a,
	0xab, 0x7a, 0xf8, 0xf4, 0x3f, 0x74, 0xfb, 0x02, 0xd5, 0x5e, 0xbf, 0x9a, 0x6b, 0xc0, 0x2d, 0x38,
	0x48, 0x87, 0x55, 0x1f, 0xc7, 0x3e, 0x09, 0x43, 0x3c, 0x0a, 0x89, 0xb2, 0xa1, 0x83, 0xd6, 0xa6,
	0xbb, 0x5c, 0x42, 0xdf, 0x41, 0x98, 0x72, 0x3c, 0xe1, 0xc3, 0x2c, 0x34, 0x65, 0x53, 0x18, 0xd8,
	0x30, 0x64, 0xa2, 0x46, 0x91, 0xa8, 0xe1, 0x15, 0x89, 0xb6, 0x9f, 0xe6, 0x0e, 0x3e, 0x2a, 0x2e,
	0x5e, 0x70, 0x9b, 0xaf, 0xde, 0x68, 0xc0, 0xdd, 0x12, 0x85, 0x0c, 0x8e, 0x5c, 0xb8, 0x49, 0xe2,
	0x40, 0xea, 0x6e, 0x3d, 0xa8, 0xfb, 0x38, 0xd7, 0xad, 0x49, 0xdd, 0x82, 0x29, 0x55, 0x37, 0x48,
	0x1c, 0x08, 0xcd, 0x5f, 0x00, 0xac, 0xc9, 0xb6, 0x49, 0x50, 0x84, 0x0e, 0x1f, 0x0a, 0xbd, 0x9b,
	0x4b, 0xef, 0x2f, 0x7b, 0x7d, 0xc7, 0xff, 0x5f, 0xb9, 0xef, 0x16, 0xe4, 0x3c, 0xfa, 0x97, 0xb0,
	0x16, 0xe2, 0x94, 0x0f, 0xfd, 0x10, 0x53, 0x21, 0xc9, 0x95, 0xea, 0x83, 0xad, 0x36, 0xdf, 0xbe,
	0xcf, 0xbf, 0x04, 0x64, 0xc7, 0x3b, 0x59, 0xd5, 0x96, 0x45, 0x8b, 0xa3, 0xaf, 0x61, 0x25, 0xe5,
	0x98, 0x4f, 0x53, 0x65, 0x5b, 0x4c, 0xd7, 0x87, 0xf7, 0x4e, 0xd7, 0x40, 0x40, 0xdd, 0x9c, 0xd2,
	0x0c, 0x60, 0x45, 0x4e, 0x07, 0xfa, 0x04, 0x56, 0x72, 0xd3, 0xb2, 0xa7, 0xb7, 0xd6, 0x7e, 0x74,
	0x3b, 0xd7, 0x76, 0xe4, 0x2d, 0x72, 0x33, 0xdc, 0x1c, 0x80, 0x4c, 0xb8, 0x19, 0x4c, 0x27, 0x98,
	0x53, 0x16, 0x8b, 0x37, 0xb9, 0xd6, 0xde, 0x5b, 0xa4, 0x53, 0x9c, 0x34, 0xdd, 0x3b, 0xd0, 0xb3,
	0xdf, 0x01, 0x84, 0x8b, 0xe1, 0x46, 0x9f, 0xc1, 0xc7, 0x03, 0xcf, 0x75, 0xac, 0xe3, 0x61, 0xdf,
	0x3a, 0x3b, 0x76, 0x4e, 0xbc, 0xa1, 0x77, 0xd6, 0x77, 0x86, 0x1d, 0xe7, 0xc8, 0x3a, 0x73, 0x3a,
	0xf5, 0x52, 0xa3, 0x76, 0x71, 0xa9, 0x57, 0x33, 0x68, 0x87, 0x84, 0x78, 0x46, 0x02, 0xf4, 0x05,
	0x54, 0x57, 0x31, 0xec, 0xde, 0x89, 0xd7, 0x3d, 0x39, 0xed, 0x9d, 0x0e, 0xea, 0xa0, 0x81, 0x2e,
	0x2e, 0xf5, 0xdd, 0x8c, 0x64, 0xb3, 0x98, 0xd3, 0x78, 0xca, 0xa6, 0x29, 0x3a, 0x84, 0x4f, 0x56,
	0xf1, 0xfa, 0x8e, 0xdb, 0xed, 0x75, 0xba, 0x76, 0xbd, 0xdc, 0xa8, 0x5f, 0x5c, 0xea, 0xdb, 0x19,
	0x4b, 0xda, 0x40, 0xfd, 0xc6, 0xfa, 0xcf, 0xbf, 0xaa, 0xa5, 0x67, 0x7f, 0x00, 0xb8, 0xbd, 0xec,
	0xd8, 0x0a, 0xa9, 0x81, 0x67, 0x79, 0xa7, 0x83, 0xa1, 0x65, 0x7b, 0xdd, 0x6f, 0x9d, 0x7a, 0x49,
	0x4a, 0x49, 0xb4, 0xe5, 0x73, 0x7a, 0x4e, 0xd0, 0x97, 0x50, 0x5b, 0xcd, 0xb1, 0x7b, 0xc7, 0xfd,
	0x23, 0xc7, 0x73, 0x3a, 0x75, 0xd0, 0xd8, 0xbb, 0xb8, 0xd4, 0x6b, 0x92, 0x66, 0xb3, 0x28, 0x09,
	0x09, 0x27, 0xc1, 0x3d, 0x4c, 0xeb, 0xc4, 0x76, 0x8e, 0x8e, 0x9c, 0x4e, 0xbd, 0xfc, 0x16, 0x53,
	0x3e, 0x5b, 0x12, 0xc8, 0xeb, 0xb7, 0xbb, 0x57, 0xd7, 0x2a, 0x78, 0x7d, 0xad, 0x82, 0xbf, 0xae,
	0x55, 0xf0, 0xea, 0x46, 0x2d, 0xbd, 0xbe, 0x51, 0x4b, 0x7f, 0xde, 0xa8, 0xa5, 0x17, 0xe6, 0xd2,
	0x3c, 0xaf, 0xf8, 0x20, 0x2c, 0xaf, 0xc5, 0x70, 0x8f, 0x2a, 0x62, 0x4c, 0x3f, 0xff, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x83, 0xab, 0x47, 0x3f, 0x06, 0x00, 0x00,
}

func (m *StreamPayment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamPayment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamPayment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintStreampay(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x60
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastClaimedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastClaimedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStreampay(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x5a
	{
		size, err := m.StreamedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStreampay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintStreampay(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x4a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintStreampay(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x42
	if m.Cancellable {
		i--
		if m.Cancellable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Periods) > 0 {
		for iNdEx := len(m.Periods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Periods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStreampay(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.StreamType != 0 {
		i = encodeVarintStreampay(dAtA, i, uint64(m.StreamType))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TotalAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStreampay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintStreampay(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintStreampay(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintStreampay(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Duration != 0 {
		i = encodeVarintStreampay(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x10
	}
	if m.Amount != 0 {
		i = encodeVarintStreampay(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStreampay(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreampay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamPayment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovStreampay(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovStreampay(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovStreampay(uint64(l))
	}
	l = m.TotalAmount.Size()
	n += 1 + l + sovStreampay(uint64(l))
	if m.StreamType != 0 {
		n += 1 + sovStreampay(uint64(m.StreamType))
	}
	if len(m.Periods) > 0 {
		for _, e := range m.Periods {
			l = e.Size()
			n += 1 + l + sovStreampay(uint64(l))
		}
	}
	if m.Cancellable {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovStreampay(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovStreampay(uint64(l))
	l = m.StreamedAmount.Size()
	n += 1 + l + sovStreampay(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastClaimedAt)
	n += 1 + l + sovStreampay(uint64(l))
	if m.Status != 0 {
		n += 1 + sovStreampay(uint64(m.Status))
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovStreampay(uint64(m.Amount))
	}
	if m.Duration != 0 {
		n += 1 + sovStreampay(uint64(m.Duration))
	}
	return n
}

func sovStreampay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreampay(x uint64) (n int) {
	return sovStreampay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamPayment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreampay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamPayment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamPayment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamType", wireType)
			}
			m.StreamType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamType |= StreamType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Periods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Periods = append(m.Periods, &Period{})
			if err := m.Periods[len(m.Periods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancellable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Cancellable = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StreamedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastClaimedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastClaimedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= StreamStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStreampay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreampay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreampay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStreampay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreampay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStreampay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreampay
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStreampay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreampay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreampay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreampay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreampay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreampay = fmt.Errorf("proto: unexpected end of group")
)
