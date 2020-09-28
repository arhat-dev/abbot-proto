// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package abbotgopb

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
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

type RequestType int32

const (
	_INVALID_REQUEST              RequestType = 0
	REQ_ENSURE_CTR_NETWORK_CONFIG RequestType = 1
	REQ_ENSURE_CTR_NETWORK        RequestType = 2
	REQ_RESTORE_CTR_NETWORK       RequestType = 3
	REQ_DELETE_CTR_NETWORK        RequestType = 4
	REQ_QUERY_CTR_NETWORK         RequestType = 5
)

var RequestType_name = map[int32]string{
	0: "_INVALID_REQUEST",
	1: "REQ_ENSURE_CTR_NETWORK_CONFIG",
	2: "REQ_ENSURE_CTR_NETWORK",
	3: "REQ_RESTORE_CTR_NETWORK",
	4: "REQ_DELETE_CTR_NETWORK",
	5: "REQ_QUERY_CTR_NETWORK",
}

var RequestType_value = map[string]int32{
	"_INVALID_REQUEST":              0,
	"REQ_ENSURE_CTR_NETWORK_CONFIG": 1,
	"REQ_ENSURE_CTR_NETWORK":        2,
	"REQ_RESTORE_CTR_NETWORK":       3,
	"REQ_DELETE_CTR_NETWORK":        4,
	"REQ_QUERY_CTR_NETWORK":         5,
}

func (RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

type ResponseType int32

const (
	_INVALID_RESPONSE       ResponseType = 0
	RESP_DONE               ResponseType = 1
	RESP_CTR_NETWORK_STATUS ResponseType = 2
)

var ResponseType_name = map[int32]string{
	0: "_INVALID_RESPONSE",
	1: "RESP_DONE",
	2: "RESP_CTR_NETWORK_STATUS",
}

var ResponseType_value = map[string]int32{
	"_INVALID_RESPONSE":       0,
	"RESP_DONE":               1,
	"RESP_CTR_NETWORK_STATUS": 2,
}

func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

type Request struct {
	Kind RequestType `protobuf:"varint,1,opt,name=kind,proto3,enum=abbot.RequestType" json:"kind,omitempty"`
	Body []byte      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Request) Reset()      { *m = Request{} }
func (*Request) ProtoMessage() {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetKind() RequestType {
	if m != nil {
		return m.Kind
	}
	return _INVALID_REQUEST
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Kind ResponseType `protobuf:"varint,1,opt,name=kind,proto3,enum=abbot.ResponseType" json:"kind,omitempty"`
	Body []byte       `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()      { *m = Response{} }
func (*Response) ProtoMessage() {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetKind() ResponseType {
	if m != nil {
		return m.Kind
	}
	return _INVALID_RESPONSE
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("abbot.RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("abbot.ResponseType", ResponseType_name, ResponseType_value)
	proto.RegisterType((*Request)(nil), "abbot.Request")
	proto.RegisterType((*Response)(nil), "abbot.Response")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x6e, 0xda, 0x30,
	0x1c, 0xc7, 0x6d, 0x06, 0xfb, 0x63, 0xd8, 0x94, 0x79, 0x63, 0x63, 0x4c, 0xf3, 0x18, 0x87, 0x0d,
	0x21, 0x11, 0xa4, 0xed, 0x09, 0x18, 0x78, 0x08, 0x0d, 0x25, 0xc4, 0x76, 0x36, 0x6d, 0x97, 0x88,
	0x94, 0x34, 0xa0, 0xaa, 0x38, 0x85, 0x50, 0x89, 0x5b, 0x1f, 0xa1, 0x8f, 0xd1, 0x27, 0xe8, 0x33,
	0xf4, 0xc8, 0x91, 0x63, 0x09, 0x97, 0x1e, 0x79, 0x84, 0x0a, 0xa7, 0xad, 0x48, 0x55, 0xf5, 0x62,
	0x7d, 0xbf, 0xfe, 0xf8, 0xf7, 0xf5, 0xef, 0x67, 0xa3, 0x6c, 0x30, 0x91, 0xa1, 0xd4, 0xd5, 0x8a,
	0x33, 0x7d, 0xd7, 0x95, 0x61, 0xb1, 0xe6, 0x8f, 0xc2, 0xe1, 0xcc, 0xd5, 0xf7, 0xe4, 0x61, 0xdd,
	0x97, 0xbe, 0xac, 0x2b, 0xea, 0xce, 0xf6, 0x95, 0x53, 0x46, 0xa9, 0xb8, 0xaa, 0x4c, 0xd1, 0x33,
	0xe6, 0x1d, 0xcd, 0xbc, 0x69, 0x88, 0xbf, 0xa2, 0xf4, 0xc1, 0x68, 0x3c, 0x28, 0xc0, 0x12, 0xac,
	0xbc, 0xfa, 0x8e, 0x75, 0x95, 0xa7, 0xdf, 0x50, 0x31, 0x0f, 0x3c, 0xa6, 0x38, 0xc6, 0x28, 0xed,
	0xca, 0xc1, 0xbc, 0x90, 0x2a, 0xc1, 0x4a, 0x8e, 0x29, 0x5d, 0x6e, 0xa3, 0xe7, 0xcc, 0x9b, 0x06,
	0x72, 0x3c, 0xf5, 0xf0, 0xb7, 0x44, 0xce, 0x9b, 0xbb, 0x9c, 0x18, 0x3f, 0x1e, 0x54, 0x3d, 0x87,
	0x28, 0xbb, 0x73, 0x25, 0x7e, 0x8b, 0x34, 0xa7, 0x63, 0xfc, 0x69, 0x74, 0x3b, 0x2d, 0x87, 0x51,
	0xcb, 0xa6, 0x5c, 0x68, 0x00, 0x7f, 0x41, 0x9f, 0x18, 0xb5, 0x1c, 0x6a, 0x70, 0x9b, 0x51, 0xa7,
	0x29, 0x98, 0x63, 0x50, 0xf1, 0xd7, 0x64, 0xbf, 0x9d, 0xa6, 0x69, 0xfc, 0xea, 0xb4, 0x35, 0x88,
	0x8b, 0xe8, 0xdd, 0xc3, 0x47, 0xb4, 0x14, 0xfe, 0x88, 0xde, 0x6f, 0x19, 0xa3, 0x5c, 0x98, 0xf7,
	0xe0, 0x93, 0xdb, 0xc2, 0x16, 0xed, 0x52, 0x91, 0x64, 0x69, 0xfc, 0x01, 0xe5, 0xb7, 0xcc, 0xb2,
	0x29, 0xfb, 0x97, 0x40, 0x99, 0xaa, 0x85, 0x72, 0xbb, 0x23, 0xe2, 0x3c, 0x7a, 0xbd, 0xd3, 0x38,
	0xef, 0x99, 0x06, 0xa7, 0x1a, 0xc0, 0x2f, 0xd1, 0x8b, 0xad, 0x73, 0x5a, 0xa6, 0x41, 0x35, 0x18,
	0x77, 0xc2, 0x7b, 0x89, 0x11, 0xb8, 0x68, 0x08, 0x9b, 0x6b, 0xa9, 0x9f, 0xf6, 0x62, 0x45, 0xc0,
	0x72, 0x45, 0xc0, 0x66, 0x45, 0xe0, 0x49, 0x44, 0xe0, 0x59, 0x44, 0xe0, 0x45, 0x44, 0xe0, 0x22,
	0x22, 0xf0, 0x32, 0x22, 0xf0, 0x2a, 0x22, 0x60, 0x13, 0x11, 0x78, 0xba, 0x26, 0x60, 0xb1, 0x26,
	0x60, 0xb9, 0x26, 0xe0, 0xff, 0xe7, 0xfe, 0x64, 0xd8, 0x0f, 0xf5, 0x81, 0x77, 0x5c, 0x57, 0x2f,
	0x5f, 0x8b, 0xff, 0x5c, 0x69, 0x5f, 0x06, 0xae, 0xfb, 0x54, 0x6d, 0xfc, 0xb8, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0xc1, 0xb8, 0x9b, 0x3e, 0x02, 0x00, 0x00,
}

func (x RequestType) String() string {
	s, ok := RequestType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ResponseType) String() string {
	s, ok := ResponseType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request)
	if !ok {
		that2, ok := that.(Request)
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
	if this.Kind != that1.Kind {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response)
	if !ok {
		that2, ok := that.(Response)
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
	if this.Kind != that1.Kind {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *Request) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&abbotgopb.Request{")
	s = append(s, "Kind: "+fmt.Sprintf("%#v", this.Kind)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Response) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&abbotgopb.Response{")
	s = append(s, "Kind: "+fmt.Sprintf("%#v", this.Kind)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProto(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.Kind != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.Kind != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProto(dAtA []byte, offset int, v uint64) int {
	offset -= sovProto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovProto(uint64(m.Kind))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovProto(uint64(m.Kind))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func sovProto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProto(x uint64) (n int) {
	return sovProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Request{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Response{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProto(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= RequestType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= ResponseType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
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
func skipProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
				return 0, ErrInvalidLengthProto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProto = fmt.Errorf("proto: unexpected end of group")
)
