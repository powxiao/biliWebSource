// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/identify-game/model/token.proto

package model

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

// result of http access info.
type AccessInfo struct {
	Mid                  int64    `protobuf:"varint,1,opt,name=Mid,proto3" json:"mid"`
	AppID                int64    `protobuf:"varint,2,opt,name=AppID,proto3" json:"appid"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"access_key"`
	CreateAt             int64    `protobuf:"varint,4,opt,name=CreateAt,proto3" json:"create_at"`
	UserID               string   `protobuf:"bytes,5,opt,name=UserID,proto3" json:"userid"`
	Name                 string   `protobuf:"bytes,6,opt,name=Name,proto3" json:"uname"`
	Expires              int64    `protobuf:"varint,7,opt,name=Expires,proto3" json:"expires"`
	Permission           string   `protobuf:"bytes,8,opt,name=Permission,proto3" json:"permission"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessInfo) Reset()         { *m = AccessInfo{} }
func (m *AccessInfo) String() string { return proto.CompactTextString(m) }
func (*AccessInfo) ProtoMessage()    {}
func (*AccessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_a3d7e52fbb774120, []int{0}
}
func (m *AccessInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AccessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessInfo.Merge(dst, src)
}
func (m *AccessInfo) XXX_Size() int {
	return m.Size()
}
func (m *AccessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccessInfo proto.InternalMessageInfo

func (m *AccessInfo) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *AccessInfo) GetAppID() int64 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *AccessInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccessInfo) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *AccessInfo) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AccessInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccessInfo) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *AccessInfo) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

type RenewInfo struct {
	Expires              int64    `protobuf:"varint,1,opt,name=Expires,proto3" json:"expires"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewInfo) Reset()         { *m = RenewInfo{} }
func (m *RenewInfo) String() string { return proto.CompactTextString(m) }
func (*RenewInfo) ProtoMessage()    {}
func (*RenewInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_a3d7e52fbb774120, []int{1}
}
func (m *RenewInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RenewInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RenewInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RenewInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewInfo.Merge(dst, src)
}
func (m *RenewInfo) XXX_Size() int {
	return m.Size()
}
func (m *RenewInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RenewInfo proto.InternalMessageInfo

func (m *RenewInfo) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*AccessInfo)(nil), "model.accessInfo")
	proto.RegisterType((*RenewInfo)(nil), "model.renewInfo")
}
func (m *AccessInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintToken(dAtA, i, uint64(m.Mid))
	}
	if m.AppID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintToken(dAtA, i, uint64(m.AppID))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintToken(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.CreateAt != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintToken(dAtA, i, uint64(m.CreateAt))
	}
	if len(m.UserID) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintToken(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintToken(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Expires != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintToken(dAtA, i, uint64(m.Expires))
	}
	if len(m.Permission) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintToken(dAtA, i, uint64(len(m.Permission)))
		i += copy(dAtA[i:], m.Permission)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RenewInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RenewInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Expires != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintToken(dAtA, i, uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AccessInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovToken(uint64(m.Mid))
	}
	if m.AppID != 0 {
		n += 1 + sovToken(uint64(m.AppID))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovToken(uint64(m.CreateAt))
	}
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovToken(uint64(m.Expires))
	}
	l = len(m.Permission)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RenewInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Expires != 0 {
		n += 1 + sovToken(uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovToken(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: accessInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: accessInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			m.AppID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permission = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RenewInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: renewInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: renewInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowToken
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
				next, err := skipToken(dAtA[start:])
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
	ErrInvalidLengthToken = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/service/main/identify-game/model/token.proto", fileDescriptor_token_a3d7e52fbb774120)
}

var fileDescriptor_token_a3d7e52fbb774120 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4e, 0xeb, 0x30,
	0x10, 0x87, 0x95, 0xa6, 0x49, 0x1a, 0xbf, 0x3f, 0x0b, 0xaf, 0xfc, 0x9e, 0x44, 0x53, 0x55, 0x20,
	0x95, 0x45, 0x1b, 0x04, 0x27, 0x68, 0x29, 0x8b, 0x2e, 0x40, 0xc8, 0x82, 0x75, 0xe5, 0x26, 0xd3,
	0x60, 0x15, 0xc7, 0x96, 0x9d, 0x00, 0x3d, 0x17, 0x97, 0x60, 0xc9, 0x09, 0x22, 0xd4, 0x65, 0x4e,
	0x81, 0xea, 0x94, 0xd2, 0x0d, 0x3b, 0xcf, 0xf7, 0x8d, 0x67, 0x7e, 0x1a, 0x74, 0xc6, 0x94, 0x8a,
	0x0d, 0xe8, 0x27, 0x9e, 0x40, 0x2c, 0x18, 0xcf, 0x63, 0x9e, 0x42, 0x5e, 0xf0, 0xe5, 0x7a, 0x98,
	0x31, 0x01, 0xb1, 0x90, 0x29, 0x3c, 0xc6, 0x85, 0x5c, 0x41, 0x3e, 0x52, 0x5a, 0x16, 0x12, 0x7b,
	0x16, 0xfd, 0x1f, 0x66, 0xbc, 0x78, 0x28, 0x17, 0xa3, 0x44, 0x8a, 0x38, 0x93, 0x99, 0x8c, 0xad,
	0x5d, 0x94, 0x4b, 0x5b, 0xd9, 0xc2, 0xbe, 0x9a, 0x5f, 0xfd, 0xd7, 0x16, 0x42, 0x2c, 0x49, 0xc0,
	0x98, 0x59, 0xbe, 0x94, 0xf8, 0x1f, 0x72, 0xaf, 0x79, 0x4a, 0x9c, 0x9e, 0x33, 0x70, 0x27, 0x41,
	0x5d, 0x45, 0xae, 0xe0, 0x29, 0xdd, 0x32, 0x1c, 0x21, 0x6f, 0xac, 0xd4, 0x6c, 0x4a, 0x5a, 0x56,
	0x86, 0x75, 0x15, 0x79, 0x4c, 0x29, 0x9e, 0xd2, 0x86, 0xe3, 0x63, 0xe4, 0xdd, 0x6d, 0xf3, 0x10,
	0xb7, 0xe7, 0x0c, 0xc2, 0xc9, 0xdf, 0xba, 0x8a, 0x76, 0xa3, 0xe7, 0x2b, 0x58, 0xd3, 0x46, 0xe2,
	0x53, 0xd4, 0xb9, 0xd4, 0xc0, 0x0a, 0x18, 0x17, 0xa4, 0x6d, 0x27, 0xfd, 0xa9, 0xab, 0x28, 0x4c,
	0x2c, 0x9b, 0xb3, 0x82, 0xee, 0x35, 0xee, 0x23, 0xff, 0xde, 0x80, 0x9e, 0x4d, 0x89, 0x67, 0x27,
	0xa2, 0xba, 0x8a, 0xfc, 0xd2, 0x80, 0xe6, 0x29, 0xdd, 0x19, 0x7c, 0x84, 0xda, 0x37, 0x4c, 0x00,
	0xf1, 0x6d, 0x87, 0x0d, 0x55, 0xe6, 0x4c, 0x00, 0xb5, 0x18, 0x9f, 0xa0, 0xe0, 0xea, 0x45, 0x71,
	0x0d, 0x86, 0x04, 0x76, 0xd9, 0xaf, 0xba, 0x8a, 0x02, 0x68, 0x10, 0xfd, 0x72, 0x78, 0x84, 0xd0,
	0x2d, 0x68, 0xc1, 0x8d, 0xe1, 0x32, 0x27, 0x9d, 0xef, 0xfc, 0x6a, 0x4f, 0xe9, 0x41, 0x47, 0xff,
	0x1c, 0x85, 0x1a, 0x72, 0x78, 0xb6, 0x37, 0x3b, 0xd8, 0xe1, 0xfc, 0xbc, 0x63, 0xf2, 0xfb, 0x6d,
	0xd3, 0x75, 0xde, 0x37, 0x5d, 0xe7, 0x63, 0xd3, 0x75, 0x16, 0xbe, 0x3d, 0xff, 0xc5, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe3, 0x2b, 0xe9, 0x37, 0xe8, 0x01, 0x00, 0x00,
}
