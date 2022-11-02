// Code generated by protoc-gen-go.
// source: polaris_ratelimit.proto
// DO NOT EDIT!

/*
Package metric is a generated protocol buffer package.

It is generated from these files:

	polaris_ratelimit.proto

It has these top-level messages:

	RateLimitRequest
	RateLimitResponse
	Limiter
*/
package metric

import (
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"
	google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RateLimitRequest struct {
	Key       *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Service   *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Namespace *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Timestamp *google_protobuf.Int64Value  `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Totals    []*Limiter                   `protobuf:"bytes,5,rep,name=totals" json:"totals,omitempty"`
	Useds     []*Limiter                   `protobuf:"bytes,6,rep,name=useds" json:"useds,omitempty"`
}

func (m *RateLimitRequest) Reset()                    { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string            { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()               {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *RateLimitRequest) GetKey() *google_protobuf.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RateLimitRequest) GetService() *google_protobuf.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *RateLimitRequest) GetNamespace() *google_protobuf.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *RateLimitRequest) GetTimestamp() *google_protobuf.Int64Value {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *RateLimitRequest) GetTotals() []*Limiter {
	if m != nil {
		return m.Totals
	}
	return nil
}

func (m *RateLimitRequest) GetUseds() []*Limiter {
	if m != nil {
		return m.Useds
	}
	return nil
}

type RateLimitResponse struct {
	Code      *google_protobuf.UInt32Value `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Info      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Key       *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Timestamp *google_protobuf.Int64Value  `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	SumUseds  []*Limiter                   `protobuf:"bytes,5,rep,name=sum_useds,json=sumUseds" json:"sum_useds,omitempty"`
}

func (m *RateLimitResponse) Reset()                    { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string            { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()               {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RateLimitResponse) GetCode() *google_protobuf.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *RateLimitResponse) GetInfo() *google_protobuf.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *RateLimitResponse) GetKey() *google_protobuf.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RateLimitResponse) GetTimestamp() *google_protobuf.Int64Value {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *RateLimitResponse) GetSumUseds() []*Limiter {
	if m != nil {
		return m.SumUseds
	}
	return nil
}

type Limiter struct {
	Amount   *google_protobuf.UInt32Value `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
	Duration *google_protobuf1.Duration   `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
	Limited  *google_protobuf.UInt32Value `protobuf:"bytes,3,opt,name=limited" json:"limited,omitempty"`
}

func (m *Limiter) Reset()                    { *m = Limiter{} }
func (m *Limiter) String() string            { return proto.CompactTextString(m) }
func (*Limiter) ProtoMessage()               {}
func (*Limiter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Limiter) GetAmount() *google_protobuf.UInt32Value {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Limiter) GetDuration() *google_protobuf1.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Limiter) GetLimited() *google_protobuf.UInt32Value {
	if m != nil {
		return m.Limited
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitRequest)(nil), "metric.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "metric.RateLimitResponse")
	proto.RegisterType((*Limiter)(nil), "metric.Limiter")
}

func init() { proto.RegisterFile("polaris_ratelimit.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0xd1, 0x68, 0xd4, 0x71, 0x71, 0xef, 0x9d, 0xcd, 0xcd, 0xfd, 0x43, 0x11, 0xa1, 0xd4,
	0x45, 0x89, 0x45, 0xad, 0xd0, 0xae, 0xbb, 0x11, 0xba, 0x4a, 0xb1, 0x5b, 0x19, 0x93, 0xa3, 0x0c,
	0xcd, 0xcc, 0xa4, 0x33, 0x67, 0x5a, 0xfa, 0x16, 0x7d, 0x92, 0xbe, 0x43, 0xdf, 0xac, 0x64, 0x92,
	0x58, 0xc1, 0x0a, 0x81, 0x6e, 0xe7, 0xfb, 0x7d, 0x27, 0xe1, 0x37, 0x67, 0xc8, 0xef, 0x4c, 0xa5,
	0x4c, 0x73, 0xb3, 0xd2, 0x0c, 0x21, 0xe5, 0x82, 0x63, 0x98, 0x69, 0x85, 0x8a, 0xfa, 0x02, 0x50,
	0xf3, 0xf8, 0xef, 0xc9, 0x56, 0xa9, 0x6d, 0x0a, 0x63, 0x77, 0xba, 0xb6, 0x9b, 0xf1, 0xb3, 0x66,
	0x59, 0x06, 0xda, 0x14, 0xdc, 0x61, 0x9e, 0x58, 0xcd, 0x90, 0x2b, 0x59, 0xe4, 0xc3, 0xf7, 0x26,
	0xf9, 0x19, 0x31, 0x84, 0xdb, 0x7c, 0x76, 0x04, 0x8f, 0x16, 0x0c, 0xd2, 0x90, 0x78, 0x0f, 0xf0,
	0x12, 0x34, 0x06, 0x8d, 0x51, 0x7f, 0xf2, 0x3f, 0x2c, 0x46, 0x84, 0xd5, 0x88, 0xf0, 0x0e, 0x35,
	0x97, 0xdb, 0x7b, 0x96, 0x5a, 0x88, 0x72, 0x90, 0xce, 0x49, 0xc7, 0x80, 0x7e, 0xe2, 0x31, 0x04,
	0xcd, 0x1a, 0x9d, 0x0a, 0xa6, 0xd7, 0xa4, 0x27, 0x99, 0x00, 0x93, 0xb1, 0x18, 0x02, 0xaf, 0x46,
	0xf3, 0x13, 0xa7, 0x57, 0xa4, 0x87, 0x5c, 0x80, 0x41, 0x26, 0xb2, 0xa0, 0xe5, 0xba, 0xff, 0x0e,
	0xba, 0x0b, 0x89, 0xf3, 0x59, 0x59, 0xdd, 0xd1, 0xf4, 0x8c, 0xf8, 0xa8, 0x90, 0xa5, 0x26, 0x68,
	0x0f, 0xbc, 0x51, 0x7f, 0xf2, 0x23, 0x2c, 0x64, 0x86, 0x4e, 0x02, 0xe8, 0xa8, 0x8c, 0xe9, 0x29,
	0x69, 0x5b, 0x03, 0x89, 0x09, 0xfc, 0xaf, 0xb9, 0x22, 0x1d, 0xbe, 0x36, 0xc9, 0xaf, 0x3d, 0x87,
	0x26, 0x53, 0xd2, 0x00, 0xbd, 0x20, 0xad, 0x58, 0x25, 0x70, 0xd4, 0xe2, 0x72, 0x21, 0x71, 0x3a,
	0x29, 0x7e, 0xce, 0x91, 0x79, 0x83, 0xcb, 0x8d, 0xaa, 0xe5, 0xd0, 0x91, 0xd5, 0x45, 0x79, 0x75,
	0x2f, 0xea, 0x1b, 0xd2, 0xce, 0x49, 0xcf, 0x58, 0xb1, 0x2a, 0x7c, 0x1c, 0xf1, 0xd6, 0x35, 0x56,
	0x2c, 0x9d, 0x92, 0xb7, 0x06, 0xe9, 0x94, 0xa7, 0x74, 0x46, 0x7c, 0x26, 0x94, 0x95, 0x58, 0x4b,
	0x45, 0xc9, 0xd2, 0x4b, 0xd2, 0xad, 0x56, 0xb5, 0x14, 0xf2, 0xe7, 0xa0, 0x77, 0x53, 0x02, 0xd1,
	0x0e, 0xcd, 0x57, 0xd1, 0x3d, 0x13, 0x48, 0x8e, 0x5a, 0xd9, 0xff, 0x5a, 0x05, 0xaf, 0x7d, 0x17,
	0x4f, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x35, 0x57, 0x65, 0xce, 0x71, 0x03, 0x00, 0x00,
}
