// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lokahi.proto

/*
Package lokahi is a generated protocol buffer package.

lokahi is a HTTP health checking and response time monitoring service.

It is generated from these files:
	lokahi.proto

It has these top-level messages:
	CreateOpts
	CheckID
	Check
	CheckStatus
	ListOpts
	ChecksPage
	HistogramData
*/
package lokahi

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

// State is the finite state machine state.
type Check_State int32

const (
	Check_INIT  Check_State = 0
	Check_UP    Check_State = 1
	Check_DOWN  Check_State = 2
	Check_ERROR Check_State = 3
)

var Check_State_name = map[int32]string{
	0: "INIT",
	1: "UP",
	2: "DOWN",
	3: "ERROR",
}
var Check_State_value = map[string]int32{
	"INIT":  0,
	"UP":    1,
	"DOWN":  2,
	"ERROR": 3,
}

func (x Check_State) String() string {
	return proto.EnumName(Check_State_name, int32(x))
}
func (Check_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// CreateOpts contains arguments used to construct a Check.
type CreateOpts struct {
	// url is the HTTP/S url that will be monitored.
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// webhook_urls are the HTTP/S urls that will be HTTPS POST-ed to with the body
	// containing a protobuf-encoded CheckStatus message.
	WebhookUrls []string `protobuf:"bytes,2,rep,name=webhook_urls,json=webhookUrls" json:"webhook_urls,omitempty"`
	// health checks happen every number of seconds (minimum: 60, maximum: 600).
	Every int32 `protobuf:"varint,3,opt,name=every" json:"every,omitempty"`
	// playbook_url is the playbook URL that will be passed to status webhooks.
	PlaybookUrl string `protobuf:"bytes,4,opt,name=playbook_url,json=playbookUrl" json:"playbook_url,omitempty"`
}

func (m *CreateOpts) Reset()                    { *m = CreateOpts{} }
func (m *CreateOpts) String() string            { return proto.CompactTextString(m) }
func (*CreateOpts) ProtoMessage()               {}
func (*CreateOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateOpts) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateOpts) GetWebhookUrls() []string {
	if m != nil {
		return m.WebhookUrls
	}
	return nil
}

func (m *CreateOpts) GetEvery() int32 {
	if m != nil {
		return m.Every
	}
	return 0
}

func (m *CreateOpts) GetPlaybookUrl() string {
	if m != nil {
		return m.PlaybookUrl
	}
	return ""
}

// CheckID is a small wrapper around a Check ID.
type CheckID struct {
	// id is the Check id.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CheckID) Reset()                    { *m = CheckID{} }
func (m *CheckID) String() string            { return proto.CompactTextString(m) }
func (*CheckID) ProtoMessage()               {}
func (*CheckID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Check is an individual HTTP check.
type Check struct {
	// id is the unique id of this check.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// url is the HTTP/S url that will be monitored.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	// webhook_url is the HTTP/S url that will be HTTPS POST-ed to with the body
	// containing a protobuf-encoded CheckStatus message.
	WebhookUrl string `protobuf:"bytes,3,opt,name=webhook_url,json=webhookUrl" json:"webhook_url,omitempty"`
	// webhook_response_time_nanoseconds is last the response time of the webhook.
	WebhookResponseTimeNanoseconds int64 `protobuf:"varint,4,opt,name=webhook_response_time_nanoseconds,json=webhookResponseTimeNanoseconds" json:"webhook_response_time_nanoseconds,omitempty"`
	// health checks happen every number of seconds (minimum: 60, maximum: 600).
	Every int32 `protobuf:"varint,5,opt,name=every" json:"every,omitempty"`
	// playbook_url is the playbook URL that will be passed to status webhooks.
	PlaybookUrl string `protobuf:"bytes,6,opt,name=playbook_url,json=playbookUrl" json:"playbook_url,omitempty"`
	// state is the current state of this uptime check.
	State Check_State `protobuf:"varint,7,opt,name=state,enum=github.xe.lokahi.Check_State" json:"state,omitempty"`
}

func (m *Check) Reset()                    { *m = Check{} }
func (m *Check) String() string            { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()               {}
func (*Check) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Check) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Check) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Check) GetWebhookUrl() string {
	if m != nil {
		return m.WebhookUrl
	}
	return ""
}

func (m *Check) GetWebhookResponseTimeNanoseconds() int64 {
	if m != nil {
		return m.WebhookResponseTimeNanoseconds
	}
	return 0
}

func (m *Check) GetEvery() int32 {
	if m != nil {
		return m.Every
	}
	return 0
}

func (m *Check) GetPlaybookUrl() string {
	if m != nil {
		return m.PlaybookUrl
	}
	return ""
}

func (m *Check) GetState() Check_State {
	if m != nil {
		return m.State
	}
	return Check_INIT
}

// CheckStatus contains detailed information about the uptime status of a Check.
// This is POST-ed to webhook recipients.
type CheckStatus struct {
	// check is the information for the relevant Check.
	Check *Check `protobuf:"bytes,1,opt,name=check" json:"check,omitempty"`
	// last_response_time_nanoseconds is the last http web response time from the
	// Check's monitoring URL in nanoseconds.
	LastResponseTimeNanoseconds int64 `protobuf:"varint,2,opt,name=last_response_time_nanoseconds,json=lastResponseTimeNanoseconds" json:"last_response_time_nanoseconds,omitempty"`
	// histogram_data is the detailed histogram data for this check.
	HistogramData *HistogramData `protobuf:"bytes,3,opt,name=histogram_data,json=histogramData" json:"histogram_data,omitempty"`
}

func (m *CheckStatus) Reset()                    { *m = CheckStatus{} }
func (m *CheckStatus) String() string            { return proto.CompactTextString(m) }
func (*CheckStatus) ProtoMessage()               {}
func (*CheckStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckStatus) GetCheck() *Check {
	if m != nil {
		return m.Check
	}
	return nil
}

func (m *CheckStatus) GetLastResponseTimeNanoseconds() int64 {
	if m != nil {
		return m.LastResponseTimeNanoseconds
	}
	return 0
}

func (m *CheckStatus) GetHistogramData() *HistogramData {
	if m != nil {
		return m.HistogramData
	}
	return nil
}

// ListOpts contains options to the service Checks method List.
type ListOpts struct {
	// count is the number of checks that will be returned.
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// offset is the number of checks that will be offset.
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// include_status includes the histogram data for each check result.
	// This is O(scary). Use this with care.
	IncludeStatus bool `protobuf:"varint,3,opt,name=include_status,json=includeStatus" json:"include_status,omitempty"`
}

func (m *ListOpts) Reset()                    { *m = ListOpts{} }
func (m *ListOpts) String() string            { return proto.CompactTextString(m) }
func (*ListOpts) ProtoMessage()               {}
func (*ListOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListOpts) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListOpts) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListOpts) GetIncludeStatus() bool {
	if m != nil {
		return m.IncludeStatus
	}
	return false
}

// ChecksPage is a paginated Check list response.
type ChecksPage struct {
	Results []*ChecksPage_Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ChecksPage) Reset()                    { *m = ChecksPage{} }
func (m *ChecksPage) String() string            { return proto.CompactTextString(m) }
func (*ChecksPage) ProtoMessage()               {}
func (*ChecksPage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChecksPage) GetResults() []*ChecksPage_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// Result is an individual List result.
type ChecksPage_Result struct {
	// check is the individual check being listed.
	Check *Check `protobuf:"bytes,1,opt,name=check" json:"check,omitempty"`
	// histogram_data is the detailed histogram data for this check, this is
	// nornally not set unless include_status is set in ListOpts.
	HistogramData *HistogramData `protobuf:"bytes,2,opt,name=histogram_data,json=histogramData" json:"histogram_data,omitempty"`
}

func (m *ChecksPage_Result) Reset()                    { *m = ChecksPage_Result{} }
func (m *ChecksPage_Result) String() string            { return proto.CompactTextString(m) }
func (*ChecksPage_Result) ProtoMessage()               {}
func (*ChecksPage_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ChecksPage_Result) GetCheck() *Check {
	if m != nil {
		return m.Check
	}
	return nil
}

func (m *ChecksPage_Result) GetHistogramData() *HistogramData {
	if m != nil {
		return m.HistogramData
	}
	return nil
}

// HistogramData contains information from the HDR histogram maintained for
// each check.
type HistogramData struct {
	// max_nanoseconds is the maximum http web response time in nanoseconds.
	MaxNanoseconds float64 `protobuf:"fixed64,1,opt,name=max_nanoseconds,json=maxNanoseconds" json:"max_nanoseconds,omitempty"`
	// min_nanoseconds is the minimum http web response time in nanoseconds.
	MinNanoseconds float64 `protobuf:"fixed64,2,opt,name=min_nanoseconds,json=minNanoseconds" json:"min_nanoseconds,omitempty"`
	// mean_nanoseconds is the mean http web response time in nanoseconds.
	MeanNanoseconds float64 `protobuf:"fixed64,3,opt,name=mean_nanoseconds,json=meanNanoseconds" json:"mean_nanoseconds,omitempty"`
	// stddev is the standard deviation from the mean.
	Stddev float64 `protobuf:"fixed64,4,opt,name=stddev" json:"stddev,omitempty"`
	// p50_nanoseconds is the 50th percentile of the http web response times in
	// nanoseconds.
	P50Nanoseconds float64 `protobuf:"fixed64,5,opt,name=p50_nanoseconds,json=p50Nanoseconds" json:"p50_nanoseconds,omitempty"`
	// p75_nanoseconds is the 75th percentile of the http web response times in
	// nanoseconds.
	P75Nanoseconds float64 `protobuf:"fixed64,6,opt,name=p75_nanoseconds,json=p75Nanoseconds" json:"p75_nanoseconds,omitempty"`
	// p95_nanoseconds is the 95th percentile of the http web response times in
	// nanoseconds.
	P95Nanoseconds float64 `protobuf:"fixed64,7,opt,name=p95_nanoseconds,json=p95Nanoseconds" json:"p95_nanoseconds,omitempty"`
	// p99_nanoseconds is the 95th percentile of the http web response times in
	// nanoseconds.
	P99Nanoseconds float64 `protobuf:"fixed64,8,opt,name=p99_nanoseconds,json=p99Nanoseconds" json:"p99_nanoseconds,omitempty"`
}

func (m *HistogramData) Reset()                    { *m = HistogramData{} }
func (m *HistogramData) String() string            { return proto.CompactTextString(m) }
func (*HistogramData) ProtoMessage()               {}
func (*HistogramData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HistogramData) GetMaxNanoseconds() float64 {
	if m != nil {
		return m.MaxNanoseconds
	}
	return 0
}

func (m *HistogramData) GetMinNanoseconds() float64 {
	if m != nil {
		return m.MinNanoseconds
	}
	return 0
}

func (m *HistogramData) GetMeanNanoseconds() float64 {
	if m != nil {
		return m.MeanNanoseconds
	}
	return 0
}

func (m *HistogramData) GetStddev() float64 {
	if m != nil {
		return m.Stddev
	}
	return 0
}

func (m *HistogramData) GetP50Nanoseconds() float64 {
	if m != nil {
		return m.P50Nanoseconds
	}
	return 0
}

func (m *HistogramData) GetP75Nanoseconds() float64 {
	if m != nil {
		return m.P75Nanoseconds
	}
	return 0
}

func (m *HistogramData) GetP95Nanoseconds() float64 {
	if m != nil {
		return m.P95Nanoseconds
	}
	return 0
}

func (m *HistogramData) GetP99Nanoseconds() float64 {
	if m != nil {
		return m.P99Nanoseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateOpts)(nil), "github.xe.lokahi.CreateOpts")
	proto.RegisterType((*CheckID)(nil), "github.xe.lokahi.CheckID")
	proto.RegisterType((*Check)(nil), "github.xe.lokahi.Check")
	proto.RegisterType((*CheckStatus)(nil), "github.xe.lokahi.CheckStatus")
	proto.RegisterType((*ListOpts)(nil), "github.xe.lokahi.ListOpts")
	proto.RegisterType((*ChecksPage)(nil), "github.xe.lokahi.ChecksPage")
	proto.RegisterType((*ChecksPage_Result)(nil), "github.xe.lokahi.ChecksPage.Result")
	proto.RegisterType((*HistogramData)(nil), "github.xe.lokahi.HistogramData")
	proto.RegisterEnum("github.xe.lokahi.Check_State", Check_State_name, Check_State_value)
}

func init() { proto.RegisterFile("lokahi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x51, 0x6b, 0xdb, 0x48,
	0x10, 0xc7, 0x4f, 0x52, 0x24, 0xdb, 0xa3, 0xc4, 0x11, 0x4b, 0xc8, 0x39, 0xbe, 0x5c, 0xe2, 0xe8,
	0x38, 0xea, 0x3e, 0x54, 0x04, 0x07, 0x13, 0x0c, 0x2d, 0x2d, 0x89, 0x5b, 0x6a, 0x28, 0x49, 0xd8,
	0x26, 0x14, 0xfa, 0x62, 0xd6, 0xd6, 0x26, 0x16, 0x91, 0x25, 0xa1, 0x5d, 0xa5, 0x49, 0x5f, 0xfa,
	0xa1, 0xfa, 0x1d, 0x4a, 0x0b, 0xfd, 0x50, 0x65, 0x57, 0x52, 0x6d, 0xd9, 0x95, 0xa1, 0x7d, 0xf3,
	0x8c, 0x7e, 0xa3, 0xf9, 0xcf, 0x7f, 0xb4, 0x6b, 0x58, 0xf7, 0xc3, 0x5b, 0x32, 0xf1, 0x9c, 0x28,
	0x0e, 0x79, 0x88, 0xac, 0x1b, 0x8f, 0x4f, 0x92, 0x91, 0x73, 0x4f, 0x9d, 0x34, 0x6f, 0x7f, 0x04,
	0x38, 0x8d, 0x29, 0xe1, 0xf4, 0x3c, 0xe2, 0x0c, 0x59, 0xa0, 0x25, 0xb1, 0xdf, 0x50, 0x5a, 0x4a,
	0xbb, 0x86, 0xc5, 0x4f, 0x74, 0x00, 0xeb, 0x1f, 0xe8, 0x68, 0x12, 0x86, 0xb7, 0xc3, 0x24, 0xf6,
	0x59, 0x43, 0x6d, 0x69, 0xed, 0x1a, 0x36, 0xb3, 0xdc, 0x55, 0xec, 0x33, 0xb4, 0x05, 0x3a, 0xbd,
	0xa3, 0xf1, 0x43, 0x43, 0x6b, 0x29, 0x6d, 0x1d, 0xa7, 0x81, 0x28, 0x8c, 0x7c, 0xf2, 0x30, 0xca,
	0x2a, 0x1b, 0x6b, 0xf2, 0x9d, 0x66, 0x9e, 0xbb, 0x8a, 0x7d, 0x7b, 0x07, 0x2a, 0xa7, 0x13, 0x3a,
	0xbe, 0x1d, 0xf4, 0x51, 0x1d, 0x54, 0xcf, 0xcd, 0xfa, 0xaa, 0x9e, 0x6b, 0x7f, 0x56, 0x41, 0x97,
	0xcf, 0x16, 0x9f, 0xe4, 0x12, 0xd5, 0x99, 0xc4, 0x7d, 0x30, 0xe7, 0x24, 0x4a, 0x15, 0x35, 0x0c,
	0x33, 0x85, 0x68, 0x00, 0x07, 0x39, 0x10, 0x53, 0x16, 0x85, 0x01, 0xa3, 0x43, 0xee, 0x4d, 0xe9,
	0x30, 0x20, 0x41, 0xc8, 0xe8, 0x38, 0x0c, 0x5c, 0x26, 0xf5, 0x69, 0x78, 0x2f, 0x03, 0x71, 0xc6,
	0x5d, 0x7a, 0x53, 0x7a, 0x36, 0xa3, 0x66, 0xb3, 0xea, 0xab, 0x66, 0x35, 0x96, 0x66, 0x45, 0x47,
	0xa0, 0x33, 0x4e, 0x38, 0x6d, 0x54, 0x5a, 0x4a, 0xbb, 0xde, 0xf9, 0xd7, 0x59, 0xdc, 0x84, 0x23,
	0xc7, 0x75, 0xde, 0x0a, 0x08, 0xa7, 0xac, 0xed, 0x80, 0x2e, 0x63, 0x54, 0x85, 0xb5, 0xc1, 0xd9,
	0xe0, 0xd2, 0xfa, 0x0b, 0x19, 0xa0, 0x5e, 0x5d, 0x58, 0x8a, 0xc8, 0xf4, 0xcf, 0xdf, 0x9d, 0x59,
	0x2a, 0xaa, 0x81, 0xfe, 0x12, 0xe3, 0x73, 0x6c, 0x69, 0xf6, 0x37, 0x05, 0x4c, 0xf9, 0x1a, 0x51,
	0x95, 0x30, 0xf4, 0x04, 0xf4, 0xb1, 0x08, 0xa5, 0x7d, 0x66, 0xe7, 0xef, 0x92, 0xa6, 0x38, 0xa5,
	0xd0, 0x29, 0xec, 0xf9, 0x84, 0xf1, 0x15, 0x26, 0xa9, 0xd2, 0xa4, 0x7f, 0x04, 0x55, 0xe6, 0xd0,
	0x2b, 0xa8, 0x4f, 0x3c, 0xc6, 0xc3, 0x9b, 0x98, 0x4c, 0x87, 0x2e, 0xe1, 0x44, 0x2e, 0xc4, 0xec,
	0xec, 0x2f, 0x37, 0x7f, 0x9d, 0x73, 0x7d, 0xc2, 0x09, 0xde, 0x98, 0xcc, 0x87, 0xf6, 0x10, 0xaa,
	0x6f, 0x3c, 0xc6, 0xe5, 0x67, 0xb9, 0x05, 0xfa, 0x38, 0x4c, 0x02, 0x2e, 0xe7, 0xd0, 0x71, 0x1a,
	0xa0, 0x6d, 0x30, 0xc2, 0xeb, 0x6b, 0x46, 0xb9, 0x94, 0xa5, 0xe3, 0x2c, 0x42, 0xff, 0x43, 0xdd,
	0x0b, 0xc6, 0x7e, 0xe2, 0xd2, 0x21, 0x93, 0x3e, 0x48, 0x05, 0x55, 0xbc, 0x91, 0x65, 0x53, 0x73,
	0xec, 0xef, 0x0a, 0x80, 0x1c, 0x9f, 0x5d, 0x90, 0x1b, 0x8a, 0x9e, 0x41, 0x25, 0xa6, 0x2c, 0xf1,
	0x39, 0x6b, 0x28, 0x2d, 0xad, 0x6d, 0x76, 0xfe, 0x2b, 0x71, 0x4b, 0xe2, 0x0e, 0x96, 0x2c, 0xce,
	0x6b, 0x9a, 0x9f, 0xc0, 0x48, 0x53, 0xbf, 0x6b, 0xfa, 0xb2, 0x5f, 0xea, 0x1f, 0xf9, 0xf5, 0x55,
	0x85, 0x8d, 0x02, 0x80, 0x1e, 0xc1, 0xe6, 0x94, 0xdc, 0x17, 0xf6, 0x27, 0x24, 0x29, 0xb8, 0x3e,
	0x25, 0xf7, 0xf3, 0x2b, 0x13, 0xa0, 0x17, 0x2c, 0x2d, 0x5a, 0x80, 0x5e, 0x30, 0x0f, 0x3e, 0x06,
	0x6b, 0x4a, 0x49, 0x91, 0xd4, 0x24, 0xb9, 0x29, 0xf2, 0xf3, 0xe8, 0x36, 0x18, 0x8c, 0xbb, 0x2e,
	0xbd, 0x93, 0x07, 0x4b, 0xc1, 0x59, 0x24, 0x7a, 0x45, 0xdd, 0xc3, 0xc2, 0x1b, 0xf4, 0xb4, 0x57,
	0xd4, 0x3d, 0x5c, 0x10, 0x15, 0x1d, 0x77, 0x0b, 0xa0, 0x91, 0x81, 0xc7, 0xdd, 0x45, 0xb0, 0x57,
	0x04, 0x2b, 0x19, 0xd8, 0x5b, 0x06, 0x7b, 0x05, 0xb0, 0x9a, 0x83, 0xbd, 0x39, 0xb0, 0xf3, 0x45,
	0x05, 0x23, 0x5d, 0x35, 0x7a, 0x0e, 0x46, 0x7a, 0x3d, 0xa2, 0xdd, 0x5f, 0xec, 0xf1, 0xe7, 0xc5,
	0xd9, 0x2c, 0xdb, 0x32, 0x7a, 0x0a, 0x46, 0x9f, 0xfa, 0x94, 0x53, 0xb4, 0x53, 0x82, 0x0c, 0xfa,
	0xe5, 0xd5, 0x2f, 0x60, 0x4d, 0x1c, 0x02, 0xd4, 0x5c, 0x06, 0xf2, 0xc3, 0xd1, 0xdc, 0x5d, 0xf5,
	0x9d, 0xa2, 0x63, 0xd0, 0x2e, 0x12, 0x8e, 0xca, 0x3a, 0x94, 0xb7, 0x3e, 0x01, 0x23, 0xbb, 0x45,
	0x56, 0x08, 0x2f, 0xbb, 0xc6, 0xd2, 0xca, 0x93, 0xea, 0x7b, 0x23, 0xcd, 0x8e, 0x0c, 0xf9, 0xff,
	0x73, 0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xe2, 0xc6, 0x14, 0x8f, 0x06, 0x00, 0x00,
}