// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ticket_updated_event.proto

package types

import (
	encoding_binary "encoding/binary"
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

type TicketUpdatedEvent struct {
	ID      int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Title   string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Price   float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	UserID  int64   `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderID int64   `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (m *TicketUpdatedEvent) Reset()         { *m = TicketUpdatedEvent{} }
func (m *TicketUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*TicketUpdatedEvent) ProtoMessage()    {}
func (*TicketUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e1703c37432f0ee, []int{0}
}
func (m *TicketUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TicketUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TicketUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TicketUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketUpdatedEvent.Merge(m, src)
}
func (m *TicketUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *TicketUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TicketUpdatedEvent proto.InternalMessageInfo

func (m *TicketUpdatedEvent) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TicketUpdatedEvent) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TicketUpdatedEvent) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TicketUpdatedEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TicketUpdatedEvent) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *TicketUpdatedEvent) GetOrderID() int64 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

func init() {
	proto.RegisterType((*TicketUpdatedEvent)(nil), "types.TicketUpdatedEvent")
}

func init() { proto.RegisterFile("ticket_updated_event.proto", fileDescriptor_2e1703c37432f0ee) }

var fileDescriptor_2e1703c37432f0ee = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc6, 0x2f, 0x3d, 0xdb, 0x6a, 0xdc, 0x82, 0x48, 0xb8, 0x21, 0x2d, 0x0a, 0xd2, 0xc5, 0xbb,
	0xc1, 0x37, 0x38, 0xea, 0xd0, 0x49, 0x08, 0xde, 0x5c, 0x6c, 0xf3, 0xb7, 0x06, 0xf5, 0x52, 0xd2,
	0xf4, 0xc0, 0xb7, 0xf0, 0x8d, 0x5c, 0x1d, 0x6f, 0x74, 0x2a, 0x92, 0xbe, 0x88, 0x24, 0xc1, 0x2d,
	0xbf, 0xef, 0xf7, 0xe5, 0x0b, 0xc1, 0x2b, 0x23, 0xdb, 0x57, 0x30, 0xf5, 0xd8, 0x8b, 0x27, 0x03,
	0xa2, 0x86, 0x03, 0xec, 0xcd, 0xba, 0xd7, 0xca, 0x28, 0x12, 0x9b, 0x8f, 0x1e, 0x86, 0xd5, 0x6d,
	0x27, 0xcd, 0xcb, 0xd8, 0xac, 0x5b, 0xf5, 0xbe, 0xe9, 0x54, 0xa7, 0x36, 0xde, 0x36, 0xe3, 0xb3,
	0x27, 0x0f, 0xfe, 0x14, 0x6e, 0x5d, 0x7d, 0x21, 0x4c, 0x1e, 0xfd, 0xe8, 0x2e, 0x6c, 0xde, 0xbb,
	0x49, 0x72, 0x89, 0x23, 0x29, 0x28, 0xca, 0x51, 0xb1, 0xdc, 0x26, 0x76, 0xca, 0xa2, 0xaa, 0xe4,
	0x91, 0x14, 0x84, 0xe2, 0xf4, 0x00, 0x7a, 0x90, 0x6a, 0x4f, 0x23, 0x27, 0xf9, 0x3f, 0x92, 0x0b,
	0x1c, 0x1b, 0x69, 0xde, 0x80, 0x2e, 0x73, 0x54, 0x9c, 0xf1, 0x00, 0x2e, 0xed, 0xb5, 0x6c, 0x81,
	0x9e, 0xe4, 0xa8, 0x40, 0x3c, 0x00, 0xb9, 0xc6, 0xe9, 0x38, 0x80, 0xae, 0xa5, 0xa0, 0xb1, 0x7f,
	0x02, 0xdb, 0x29, 0x4b, 0x76, 0x03, 0xe8, 0xaa, 0xe4, 0x89, 0x53, 0x95, 0x20, 0x37, 0xf8, 0x54,
	0x69, 0x11, 0x5a, 0x89, 0x6f, 0x9d, 0xdb, 0x29, 0x4b, 0x1f, 0x5c, 0x56, 0x95, 0x3c, 0xf5, 0xb2,
	0x12, 0x5b, 0xfa, 0x6d, 0x19, 0x3a, 0x5a, 0x86, 0x7e, 0x2d, 0x43, 0x9f, 0x33, 0x5b, 0x1c, 0x67,
	0xb6, 0xf8, 0x99, 0xd9, 0xa2, 0x49, 0xfc, 0x17, 0xef, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa6,
	0x52, 0x6a, 0x9a, 0x36, 0x01, 0x00, 0x00,
}

func (m *TicketUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TicketUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TicketUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderID != 0 {
		i = encodeVarintTicketUpdatedEvent(dAtA, i, uint64(m.OrderID))
		i--
		dAtA[i] = 0x30
	}
	if m.UserID != 0 {
		i = encodeVarintTicketUpdatedEvent(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x28
	}
	if m.Price != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Price))))
		i--
		dAtA[i] = 0x21
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTicketUpdatedEvent(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Version != 0 {
		i = encodeVarintTicketUpdatedEvent(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintTicketUpdatedEvent(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicketUpdatedEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicketUpdatedEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TicketUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTicketUpdatedEvent(uint64(m.ID))
	}
	if m.Version != 0 {
		n += 1 + sovTicketUpdatedEvent(uint64(m.Version))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTicketUpdatedEvent(uint64(l))
	}
	if m.Price != 0 {
		n += 9
	}
	if m.UserID != 0 {
		n += 1 + sovTicketUpdatedEvent(uint64(m.UserID))
	}
	if m.OrderID != 0 {
		n += 1 + sovTicketUpdatedEvent(uint64(m.OrderID))
	}
	return n
}

func sovTicketUpdatedEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicketUpdatedEvent(x uint64) (n int) {
	return sovTicketUpdatedEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TicketUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicketUpdatedEvent
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
			return fmt.Errorf("proto: TicketUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TicketUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicketUpdatedEvent
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicketUpdatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicketUpdatedEvent
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
				return ErrInvalidLengthTicketUpdatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicketUpdatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Price = float64(math.Float64frombits(v))
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicketUpdatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			m.OrderID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicketUpdatedEvent
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
			skippy, err := skipTicketUpdatedEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicketUpdatedEvent
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
func skipTicketUpdatedEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicketUpdatedEvent
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
					return 0, ErrIntOverflowTicketUpdatedEvent
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
					return 0, ErrIntOverflowTicketUpdatedEvent
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
				return 0, ErrInvalidLengthTicketUpdatedEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicketUpdatedEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicketUpdatedEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicketUpdatedEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicketUpdatedEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicketUpdatedEvent = fmt.Errorf("proto: unexpected end of group")
)
