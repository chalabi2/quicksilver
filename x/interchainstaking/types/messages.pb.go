// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/interchainstaking/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgRegisterZone represents a message to send coins from one account to
// another.
type MsgRegisterZone struct {
	Identifier  string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty" yaml:"zone_name"`
	ChainId     string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	LocalDenom  string `protobuf:"bytes,3,opt,name=local_denom,json=localDenom,proto3" json:"local_denom,omitempty" yaml:"local_denom"`
	RemoteDenom string `protobuf:"bytes,4,opt,name=remote_denom,json=remoteDenom,proto3" json:"remote_denom,omitempty" yaml:"remote_denom"`
	FromAddress string `protobuf:"bytes,5,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgRegisterZone) Reset()         { *m = MsgRegisterZone{} }
func (m *MsgRegisterZone) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterZone) ProtoMessage()    {}
func (*MsgRegisterZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee484030fa140a82, []int{0}
}
func (m *MsgRegisterZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterZone.Merge(m, src)
}
func (m *MsgRegisterZone) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterZone) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterZone.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterZone proto.InternalMessageInfo

// MsgRegisterZoneResponse defines the Msg/Send response type.
type MsgRegisterZoneResponse struct {
}

func (m *MsgRegisterZoneResponse) Reset()         { *m = MsgRegisterZoneResponse{} }
func (m *MsgRegisterZoneResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterZoneResponse) ProtoMessage()    {}
func (*MsgRegisterZoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee484030fa140a82, []int{1}
}
func (m *MsgRegisterZoneResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterZoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterZoneResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterZoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterZoneResponse.Merge(m, src)
}
func (m *MsgRegisterZoneResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterZoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterZoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterZoneResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterZone)(nil), "quicksilver.interchainstaking.v1.MsgRegisterZone")
	proto.RegisterType((*MsgRegisterZoneResponse)(nil), "quicksilver.interchainstaking.v1.MsgRegisterZoneResponse")
}

func init() {
	proto.RegisterFile("quicksilver/interchainstaking/v1/messages.proto", fileDescriptor_ee484030fa140a82)
}

var fileDescriptor_ee484030fa140a82 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0x87, 0xed, 0x3b, 0xfe, 0x1c, 0x9b, 0x48, 0x87, 0x7c, 0x11, 0xe7, 0x4b, 0x61, 0x9f, 0x5c,
	0xd1, 0x9c, 0xad, 0x00, 0x12, 0x22, 0x34, 0x70, 0xa2, 0xa1, 0xb8, 0xc6, 0x74, 0x69, 0x2c, 0xc7,
	0x9e, 0x6c, 0x56, 0xf1, 0xee, 0x86, 0xdd, 0x75, 0x44, 0x10, 0x1d, 0x0d, 0x25, 0x8f, 0x90, 0x77,
	0x80, 0x87, 0xa0, 0x8c, 0xa8, 0xa8, 0x2c, 0x94, 0x34, 0xd4, 0x7e, 0x02, 0x94, 0xdd, 0x44, 0x72,
	0x02, 0x12, 0xa2, 0x9b, 0x9f, 0xbf, 0xf9, 0xac, 0xd1, 0xec, 0xa0, 0xe8, 0x6d, 0x49, 0xb2, 0x89,
	0x24, 0xc5, 0x0c, 0x44, 0x44, 0x98, 0x02, 0x91, 0x8d, 0x53, 0xc2, 0xa4, 0x4a, 0x27, 0x84, 0xe1,
	0x68, 0xd6, 0x8b, 0x28, 0x48, 0x99, 0x62, 0x90, 0xe1, 0x54, 0x70, 0xc5, 0x9d, 0xcb, 0x86, 0x10,
	0xfe, 0x21, 0x84, 0xb3, 0x5e, 0xb7, 0x83, 0x39, 0xe6, 0xba, 0x39, 0xda, 0x54, 0xc6, 0xeb, 0x5e,
	0x64, 0x5c, 0x52, 0x2e, 0x13, 0x03, 0x4c, 0x30, 0x28, 0xf8, 0x72, 0x84, 0x4e, 0x6f, 0x24, 0x8e,
	0x01, 0x13, 0xa9, 0x40, 0x0c, 0x38, 0x03, 0xe7, 0x09, 0x42, 0x24, 0x07, 0xa6, 0xc8, 0x88, 0x80,
	0x70, 0xed, 0x4b, 0xfb, 0xe1, 0xbd, 0xeb, 0x4e, 0x5d, 0xf9, 0xf7, 0xe7, 0x29, 0x2d, 0xfa, 0xc1,
	0x7b, 0xce, 0x20, 0x61, 0x29, 0x85, 0x20, 0x6e, 0xf4, 0x39, 0x21, 0x3a, 0xd1, 0xd3, 0x24, 0x24,
	0x77, 0x8f, 0xb4, 0x73, 0x56, 0x57, 0xfe, 0xa9, 0x71, 0x76, 0x24, 0x88, 0xef, 0xea, 0xf2, 0x75,
	0xee, 0x3c, 0x45, 0xad, 0x82, 0x67, 0x69, 0x91, 0xe4, 0xc0, 0x38, 0x75, 0x8f, 0xb5, 0xf2, 0xa0,
	0xae, 0x7c, 0xc7, 0x28, 0x0d, 0x18, 0xc4, 0x48, 0xa7, 0x57, 0x9b, 0xe0, 0xf4, 0x51, 0x5b, 0x00,
	0xe5, 0x0a, 0xb6, 0xe6, 0x2d, 0x6d, 0x9e, 0xd7, 0x95, 0x7f, 0x66, 0xcc, 0x26, 0x0d, 0xe2, 0x96,
	0x89, 0xc6, 0x7d, 0x8e, 0xda, 0x23, 0xc1, 0x69, 0x92, 0xe6, 0xb9, 0x00, 0x29, 0xdd, 0xdb, 0xda,
	0x75, 0xbf, 0x7f, 0xbd, 0xea, 0x6c, 0xd7, 0xf2, 0xd2, 0x90, 0x37, 0x4a, 0x10, 0x86, 0xe3, 0xd6,
	0xa6, 0x7b, 0xfb, 0xa9, 0x7f, 0xf2, 0x69, 0xe1, 0x5b, 0xbf, 0x16, 0xbe, 0x15, 0x5c, 0xa0, 0xf3,
	0x83, 0xa5, 0xc5, 0x20, 0xa7, 0x9c, 0x49, 0x78, 0xf4, 0xd1, 0x46, 0xc7, 0x37, 0x12, 0x3b, 0x1f,
	0x50, 0x7b, 0x6f, 0xa9, 0xbd, 0xf0, 0x5f, 0x8f, 0x17, 0x1e, 0xfc, 0xb2, 0xfb, 0xec, 0xbf, 0x95,
	0xdd, 0x14, 0xd7, 0x83, 0x6f, 0x2b, 0xcf, 0x5e, 0xae, 0x3c, 0xfb, 0xe7, 0xca, 0xb3, 0x3f, 0xaf,
	0x3d, 0x6b, 0xb9, 0xf6, 0xac, 0x1f, 0x6b, 0xcf, 0x1a, 0xbc, 0xc0, 0x44, 0x8d, 0xcb, 0x61, 0x98,
	0x71, 0x1a, 0x11, 0x86, 0x81, 0x95, 0x44, 0xcd, 0xaf, 0x86, 0x25, 0x29, 0xf2, 0xbd, 0x7b, 0x7c,
	0xf7, 0x97, 0x8b, 0x54, 0xf3, 0x29, 0xc8, 0xe1, 0x1d, 0x7d, 0x39, 0x8f, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0xed, 0xd2, 0xa6, 0xbf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterZone defines a method for sending coins from one account to another
	// account.
	RegisterZone(ctx context.Context, in *MsgRegisterZone, opts ...grpc.CallOption) (*MsgRegisterZoneResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterZone(ctx context.Context, in *MsgRegisterZone, opts ...grpc.CallOption) (*MsgRegisterZoneResponse, error) {
	out := new(MsgRegisterZoneResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainstaking.v1.Msg/RegisterZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterZone defines a method for sending coins from one account to another
	// account.
	RegisterZone(context.Context, *MsgRegisterZone) (*MsgRegisterZoneResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterZone(ctx context.Context, req *MsgRegisterZone) (*MsgRegisterZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterZone not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainstaking.v1.Msg/RegisterZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterZone(ctx, req.(*MsgRegisterZone))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.interchainstaking.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterZone",
			Handler:    _Msg_RegisterZone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/interchainstaking/v1/messages.proto",
}

func (m *MsgRegisterZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RemoteDenom) > 0 {
		i -= len(m.RemoteDenom)
		copy(dAtA[i:], m.RemoteDenom)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.RemoteDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LocalDenom) > 0 {
		i -= len(m.LocalDenom)
		copy(dAtA[i:], m.LocalDenom)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.LocalDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterZoneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterZoneResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterZoneResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.LocalDenom)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.RemoteDenom)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgRegisterZoneResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgRegisterZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgRegisterZoneResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgRegisterZoneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterZoneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
