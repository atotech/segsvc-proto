// Code generated by protoc-gen-gogo.
// source: segsvc.proto
// DO NOT EDIT!

/*
	Package segsvc is a generated protocol buffer package.

	It is generated from these files:
		segsvc.proto

	It has these top-level messages:
		User
		GetResponse
		ScoreData
		GeoData
*/
package segsvc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Reward int32

const (
	ERR   Reward = 0
	ZERO  Reward = 1
	ONE   Reward = 2
	TWO   Reward = 3
	THREE Reward = 4
	FOUR  Reward = 5
	FIVE  Reward = 6
)

var Reward_name = map[int32]string{
	0: "ERR",
	1: "ZERO",
	2: "ONE",
	3: "TWO",
	4: "THREE",
	5: "FOUR",
	6: "FIVE",
}
var Reward_value = map[string]int32{
	"ERR":   0,
	"ZERO":  1,
	"ONE":   2,
	"TWO":   3,
	"THREE": 4,
	"FOUR":  5,
	"FIVE":  6,
}

func (Reward) EnumDescriptor() ([]byte, []int) { return fileDescriptorSegsvc, []int{0} }

type User struct {
	UserID string     `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	URI    string     `protobuf:"bytes,2,opt,name=URI,proto3" json:"URI,omitempty"`
	Scores *ScoreData `protobuf:"bytes,3,opt,name=Scores" json:"Scores,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorSegsvc, []int{0} }

func (m *User) GetScores() *ScoreData {
	if m != nil {
		return m.Scores
	}
	return nil
}

type GetResponse struct {
	Segment string `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptorSegsvc, []int{1} }

type ScoreData struct {
	Segment     map[string]int64 `protobuf:"bytes,1,rep,name=segment" json:"segment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Geo         *GeoData         `protobuf:"bytes,2,opt,name=geo" json:"geo,omitempty"`
	RewardLevel Reward           `protobuf:"varint,3,opt,name=rewardLevel,proto3,enum=Reward" json:"rewardLevel,omitempty"`
	CRM         string           `protobuf:"bytes,4,opt,name=CRM,proto3" json:"CRM,omitempty"`
}

func (m *ScoreData) Reset()                    { *m = ScoreData{} }
func (*ScoreData) ProtoMessage()               {}
func (*ScoreData) Descriptor() ([]byte, []int) { return fileDescriptorSegsvc, []int{2} }

func (m *ScoreData) GetSegment() map[string]int64 {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *ScoreData) GetGeo() *GeoData {
	if m != nil {
		return m.Geo
	}
	return nil
}

type GeoData struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	ZipCode   string `protobuf:"bytes,2,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	DMZ       string `protobuf:"bytes,3,opt,name=DMZ,proto3" json:"DMZ,omitempty"`
}

func (m *GeoData) Reset()                    { *m = GeoData{} }
func (*GeoData) ProtoMessage()               {}
func (*GeoData) Descriptor() ([]byte, []int) { return fileDescriptorSegsvc, []int{3} }

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterType((*ScoreData)(nil), "ScoreData")
	proto.RegisterType((*GeoData)(nil), "GeoData")
	proto.RegisterEnum("Reward", Reward_name, Reward_value)
}
func (x Reward) String() string {
	s, ok := Reward_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *User) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UserID != that1.UserID {
		return false
	}
	if this.URI != that1.URI {
		return false
	}
	if !this.Scores.Equal(that1.Scores) {
		return false
	}
	return true
}
func (this *GetResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetResponse)
	if !ok {
		that2, ok := that.(GetResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Segment != that1.Segment {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *ScoreData) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ScoreData)
	if !ok {
		that2, ok := that.(ScoreData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Segment) != len(that1.Segment) {
		return false
	}
	for i := range this.Segment {
		if this.Segment[i] != that1.Segment[i] {
			return false
		}
	}
	if !this.Geo.Equal(that1.Geo) {
		return false
	}
	if this.RewardLevel != that1.RewardLevel {
		return false
	}
	if this.CRM != that1.CRM {
		return false
	}
	return true
}
func (this *GeoData) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GeoData)
	if !ok {
		that2, ok := that.(GeoData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.IpAddress != that1.IpAddress {
		return false
	}
	if this.ZipCode != that1.ZipCode {
		return false
	}
	if this.DMZ != that1.DMZ {
		return false
	}
	return true
}
func (this *User) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&segsvc.User{")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "URI: "+fmt.Sprintf("%#v", this.URI)+",\n")
	if this.Scores != nil {
		s = append(s, "Scores: "+fmt.Sprintf("%#v", this.Scores)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&segsvc.GetResponse{")
	s = append(s, "Segment: "+fmt.Sprintf("%#v", this.Segment)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ScoreData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&segsvc.ScoreData{")
	keysForSegment := make([]string, 0, len(this.Segment))
	for k, _ := range this.Segment {
		keysForSegment = append(keysForSegment, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForSegment)
	mapStringForSegment := "map[string]int64{"
	for _, k := range keysForSegment {
		mapStringForSegment += fmt.Sprintf("%#v: %#v,", k, this.Segment[k])
	}
	mapStringForSegment += "}"
	if this.Segment != nil {
		s = append(s, "Segment: "+mapStringForSegment+",\n")
	}
	if this.Geo != nil {
		s = append(s, "Geo: "+fmt.Sprintf("%#v", this.Geo)+",\n")
	}
	s = append(s, "RewardLevel: "+fmt.Sprintf("%#v", this.RewardLevel)+",\n")
	s = append(s, "CRM: "+fmt.Sprintf("%#v", this.CRM)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GeoData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&segsvc.GeoData{")
	s = append(s, "IpAddress: "+fmt.Sprintf("%#v", this.IpAddress)+",\n")
	s = append(s, "ZipCode: "+fmt.Sprintf("%#v", this.ZipCode)+",\n")
	s = append(s, "DMZ: "+fmt.Sprintf("%#v", this.DMZ)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSegsvc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringSegsvc(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SegSvc service

type SegSvcClient interface {
	Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*GetResponse, error)
}

type segSvcClient struct {
	cc *grpc.ClientConn
}

func NewSegSvcClient(cc *grpc.ClientConn) SegSvcClient {
	return &segSvcClient{cc}
}

func (c *segSvcClient) Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/SegSvc/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SegSvc service

type SegSvcServer interface {
	Get(context.Context, *User) (*GetResponse, error)
}

func RegisterSegSvcServer(s *grpc.Server, srv SegSvcServer) {
	s.RegisterService(&_SegSvc_serviceDesc, srv)
}

func _SegSvc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegSvcServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SegSvc/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegSvcServer).Get(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _SegSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SegSvc",
	HandlerType: (*SegSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SegSvc_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "segsvc.proto",
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if len(m.URI) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.URI)))
		i += copy(dAtA[i:], m.URI)
	}
	if m.Scores != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(m.Scores.Size()))
		n1, err := m.Scores.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *GetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Segment) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.Segment)))
		i += copy(dAtA[i:], m.Segment)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *ScoreData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScoreData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Segment) > 0 {
		for k, _ := range m.Segment {
			dAtA[i] = 0xa
			i++
			v := m.Segment[k]
			mapSize := 1 + len(k) + sovSegsvc(uint64(len(k))) + 1 + sovSegsvc(uint64(v))
			i = encodeVarintSegsvc(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintSegsvc(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintSegsvc(dAtA, i, uint64(v))
		}
	}
	if m.Geo != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(m.Geo.Size()))
		n2, err := m.Geo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.RewardLevel != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(m.RewardLevel))
	}
	if len(m.CRM) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.CRM)))
		i += copy(dAtA[i:], m.CRM)
	}
	return i, nil
}

func (m *GeoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GeoData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IpAddress) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.IpAddress)))
		i += copy(dAtA[i:], m.IpAddress)
	}
	if len(m.ZipCode) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.ZipCode)))
		i += copy(dAtA[i:], m.ZipCode)
	}
	if len(m.DMZ) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSegsvc(dAtA, i, uint64(len(m.DMZ)))
		i += copy(dAtA[i:], m.DMZ)
	}
	return i, nil
}

func encodeFixed64Segsvc(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Segsvc(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSegsvc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	if m.Scores != nil {
		l = m.Scores.Size()
		n += 1 + l + sovSegsvc(uint64(l))
	}
	return n
}

func (m *GetResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Segment)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	return n
}

func (m *ScoreData) Size() (n int) {
	var l int
	_ = l
	if len(m.Segment) > 0 {
		for k, v := range m.Segment {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovSegsvc(uint64(len(k))) + 1 + sovSegsvc(uint64(v))
			n += mapEntrySize + 1 + sovSegsvc(uint64(mapEntrySize))
		}
	}
	if m.Geo != nil {
		l = m.Geo.Size()
		n += 1 + l + sovSegsvc(uint64(l))
	}
	if m.RewardLevel != 0 {
		n += 1 + sovSegsvc(uint64(m.RewardLevel))
	}
	l = len(m.CRM)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	return n
}

func (m *GeoData) Size() (n int) {
	var l int
	_ = l
	l = len(m.IpAddress)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	l = len(m.ZipCode)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	l = len(m.DMZ)
	if l > 0 {
		n += 1 + l + sovSegsvc(uint64(l))
	}
	return n
}

func sovSegsvc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSegsvc(x uint64) (n int) {
	return sovSegsvc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *User) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&User{`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`URI:` + fmt.Sprintf("%v", this.URI) + `,`,
		`Scores:` + strings.Replace(fmt.Sprintf("%v", this.Scores), "ScoreData", "ScoreData", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetResponse{`,
		`Segment:` + fmt.Sprintf("%v", this.Segment) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ScoreData) String() string {
	if this == nil {
		return "nil"
	}
	keysForSegment := make([]string, 0, len(this.Segment))
	for k, _ := range this.Segment {
		keysForSegment = append(keysForSegment, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForSegment)
	mapStringForSegment := "map[string]int64{"
	for _, k := range keysForSegment {
		mapStringForSegment += fmt.Sprintf("%v: %v,", k, this.Segment[k])
	}
	mapStringForSegment += "}"
	s := strings.Join([]string{`&ScoreData{`,
		`Segment:` + mapStringForSegment + `,`,
		`Geo:` + strings.Replace(fmt.Sprintf("%v", this.Geo), "GeoData", "GeoData", 1) + `,`,
		`RewardLevel:` + fmt.Sprintf("%v", this.RewardLevel) + `,`,
		`CRM:` + fmt.Sprintf("%v", this.CRM) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GeoData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GeoData{`,
		`IpAddress:` + fmt.Sprintf("%v", this.IpAddress) + `,`,
		`ZipCode:` + fmt.Sprintf("%v", this.ZipCode) + `,`,
		`DMZ:` + fmt.Sprintf("%v", this.DMZ) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSegsvc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSegsvc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Scores == nil {
				m.Scores = &ScoreData{}
			}
			if err := m.Scores.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSegsvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSegsvc
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
func (m *GetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSegsvc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Segment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Segment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSegsvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSegsvc
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
func (m *ScoreData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSegsvc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScoreData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScoreData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Segment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthSegsvc
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Segment == nil {
				m.Segment = make(map[string]int64)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSegsvc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvalue int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSegsvc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Segment[mapkey] = mapvalue
			} else {
				var mapvalue int64
				m.Segment[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Geo == nil {
				m.Geo = &GeoData{}
			}
			if err := m.Geo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardLevel", wireType)
			}
			m.RewardLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardLevel |= (Reward(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CRM", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CRM = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSegsvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSegsvc
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
func (m *GeoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSegsvc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GeoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GeoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZipCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZipCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DMZ", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSegsvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DMZ = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSegsvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSegsvc
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
func skipSegsvc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSegsvc
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
					return 0, ErrIntOverflowSegsvc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSegsvc
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSegsvc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSegsvc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSegsvc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSegsvc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSegsvc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("segsvc.proto", fileDescriptorSegsvc) }

var fileDescriptorSegsvc = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xce, 0x74, 0xda, 0x74, 0xf3, 0x52, 0x64, 0x18, 0x44, 0x43, 0x59, 0x86, 0x25, 0x07, 0xa9,
	0x22, 0x01, 0xe3, 0x45, 0x16, 0x3c, 0x68, 0x1b, 0x6b, 0x61, 0x6b, 0x61, 0xd2, 0x2a, 0xec, 0xad,
	0xb6, 0x8f, 0xb2, 0xb8, 0x36, 0x25, 0x13, 0x2b, 0xeb, 0xc9, 0x9f, 0xe0, 0xcf, 0xf0, 0xa7, 0x78,
	0x5c, 0xf0, 0xe2, 0xd1, 0xc6, 0x8b, 0xc7, 0xfd, 0x09, 0x32, 0x93, 0xa9, 0xcd, 0x9e, 0xf2, 0x7d,
	0xdf, 0x9b, 0x79, 0xef, 0xfb, 0x26, 0x0f, 0x3a, 0x0a, 0x57, 0x6a, 0xbb, 0x88, 0x36, 0x79, 0x56,
	0x64, 0xe1, 0x14, 0x9a, 0x33, 0x85, 0x39, 0xbf, 0x07, 0xae, 0xfe, 0x8e, 0x06, 0x01, 0x39, 0x21,
	0x3d, 0x4f, 0x5a, 0xc6, 0x19, 0xd0, 0x99, 0x1c, 0x05, 0x0d, 0x23, 0x6a, 0xc8, 0x43, 0x70, 0xd3,
	0x45, 0x96, 0xa3, 0x0a, 0xe8, 0x09, 0xe9, 0xf9, 0x31, 0x44, 0x86, 0x0e, 0xe6, 0xc5, 0x5c, 0xda,
	0x4a, 0xf8, 0x1c, 0xfc, 0x21, 0x16, 0x12, 0xd5, 0x26, 0x5b, 0x2b, 0xe4, 0x01, 0xb4, 0x15, 0xae,
	0x3e, 0xe2, 0xba, 0xb0, 0xdd, 0xf7, 0x94, 0xdf, 0x85, 0xd6, 0x76, 0x7e, 0xf9, 0x09, 0xed, 0x80,
	0x8a, 0x84, 0x3f, 0x09, 0x78, 0xff, 0x9b, 0xf2, 0x27, 0xf5, 0xdb, 0xb4, 0xe7, 0xc7, 0xf7, 0x0f,
	0x13, 0xa3, 0xb4, 0xaa, 0x24, 0xeb, 0x22, 0xbf, 0x3a, 0xb4, 0xed, 0x02, 0x5d, 0x61, 0x66, 0x9a,
	0xfa, 0xf1, 0x51, 0x34, 0xc4, 0xcc, 0xd8, 0xd3, 0x22, 0x7f, 0x08, 0x7e, 0x8e, 0x9f, 0xe7, 0xf9,
	0xf2, 0x0c, 0xb7, 0x78, 0x69, 0x42, 0xdc, 0x89, 0xdb, 0x91, 0x34, 0x9a, 0xac, 0xd7, 0x74, 0xf8,
	0xbe, 0x1c, 0x07, 0xcd, 0x2a, 0x7c, 0x5f, 0x8e, 0xbb, 0xa7, 0xd0, 0xa9, 0x4f, 0xd4, 0x27, 0x3e,
	0xe0, 0x95, 0x4d, 0xa5, 0xe1, 0xed, 0x44, 0xd4, 0x26, 0x3a, 0x6d, 0x3c, 0x23, 0x61, 0x0a, 0x6d,
	0x6b, 0x84, 0x1f, 0x83, 0x77, 0xb1, 0x79, 0xb1, 0x5c, 0xe6, 0xa8, 0x94, 0xbd, 0x7c, 0x10, 0xf4,
	0x73, 0x7d, 0xb9, 0xd8, 0xf4, 0xb3, 0xe5, 0xfe, 0x59, 0xf6, 0x54, 0x8f, 0x1b, 0x8c, 0xcf, 0x8d,
	0x67, 0x4f, 0x6a, 0xf8, 0xe8, 0x0c, 0xdc, 0xca, 0x39, 0x6f, 0x03, 0x4d, 0xa4, 0x64, 0x0e, 0x3f,
	0x82, 0xe6, 0x79, 0x22, 0x27, 0x8c, 0x68, 0x69, 0xf2, 0x26, 0x61, 0x0d, 0x0d, 0xa6, 0xef, 0x26,
	0x8c, 0x72, 0x0f, 0x5a, 0xd3, 0xd7, 0x32, 0x49, 0x58, 0x53, 0x1f, 0x7b, 0x35, 0x99, 0x49, 0xd6,
	0x32, 0x68, 0xf4, 0x36, 0x61, 0x6e, 0xfc, 0x00, 0xdc, 0x14, 0x57, 0xe9, 0x76, 0xc1, 0x8f, 0x81,
	0x0e, 0xb1, 0xe0, 0xad, 0x48, 0xef, 0x41, 0xb7, 0x13, 0xd5, 0x7e, 0x67, 0xe8, 0xbc, 0x7c, 0x7c,
	0xbd, 0x13, 0xce, 0xaf, 0x9d, 0x70, 0x6e, 0x76, 0x82, 0x7c, 0x2d, 0x05, 0xf9, 0x5e, 0x0a, 0xf2,
	0xa3, 0x14, 0xe4, 0xba, 0x14, 0xe4, 0x77, 0x29, 0xc8, 0xdf, 0x52, 0x38, 0x37, 0xa5, 0x20, 0xdf,
	0xfe, 0x08, 0xe7, 0xbd, 0x6b, 0x56, 0xed, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xff,
	0x3d, 0xae, 0x7a, 0x02, 0x00, 0x00,
}
