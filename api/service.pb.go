// Code generated by protoc-gen-gogo.
// source: service.proto
// DO NOT EDIT!

/*
	Package api is a generated protocol buffer package.

	It is generated from these files:
		service.proto
		types.proto
		swarm.proto
		master.proto
		agent.proto

	It has these top-level messages:
		ResultParameters
		Node
		Meta
		ImageSpec
		ContainerSpec
		ContainerConfig
		Bundle
		PODSpec
		Spec
		TaskStatus
		Task
		Job
		ListNodesRequest
		ListNodesResponse
		DrainNodeRequest
		DrainNodeResponse
		CreateTaskRequest
		CreateTaskResponse
		GetTasksRequest
		GetTasksResponse
		RemoveTaskRequest
		RemoveTaskResponse
		ListTasksRequest
		ListTasksResponse
		CreateJobRequest
		CreateJobResponse
		GetJobRequest
		GetJobResponse
		UpdateJobRequest
		UpdateJobResponse
		RemoveJobRequest
		RemoveJobResponse
		ListJobsRequest
		ListJobsResponse
		Update
		RegisterNodeRequest
		RegisterNodeResponse
		UpdateNodeStatusRequest
		UpdateNodeStatusResponse
		UpdateTaskStatusRequest
		UpdateTaskStatusResponse
		WatchTasksRequest
		WatchTasksResponse
*/
package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ResultParameters defines commn fields for controlling response formats.
// Fields may be used on requests or responses, depending on the role.
//
// All service request and response types that query large datasets should have
// the ResultParameters type available.
//
// As a convention, this should always be field id 64 for request/responses.
type ResultParameters struct {
	// Next provides a parameter specification based on the current one to get
	// the next set of results, pre-calculated by the server. This will be
	// empty if the result set is complete as of the containing message.
	Next *ResultParameters `protobuf:"bytes,1,opt,name=next" json:"next,omitempty"`
	// Total provides the total size of a given result set. Usually only set on
	// responses.
	Total uint64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	// Offset specifies the position to start from when returning a given
	// resultset.
	Offset uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit bounds the number of records in a response to the provided value.
	// Use with offset to provide close-open interval semantics over queries.
	Limit uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// ID may be set to select a particular result set that cannot be
	// expressed as an offset, limit tuple. The only source for this parameter
	// should be the "next" field but we reserve the possibility to define
	// future semantics.
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ResultParameters) Reset()      { *m = ResultParameters{} }
func (*ResultParameters) ProtoMessage() {}

func init() {
	proto.RegisterType((*ResultParameters)(nil), "api.ResultParameters")
}
func (this *ResultParameters) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&api.ResultParameters{")
	if this.Next != nil {
		s = append(s, "Next: "+fmt.Sprintf("%#v", this.Next)+",\n")
	}
	s = append(s, "Total: "+fmt.Sprintf("%#v", this.Total)+",\n")
	s = append(s, "Offset: "+fmt.Sprintf("%#v", this.Offset)+",\n")
	s = append(s, "Limit: "+fmt.Sprintf("%#v", this.Limit)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringService(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *ResultParameters) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ResultParameters) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Next != nil {
		data[i] = 0xa
		i++
		i = encodeVarintService(data, i, uint64(m.Next.Size()))
		n1, err := m.Next.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Total != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintService(data, i, uint64(m.Total))
	}
	if m.Offset != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintService(data, i, uint64(m.Offset))
	}
	if m.Limit != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintService(data, i, uint64(m.Limit))
	}
	if len(m.Id) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintService(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	return i, nil
}

func encodeFixed64Service(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Service(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintService(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ResultParameters) Size() (n int) {
	var l int
	_ = l
	if m.Next != nil {
		l = m.Next.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.Total != 0 {
		n += 1 + sovService(uint64(m.Total))
	}
	if m.Offset != 0 {
		n += 1 + sovService(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovService(uint64(m.Limit))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ResultParameters) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResultParameters{`,
		`Next:` + strings.Replace(fmt.Sprintf("%v", this.Next), "ResultParameters", "ResultParameters", 1) + `,`,
		`Total:` + fmt.Sprintf("%v", this.Total) + `,`,
		`Offset:` + fmt.Sprintf("%v", this.Offset) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ResultParameters) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResultParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Next", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Next == nil {
				m.Next = &ResultParameters{}
			}
			if err := m.Next.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Total |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Limit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func skipService(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipService(data[start:])
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
	ErrInvalidLengthService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService   = fmt.Errorf("proto: integer overflow")
)