// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: driver_wireguard.proto

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

type DriverWireguard struct {
	// LogLevel for wireguard output
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// PrivateKey specifies a private key configuration
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// ListenPort specifies a device's listening port, if >= 0.
	ListenPort int32                    `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	Routing    *DriverWireguard_Routing `protobuf:"bytes,4,opt,name=routing,proto3" json:"routing,omitempty"`
	// Peers specifies a list of peer configurations to apply to a device.
	Peers []*DriverWireguard_Peer `protobuf:"bytes,5,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (m *DriverWireguard) Reset()      { *m = DriverWireguard{} }
func (*DriverWireguard) ProtoMessage() {}
func (*DriverWireguard) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dcf6bd0cf28bcc1, []int{0}
}
func (m *DriverWireguard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverWireguard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverWireguard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverWireguard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverWireguard.Merge(m, src)
}
func (m *DriverWireguard) XXX_Size() int {
	return m.Size()
}
func (m *DriverWireguard) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverWireguard.DiscardUnknown(m)
}

var xxx_messageInfo_DriverWireguard proto.InternalMessageInfo

func (m *DriverWireguard) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *DriverWireguard) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *DriverWireguard) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *DriverWireguard) GetRouting() *DriverWireguard_Routing {
	if m != nil {
		return m.Routing
	}
	return nil
}

func (m *DriverWireguard) GetPeers() []*DriverWireguard_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type DriverWireguard_Peer struct {
	PublicKey    string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PreSharedKey string `protobuf:"bytes,2,opt,name=pre_shared_key,json=preSharedKey,proto3" json:"pre_shared_key,omitempty"`
	Endpoint     string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// duration in nanoseconds
	PersistentKeepaliveInterval int64    `protobuf:"varint,4,opt,name=persistent_keepalive_interval,json=persistentKeepaliveInterval,proto3" json:"persistent_keepalive_interval,omitempty"`
	AllowedIps                  []string `protobuf:"bytes,5,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
}

func (m *DriverWireguard_Peer) Reset()      { *m = DriverWireguard_Peer{} }
func (*DriverWireguard_Peer) ProtoMessage() {}
func (*DriverWireguard_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dcf6bd0cf28bcc1, []int{0, 0}
}
func (m *DriverWireguard_Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverWireguard_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverWireguard_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverWireguard_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverWireguard_Peer.Merge(m, src)
}
func (m *DriverWireguard_Peer) XXX_Size() int {
	return m.Size()
}
func (m *DriverWireguard_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverWireguard_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_DriverWireguard_Peer proto.InternalMessageInfo

func (m *DriverWireguard_Peer) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *DriverWireguard_Peer) GetPreSharedKey() string {
	if m != nil {
		return m.PreSharedKey
	}
	return ""
}

func (m *DriverWireguard_Peer) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *DriverWireguard_Peer) GetPersistentKeepaliveInterval() int64 {
	if m != nil {
		return m.PersistentKeepaliveInterval
	}
	return 0
}

func (m *DriverWireguard_Peer) GetAllowedIps() []string {
	if m != nil {
		return m.AllowedIps
	}
	return nil
}

type DriverWireguard_Routing struct {
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Table to add routes, if not set (0), will use default route table
	Table int32 `protobuf:"varint,2,opt,name=table,proto3" json:"table,omitempty"`
	// FirewallMark specifies a device's firewall mark
	FirewallMark int32 `protobuf:"varint,3,opt,name=firewall_mark,json=firewallMark,proto3" json:"firewall_mark,omitempty"`
}

func (m *DriverWireguard_Routing) Reset()      { *m = DriverWireguard_Routing{} }
func (*DriverWireguard_Routing) ProtoMessage() {}
func (*DriverWireguard_Routing) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dcf6bd0cf28bcc1, []int{0, 1}
}
func (m *DriverWireguard_Routing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverWireguard_Routing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverWireguard_Routing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverWireguard_Routing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverWireguard_Routing.Merge(m, src)
}
func (m *DriverWireguard_Routing) XXX_Size() int {
	return m.Size()
}
func (m *DriverWireguard_Routing) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverWireguard_Routing.DiscardUnknown(m)
}

var xxx_messageInfo_DriverWireguard_Routing proto.InternalMessageInfo

func (m *DriverWireguard_Routing) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DriverWireguard_Routing) GetTable() int32 {
	if m != nil {
		return m.Table
	}
	return 0
}

func (m *DriverWireguard_Routing) GetFirewallMark() int32 {
	if m != nil {
		return m.FirewallMark
	}
	return 0
}

func init() {
	proto.RegisterType((*DriverWireguard)(nil), "abbot.DriverWireguard")
	proto.RegisterType((*DriverWireguard_Peer)(nil), "abbot.DriverWireguard.Peer")
	proto.RegisterType((*DriverWireguard_Routing)(nil), "abbot.DriverWireguard.Routing")
}

func init() { proto.RegisterFile("driver_wireguard.proto", fileDescriptor_5dcf6bd0cf28bcc1) }

var fileDescriptor_5dcf6bd0cf28bcc1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x8e, 0xd3, 0x30,
	0x18, 0xc7, 0x63, 0xda, 0xd0, 0xd6, 0x3d, 0x40, 0xb2, 0x10, 0x8a, 0x5a, 0x9d, 0xaf, 0x02, 0x86,
	0x2e, 0x04, 0x71, 0x2c, 0xcc, 0x27, 0x96, 0xd3, 0x81, 0x74, 0x32, 0x42, 0x48, 0x0c, 0x44, 0x0e,
	0xf9, 0xc8, 0x59, 0x35, 0xb1, 0xf5, 0xc5, 0x4d, 0x75, 0x1b, 0x0f, 0xc0, 0xc0, 0x63, 0xf0, 0x28,
	0x0c, 0x0c, 0x1d, 0x6f, 0xa4, 0xe9, 0xc2, 0x78, 0x8f, 0x80, 0xe2, 0x24, 0x80, 0x90, 0xd8, 0xfc,
	0xff, 0xe7, 0x17, 0xe9, 0xfb, 0xf9, 0x33, 0xbd, 0x97, 0xa1, 0xaa, 0x00, 0x93, 0x8d, 0x42, 0xc8,
	0xd7, 0x12, 0xb3, 0xd8, 0xa2, 0x71, 0x86, 0x85, 0x32, 0x4d, 0x8d, 0xbb, 0xff, 0x79, 0x48, 0xef,
	0x3c, 0xf7, 0xc4, 0x9b, 0x1e, 0x60, 0x73, 0x3a, 0xd1, 0x26, 0x4f, 0x34, 0x54, 0xa0, 0x23, 0xb2,
	0x20, 0xcb, 0x89, 0x18, 0x6b, 0x93, 0xbf, 0x68, 0x32, 0x3b, 0xa2, 0x53, 0x8b, 0xaa, 0x92, 0x0e,
	0x92, 0x15, 0x5c, 0x46, 0x37, 0xfc, 0x67, 0xda, 0x55, 0x67, 0x70, 0xd9, 0x00, 0x5a, 0x95, 0x0e,
	0x8a, 0xc4, 0x1a, 0x74, 0xd1, 0x60, 0x41, 0x96, 0xa1, 0xa0, 0x6d, 0x75, 0x6e, 0xd0, 0xb1, 0x67,
	0x74, 0x84, 0x66, 0xed, 0x54, 0x91, 0x47, 0xc3, 0x05, 0x59, 0x4e, 0x8f, 0x79, 0xec, 0x67, 0x89,
	0xff, 0x99, 0x23, 0x16, 0x2d, 0x25, 0x7a, 0x9c, 0x3d, 0xa1, 0xa1, 0x05, 0xc0, 0x32, 0x0a, 0x17,
	0x83, 0xe5, 0xf4, 0x78, 0xfe, 0x9f, 0xff, 0xce, 0x01, 0x50, 0xb4, 0xe4, 0xec, 0x3b, 0xa1, 0xc3,
	0x26, 0xb3, 0x43, 0x4a, 0xed, 0x3a, 0xd5, 0xea, 0xbd, 0x1f, 0xbb, 0xb5, 0x9a, 0xb4, 0x4d, 0x33,
	0xf5, 0x43, 0x7a, 0xdb, 0x22, 0x24, 0xe5, 0x85, 0x44, 0xc8, 0xfe, 0x32, 0x3b, 0xb0, 0x08, 0xaf,
	0x7c, 0xd9, 0x50, 0x33, 0x3a, 0x86, 0x22, 0xb3, 0x46, 0x15, 0xad, 0xd8, 0x44, 0xfc, 0xce, 0xec,
	0x84, 0x1e, 0x5a, 0xc0, 0xd2, 0x7b, 0xba, 0x64, 0x05, 0x60, 0xa5, 0x56, 0x15, 0x24, 0xaa, 0x70,
	0x80, 0x95, 0xd4, 0x5e, 0x76, 0x20, 0xe6, 0x7f, 0xa0, 0xb3, 0x9e, 0x39, 0xed, 0x90, 0xe6, 0xee,
	0xa4, 0xd6, 0x66, 0x03, 0x59, 0xa2, 0x6c, 0xab, 0x39, 0x11, 0xb4, 0xab, 0x4e, 0x6d, 0x39, 0x7b,
	0x47, 0x47, 0xdd, 0xad, 0xb0, 0x88, 0x8e, 0xa0, 0x90, 0xa9, 0x86, 0xcc, 0xdb, 0x8c, 0x45, 0x1f,
	0xd9, 0x5d, 0x1a, 0xba, 0xe6, 0xe4, 0x15, 0x42, 0xd1, 0x06, 0xf6, 0x80, 0xde, 0xfa, 0xa0, 0x10,
	0x36, 0x52, 0xeb, 0xe4, 0xa3, 0xc4, 0x55, 0xb7, 0x99, 0x83, 0xbe, 0x7c, 0x29, 0x71, 0x75, 0xf2,
	0x7a, 0xbb, 0xe3, 0xc1, 0xd5, 0x8e, 0x07, 0xd7, 0x3b, 0x4e, 0x3e, 0xd5, 0x9c, 0x7c, 0xad, 0x39,
	0xf9, 0x56, 0x73, 0xb2, 0xad, 0x39, 0xf9, 0x51, 0x73, 0xf2, 0xb3, 0xe6, 0xc1, 0x75, 0xcd, 0xc9,
	0x97, 0x3d, 0x0f, 0xb6, 0x7b, 0x1e, 0x5c, 0xed, 0x79, 0xf0, 0xf6, 0x48, 0xe2, 0x85, 0x74, 0x71,
	0x06, 0xd5, 0x63, 0xbf, 0x91, 0x47, 0xfe, 0x89, 0xb5, 0xe7, 0xdc, 0xd8, 0x34, 0xbd, 0xe9, 0x8b,
	0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x69, 0x9a, 0xa9, 0x67, 0x8d, 0x02, 0x00, 0x00,
}

func (this *DriverWireguard) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DriverWireguard)
	if !ok {
		that2, ok := that.(DriverWireguard)
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
	if this.LogLevel != that1.LogLevel {
		return false
	}
	if this.PrivateKey != that1.PrivateKey {
		return false
	}
	if this.ListenPort != that1.ListenPort {
		return false
	}
	if !this.Routing.Equal(that1.Routing) {
		return false
	}
	if len(this.Peers) != len(that1.Peers) {
		return false
	}
	for i := range this.Peers {
		if !this.Peers[i].Equal(that1.Peers[i]) {
			return false
		}
	}
	return true
}
func (this *DriverWireguard_Peer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DriverWireguard_Peer)
	if !ok {
		that2, ok := that.(DriverWireguard_Peer)
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
	if this.PublicKey != that1.PublicKey {
		return false
	}
	if this.PreSharedKey != that1.PreSharedKey {
		return false
	}
	if this.Endpoint != that1.Endpoint {
		return false
	}
	if this.PersistentKeepaliveInterval != that1.PersistentKeepaliveInterval {
		return false
	}
	if len(this.AllowedIps) != len(that1.AllowedIps) {
		return false
	}
	for i := range this.AllowedIps {
		if this.AllowedIps[i] != that1.AllowedIps[i] {
			return false
		}
	}
	return true
}
func (this *DriverWireguard_Routing) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DriverWireguard_Routing)
	if !ok {
		that2, ok := that.(DriverWireguard_Routing)
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
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.Table != that1.Table {
		return false
	}
	if this.FirewallMark != that1.FirewallMark {
		return false
	}
	return true
}
func (this *DriverWireguard) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&abbotgopb.DriverWireguard{")
	s = append(s, "LogLevel: "+fmt.Sprintf("%#v", this.LogLevel)+",\n")
	s = append(s, "PrivateKey: "+fmt.Sprintf("%#v", this.PrivateKey)+",\n")
	s = append(s, "ListenPort: "+fmt.Sprintf("%#v", this.ListenPort)+",\n")
	if this.Routing != nil {
		s = append(s, "Routing: "+fmt.Sprintf("%#v", this.Routing)+",\n")
	}
	if this.Peers != nil {
		s = append(s, "Peers: "+fmt.Sprintf("%#v", this.Peers)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DriverWireguard_Peer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&abbotgopb.DriverWireguard_Peer{")
	s = append(s, "PublicKey: "+fmt.Sprintf("%#v", this.PublicKey)+",\n")
	s = append(s, "PreSharedKey: "+fmt.Sprintf("%#v", this.PreSharedKey)+",\n")
	s = append(s, "Endpoint: "+fmt.Sprintf("%#v", this.Endpoint)+",\n")
	s = append(s, "PersistentKeepaliveInterval: "+fmt.Sprintf("%#v", this.PersistentKeepaliveInterval)+",\n")
	s = append(s, "AllowedIps: "+fmt.Sprintf("%#v", this.AllowedIps)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DriverWireguard_Routing) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&abbotgopb.DriverWireguard_Routing{")
	s = append(s, "Enabled: "+fmt.Sprintf("%#v", this.Enabled)+",\n")
	s = append(s, "Table: "+fmt.Sprintf("%#v", this.Table)+",\n")
	s = append(s, "FirewallMark: "+fmt.Sprintf("%#v", this.FirewallMark)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDriverWireguard(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DriverWireguard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverWireguard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverWireguard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Peers) > 0 {
		for iNdEx := len(m.Peers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Peers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDriverWireguard(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Routing != nil {
		{
			size, err := m.Routing.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDriverWireguard(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ListenPort != 0 {
		i = encodeVarintDriverWireguard(dAtA, i, uint64(m.ListenPort))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PrivateKey) > 0 {
		i -= len(m.PrivateKey)
		copy(dAtA[i:], m.PrivateKey)
		i = encodeVarintDriverWireguard(dAtA, i, uint64(len(m.PrivateKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LogLevel) > 0 {
		i -= len(m.LogLevel)
		copy(dAtA[i:], m.LogLevel)
		i = encodeVarintDriverWireguard(dAtA, i, uint64(len(m.LogLevel)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DriverWireguard_Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverWireguard_Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverWireguard_Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedIps) > 0 {
		for iNdEx := len(m.AllowedIps) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedIps[iNdEx])
			copy(dAtA[i:], m.AllowedIps[iNdEx])
			i = encodeVarintDriverWireguard(dAtA, i, uint64(len(m.AllowedIps[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PersistentKeepaliveInterval != 0 {
		i = encodeVarintDriverWireguard(dAtA, i, uint64(m.PersistentKeepaliveInterval))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintDriverWireguard(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PreSharedKey) > 0 {
		i -= len(m.PreSharedKey)
		copy(dAtA[i:], m.PreSharedKey)
		i = encodeVarintDriverWireguard(dAtA, i, uint64(len(m.PreSharedKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintDriverWireguard(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DriverWireguard_Routing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverWireguard_Routing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverWireguard_Routing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FirewallMark != 0 {
		i = encodeVarintDriverWireguard(dAtA, i, uint64(m.FirewallMark))
		i--
		dAtA[i] = 0x18
	}
	if m.Table != 0 {
		i = encodeVarintDriverWireguard(dAtA, i, uint64(m.Table))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDriverWireguard(dAtA []byte, offset int, v uint64) int {
	offset -= sovDriverWireguard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DriverWireguard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LogLevel)
	if l > 0 {
		n += 1 + l + sovDriverWireguard(uint64(l))
	}
	l = len(m.PrivateKey)
	if l > 0 {
		n += 1 + l + sovDriverWireguard(uint64(l))
	}
	if m.ListenPort != 0 {
		n += 1 + sovDriverWireguard(uint64(m.ListenPort))
	}
	if m.Routing != nil {
		l = m.Routing.Size()
		n += 1 + l + sovDriverWireguard(uint64(l))
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovDriverWireguard(uint64(l))
		}
	}
	return n
}

func (m *DriverWireguard_Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovDriverWireguard(uint64(l))
	}
	l = len(m.PreSharedKey)
	if l > 0 {
		n += 1 + l + sovDriverWireguard(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovDriverWireguard(uint64(l))
	}
	if m.PersistentKeepaliveInterval != 0 {
		n += 1 + sovDriverWireguard(uint64(m.PersistentKeepaliveInterval))
	}
	if len(m.AllowedIps) > 0 {
		for _, s := range m.AllowedIps {
			l = len(s)
			n += 1 + l + sovDriverWireguard(uint64(l))
		}
	}
	return n
}

func (m *DriverWireguard_Routing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.Table != 0 {
		n += 1 + sovDriverWireguard(uint64(m.Table))
	}
	if m.FirewallMark != 0 {
		n += 1 + sovDriverWireguard(uint64(m.FirewallMark))
	}
	return n
}

func sovDriverWireguard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDriverWireguard(x uint64) (n int) {
	return sovDriverWireguard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DriverWireguard) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPeers := "[]*DriverWireguard_Peer{"
	for _, f := range this.Peers {
		repeatedStringForPeers += strings.Replace(fmt.Sprintf("%v", f), "DriverWireguard_Peer", "DriverWireguard_Peer", 1) + ","
	}
	repeatedStringForPeers += "}"
	s := strings.Join([]string{`&DriverWireguard{`,
		`LogLevel:` + fmt.Sprintf("%v", this.LogLevel) + `,`,
		`PrivateKey:` + fmt.Sprintf("%v", this.PrivateKey) + `,`,
		`ListenPort:` + fmt.Sprintf("%v", this.ListenPort) + `,`,
		`Routing:` + strings.Replace(fmt.Sprintf("%v", this.Routing), "DriverWireguard_Routing", "DriverWireguard_Routing", 1) + `,`,
		`Peers:` + repeatedStringForPeers + `,`,
		`}`,
	}, "")
	return s
}
func (this *DriverWireguard_Peer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DriverWireguard_Peer{`,
		`PublicKey:` + fmt.Sprintf("%v", this.PublicKey) + `,`,
		`PreSharedKey:` + fmt.Sprintf("%v", this.PreSharedKey) + `,`,
		`Endpoint:` + fmt.Sprintf("%v", this.Endpoint) + `,`,
		`PersistentKeepaliveInterval:` + fmt.Sprintf("%v", this.PersistentKeepaliveInterval) + `,`,
		`AllowedIps:` + fmt.Sprintf("%v", this.AllowedIps) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DriverWireguard_Routing) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DriverWireguard_Routing{`,
		`Enabled:` + fmt.Sprintf("%v", this.Enabled) + `,`,
		`Table:` + fmt.Sprintf("%v", this.Table) + `,`,
		`FirewallMark:` + fmt.Sprintf("%v", this.FirewallMark) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDriverWireguard(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DriverWireguard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriverWireguard
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
			return fmt.Errorf("proto: DriverWireguard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DriverWireguard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogLevel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogLevel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenPort", wireType)
			}
			m.ListenPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ListenPort |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Routing == nil {
				m.Routing = &DriverWireguard_Routing{}
			}
			if err := m.Routing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &DriverWireguard_Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriverWireguard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDriverWireguard
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
func (m *DriverWireguard_Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriverWireguard
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreSharedKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreSharedKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PersistentKeepaliveInterval", wireType)
			}
			m.PersistentKeepaliveInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PersistentKeepaliveInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedIps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
				return ErrInvalidLengthDriverWireguard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedIps = append(m.AllowedIps, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriverWireguard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDriverWireguard
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
func (m *DriverWireguard_Routing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriverWireguard
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
			return fmt.Errorf("proto: Routing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Routing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Table", wireType)
			}
			m.Table = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Table |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirewallMark", wireType)
			}
			m.FirewallMark = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriverWireguard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirewallMark |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDriverWireguard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDriverWireguard
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDriverWireguard
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
func skipDriverWireguard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDriverWireguard
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
					return 0, ErrIntOverflowDriverWireguard
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
					return 0, ErrIntOverflowDriverWireguard
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
				return 0, ErrInvalidLengthDriverWireguard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDriverWireguard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDriverWireguard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDriverWireguard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDriverWireguard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDriverWireguard = fmt.Errorf("proto: unexpected end of group")
)
