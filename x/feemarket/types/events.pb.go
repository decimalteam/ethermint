// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/feemarket/v1/events.proto

package types

import (
	fmt "fmt"
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

// EventFeeMarket is the event type for the fee market module
type EventFeeMarket struct {
	// base_fee for EIP-1559 blocks
	BaseFee string `protobuf:"bytes,1,opt,name=base_fee,json=baseFee,proto3" json:"base_fee,omitempty"`
}

func (m *EventFeeMarket) Reset()         { *m = EventFeeMarket{} }
func (m *EventFeeMarket) String() string { return proto.CompactTextString(m) }
func (*EventFeeMarket) ProtoMessage()    {}
func (*EventFeeMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6edce8d670faff7, []int{0}
}
func (m *EventFeeMarket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFeeMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFeeMarket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFeeMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFeeMarket.Merge(m, src)
}
func (m *EventFeeMarket) XXX_Size() int {
	return m.Size()
}
func (m *EventFeeMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFeeMarket.DiscardUnknown(m)
}

var xxx_messageInfo_EventFeeMarket proto.InternalMessageInfo

func (m *EventFeeMarket) GetBaseFee() string {
	if m != nil {
		return m.BaseFee
	}
	return ""
}

// EventBlockGas defines an Ethereum block gas event
type EventBlockGas struct {
	// height of the block
	Height string `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
	// amount of gas wanted by the block
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventBlockGas) Reset()         { *m = EventBlockGas{} }
func (m *EventBlockGas) String() string { return proto.CompactTextString(m) }
func (*EventBlockGas) ProtoMessage()    {}
func (*EventBlockGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6edce8d670faff7, []int{1}
}
func (m *EventBlockGas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBlockGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBlockGas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBlockGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBlockGas.Merge(m, src)
}
func (m *EventBlockGas) XXX_Size() int {
	return m.Size()
}
func (m *EventBlockGas) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBlockGas.DiscardUnknown(m)
}

var xxx_messageInfo_EventBlockGas proto.InternalMessageInfo

func (m *EventBlockGas) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *EventBlockGas) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*EventFeeMarket)(nil), "ethermint.feemarket.v1.EventFeeMarket")
	proto.RegisterType((*EventBlockGas)(nil), "ethermint.feemarket.v1.EventBlockGas")
}

func init() {
	proto.RegisterFile("ethermint/feemarket/v1/events.proto", fileDescriptor_c6edce8d670faff7)
}

var fileDescriptor_c6edce8d670faff7 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2d, 0xc9, 0x48,
	0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0x4b, 0x4d, 0xcd, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd1,
	0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x83, 0x2b, 0xd2, 0x83, 0x2b, 0xd2, 0x2b, 0x33, 0x54, 0xd2, 0xe6, 0xe2, 0x73, 0x05, 0xa9,
	0x73, 0x4b, 0x4d, 0xf5, 0x05, 0x0b, 0x0a, 0x49, 0x72, 0x71, 0x24, 0x25, 0x16, 0xa7, 0xc6, 0xa7,
	0xa5, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x83, 0xf8, 0x6e, 0xa9, 0xa9, 0x4a,
	0xf6, 0x5c, 0xbc, 0x60, 0xc5, 0x4e, 0x39, 0xf9, 0xc9, 0xd9, 0xee, 0x89, 0xc5, 0x42, 0x62, 0x5c,
	0x6c, 0x19, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x50, 0x95, 0x50, 0x1e, 0x48, 0x3c, 0x31, 0x37, 0xbf,
	0x34, 0xaf, 0x44, 0x82, 0x09, 0x22, 0x0e, 0xe1, 0x39, 0xf9, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x4a, 0x6a, 0x72, 0x66, 0x6e, 0x62, 0x4e, 0x49, 0x6a, 0x62, 0xae, 0x3e, 0xc2, 0x6f, 0x15,
	0x48, 0xbe, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x7b, 0xcd, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0x0d, 0xc6, 0xea, 0x01, 0x01, 0x00, 0x00,
}

func (m *EventFeeMarket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFeeMarket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFeeMarket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BaseFee) > 0 {
		i -= len(m.BaseFee)
		copy(dAtA[i:], m.BaseFee)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BaseFee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBlockGas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBlockGas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBlockGas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Height) > 0 {
		i -= len(m.Height)
		copy(dAtA[i:], m.Height)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Height)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventFeeMarket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseFee)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventBlockGas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Height)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventFeeMarket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFeeMarket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFeeMarket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBlockGas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBlockGas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBlockGas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
