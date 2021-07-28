// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store/statspb/stats.proto

package statspb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Statistics struct {
	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Nodes []*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes"`
}

func (m *Statistics) Reset()         { *m = Statistics{} }
func (m *Statistics) String() string { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()    {}
func (*Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_486eb8d0e3e11897, []int{0}
}
func (m *Statistics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Statistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Statistics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Statistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statistics.Merge(m, src)
}
func (m *Statistics) XXX_Size() int {
	return m.Size()
}
func (m *Statistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Statistics.DiscardUnknown(m)
}

var xxx_messageInfo_Statistics proto.InternalMessageInfo

type Node struct {
	Name         string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Nodes        []*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes"`
	NetRecvBytes int64   `protobuf:"varint,3,opt,name=net_recv_bytes,json=netRecvBytes,proto3" json:"net_recv_bytes"`
	Samples      int64   `protobuf:"varint,4,opt,name=samples,proto3" json:"samples"`
	Series       int64   `protobuf:"varint,5,opt,name=series,proto3" json:"series"`
	Chunks       int64   `protobuf:"varint,6,opt,name=chunks,proto3" json:"chunks"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_486eb8d0e3e11897, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Statistics)(nil), "stats.Statistics")
	proto.RegisterType((*Node)(nil), "stats.Node")
}

func init() { proto.RegisterFile("store/statspb/stats.proto", fileDescriptor_486eb8d0e3e11897) }

var fileDescriptor_486eb8d0e3e11897 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0xfa, 0x47, 0x1d, 0xc4, 0x60, 0x31, 0x18, 0x84, 0x9c, 0x28, 0x52, 0xa5, 0x0c,
	0xa8, 0x95, 0x60, 0x61, 0xce, 0x03, 0x30, 0x98, 0x05, 0xb1, 0x54, 0x49, 0x7a, 0x55, 0x22, 0x68,
	0x1c, 0xe5, 0x9a, 0x4a, 0xbc, 0x05, 0x8f, 0xd5, 0xb1, 0x23, 0x53, 0x04, 0xc9, 0x96, 0x57, 0x60,
	0x41, 0xb1, 0x53, 0x21, 0x31, 0xb3, 0xf8, 0xdc, 0x73, 0xbe, 0x63, 0x7b, 0xb8, 0xf4, 0x1c, 0xb5,
	0x2a, 0x61, 0x81, 0x3a, 0xd6, 0x58, 0x24, 0x56, 0xe7, 0x45, 0xa9, 0xb4, 0x62, 0x23, 0x63, 0x2e,
	0xce, 0xd6, 0x6a, 0xad, 0x4c, 0xb2, 0xe8, 0x26, 0x0b, 0x83, 0x07, 0x4a, 0xef, 0x75, 0xac, 0x33,
	0xd4, 0x59, 0x8a, 0xec, 0x92, 0x0e, 0xf3, 0x78, 0x03, 0x9c, 0xf8, 0x24, 0x9c, 0x46, 0xc7, 0x6d,
	0xe5, 0x19, 0x2f, 0xcd, 0xc9, 0xae, 0xe8, 0x28, 0x57, 0x2b, 0x40, 0x7e, 0xe4, 0x0f, 0x42, 0xf7,
	0xda, 0x9d, 0xdb, 0x5f, 0xee, 0xd4, 0x0a, 0xa2, 0x69, 0x5b, 0x79, 0x96, 0x4a, 0x2b, 0xc1, 0x37,
	0xa1, 0xc3, 0x0e, 0xfd, 0xe7, 0xa3, 0xec, 0x96, 0x9e, 0xe6, 0xa0, 0x97, 0x25, 0xa4, 0xdb, 0x65,
	0xf2, 0xa6, 0x01, 0xf9, 0xc0, 0x27, 0xe1, 0x20, 0x62, 0x6d, 0xe5, 0xfd, 0x21, 0xf2, 0x24, 0x07,
	0x2d, 0x21, 0xdd, 0x46, 0x9d, 0x63, 0x33, 0x3a, 0xc1, 0x78, 0x53, 0xbc, 0x00, 0xf2, 0xa1, 0xb9,
	0xe2, 0xb6, 0x95, 0x77, 0x88, 0xe4, 0x61, 0x60, 0x01, 0x1d, 0x23, 0x94, 0x19, 0x20, 0x1f, 0x99,
	0x16, 0x6d, 0x2b, 0xaf, 0x4f, 0x64, 0xaf, 0x5d, 0x27, 0x7d, 0x7a, 0xcd, 0x9f, 0x91, 0x8f, 0x7f,
	0x3b, 0x36, 0x91, 0xbd, 0x46, 0xb3, 0xdd, 0x97, 0x70, 0x76, 0xb5, 0x20, 0xfb, 0x5a, 0x90, 0xcf,
	0x5a, 0x90, 0xf7, 0x46, 0x38, 0xfb, 0x46, 0x38, 0x1f, 0x8d, 0x70, 0x1e, 0x27, 0xfd, 0x9e, 0x92,
	0xb1, 0xd9, 0xc2, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x02, 0x3e, 0x3b, 0xbf, 0x01,
	0x00, 0x00,
}

func (m *Statistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Statistics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Statistics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStats(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Chunks != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.Chunks))
		i--
		dAtA[i] = 0x30
	}
	if m.Series != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.Series))
		i--
		dAtA[i] = 0x28
	}
	if m.Samples != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.Samples))
		i--
		dAtA[i] = 0x20
	}
	if m.NetRecvBytes != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.NetRecvBytes))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStats(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Statistics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	return n
}

func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	if m.NetRecvBytes != 0 {
		n += 1 + sovStats(uint64(m.NetRecvBytes))
	}
	if m.Samples != 0 {
		n += 1 + sovStats(uint64(m.Samples))
	}
	if m.Series != 0 {
		n += 1 + sovStats(uint64(m.Series))
	}
	if m.Chunks != 0 {
		n += 1 + sovStats(uint64(m.Chunks))
	}
	return n
}

func sovStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Statistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: Statistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Statistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetRecvBytes", wireType)
			}
			m.NetRecvBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetRecvBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Samples", wireType)
			}
			m.Samples = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Samples |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Series", wireType)
			}
			m.Series = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Series |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			m.Chunks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chunks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStats
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
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
				return 0, ErrInvalidLengthStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStats = fmt.Errorf("proto: unexpected end of group")
)
