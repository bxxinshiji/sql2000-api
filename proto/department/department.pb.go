// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/department/department.proto

package department

import (
	fmt "fmt"
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

type Request struct {
	StartDate  string  `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    string  `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Where      string  `protobuf:"bytes,3,opt,name=where,proto3" json:"where,omitempty"`
	Department []int64 `protobuf:"varint,4,rep,packed,name=department,proto3" json:"department,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{0}
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

func (m *Request) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Request) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *Request) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

func (m *Request) GetDepartment() []int64 {
	if m != nil {
		return m.Department
	}
	return nil
}

type Response struct {
	Valid bool  `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{1}
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

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "department.Request")
	proto.RegisterType((*Response)(nil), "department.Response")
}

func init() { proto.RegisterFile("proto/department/department.proto", fileDescriptor_a231da9bf7b7101c) }

var fileDescriptor_a231da9bf7b7101c = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x27, 0xa6, 0xda, 0xe9, 0x5b, 0xc6, 0x2e, 0xa2, 0x60, 0xa8, 0x5d, 0x75, 0x55, 0xc1,
	0x82, 0x7b, 0xa5, 0x27, 0x88, 0x07, 0x90, 0x48, 0x1e, 0x28, 0x8c, 0xc9, 0x98, 0x3c, 0x15, 0x6f,
	0xe1, 0xb1, 0x5c, 0x76, 0xe9, 0x52, 0x66, 0x2e, 0x22, 0xbe, 0x88, 0xc6, 0x5d, 0xfe, 0xef, 0x87,
	0xf7, 0xf1, 0x07, 0x4e, 0xfb, 0x14, 0x29, 0x9e, 0x79, 0xec, 0x5d, 0xa2, 0x07, 0x0c, 0x54, 0x3d,
	0xd7, 0xdc, 0x29, 0xf8, 0x23, 0xcb, 0x57, 0x98, 0x5a, 0x7c, 0x7c, 0xc2, 0x4c, 0xea, 0x04, 0x20,
	0x93, 0x4b, 0x74, 0xe3, 0x1d, 0xa1, 0x16, 0x0b, 0xb1, 0x9a, 0xd9, 0x19, 0x93, 0xad, 0x23, 0x54,
	0x47, 0xd0, 0x62, 0xf0, 0xa5, 0xdc, 0xe3, 0x72, 0x8a, 0xc1, 0x73, 0x35, 0x87, 0xfd, 0x97, 0x3b,
	0x4c, 0xa8, 0x25, 0xf3, 0x12, 0x94, 0x81, 0x4a, 0xa4, 0x27, 0x0b, 0xb9, 0x92, 0xb6, 0x56, 0x5f,
	0x40, 0x6b, 0x31, 0xf7, 0x31, 0x64, 0xbe, 0xf0, 0xec, 0xba, 0x7b, 0xcf, 0xda, 0xd6, 0x96, 0xf0,
	0x4d, 0x29, 0x92, 0xeb, 0xd8, 0x27, 0x6d, 0x09, 0xe7, 0x97, 0x00, 0xdb, 0xdf, 0x2b, 0x6a, 0x03,
	0x93, 0x6b, 0xd7, 0xa1, 0x3a, 0x5c, 0x57, 0x3b, 0x7f, 0x26, 0x1d, 0xcf, 0xff, 0xc3, 0x22, 0x5b,
	0x36, 0x57, 0xfa, 0x7d, 0x30, 0x62, 0x37, 0x18, 0xf1, 0x39, 0x18, 0xf1, 0x36, 0x9a, 0x66, 0x37,
	0x9a, 0xe6, 0x63, 0x34, 0xcd, 0xed, 0x01, 0x7f, 0xd1, 0xe6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb2,
	0x86, 0x7b, 0x1d, 0x47, 0x01, 0x00, 0x00,
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
	if len(m.Department) > 0 {
		dAtA2 := make([]byte, len(m.Department)*10)
		var j1 int
		for _, num1 := range m.Department {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintDepartment(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Where) > 0 {
		i -= len(m.Where)
		copy(dAtA[i:], m.Where)
		i = encodeVarintDepartment(dAtA, i, uint64(len(m.Where)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EndDate) > 0 {
		i -= len(m.EndDate)
		copy(dAtA[i:], m.EndDate)
		i = encodeVarintDepartment(dAtA, i, uint64(len(m.EndDate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintDepartment(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0xa
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
	if m.Total != 0 {
		i = encodeVarintDepartment(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepartment(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepartment(v)
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
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovDepartment(uint64(l))
	}
	l = len(m.EndDate)
	if l > 0 {
		n += 1 + l + sovDepartment(uint64(l))
	}
	l = len(m.Where)
	if l > 0 {
		n += 1 + l + sovDepartment(uint64(l))
	}
	if len(m.Department) > 0 {
		l = 0
		for _, e := range m.Department {
			l += sovDepartment(uint64(e))
		}
		n += 1 + sovDepartment(uint64(l)) + l
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valid {
		n += 2
	}
	if m.Total != 0 {
		n += 1 + sovDepartment(uint64(m.Total))
	}
	return n
}

func sovDepartment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepartment(x uint64) (n int) {
	return sovDepartment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepartment
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepartment
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
				return ErrInvalidLengthDepartment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepartment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepartment
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
				return ErrInvalidLengthDepartment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepartment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Where", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepartment
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
				return ErrInvalidLengthDepartment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepartment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Where = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDepartment
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Department = append(m.Department, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDepartment
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthDepartment
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthDepartment
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Department) == 0 {
					m.Department = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDepartment
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Department = append(m.Department, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Department", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDepartment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDepartment
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDepartment
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
				return ErrIntOverflowDepartment
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
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepartment
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
			m.Valid = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepartment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDepartment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDepartment
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDepartment
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
func skipDepartment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepartment
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
					return 0, ErrIntOverflowDepartment
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
					return 0, ErrIntOverflowDepartment
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
				return 0, ErrInvalidLengthDepartment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepartment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepartment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepartment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepartment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepartment = fmt.Errorf("proto: unexpected end of group")
)
