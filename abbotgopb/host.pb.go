// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: host.proto

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

type HostNetworkInterface struct {
	Metadata *NetworkInterface `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*HostNetworkInterface_Unknown
	//	*HostNetworkInterface_Bridge
	//	*HostNetworkInterface_Wireguard
	Config isHostNetworkInterface_Config `protobuf_oneof:"config"`
}

func (m *HostNetworkInterface) Reset()      { *m = HostNetworkInterface{} }
func (*HostNetworkInterface) ProtoMessage() {}
func (*HostNetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e40b83b4d50a8d, []int{0}
}
func (m *HostNetworkInterface) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostNetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostNetworkInterface.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostNetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostNetworkInterface.Merge(m, src)
}
func (m *HostNetworkInterface) XXX_Size() int {
	return m.Size()
}
func (m *HostNetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_HostNetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_HostNetworkInterface proto.InternalMessageInfo

type isHostNetworkInterface_Config interface {
	isHostNetworkInterface_Config()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type HostNetworkInterface_Unknown struct {
	Unknown *DriverUnknown `protobuf:"bytes,10,opt,name=unknown,proto3,oneof" json:"unknown,omitempty"`
}
type HostNetworkInterface_Bridge struct {
	Bridge *DriverBridge `protobuf:"bytes,11,opt,name=bridge,proto3,oneof" json:"bridge,omitempty"`
}
type HostNetworkInterface_Wireguard struct {
	Wireguard *DriverWireguard `protobuf:"bytes,12,opt,name=wireguard,proto3,oneof" json:"wireguard,omitempty"`
}

func (*HostNetworkInterface_Unknown) isHostNetworkInterface_Config()   {}
func (*HostNetworkInterface_Bridge) isHostNetworkInterface_Config()    {}
func (*HostNetworkInterface_Wireguard) isHostNetworkInterface_Config() {}

func (m *HostNetworkInterface) GetConfig() isHostNetworkInterface_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *HostNetworkInterface) GetMetadata() *NetworkInterface {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *HostNetworkInterface) GetUnknown() *DriverUnknown {
	if x, ok := m.GetConfig().(*HostNetworkInterface_Unknown); ok {
		return x.Unknown
	}
	return nil
}

func (m *HostNetworkInterface) GetBridge() *DriverBridge {
	if x, ok := m.GetConfig().(*HostNetworkInterface_Bridge); ok {
		return x.Bridge
	}
	return nil
}

func (m *HostNetworkInterface) GetWireguard() *DriverWireguard {
	if x, ok := m.GetConfig().(*HostNetworkInterface_Wireguard); ok {
		return x.Wireguard
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HostNetworkInterface) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HostNetworkInterface_Unknown)(nil),
		(*HostNetworkInterface_Bridge)(nil),
		(*HostNetworkInterface_Wireguard)(nil),
	}
}

type HostNetworkConfigEnsureRequest struct {
	Expected []*HostNetworkInterface `protobuf:"bytes,1,rep,name=expected,proto3" json:"expected,omitempty"`
}

func (m *HostNetworkConfigEnsureRequest) Reset()      { *m = HostNetworkConfigEnsureRequest{} }
func (*HostNetworkConfigEnsureRequest) ProtoMessage() {}
func (*HostNetworkConfigEnsureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e40b83b4d50a8d, []int{1}
}
func (m *HostNetworkConfigEnsureRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostNetworkConfigEnsureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostNetworkConfigEnsureRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostNetworkConfigEnsureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostNetworkConfigEnsureRequest.Merge(m, src)
}
func (m *HostNetworkConfigEnsureRequest) XXX_Size() int {
	return m.Size()
}
func (m *HostNetworkConfigEnsureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HostNetworkConfigEnsureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HostNetworkConfigEnsureRequest proto.InternalMessageInfo

func (m *HostNetworkConfigEnsureRequest) GetExpected() []*HostNetworkInterface {
	if m != nil {
		return m.Expected
	}
	return nil
}

type HostNetworkStatusResponse struct {
	Actual []*HostNetworkInterface `protobuf:"bytes,1,rep,name=actual,proto3" json:"actual,omitempty"`
}

func (m *HostNetworkStatusResponse) Reset()      { *m = HostNetworkStatusResponse{} }
func (*HostNetworkStatusResponse) ProtoMessage() {}
func (*HostNetworkStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e40b83b4d50a8d, []int{2}
}
func (m *HostNetworkStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostNetworkStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostNetworkStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostNetworkStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostNetworkStatusResponse.Merge(m, src)
}
func (m *HostNetworkStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *HostNetworkStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HostNetworkStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HostNetworkStatusResponse proto.InternalMessageInfo

func (m *HostNetworkStatusResponse) GetActual() []*HostNetworkInterface {
	if m != nil {
		return m.Actual
	}
	return nil
}

func init() {
	proto.RegisterType((*HostNetworkInterface)(nil), "abbot.HostNetworkInterface")
	proto.RegisterType((*HostNetworkConfigEnsureRequest)(nil), "abbot.HostNetworkConfigEnsureRequest")
	proto.RegisterType((*HostNetworkStatusResponse)(nil), "abbot.HostNetworkStatusResponse")
}

func init() { proto.RegisterFile("host.proto", fileDescriptor_85e40b83b4d50a8d) }

var fileDescriptor_85e40b83b4d50a8d = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xe2, 0x40,
	0x18, 0xc7, 0x67, 0xb2, 0xd9, 0x2e, 0x3b, 0xec, 0xa9, 0x10, 0xb6, 0xcb, 0x26, 0xb3, 0x84, 0x13,
	0x17, 0xba, 0x1b, 0x48, 0xd6, 0x3b, 0x6a, 0x52, 0x2f, 0xc6, 0xd4, 0x10, 0xa3, 0x17, 0x33, 0x6d,
	0x3f, 0x4a, 0x83, 0x76, 0xea, 0x74, 0x0a, 0x1e, 0x7d, 0x04, 0x1f, 0xc3, 0x47, 0xf1, 0xc8, 0x91,
	0xa3, 0x94, 0x8b, 0x27, 0xc3, 0x23, 0x18, 0xa6, 0x03, 0x82, 0xf1, 0xe0, 0xad, 0xf9, 0x7d, 0xbf,
	0x7f, 0xff, 0xdf, 0xcc, 0x10, 0x32, 0xe4, 0xa9, 0xb4, 0x13, 0xc1, 0x25, 0x37, 0xbf, 0x32, 0xcf,
	0xe3, 0xb2, 0x4e, 0xae, 0x41, 0xb2, 0x02, 0xd5, 0xab, 0x81, 0x88, 0xc6, 0x20, 0x2e, 0xb3, 0x78,
	0x14, 0xf3, 0x49, 0xac, 0x69, 0x45, 0x53, 0x4f, 0x44, 0x41, 0x08, 0x1a, 0xd6, 0x34, 0x9c, 0x44,
	0x02, 0xc2, 0x8c, 0x89, 0xa0, 0xe0, 0xcd, 0x17, 0x4c, 0xaa, 0x0e, 0x4f, 0xe5, 0x31, 0xc8, 0x09,
	0x17, 0xa3, 0xa3, 0x58, 0x82, 0x18, 0x30, 0x1f, 0xcc, 0x2e, 0x29, 0xad, 0x9a, 0x02, 0x26, 0x99,
	0x85, 0x1b, 0xb8, 0x55, 0xee, 0xfc, 0xb4, 0xd5, 0x06, 0xf6, 0x7b, 0xd5, 0xdd, 0x88, 0xe6, 0x3f,
	0xf2, 0x4d, 0xef, 0x62, 0x11, 0x95, 0xa9, 0xea, 0xcc, 0x81, 0x6a, 0xef, 0x17, 0x33, 0x07, 0xb9,
	0x6b, 0xcd, 0x6c, 0x13, 0xa3, 0xd8, 0xd3, 0x2a, 0xab, 0x40, 0x65, 0x27, 0xd0, 0x53, 0x23, 0x07,
	0xb9, 0x5a, 0x32, 0xff, 0x93, 0xef, 0x9b, 0x13, 0x58, 0x3f, 0x54, 0xa2, 0xb6, 0x93, 0x38, 0x5b,
	0x4f, 0x1d, 0xe4, 0xbe, 0xa9, 0xbd, 0x12, 0x31, 0x7c, 0x1e, 0x0f, 0xa2, 0xb0, 0x79, 0x4e, 0xe8,
	0xd6, 0x79, 0xf7, 0x15, 0x3c, 0x8c, 0xd3, 0x4c, 0x80, 0x0b, 0x37, 0x19, 0xa4, 0xd2, 0xdc, 0x23,
	0x25, 0xb8, 0x4d, 0xc0, 0x97, 0x10, 0x58, 0xb8, 0xf1, 0xa5, 0x55, 0xee, 0xfc, 0xd6, 0x15, 0x1f,
	0x5d, 0x94, 0xbb, 0x91, 0x9b, 0x27, 0xe4, 0xd7, 0x96, 0x71, 0x2a, 0x99, 0xcc, 0x52, 0x17, 0xd2,
	0x84, 0xc7, 0xe9, 0xea, 0x3e, 0x0d, 0xe6, 0xcb, 0x8c, 0x5d, 0x7d, 0xe6, 0x9f, 0x5a, 0xed, 0xf5,
	0xa7, 0x73, 0x8a, 0x66, 0x73, 0x8a, 0x96, 0x73, 0x8a, 0xef, 0x72, 0x8a, 0x1f, 0x72, 0x8a, 0x1f,
	0x73, 0x8a, 0xa7, 0x39, 0xc5, 0x4f, 0x39, 0xc5, 0xcf, 0x39, 0x45, 0xcb, 0x9c, 0xe2, 0xfb, 0x05,
	0x45, 0xd3, 0x05, 0x45, 0xb3, 0x05, 0x45, 0x17, 0x7f, 0x98, 0x18, 0x32, 0x69, 0x07, 0x30, 0xfe,
	0xab, 0x3a, 0xda, 0xea, 0xa9, 0x8b, 0xef, 0x90, 0x27, 0x9e, 0x67, 0x28, 0xd0, 0x7d, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0x8c, 0xfa, 0xf4, 0x5f, 0x02, 0x00, 0x00,
}

func (this *HostNetworkInterface) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HostNetworkInterface)
	if !ok {
		that2, ok := that.(HostNetworkInterface)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if that1.Config == nil {
		if this.Config != nil {
			return false
		}
	} else if this.Config == nil {
		return false
	} else if !this.Config.Equal(that1.Config) {
		return false
	}
	return true
}
func (this *HostNetworkInterface_Unknown) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HostNetworkInterface_Unknown)
	if !ok {
		that2, ok := that.(HostNetworkInterface_Unknown)
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
	if !this.Unknown.Equal(that1.Unknown) {
		return false
	}
	return true
}
func (this *HostNetworkInterface_Bridge) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HostNetworkInterface_Bridge)
	if !ok {
		that2, ok := that.(HostNetworkInterface_Bridge)
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
	if !this.Bridge.Equal(that1.Bridge) {
		return false
	}
	return true
}
func (this *HostNetworkInterface_Wireguard) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HostNetworkInterface_Wireguard)
	if !ok {
		that2, ok := that.(HostNetworkInterface_Wireguard)
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
	if !this.Wireguard.Equal(that1.Wireguard) {
		return false
	}
	return true
}
func (this *HostNetworkConfigEnsureRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HostNetworkConfigEnsureRequest)
	if !ok {
		that2, ok := that.(HostNetworkConfigEnsureRequest)
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
	if len(this.Expected) != len(that1.Expected) {
		return false
	}
	for i := range this.Expected {
		if !this.Expected[i].Equal(that1.Expected[i]) {
			return false
		}
	}
	return true
}
func (this *HostNetworkStatusResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HostNetworkStatusResponse)
	if !ok {
		that2, ok := that.(HostNetworkStatusResponse)
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
	if len(this.Actual) != len(that1.Actual) {
		return false
	}
	for i := range this.Actual {
		if !this.Actual[i].Equal(that1.Actual[i]) {
			return false
		}
	}
	return true
}
func (this *HostNetworkInterface) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&abbotgopb.HostNetworkInterface{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.Config != nil {
		s = append(s, "Config: "+fmt.Sprintf("%#v", this.Config)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HostNetworkInterface_Unknown) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&abbotgopb.HostNetworkInterface_Unknown{` +
		`Unknown:` + fmt.Sprintf("%#v", this.Unknown) + `}`}, ", ")
	return s
}
func (this *HostNetworkInterface_Bridge) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&abbotgopb.HostNetworkInterface_Bridge{` +
		`Bridge:` + fmt.Sprintf("%#v", this.Bridge) + `}`}, ", ")
	return s
}
func (this *HostNetworkInterface_Wireguard) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&abbotgopb.HostNetworkInterface_Wireguard{` +
		`Wireguard:` + fmt.Sprintf("%#v", this.Wireguard) + `}`}, ", ")
	return s
}
func (this *HostNetworkConfigEnsureRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&abbotgopb.HostNetworkConfigEnsureRequest{")
	if this.Expected != nil {
		s = append(s, "Expected: "+fmt.Sprintf("%#v", this.Expected)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HostNetworkStatusResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&abbotgopb.HostNetworkStatusResponse{")
	if this.Actual != nil {
		s = append(s, "Actual: "+fmt.Sprintf("%#v", this.Actual)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHost(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HostNetworkInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostNetworkInterface) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostNetworkInterface) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size := m.Config.Size()
			i -= size
			if _, err := m.Config.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HostNetworkInterface_Unknown) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostNetworkInterface_Unknown) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Unknown != nil {
		{
			size, err := m.Unknown.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *HostNetworkInterface_Bridge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostNetworkInterface_Bridge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Bridge != nil {
		{
			size, err := m.Bridge.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *HostNetworkInterface_Wireguard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostNetworkInterface_Wireguard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Wireguard != nil {
		{
			size, err := m.Wireguard.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *HostNetworkConfigEnsureRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostNetworkConfigEnsureRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostNetworkConfigEnsureRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Expected) > 0 {
		for iNdEx := len(m.Expected) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Expected[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHost(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *HostNetworkStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostNetworkStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostNetworkStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Actual) > 0 {
		for iNdEx := len(m.Actual) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actual[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHost(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	offset -= sovHost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostNetworkInterface) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	if m.Config != nil {
		n += m.Config.Size()
	}
	return n
}

func (m *HostNetworkInterface_Unknown) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Unknown != nil {
		l = m.Unknown.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}
func (m *HostNetworkInterface_Bridge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bridge != nil {
		l = m.Bridge.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}
func (m *HostNetworkInterface_Wireguard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Wireguard != nil {
		l = m.Wireguard.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}
func (m *HostNetworkConfigEnsureRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Expected) > 0 {
		for _, e := range m.Expected {
			l = e.Size()
			n += 1 + l + sovHost(uint64(l))
		}
	}
	return n
}

func (m *HostNetworkStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Actual) > 0 {
		for _, e := range m.Actual {
			l = e.Size()
			n += 1 + l + sovHost(uint64(l))
		}
	}
	return n
}

func sovHost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HostNetworkInterface) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HostNetworkInterface{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "NetworkInterface", "NetworkInterface", 1) + `,`,
		`Config:` + fmt.Sprintf("%v", this.Config) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HostNetworkInterface_Unknown) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HostNetworkInterface_Unknown{`,
		`Unknown:` + strings.Replace(fmt.Sprintf("%v", this.Unknown), "DriverUnknown", "DriverUnknown", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HostNetworkInterface_Bridge) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HostNetworkInterface_Bridge{`,
		`Bridge:` + strings.Replace(fmt.Sprintf("%v", this.Bridge), "DriverBridge", "DriverBridge", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HostNetworkInterface_Wireguard) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HostNetworkInterface_Wireguard{`,
		`Wireguard:` + strings.Replace(fmt.Sprintf("%v", this.Wireguard), "DriverWireguard", "DriverWireguard", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HostNetworkConfigEnsureRequest) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForExpected := "[]*HostNetworkInterface{"
	for _, f := range this.Expected {
		repeatedStringForExpected += strings.Replace(f.String(), "HostNetworkInterface", "HostNetworkInterface", 1) + ","
	}
	repeatedStringForExpected += "}"
	s := strings.Join([]string{`&HostNetworkConfigEnsureRequest{`,
		`Expected:` + repeatedStringForExpected + `,`,
		`}`,
	}, "")
	return s
}
func (this *HostNetworkStatusResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForActual := "[]*HostNetworkInterface{"
	for _, f := range this.Actual {
		repeatedStringForActual += strings.Replace(f.String(), "HostNetworkInterface", "HostNetworkInterface", 1) + ","
	}
	repeatedStringForActual += "}"
	s := strings.Join([]string{`&HostNetworkStatusResponse{`,
		`Actual:` + repeatedStringForActual + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHost(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HostNetworkInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostNetworkInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostNetworkInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &NetworkInterface{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unknown", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DriverUnknown{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &HostNetworkInterface_Unknown{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bridge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DriverBridge{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &HostNetworkInterface_Bridge{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wireguard", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DriverWireguard{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &HostNetworkInterface_Wireguard{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func (m *HostNetworkConfigEnsureRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostNetworkConfigEnsureRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostNetworkConfigEnsureRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expected", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expected = append(m.Expected, &HostNetworkInterface{})
			if err := m.Expected[len(m.Expected)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func (m *HostNetworkStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostNetworkStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostNetworkStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actual", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actual = append(m.Actual, &HostNetworkInterface{})
			if err := m.Actual[len(m.Actual)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHost = fmt.Errorf("proto: unexpected end of group")
)
