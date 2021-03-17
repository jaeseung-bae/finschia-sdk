// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/vesting/v1beta1/vesting.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_line_lbm_sdk_v2_types "github.com/line/lbm-sdk/v2/types"
	types1 "github.com/line/lbm-sdk/v2/types"
	types "github.com/line/lbm-sdk/v2/x/auth/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BaseVestingAccount implements the VestingAccount interface. It contains all
// the necessary fields needed for any vesting account implementation.
type BaseVestingAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	OriginalVesting    github_com_line_lbm_sdk_v2_types.Coins `protobuf:"bytes,2,rep,name=original_vesting,json=originalVesting,proto3,castrepeated=github.com/line/lbm-sdk/v2/types.Coins" json:"original_vesting" yaml:"original_vesting"`
	DelegatedFree      github_com_line_lbm_sdk_v2_types.Coins `protobuf:"bytes,3,rep,name=delegated_free,json=delegatedFree,proto3,castrepeated=github.com/line/lbm-sdk/v2/types.Coins" json:"delegated_free" yaml:"delegated_free"`
	DelegatedVesting   github_com_line_lbm_sdk_v2_types.Coins `protobuf:"bytes,4,rep,name=delegated_vesting,json=delegatedVesting,proto3,castrepeated=github.com/line/lbm-sdk/v2/types.Coins" json:"delegated_vesting" yaml:"delegated_vesting"`
	EndTime            int64                                  `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
}

func (m *BaseVestingAccount) Reset()      { *m = BaseVestingAccount{} }
func (*BaseVestingAccount) ProtoMessage() {}
func (*BaseVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{0}
}
func (m *BaseVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseVestingAccount.Merge(m, src)
}
func (m *BaseVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseVestingAccount proto.InternalMessageInfo

// ContinuousVestingAccount implements the VestingAccount interface. It
// continuously vests by unlocking coins linearly with respect to time.
type ContinuousVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	StartTime           int64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
}

func (m *ContinuousVestingAccount) Reset()      { *m = ContinuousVestingAccount{} }
func (*ContinuousVestingAccount) ProtoMessage() {}
func (*ContinuousVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{1}
}
func (m *ContinuousVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContinuousVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContinuousVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContinuousVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousVestingAccount.Merge(m, src)
}
func (m *ContinuousVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ContinuousVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousVestingAccount proto.InternalMessageInfo

// DelayedVestingAccount implements the VestingAccount interface. It vests all
// coins after a specific time, but non prior. In other words, it keeps them
// locked until a specified time.
type DelayedVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
}

func (m *DelayedVestingAccount) Reset()      { *m = DelayedVestingAccount{} }
func (*DelayedVestingAccount) ProtoMessage() {}
func (*DelayedVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{2}
}
func (m *DelayedVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelayedVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelayedVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelayedVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelayedVestingAccount.Merge(m, src)
}
func (m *DelayedVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *DelayedVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DelayedVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DelayedVestingAccount proto.InternalMessageInfo

// Period defines a length of time and amount of coins that will vest.
type Period struct {
	Length int64                                  `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Amount github_com_line_lbm_sdk_v2_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/line/lbm-sdk/v2/types.Coins" json:"amount"`
}

func (m *Period) Reset()      { *m = Period{} }
func (*Period) ProtoMessage() {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{3}
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

func (m *Period) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_line_lbm_sdk_v2_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

// PeriodicVestingAccount implements the VestingAccount interface. It
// periodically vests by unlocking coins during each specified period.
type PeriodicVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	StartTime           int64    `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
	VestingPeriods      []Period `protobuf:"bytes,3,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods" yaml:"vesting_periods"`
}

func (m *PeriodicVestingAccount) Reset()      { *m = PeriodicVestingAccount{} }
func (*PeriodicVestingAccount) ProtoMessage() {}
func (*PeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{4}
}
func (m *PeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodicVestingAccount.Merge(m, src)
}
func (m *PeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *PeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodicVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseVestingAccount)(nil), "cosmos.vesting.v1beta1.BaseVestingAccount")
	proto.RegisterType((*ContinuousVestingAccount)(nil), "cosmos.vesting.v1beta1.ContinuousVestingAccount")
	proto.RegisterType((*DelayedVestingAccount)(nil), "cosmos.vesting.v1beta1.DelayedVestingAccount")
	proto.RegisterType((*Period)(nil), "cosmos.vesting.v1beta1.Period")
	proto.RegisterType((*PeriodicVestingAccount)(nil), "cosmos.vesting.v1beta1.PeriodicVestingAccount")
}

func init() {
	proto.RegisterFile("cosmos/vesting/v1beta1/vesting.proto", fileDescriptor_89e80273ca606d6e)
}

var fileDescriptor_89e80273ca606d6e = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x4d, 0x08, 0xe5, 0x0a, 0xfd, 0x63, 0xda, 0x60, 0x3a, 0xd8, 0x91, 0x85, 0x50,
	0x85, 0x84, 0xad, 0x16, 0xa6, 0x6e, 0xb8, 0x08, 0x54, 0x26, 0x64, 0x21, 0x06, 0x96, 0xe8, 0x6c,
	0x1f, 0xee, 0x09, 0xdb, 0x17, 0xf9, 0x2e, 0x11, 0xf9, 0x00, 0x48, 0x74, 0x02, 0x31, 0x31, 0x76,
	0x61, 0xe1, 0x43, 0x30, 0x77, 0x60, 0xc8, 0xc8, 0x14, 0x50, 0xf2, 0x0d, 0xf2, 0x09, 0x90, 0xef,
	0xce, 0x0e, 0xb8, 0x40, 0x54, 0x06, 0x24, 0xb6, 0xbc, 0xf7, 0xbe, 0xef, 0x73, 0xbf, 0x7b, 0xef,
	0x39, 0x07, 0xde, 0x08, 0x29, 0x4b, 0x29, 0x73, 0x07, 0x98, 0x71, 0x92, 0xc5, 0xee, 0x60, 0x37,
	0xc0, 0x1c, 0xed, 0x96, 0xb1, 0xd3, 0xcb, 0x29, 0xa7, 0x7a, 0x5b, 0x56, 0x39, 0xe5, 0xaa, 0xaa,
	0xda, 0xde, 0x8c, 0x69, 0x4c, 0x45, 0x89, 0x5b, 0xfc, 0x92, 0xd5, 0xdb, 0xa6, 0xd2, 0x0c, 0x10,
	0xc3, 0x95, 0x60, 0x48, 0x49, 0x56, 0xcb, 0xa3, 0x3e, 0x3f, 0xaa, 0xf2, 0x45, 0x20, 0xf3, 0xf6,
	0xe7, 0x26, 0xd4, 0x3d, 0xc4, 0xf0, 0x53, 0xb9, 0xdb, 0xbd, 0x30, 0xa4, 0xfd, 0x8c, 0xeb, 0x87,
	0xf0, 0x72, 0xa1, 0xd8, 0x45, 0x32, 0x36, 0x40, 0x07, 0xec, 0xac, 0xec, 0x75, 0x1c, 0xc5, 0x26,
	0x04, 0x94, 0x9a, 0x53, 0xb4, 0xab, 0x3e, 0xaf, 0x39, 0x1a, 0x5b, 0xc0, 0x5f, 0x09, 0xe6, 0x4b,
	0xfa, 0x1b, 0x00, 0xd7, 0x69, 0x4e, 0x62, 0x92, 0xa1, 0xa4, 0xab, 0x0e, 0x65, 0x2c, 0x75, 0x1a,
	0x3b, 0x2b, 0x7b, 0xd7, 0x4b, 0xbd, 0xa2, 0xbe, 0xd2, 0x3b, 0xa0, 0x24, 0xf3, 0x0e, 0x4f, 0xc7,
	0x96, 0x36, 0x1b, 0x5b, 0xd7, 0x86, 0x28, 0x4d, 0xf6, 0xed, 0xba, 0x80, 0xfd, 0xf1, 0xab, 0x75,
	0x33, 0x26, 0xfc, 0xa8, 0x1f, 0x38, 0x21, 0x4d, 0xdd, 0x84, 0x64, 0xd8, 0x4d, 0x82, 0xf4, 0x36,
	0x8b, 0x5e, 0xb8, 0x83, 0x3d, 0x97, 0x0f, 0x7b, 0x98, 0x09, 0x25, 0xe6, 0xaf, 0x95, 0xcd, 0xea,
	0x8c, 0xfa, 0x31, 0x80, 0xab, 0x11, 0x4e, 0x70, 0x8c, 0x38, 0x8e, 0xba, 0xcf, 0x73, 0x8c, 0x8d,
	0xc6, 0x22, 0x9e, 0x87, 0x8a, 0x67, 0x4b, 0xf2, 0xfc, 0xdc, 0x7e, 0x1e, 0x9a, 0x2b, 0x55, 0xeb,
	0x83, 0x1c, 0x63, 0xfd, 0x1d, 0x80, 0x1b, 0x73, 0xb1, 0x72, 0x3c, 0xcd, 0x45, 0x38, 0x8f, 0x14,
	0x8e, 0x51, 0xc7, 0xf9, 0x8b, 0xf9, 0xac, 0x57, 0xdd, 0xe5, 0x80, 0x1c, 0xb8, 0x8c, 0xb3, 0xa8,
	0xcb, 0x49, 0x8a, 0x8d, 0x0b, 0x1d, 0xb0, 0xd3, 0xf0, 0xae, 0xce, 0xc6, 0xd6, 0x9a, 0xdc, 0xab,
	0xcc, 0xd8, 0xfe, 0x45, 0x9c, 0x45, 0x4f, 0x48, 0x8a, 0xf7, 0x97, 0x5f, 0x9f, 0x58, 0xda, 0xfb,
	0x13, 0x4b, 0xb3, 0x3f, 0x01, 0x68, 0x1c, 0xd0, 0x8c, 0x93, 0xac, 0x4f, 0xfb, 0xac, 0x66, 0xaa,
	0x00, 0x6e, 0x0a, 0x53, 0x29, 0xc6, 0x9a, 0xb9, 0x6e, 0x39, 0xbf, 0x36, 0xbe, 0x73, 0xd6, 0x9e,
	0xca, 0x66, 0x7a, 0x70, 0xd6, 0xb8, 0x77, 0x21, 0x64, 0x1c, 0xe5, 0x5c, 0xc2, 0x2f, 0x09, 0xf8,
	0xad, 0xd9, 0xd8, 0xda, 0x90, 0xf0, 0xf3, 0x9c, 0xed, 0x5f, 0x12, 0x41, 0xed, 0x00, 0xaf, 0x00,
	0xdc, 0xba, 0x8f, 0x13, 0x34, 0xac, 0xa6, 0xf1, 0x0f, 0xe9, 0x7f, 0xe0, 0x38, 0x06, 0xb0, 0xf5,
	0x18, 0xe7, 0x84, 0x46, 0x7a, 0x1b, 0xb6, 0x12, 0x9c, 0xc5, 0xfc, 0x48, 0x6c, 0xd5, 0xf0, 0x55,
	0xa4, 0x23, 0xd8, 0x42, 0xa9, 0x40, 0x58, 0xf8, 0x9a, 0x9c, 0xc2, 0x2e, 0xe7, 0xb0, 0x84, 0x12,
	0xde, 0x6f, 0x0a, 0x96, 0x0f, 0x4b, 0xb0, 0x2d, 0x59, 0x48, 0xf8, 0xbf, 0x5c, 0xa9, 0x1e, 0xc3,
	0xb5, 0x12, 0xaa, 0x27, 0xd8, 0x99, 0x7a, 0xe4, 0xe6, 0xef, 0xa0, 0xe4, 0x11, 0x3d, 0x53, 0x3d,
	0xad, 0xb6, 0x94, 0xaf, 0x89, 0xd8, 0xfe, 0xaa, 0x5a, 0x91, 0xe5, 0x6c, 0x7e, 0x67, 0xde, 0xe1,
	0xe9, 0xc4, 0x04, 0xa3, 0x89, 0x09, 0xbe, 0x4d, 0x4c, 0xf0, 0x76, 0x6a, 0x6a, 0xa3, 0xa9, 0xa9,
	0x7d, 0x99, 0x9a, 0xda, 0x33, 0xf7, 0x0f, 0x73, 0x7f, 0xa9, 0xbe, 0xcd, 0xea, 0x4f, 0x41, 0x5c,
	0x43, 0xd0, 0x12, 0x5f, 0xe7, 0x3b, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xd2, 0x3f, 0x0c,
	0x33, 0x06, 0x00, 0x00,
}

func (m *BaseVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DelegatedVesting) > 0 {
		for iNdEx := len(m.DelegatedVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DelegatedFree) > 0 {
		for iNdEx := len(m.DelegatedFree) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedFree[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OriginalVesting) > 0 {
		for iNdEx := len(m.OriginalVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginalVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContinuousVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinuousVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContinuousVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelayedVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelayedVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelayedVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
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
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Length != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.OriginalVesting) > 0 {
		for _, e := range m.OriginalVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedFree) > 0 {
		for _, e := range m.DelegatedFree {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedVesting) > 0 {
		for _, e := range m.DelegatedVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovVesting(uint64(m.EndTime))
	}
	return n
}

func (m *ContinuousVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	return n
}

func (m *DelayedVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Length != 0 {
		n += 1 + sovVesting(uint64(m.Length))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func (m *PeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: BaseVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalVesting = append(m.OriginalVesting, types1.Coin{})
			if err := m.OriginalVesting[len(m.OriginalVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedFree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedFree = append(m.DelegatedFree, types1.Coin{})
			if err := m.DelegatedFree[len(m.DelegatedFree)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedVesting = append(m.DelegatedVesting, types1.Coin{})
			if err := m.DelegatedVesting[len(m.DelegatedVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *ContinuousVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ContinuousVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinuousVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *DelayedVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: DelayedVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelayedVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
				return ErrIntOverflowVesting
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
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types1.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *PeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: PeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
