// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/payment_created_event.proto

package types

import (
	fmt "fmt"
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

type PaymentCreatedEvent struct {
	ID       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StripeID string `protobuf:"bytes,2,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	OrderID  int64  `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (m *PaymentCreatedEvent) Reset()         { *m = PaymentCreatedEvent{} }
func (m *PaymentCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*PaymentCreatedEvent) ProtoMessage()    {}
func (*PaymentCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_933184e04ebb04a5, []int{0}
}
func (m *PaymentCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentCreatedEvent.Merge(m, src)
}
func (m *PaymentCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *PaymentCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentCreatedEvent proto.InternalMessageInfo

func (m *PaymentCreatedEvent) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PaymentCreatedEvent) GetStripeID() string {
	if m != nil {
		return m.StripeID
	}
	return ""
}

func (m *PaymentCreatedEvent) GetOrderID() int64 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

func init() {
	proto.RegisterType((*PaymentCreatedEvent)(nil), "types.PaymentCreatedEvent")
}

func init() { proto.RegisterFile("types/payment_created_event.proto", fileDescriptor_933184e04ebb04a5) }

var fileDescriptor_933184e04ebb04a5 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd, 0x2b, 0x89, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0x49,
	0x4d, 0x89, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x2b, 0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5,
	0xd4, 0xc0, 0xc8, 0x25, 0x1c, 0x00, 0x31, 0xd5, 0x19, 0x62, 0xa8, 0x2b, 0xc8, 0x4c, 0x21, 0x31,
	0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x66, 0x27, 0xb6, 0x47, 0xf7, 0xe4, 0x99,
	0x3c, 0x5d, 0x82, 0x98, 0x32, 0x53, 0x84, 0x34, 0xb9, 0x38, 0x8b, 0x4b, 0x8a, 0x32, 0x0b, 0x52,
	0xe3, 0x33, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x9d, 0x78, 0x1e, 0xdd, 0x93, 0xe7, 0x08,
	0x06, 0x0b, 0x7a, 0xba, 0x04, 0x71, 0x40, 0xa4, 0x3d, 0x53, 0x84, 0xd4, 0xb8, 0x38, 0xf2, 0x8b,
	0x52, 0x52, 0x8b, 0x40, 0x2a, 0x99, 0xc1, 0x06, 0x71, 0x3f, 0xba, 0x27, 0xcf, 0xee, 0x0f, 0x12,
	0xf3, 0x74, 0x09, 0x62, 0x07, 0x4b, 0x7a, 0xa6, 0x38, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x94, 0xa8, 0x95, 0xfe, 0x00, 0x00, 0x00,
}

func (m *PaymentCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderID != 0 {
		i = encodeVarintPaymentCreatedEvent(dAtA, i, uint64(m.OrderID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.StripeID) > 0 {
		i -= len(m.StripeID)
		copy(dAtA[i:], m.StripeID)
		i = encodeVarintPaymentCreatedEvent(dAtA, i, uint64(len(m.StripeID)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintPaymentCreatedEvent(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPaymentCreatedEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovPaymentCreatedEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PaymentCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPaymentCreatedEvent(uint64(m.ID))
	}
	l = len(m.StripeID)
	if l > 0 {
		n += 1 + l + sovPaymentCreatedEvent(uint64(l))
	}
	if m.OrderID != 0 {
		n += 1 + sovPaymentCreatedEvent(uint64(m.OrderID))
	}
	return n
}

func sovPaymentCreatedEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPaymentCreatedEvent(x uint64) (n int) {
	return sovPaymentCreatedEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PaymentCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaymentCreatedEvent
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
			return fmt.Errorf("proto: PaymentCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentCreatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StripeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentCreatedEvent
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
				return ErrInvalidLengthPaymentCreatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentCreatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StripeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			m.OrderID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentCreatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPaymentCreatedEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPaymentCreatedEvent
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
func skipPaymentCreatedEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPaymentCreatedEvent
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
					return 0, ErrIntOverflowPaymentCreatedEvent
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
					return 0, ErrIntOverflowPaymentCreatedEvent
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
				return 0, ErrInvalidLengthPaymentCreatedEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPaymentCreatedEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPaymentCreatedEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPaymentCreatedEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPaymentCreatedEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPaymentCreatedEvent = fmt.Errorf("proto: unexpected end of group")
)
