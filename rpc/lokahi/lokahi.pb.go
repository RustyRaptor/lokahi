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
	Nil
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
	// webhook_url is the HTTP/S url that will be HTTPS POST-ed to with the body
	// containing a protobuf-encoded CheckStatus message.
	WebhookUrl string `protobuf:"bytes,2,opt,name=webhook_url,json=webhookUrl" json:"webhook_url,omitempty"`
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

func (m *CreateOpts) GetWebhookUrl() string {
	if m != nil {
		return m.WebhookUrl
	}
	return ""
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
	// page is the page of results to fetch.
	Page int32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
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

func (m *ListOpts) GetPage() int32 {
	if m != nil {
		return m.Page
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

type Nil struct {
}

func (m *Nil) Reset()                    { *m = Nil{} }
func (m *Nil) String() string            { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()               {}
func (*Nil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*CreateOpts)(nil), "github.xe.lokahi.CreateOpts")
	proto.RegisterType((*CheckID)(nil), "github.xe.lokahi.CheckID")
	proto.RegisterType((*Check)(nil), "github.xe.lokahi.Check")
	proto.RegisterType((*CheckStatus)(nil), "github.xe.lokahi.CheckStatus")
	proto.RegisterType((*ListOpts)(nil), "github.xe.lokahi.ListOpts")
	proto.RegisterType((*ChecksPage)(nil), "github.xe.lokahi.ChecksPage")
	proto.RegisterType((*ChecksPage_Result)(nil), "github.xe.lokahi.ChecksPage.Result")
	proto.RegisterType((*HistogramData)(nil), "github.xe.lokahi.HistogramData")
	proto.RegisterType((*Nil)(nil), "github.xe.lokahi.Nil")
	proto.RegisterEnum("github.xe.lokahi.Check_State", Check_State_name, Check_State_value)
}

func init() { proto.RegisterFile("lokahi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb1, 0x1d, 0x3b, 0xe9, 0xa4, 0x4d, 0xa3, 0x55, 0x81, 0x34, 0x94, 0x36, 0x35, 0x42,
	0x84, 0x03, 0x56, 0x95, 0x2a, 0xaa, 0x22, 0xf1, 0xa5, 0x36, 0x40, 0x23, 0xa1, 0xb4, 0x5a, 0x5a,
	0x55, 0x82, 0x43, 0xb4, 0x89, 0x57, 0x8d, 0x55, 0xc7, 0xb6, 0xbc, 0x9b, 0x92, 0x9e, 0x10, 0xcf,
	0xc4, 0x4b, 0x80, 0xc4, 0x43, 0xa1, 0x5d, 0xdb, 0x34, 0x8e, 0x71, 0x24, 0x7a, 0xf3, 0xcc, 0xfe,
	0xc6, 0x33, 0xf3, 0x1f, 0xef, 0x18, 0x56, 0x5d, 0xff, 0x8a, 0x8c, 0x1d, 0x2b, 0x08, 0x7d, 0xee,
	0xa3, 0xea, 0xa5, 0xc3, 0xc7, 0xd3, 0xa1, 0x35, 0xa3, 0x56, 0xe4, 0x37, 0x67, 0x00, 0x47, 0x21,
	0x25, 0x9c, 0x9e, 0x04, 0x9c, 0xa1, 0x2a, 0x68, 0xd3, 0xd0, 0xad, 0x29, 0x0d, 0xa5, 0xb9, 0x82,
	0xc5, 0x23, 0xda, 0x81, 0xf2, 0x57, 0x3a, 0x1c, 0xfb, 0xfe, 0xd5, 0x40, 0x9c, 0xa8, 0xf2, 0x04,
	0x62, 0xd7, 0x79, 0xe8, 0xa2, 0x0d, 0xd0, 0xe9, 0x35, 0x0d, 0x6f, 0x6a, 0x5a, 0x43, 0x69, 0xea,
	0x38, 0x32, 0xd0, 0x2e, 0xac, 0x06, 0x2e, 0xb9, 0x19, 0x26, 0x71, 0x05, 0x19, 0x57, 0x4e, 0x7c,
	0xe7, 0xa1, 0x6b, 0x6e, 0x42, 0xf1, 0x68, 0x4c, 0x47, 0x57, 0xbd, 0x2e, 0xaa, 0x80, 0xea, 0xd8,
	0x71, 0x56, 0xd5, 0xb1, 0xcd, 0x1f, 0x2a, 0xe8, 0xf2, 0x6c, 0xf1, 0x24, 0x29, 0x50, 0xcd, 0x2d,
	0x50, 0xcb, 0x14, 0xd8, 0x83, 0xdd, 0x04, 0x08, 0x29, 0x0b, 0x7c, 0x8f, 0xd1, 0x01, 0x77, 0x26,
	0x74, 0xe0, 0x11, 0xcf, 0x67, 0x74, 0xe4, 0x7b, 0x36, 0x93, 0xf5, 0x69, 0x78, 0x3b, 0x06, 0x71,
	0xcc, 0x9d, 0x39, 0x13, 0xda, 0xbf, 0xa5, 0x6e, 0x7b, 0xd5, 0x97, 0xf5, 0x6a, 0x64, 0x7a, 0x45,
	0xfb, 0xa0, 0x33, 0x4e, 0x38, 0xad, 0x15, 0x1b, 0x4a, 0xb3, 0xd2, 0x7a, 0x6c, 0x2d, 0xce, 0xc1,
	0x92, 0xed, 0x5a, 0x9f, 0x04, 0x84, 0x23, 0xd6, 0xb4, 0x40, 0x97, 0x36, 0x2a, 0x41, 0xa1, 0xd7,
	0xef, 0x9d, 0x55, 0xef, 0x21, 0x03, 0xd4, 0xf3, 0xd3, 0xaa, 0x22, 0x3c, 0xdd, 0x93, 0x8b, 0x7e,
	0x55, 0x45, 0x2b, 0xa0, 0xbf, 0xc3, 0xf8, 0x04, 0x57, 0x35, 0xf3, 0x97, 0x02, 0x65, 0xf9, 0x1a,
	0x11, 0x35, 0x65, 0xe8, 0x05, 0xe8, 0x23, 0x61, 0x4a, 0xf9, 0xca, 0xad, 0x87, 0x39, 0x49, 0x71,
	0x44, 0xa1, 0x23, 0xd8, 0x76, 0x09, 0xe3, 0x4b, 0x44, 0x52, 0xa5, 0x48, 0x8f, 0x04, 0x95, 0xa7,
	0xd0, 0x7b, 0xa8, 0x8c, 0x1d, 0xc6, 0xfd, 0xcb, 0x90, 0x4c, 0x06, 0x36, 0xe1, 0x44, 0x0e, 0xa4,
	0xdc, 0xda, 0xc9, 0x26, 0x3f, 0x4e, 0xb8, 0x2e, 0xe1, 0x04, 0xaf, 0x8d, 0xe7, 0x4d, 0xf3, 0x0b,
	0x94, 0x3e, 0x3a, 0x8c, 0xcb, 0x8f, 0x72, 0x03, 0xf4, 0x91, 0x3f, 0xf5, 0xb8, 0xec, 0x43, 0xc7,
	0x91, 0x81, 0x10, 0x14, 0x02, 0x72, 0x49, 0x65, 0x51, 0x3a, 0x96, 0xcf, 0xe8, 0x29, 0x54, 0x1c,
	0x6f, 0xe4, 0x4e, 0x6d, 0x3a, 0x60, 0x52, 0x03, 0x99, 0xbd, 0x84, 0xd7, 0x62, 0x6f, 0x24, 0x8c,
	0xf9, 0x5b, 0x01, 0x90, 0xad, 0xb3, 0x53, 0x11, 0xf5, 0x0a, 0x8a, 0x21, 0x65, 0x53, 0x97, 0xb3,
	0x9a, 0xd2, 0xd0, 0x9a, 0xe5, 0xd6, 0x93, 0x1c, 0xa5, 0x24, 0x6e, 0x61, 0xc9, 0xe2, 0x24, 0xa6,
	0xfe, 0x0d, 0x8c, 0xc8, 0xf5, 0xbf, 0x82, 0x67, 0xb5, 0x52, 0xef, 0xa4, 0xd5, 0x4f, 0x15, 0xd6,
	0x52, 0x00, 0x7a, 0x06, 0xeb, 0x13, 0x32, 0x4b, 0xcd, 0x4e, 0x94, 0xa4, 0xe0, 0xca, 0x84, 0xcc,
	0xe6, 0xc7, 0x25, 0x40, 0xc7, 0xcb, 0x0c, 0x59, 0x80, 0x8e, 0x37, 0x0f, 0x3e, 0x87, 0xea, 0x84,
	0x92, 0x34, 0xa9, 0x49, 0x72, 0x5d, 0xf8, 0xe7, 0xd1, 0x07, 0x60, 0x30, 0x6e, 0xdb, 0xf4, 0x5a,
	0x5e, 0x2a, 0x05, 0xc7, 0x96, 0xc8, 0x15, 0xb4, 0xf7, 0x52, 0x6f, 0xd0, 0xa3, 0x5c, 0x41, 0x7b,
	0x6f, 0xa1, 0xa8, 0xe0, 0xa0, 0x9d, 0x02, 0x8d, 0x18, 0x3c, 0x68, 0x2f, 0x82, 0x9d, 0x34, 0x58,
	0x8c, 0xc1, 0x4e, 0x16, 0xec, 0xa4, 0xc0, 0x52, 0x02, 0x76, 0xe6, 0x40, 0x53, 0x07, 0xad, 0xef,
	0xb8, 0xad, 0xef, 0x1a, 0x18, 0xd1, 0xc4, 0xd1, 0x1b, 0x30, 0xa2, 0xfd, 0x88, 0xb6, 0xfe, 0x31,
	0xce, 0xbf, 0x9b, 0xb3, 0x9e, 0x37, 0x6c, 0xf4, 0x12, 0x8c, 0x2e, 0x75, 0x29, 0xa7, 0x68, 0x33,
	0x07, 0xe9, 0x75, 0xf3, 0xa3, 0x3b, 0xa0, 0x7d, 0xa0, 0xfc, 0x4e, 0xa1, 0x6f, 0xa1, 0x20, 0xae,
	0x10, 0xaa, 0x67, 0x81, 0xe4, 0x6a, 0xd5, 0xb7, 0x96, 0x7d, 0xe9, 0xe8, 0x00, 0xb4, 0xd3, 0x29,
	0x47, 0x79, 0x19, 0xf2, 0x53, 0x1f, 0x82, 0x11, 0xef, 0xa0, 0x25, 0x85, 0xe7, 0x2d, 0xc1, 0x28,
	0xb2, 0xd5, 0x83, 0xe2, 0x45, 0xb4, 0x8d, 0xd1, 0x6b, 0x30, 0x8e, 0x89, 0x67, 0xbb, 0x14, 0x2d,
	0x8f, 0xa9, 0xdf, 0xcf, 0x1e, 0xf7, 0x1d, 0xf7, 0xb0, 0xf4, 0xd9, 0x88, 0xac, 0xa1, 0x21, 0x7f,
	0x83, 0xfb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xb2, 0x17, 0x25, 0x16, 0x07, 0x00, 0x00,
}
