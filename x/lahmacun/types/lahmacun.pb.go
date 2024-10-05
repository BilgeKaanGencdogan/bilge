// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bilge/lahmacun/lahmacun.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type SpiceLevel int32

const (
	SpiceLevel_MILD   SpiceLevel = 0
	SpiceLevel_MEDIUM SpiceLevel = 1
	SpiceLevel_HOT    SpiceLevel = 2
)

var SpiceLevel_name = map[int32]string{
	0: "MILD",
	1: "MEDIUM",
	2: "HOT",
}

var SpiceLevel_value = map[string]int32{
	"MILD":   0,
	"MEDIUM": 1,
	"HOT":    2,
}

func (x SpiceLevel) String() string {
	return proto.EnumName(SpiceLevel_name, int32(x))
}

func (SpiceLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17af4627968bd451, []int{0}
}

type Lahmacun struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Heat    string `protobuf:"bytes,3,opt,name=heat,proto3" json:"heat,omitempty"`
	Dough   string `protobuf:"bytes,4,opt,name=dough,proto3" json:"dough,omitempty"`
	Meat    string `protobuf:"bytes,5,opt,name=meat,proto3" json:"meat,omitempty"`
	Tomato  string `protobuf:"bytes,6,opt,name=tomato,proto3" json:"tomato,omitempty"`
	Onion   string `protobuf:"bytes,7,opt,name=onion,proto3" json:"onion,omitempty"`
}

func (m *Lahmacun) Reset()         { *m = Lahmacun{} }
func (m *Lahmacun) String() string { return proto.CompactTextString(m) }
func (*Lahmacun) ProtoMessage()    {}
func (*Lahmacun) Descriptor() ([]byte, []int) {
	return fileDescriptor_17af4627968bd451, []int{0}
}
func (m *Lahmacun) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lahmacun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lahmacun.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lahmacun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lahmacun.Merge(m, src)
}
func (m *Lahmacun) XXX_Size() int {
	return m.Size()
}
func (m *Lahmacun) XXX_DiscardUnknown() {
	xxx_messageInfo_Lahmacun.DiscardUnknown(m)
}

var xxx_messageInfo_Lahmacun proto.InternalMessageInfo

func (m *Lahmacun) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Lahmacun) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Lahmacun) GetHeat() string {
	if m != nil {
		return m.Heat
	}
	return ""
}

func (m *Lahmacun) GetDough() string {
	if m != nil {
		return m.Dough
	}
	return ""
}

func (m *Lahmacun) GetMeat() string {
	if m != nil {
		return m.Meat
	}
	return ""
}

func (m *Lahmacun) GetTomato() string {
	if m != nil {
		return m.Tomato
	}
	return ""
}

func (m *Lahmacun) GetOnion() string {
	if m != nil {
		return m.Onion
	}
	return ""
}

func init() {
	proto.RegisterEnum("bilge.lahmacun.SpiceLevel", SpiceLevel_name, SpiceLevel_value)
	proto.RegisterType((*Lahmacun)(nil), "bilge.lahmacun.Lahmacun")
}

func init() { proto.RegisterFile("bilge/lahmacun/lahmacun.proto", fileDescriptor_17af4627968bd451) }

var fileDescriptor_17af4627968bd451 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0x27, 0xd3, 0x76, 0xda, 0xff, 0x5d, 0x94, 0xfe, 0x43, 0xa9, 0xb1, 0xe2, 0x20, 0xae,
	0x44, 0xa1, 0xa3, 0xf8, 0x06, 0x52, 0xc1, 0x42, 0x8b, 0xe0, 0xc7, 0xc6, 0x8d, 0xa4, 0xd3, 0x30,
	0x0d, 0x34, 0x73, 0x87, 0x99, 0xb4, 0xe8, 0x5b, 0xf8, 0x08, 0x3e, 0x8e, 0xcb, 0x2e, 0x5d, 0x4a,
	0xfb, 0x22, 0x92, 0x8f, 0x2a, 0xb8, 0x49, 0xce, 0xf9, 0x9d, 0x93, 0x4b, 0xb8, 0x70, 0x38, 0x95,
	0x8b, 0x4c, 0x24, 0x0b, 0x3e, 0x57, 0x3c, 0x5d, 0xe6, 0x3f, 0x62, 0x50, 0x94, 0xa8, 0x91, 0xb6,
	0x6d, 0x3c, 0xd8, 0xd1, 0xfe, 0x7f, 0xae, 0x64, 0x8e, 0x89, 0x3d, 0x5d, 0xa5, 0xbf, 0x97, 0x62,
	0xa5, 0xb0, 0x4a, 0x54, 0x95, 0x25, 0xab, 0x0b, 0x73, 0xf9, 0x60, 0xdf, 0x05, 0xcf, 0xd6, 0x25,
	0xce, 0xf8, 0xa8, 0x9b, 0x61, 0x86, 0x8e, 0x1b, 0xe5, 0xe9, 0xc1, 0x9f, 0xbf, 0x14, 0xbc, 0xe4,
	0xca, 0x3f, 0x39, 0x7e, 0x27, 0xd0, 0x1a, 0xfb, 0x84, 0x32, 0x68, 0xa6, 0xa5, 0xe0, 0x1a, 0x4b,
	0x46, 0x8e, 0xc8, 0xc9, 0xbf, 0xbb, 0x9d, 0xa5, 0x6d, 0x08, 0xe5, 0x8c, 0x85, 0x16, 0x86, 0x72,
	0x46, 0x29, 0xd4, 0xe7, 0x82, 0x6b, 0x56, 0xb3, 0xc4, 0x6a, 0xda, 0x85, 0xc6, 0x0c, 0x97, 0xd9,
	0x9c, 0xd5, 0x2d, 0x74, 0xc6, 0x34, 0x95, 0x69, 0x36, 0x5c, 0xd3, 0x68, 0xda, 0x83, 0x48, 0xa3,
	0xe2, 0x1a, 0x59, 0x64, 0xa9, 0x77, 0x66, 0x02, 0xe6, 0x12, 0x73, 0xd6, 0x74, 0x13, 0xac, 0x39,
	0x3d, 0x03, 0xb8, 0x2f, 0x64, 0x2a, 0xc6, 0x62, 0x25, 0x16, 0xb4, 0x05, 0xf5, 0xc9, 0x68, 0x3c,
	0xec, 0x04, 0x14, 0x20, 0x9a, 0x5c, 0x0f, 0x47, 0x8f, 0x93, 0x0e, 0xa1, 0x4d, 0xa8, 0xdd, 0xdc,
	0x3e, 0x74, 0xc2, 0xab, 0xf3, 0x8f, 0x4d, 0x4c, 0xd6, 0x9b, 0x98, 0x7c, 0x6d, 0x62, 0xf2, 0xb6,
	0x8d, 0x83, 0xf5, 0x36, 0x0e, 0x3e, 0xb7, 0x71, 0xf0, 0xd4, 0x73, 0x6b, 0x78, 0xf9, 0x5d, 0x84,
	0x7e, 0x2d, 0x44, 0x35, 0x8d, 0xec, 0x22, 0x2e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xed,
	0x61, 0x6c, 0xb3, 0x01, 0x00, 0x00,
}

func (m *Lahmacun) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lahmacun) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lahmacun) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Onion) > 0 {
		i -= len(m.Onion)
		copy(dAtA[i:], m.Onion)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Onion)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Tomato) > 0 {
		i -= len(m.Tomato)
		copy(dAtA[i:], m.Tomato)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Tomato)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Meat) > 0 {
		i -= len(m.Meat)
		copy(dAtA[i:], m.Meat)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Meat)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Dough) > 0 {
		i -= len(m.Dough)
		copy(dAtA[i:], m.Dough)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Dough)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Heat) > 0 {
		i -= len(m.Heat)
		copy(dAtA[i:], m.Heat)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Heat)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLahmacun(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLahmacun(dAtA []byte, offset int, v uint64) int {
	offset -= sovLahmacun(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lahmacun) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	l = len(m.Heat)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	l = len(m.Dough)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	l = len(m.Meat)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	l = len(m.Tomato)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	l = len(m.Onion)
	if l > 0 {
		n += 1 + l + sovLahmacun(uint64(l))
	}
	return n
}

func sovLahmacun(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLahmacun(x uint64) (n int) {
	return sovLahmacun(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lahmacun) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLahmacun
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
			return fmt.Errorf("proto: Lahmacun: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lahmacun: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Heat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Heat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dough", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dough = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tomato", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tomato = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Onion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLahmacun
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
				return ErrInvalidLengthLahmacun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLahmacun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Onion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLahmacun(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLahmacun
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
func skipLahmacun(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLahmacun
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
					return 0, ErrIntOverflowLahmacun
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
					return 0, ErrIntOverflowLahmacun
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
				return 0, ErrInvalidLengthLahmacun
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLahmacun
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLahmacun
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLahmacun        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLahmacun          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLahmacun = fmt.Errorf("proto: unexpected end of group")
)
