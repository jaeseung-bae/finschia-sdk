// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/foundation/v1/authz.proto

package foundation

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	proto "github.com/cosmos/gogoproto/proto"
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

// ReceiveFromTreasuryAuthorization allows the grantee to receive coins
// up to receive_limit from the treasury.
type ReceiveFromTreasuryAuthorization struct {
}

func (m *ReceiveFromTreasuryAuthorization) Reset()         { *m = ReceiveFromTreasuryAuthorization{} }
func (m *ReceiveFromTreasuryAuthorization) String() string { return proto.CompactTextString(m) }
func (*ReceiveFromTreasuryAuthorization) ProtoMessage()    {}
func (*ReceiveFromTreasuryAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bdb89c90659aa0e, []int{0}
}
func (m *ReceiveFromTreasuryAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiveFromTreasuryAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiveFromTreasuryAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceiveFromTreasuryAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveFromTreasuryAuthorization.Merge(m, src)
}
func (m *ReceiveFromTreasuryAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *ReceiveFromTreasuryAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveFromTreasuryAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveFromTreasuryAuthorization proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReceiveFromTreasuryAuthorization)(nil), "lbm.foundation.v1.ReceiveFromTreasuryAuthorization")
}

func init() { proto.RegisterFile("lbm/foundation/v1/authz.proto", fileDescriptor_8bdb89c90659aa0e) }

var fileDescriptor_8bdb89c90659aa0e = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x49, 0xca, 0xd5,
	0x4f, 0xcb, 0x2f, 0xcd, 0x4b, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x33, 0xd4, 0x4f, 0x2c,
	0x2d, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x49, 0xca, 0xd5, 0x43,
	0x48, 0xeb, 0x95, 0x19, 0x4a, 0x49, 0x26, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xc7, 0x83, 0x15, 0xe8,
	0x43, 0x38, 0x10, 0xd5, 0x52, 0x82, 0x89, 0xb9, 0x99, 0x79, 0xf9, 0xfa, 0x60, 0x12, 0x22, 0xa4,
	0x34, 0x81, 0x91, 0x4b, 0x21, 0x28, 0x35, 0x39, 0x35, 0xb3, 0x2c, 0xd5, 0xad, 0x28, 0x3f, 0x37,
	0xa4, 0x28, 0x35, 0xb1, 0xb8, 0xb4, 0xa8, 0xd2, 0xb1, 0xb4, 0x24, 0x23, 0xbf, 0x28, 0xb3, 0x0a,
	0x6c, 0xaa, 0x55, 0xce, 0xa9, 0x2d, 0xba, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x6e, 0x99, 0x79, 0xc5, 0xc9, 0x19, 0x99, 0x89, 0xfa, 0x69, 0x50, 0x86, 0x6e,
	0x71, 0x4a, 0xb6, 0x7e, 0x05, 0x92, 0x3b, 0xf5, 0x50, 0x0c, 0xe8, 0x7a, 0xbe, 0x41, 0x4b, 0x23,
	0x27, 0x29, 0x17, 0xac, 0x8c, 0x90, 0x6d, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xa5, 0x4b, 0x92, 0x1b, 0x92, 0xd8, 0xc0, 0x5e, 0x34, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x46, 0x80, 0x15, 0x44, 0x01, 0x00, 0x00,
}

func (m *ReceiveFromTreasuryAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiveFromTreasuryAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceiveFromTreasuryAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReceiveFromTreasuryAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReceiveFromTreasuryAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: ReceiveFromTreasuryAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiveFromTreasuryAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
