// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: driver_unknown.proto

package abbotgopb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type DriverUnknown struct {
}

func (m *DriverUnknown) Reset()      { *m = DriverUnknown{} }
func (*DriverUnknown) ProtoMessage() {}
func (*DriverUnknown) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b8b0b173a9cb832, []int{0}
}
func (m *DriverUnknown) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverUnknown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverUnknown.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverUnknown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverUnknown.Merge(m, src)
}
func (m *DriverUnknown) XXX_Size() int {
	return m.Size()
}
func (m *DriverUnknown) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverUnknown.DiscardUnknown(m)
}

var xxx_messageInfo_DriverUnknown proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DriverUnknown)(nil), "abbot.DriverUnknown")
}

func init() { proto.RegisterFile("driver_unknown.proto", fileDescriptor_1b8b0b173a9cb832) }

var fileDescriptor_1b8b0b173a9cb832 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x29, 0xca, 0x2c,
	0x4b, 0x2d, 0x8a, 0x2f, 0xcd, 0xcb, 0xce, 0xcb, 0x2f, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4d, 0x4c, 0x4a, 0xca, 0x2f, 0x51, 0xe2, 0xe7, 0xe2, 0x75, 0x01, 0x4b, 0x87, 0x42,
	0x64, 0x9d, 0x42, 0x2f, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50, 0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6,
	0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xc9, 0x27, 0x16, 0x65, 0x24,
	0x96, 0xe8, 0xa5, 0xa4, 0x96, 0xe9, 0x83, 0xcd, 0xd5, 0x05, 0x5b, 0x02, 0x61, 0xa7, 0xe7, 0x17,
	0x24, 0x25, 0xb1, 0x81, 0x05, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2e, 0x58, 0x33,
	0x8d, 0x00, 0x00, 0x00,
}

func (this *DriverUnknown) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DriverUnknown)
	if !ok {
		that2, ok := that.(DriverUnknown)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *DriverUnknown) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&abbotgopb.DriverUnknown{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDriverUnknown(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DriverUnknown) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverUnknown) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverUnknown) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintDriverUnknown(dAtA []byte, offset int, v uint64) int {
	offset -= sovDriverUnknown(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DriverUnknown) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovDriverUnknown(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDriverUnknown(x uint64) (n int) {
	return sovDriverUnknown(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DriverUnknown) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DriverUnknown{`,
		`}`,
	}, "")
	return s
}
func valueToStringDriverUnknown(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DriverUnknown) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriverUnknown
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
			return fmt.Errorf("proto: DriverUnknown: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DriverUnknown: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDriverUnknown(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriverUnknown
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
func skipDriverUnknown(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDriverUnknown
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
					return 0, ErrIntOverflowDriverUnknown
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
					return 0, ErrIntOverflowDriverUnknown
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
				return 0, ErrInvalidLengthDriverUnknown
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDriverUnknown
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDriverUnknown
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDriverUnknown        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDriverUnknown          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDriverUnknown = fmt.Errorf("proto: unexpected end of group")
)
