// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lokahiadmin.proto

/*
Package lokahiadmin is a generated protocol buffer package.

lokahiadmin is the administrative/backend API for lokahi power usage.

It is generated from these files:
	lokahiadmin.proto

It has these top-level messages:
	CheckIDs
	Run
	Nil
	HistogramData
*/
package lokahiadmin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckIDs struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *CheckIDs) Reset()                    { *m = CheckIDs{} }
func (m *CheckIDs) String() string            { return proto.CompactTextString(m) }
func (*CheckIDs) ProtoMessage()               {}
func (*CheckIDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckIDs) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Run struct {
	Id                 string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Cids               *CheckIDs              `protobuf:"bytes,2,opt,name=cids" json:"cids,omitempty"`
	Finished           bool                   `protobuf:"varint,3,opt,name=finished" json:"finished,omitempty"`
	Results            map[string]*Run_Health `protobuf:"bytes,4,rep,name=results" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	StartTimeUnix      int64                  `protobuf:"varint,5,opt,name=start_time_unix,json=startTimeUnix" json:"start_time_unix,omitempty"`
	ElapsedNanoseconds int64                  `protobuf:"varint,6,opt,name=elapsed_nanoseconds,json=elapsedNanoseconds" json:"elapsed_nanoseconds,omitempty"`
}

func (m *Run) Reset()                    { *m = Run{} }
func (m *Run) String() string            { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()               {}
func (*Run) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Run) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Run) GetCids() *CheckIDs {
	if m != nil {
		return m.Cids
	}
	return nil
}

func (m *Run) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func (m *Run) GetResults() map[string]*Run_Health {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Run) GetStartTimeUnix() int64 {
	if m != nil {
		return m.StartTimeUnix
	}
	return 0
}

func (m *Run) GetElapsedNanoseconds() int64 {
	if m != nil {
		return m.ElapsedNanoseconds
	}
	return 0
}

type Run_Health struct {
	Url                     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	ResponseTimeNanoseconds int64  `protobuf:"varint,2,opt,name=response_time_nanoseconds,json=responseTimeNanoseconds" json:"response_time_nanoseconds,omitempty"`
	StatusCode              int32  `protobuf:"varint,3,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Body                    string `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	Error                   string `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	Healthy                 bool   `protobuf:"varint,6,opt,name=healthy" json:"healthy,omitempty"`
}

func (m *Run_Health) Reset()                    { *m = Run_Health{} }
func (m *Run_Health) String() string            { return proto.CompactTextString(m) }
func (*Run_Health) ProtoMessage()               {}
func (*Run_Health) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Run_Health) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Run_Health) GetResponseTimeNanoseconds() int64 {
	if m != nil {
		return m.ResponseTimeNanoseconds
	}
	return 0
}

func (m *Run_Health) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Run_Health) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Run_Health) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Run_Health) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

type Nil struct {
}

func (m *Nil) Reset()                    { *m = Nil{} }
func (m *Nil) String() string            { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()               {}
func (*Nil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// HistogramData contains information from the HDR histogram maintained for
// each check.
type HistogramData struct {
	// max_nanoseconds is the maximum http web response time in nanoseconds.
	MaxNanoseconds int64 `protobuf:"varint,1,opt,name=max_nanoseconds,json=maxNanoseconds" json:"max_nanoseconds,omitempty"`
	// min_nanoseconds is the minimum http web response time in nanoseconds.
	MinNanoseconds int64 `protobuf:"varint,2,opt,name=min_nanoseconds,json=minNanoseconds" json:"min_nanoseconds,omitempty"`
	// mean_nanoseconds is the mean http web response time in nanoseconds.
	MeanNanoseconds int64 `protobuf:"varint,3,opt,name=mean_nanoseconds,json=meanNanoseconds" json:"mean_nanoseconds,omitempty"`
	// stddev is the standard deviation from the mean.
	Stddev int64 `protobuf:"varint,4,opt,name=stddev" json:"stddev,omitempty"`
	// p50_nanoseconds is the 50th percentile of the http web response times in
	// nanoseconds.
	P50Nanoseconds int64 `protobuf:"varint,5,opt,name=p50_nanoseconds,json=p50Nanoseconds" json:"p50_nanoseconds,omitempty"`
	// p75_nanoseconds is the 75th percentile of the http web response times in
	// nanoseconds.
	P75Nanoseconds int64 `protobuf:"varint,6,opt,name=p75_nanoseconds,json=p75Nanoseconds" json:"p75_nanoseconds,omitempty"`
	// p95_nanoseconds is the 95th percentile of the http web response times in
	// nanoseconds.
	P95Nanoseconds int64 `protobuf:"varint,7,opt,name=p95_nanoseconds,json=p95Nanoseconds" json:"p95_nanoseconds,omitempty"`
	// p99_nanoseconds is the 95th percentile of the http web response times in
	// nanoseconds.
	P99Nanoseconds int64 `protobuf:"varint,8,opt,name=p99_nanoseconds,json=p99Nanoseconds" json:"p99_nanoseconds,omitempty"`
}

func (m *HistogramData) Reset()                    { *m = HistogramData{} }
func (m *HistogramData) String() string            { return proto.CompactTextString(m) }
func (*HistogramData) ProtoMessage()               {}
func (*HistogramData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HistogramData) GetMaxNanoseconds() int64 {
	if m != nil {
		return m.MaxNanoseconds
	}
	return 0
}

func (m *HistogramData) GetMinNanoseconds() int64 {
	if m != nil {
		return m.MinNanoseconds
	}
	return 0
}

func (m *HistogramData) GetMeanNanoseconds() int64 {
	if m != nil {
		return m.MeanNanoseconds
	}
	return 0
}

func (m *HistogramData) GetStddev() int64 {
	if m != nil {
		return m.Stddev
	}
	return 0
}

func (m *HistogramData) GetP50Nanoseconds() int64 {
	if m != nil {
		return m.P50Nanoseconds
	}
	return 0
}

func (m *HistogramData) GetP75Nanoseconds() int64 {
	if m != nil {
		return m.P75Nanoseconds
	}
	return 0
}

func (m *HistogramData) GetP95Nanoseconds() int64 {
	if m != nil {
		return m.P95Nanoseconds
	}
	return 0
}

func (m *HistogramData) GetP99Nanoseconds() int64 {
	if m != nil {
		return m.P99Nanoseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckIDs)(nil), "github.xe.lokahi.admin.CheckIDs")
	proto.RegisterType((*Run)(nil), "github.xe.lokahi.admin.Run")
	proto.RegisterType((*Run_Health)(nil), "github.xe.lokahi.admin.Run.Health")
	proto.RegisterType((*Nil)(nil), "github.xe.lokahi.admin.Nil")
	proto.RegisterType((*HistogramData)(nil), "github.xe.lokahi.admin.HistogramData")
}

func init() { proto.RegisterFile("lokahiadmin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0xa4, 0xe9, 0xb2, 0x53, 0xd6, 0x0d, 0x83, 0x46, 0x28, 0x48, 0x44, 0x95, 0x80,
	0x70, 0x13, 0xa6, 0x41, 0xb5, 0x75, 0x97, 0xdb, 0x10, 0x43, 0xa0, 0x5d, 0x18, 0xb8, 0xe1, 0x82,
	0xca, 0xad, 0xcd, 0x62, 0x35, 0xb1, 0xa3, 0xd8, 0x99, 0xda, 0xb7, 0x41, 0xbc, 0x06, 0x2f, 0xc0,
	0x63, 0xa1, 0x38, 0x0d, 0x24, 0xa3, 0x1d, 0x77, 0x3e, 0x3e, 0x9f, 0x7f, 0xfb, 0xfc, 0xc7, 0x36,
	0xdc, 0x4d, 0xe4, 0x9c, 0xc4, 0x9c, 0xd0, 0x94, 0x8b, 0x28, 0xcb, 0xa5, 0x96, 0x68, 0xff, 0x8a,
	0xeb, 0xb8, 0x98, 0x46, 0x0b, 0x16, 0x55, 0xc9, 0xc8, 0x64, 0x87, 0x8f, 0xc1, 0x3b, 0x8b, 0xd9,
	0x6c, 0xfe, 0xee, 0x5c, 0xa1, 0x3d, 0x70, 0x38, 0x55, 0xbe, 0x15, 0x38, 0xe1, 0x36, 0x2e, 0x87,
	0xc3, 0x1f, 0x1d, 0x70, 0x70, 0x21, 0x50, 0x1f, 0x6c, 0x4e, 0x7d, 0x2b, 0xb0, 0xc2, 0x6d, 0x6c,
	0x73, 0x8a, 0x5e, 0x43, 0x67, 0x56, 0xa2, 0x76, 0x60, 0x85, 0xbd, 0xc3, 0x20, 0x5a, 0x2f, 0x1e,
	0xd5, 0xca, 0xd8, 0xd0, 0x68, 0x00, 0xde, 0x37, 0x2e, 0xb8, 0x8a, 0x19, 0xf5, 0x9d, 0xc0, 0x0a,
	0x3d, 0xfc, 0x27, 0x46, 0xa7, 0xb0, 0x95, 0x33, 0x55, 0x24, 0x5a, 0xf9, 0x9d, 0xc0, 0x09, 0x7b,
	0x87, 0xe1, 0x26, 0x51, 0x5c, 0x88, 0x08, 0x57, 0xe8, 0x1b, 0xa1, 0xf3, 0x25, 0xae, 0x17, 0xa2,
	0x67, 0xb0, 0xab, 0x34, 0xc9, 0xf5, 0x44, 0xf3, 0x94, 0x4d, 0x0a, 0xc1, 0x17, 0xbe, 0x1b, 0x58,
	0xa1, 0x83, 0x77, 0xcc, 0xf4, 0x27, 0x9e, 0xb2, 0xcf, 0x82, 0x2f, 0xd0, 0x4b, 0xb8, 0xc7, 0x12,
	0x92, 0x29, 0x46, 0x27, 0x82, 0x08, 0xa9, 0xd8, 0x4c, 0x0a, 0xaa, 0xfc, 0xae, 0x61, 0xd1, 0x2a,
	0x75, 0xf9, 0x37, 0x33, 0xf8, 0x69, 0x41, 0xf7, 0x82, 0x91, 0x44, 0xc7, 0xa5, 0x47, 0x45, 0x9e,
	0xac, 0xac, 0x28, 0x87, 0xe8, 0x04, 0x1e, 0xe6, 0x4c, 0x65, 0x52, 0x28, 0x56, 0x6d, 0xdc, 0xd4,
	0xb4, 0x8d, 0xe6, 0x83, 0x1a, 0x28, 0x8f, 0xd0, 0x10, 0x46, 0x4f, 0xa0, 0xa7, 0x34, 0xd1, 0x85,
	0x9a, 0xcc, 0x24, 0x65, 0xc6, 0x14, 0x17, 0x43, 0x35, 0x75, 0x26, 0x29, 0x43, 0x08, 0x3a, 0x53,
	0x49, 0x97, 0x7e, 0xc7, 0xec, 0x67, 0xc6, 0xe8, 0x3e, 0xb8, 0x2c, 0xcf, 0x65, 0x6e, 0x8a, 0xdb,
	0xc6, 0x55, 0x80, 0x7c, 0xd8, 0x8a, 0xcd, 0x11, 0x97, 0xa6, 0x10, 0x0f, 0xd7, 0xe1, 0xe0, 0x2b,
	0xdc, 0x69, 0xfa, 0x55, 0x96, 0x30, 0x67, 0xcb, 0xba, 0x84, 0x39, 0x5b, 0xa2, 0x63, 0x70, 0xaf,
	0x49, 0x52, 0xb0, 0x55, 0x3f, 0x87, 0xb7, 0x59, 0x5f, 0xf9, 0x80, 0xab, 0x05, 0x27, 0xf6, 0xb1,
	0x35, 0x74, 0xc1, 0xb9, 0xe4, 0xc9, 0xf0, 0x97, 0x0d, 0x3b, 0x17, 0x5c, 0x69, 0x79, 0x95, 0x93,
	0xf4, 0x9c, 0x68, 0x82, 0x9e, 0xc3, 0x6e, 0x4a, 0x16, 0x2d, 0x3f, 0x2c, 0xe3, 0x47, 0x3f, 0x25,
	0x8b, 0xa6, 0x0d, 0x25, 0xc8, 0xc5, 0x1a, 0xe3, 0xfa, 0x29, 0x17, 0x4d, 0xf0, 0x05, 0xec, 0xa5,
	0x8c, 0xb4, 0x49, 0xc7, 0x90, 0xbb, 0xe5, 0x7c, 0x13, 0xdd, 0x87, 0xae, 0xd2, 0x94, 0xb2, 0x6b,
	0xe3, 0x9d, 0x83, 0x57, 0x51, 0xb9, 0x57, 0x36, 0x3a, 0x68, 0x29, 0x54, 0x97, 0xa4, 0x9f, 0x8d,
	0x0e, 0x6e, 0x1c, 0x2a, 0x3b, 0x1a, 0xad, 0xb9, 0x21, 0xfd, 0xec, 0x68, 0x74, 0x13, 0x1c, 0xb7,
	0xc1, 0xad, 0x15, 0x38, 0xfe, 0x17, 0x1c, 0xb7, 0x40, 0xaf, 0x06, 0xc7, 0x0d, 0xf0, 0xf0, 0xbb,
	0x05, 0x1e, 0x2e, 0xc4, 0x07, 0x39, 0x23, 0x09, 0x7a, 0x0b, 0x5d, 0xf3, 0x8e, 0x14, 0xfa, 0xef,
	0x3b, 0x1b, 0x3c, 0xba, 0xa5, 0x73, 0xe8, 0x3d, 0xb8, 0x1f, 0x35, 0xd1, 0x0a, 0x6d, 0xa4, 0x2e,
	0x79, 0x32, 0x78, 0xba, 0x29, 0xd9, 0xea, 0xed, 0xe9, 0xce, 0x97, 0x5e, 0xe3, 0x93, 0x99, 0x76,
	0xcd, 0x2f, 0xf3, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x4d, 0x8e, 0x55, 0x7a, 0x04,
	0x00, 0x00,
}
