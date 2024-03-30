// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/forwarding/v1/account.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ForwardingAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	Channel            string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Recipient          string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	CreatedAt          int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *ForwardingAccount) Reset()         { *m = ForwardingAccount{} }
func (m *ForwardingAccount) String() string { return proto.CompactTextString(m) }
func (*ForwardingAccount) ProtoMessage()    {}
func (*ForwardingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d86db85ab0c667b, []int{0}
}
func (m *ForwardingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForwardingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForwardingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForwardingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardingAccount.Merge(m, src)
}
func (m *ForwardingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ForwardingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardingAccount proto.InternalMessageInfo

func (m *ForwardingAccount) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ForwardingAccount) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *ForwardingAccount) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*ForwardingAccount)(nil), "noble.forwarding.v1.ForwardingAccount")
}

func init() { proto.RegisterFile("noble/forwarding/v1/account.proto", fileDescriptor_9d86db85ab0c667b) }

var fileDescriptor_9d86db85ab0c667b = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x6b, 0x5a, 0x81, 0xea, 0xb2, 0x10, 0x18, 0xa2, 0x0a, 0x4c, 0x60, 0xea, 0x82, 0xad,
	0x80, 0x38, 0x40, 0x33, 0x20, 0x31, 0x21, 0x65, 0x64, 0xa9, 0x6c, 0xd7, 0x24, 0x91, 0x5a, 0x3b,
	0x8a, 0x5f, 0x02, 0xdc, 0x82, 0xcb, 0x70, 0x07, 0xc6, 0x8c, 0x4c, 0x08, 0x25, 0x17, 0x41, 0x38,
	0x81, 0x74, 0x7b, 0xff, 0xff, 0xde, 0xfb, 0x3f, 0xe9, 0xc7, 0x17, 0xda, 0x88, 0x8d, 0x62, 0x4f,
	0xa6, 0x78, 0xe6, 0xc5, 0x3a, 0xd3, 0x09, 0xab, 0x42, 0xc6, 0xa5, 0x34, 0xa5, 0x06, 0x9a, 0x17,
	0x06, 0x8c, 0x77, 0xec, 0x4e, 0xe8, 0x70, 0x42, 0xab, 0x70, 0x4e, 0xa4, 0xb1, 0x5b, 0x63, 0x19,
	0x2f, 0x21, 0x65, 0x55, 0x28, 0x14, 0xf0, 0xd0, 0x89, 0xee, 0x69, 0x7e, 0x92, 0x98, 0xc4, 0xb8,
	0x91, 0xfd, 0x4e, 0x9d, 0x7b, 0xf9, 0x8e, 0xf0, 0xd1, 0xdd, 0x7f, 0xce, 0xb2, 0xc3, 0x78, 0xf7,
	0xf8, 0x50, 0x70, 0xab, 0x56, 0x3d, 0xd6, 0x47, 0x01, 0x5a, 0xcc, 0xae, 0x03, 0xda, 0x21, 0xa8,
	0x4b, 0xed, 0x11, 0x34, 0xe2, 0x56, 0xf5, 0x7f, 0xd1, 0xa4, 0xfe, 0x3a, 0x47, 0xf1, 0x4c, 0x0c,
	0x96, 0xe7, 0xe3, 0x03, 0x99, 0x72, 0xad, 0xd5, 0xc6, 0xdf, 0x0b, 0xd0, 0x62, 0x1a, 0xff, 0x49,
	0xef, 0x14, 0x4f, 0x0b, 0x25, 0xb3, 0x3c, 0x53, 0x1a, 0xfc, 0xb1, 0xdb, 0x0d, 0x86, 0x77, 0x86,
	0xb1, 0x2c, 0x14, 0x07, 0xb5, 0x5e, 0x71, 0xf0, 0x27, 0x01, 0x5a, 0x8c, 0xe3, 0x69, 0xef, 0x2c,
	0x21, 0x7a, 0xf8, 0x68, 0x08, 0xaa, 0x1b, 0x82, 0xbe, 0x1b, 0x82, 0xde, 0x5a, 0x32, 0xaa, 0x5b,
	0x32, 0xfa, 0x6c, 0xc9, 0xe8, 0xf1, 0x36, 0xc9, 0x20, 0x2d, 0x05, 0x95, 0x66, 0xcb, 0x5c, 0x4f,
	0x57, 0xdc, 0x5a, 0x05, 0x76, 0xb7, 0xd1, 0x97, 0x5d, 0x01, 0xaf, 0xb9, 0xb2, 0x62, 0xdf, 0xf5,
	0x71, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x7a, 0x95, 0x3b, 0x7f, 0x01, 0x00, 0x00,
}

func (m *ForwardingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForwardingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ForwardingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovAccount(uint64(m.CreatedAt))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ForwardingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: ForwardingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
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
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)