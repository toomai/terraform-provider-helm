// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/configured_target.proto

package resultstore // import "google.golang.org/genproto/googleapis/devtools/resultstore/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Each ConfiguredTarget represents data for a given configuration of a given
// target in a given Invocation.
// Every ConfiguredTarget should have at least one Action as a child resource
// before the invocation is finalized. Refer to the Action's documentation for
// more info on this.
type ConfiguredTarget struct {
	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${TARGET_ID}/configuredTargets/${CONFIG_ID}
	// where ${CONFIG_ID} must match the ID of an existing Configuration under
	// this Invocation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the ConfiguredTarget. They must
	// match the resource name after proper encoding.
	Id *ConfiguredTarget_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status for this configuration of this target. If testing
	// was not requested, set this to the build status (e.g. BUILT or
	// FAILED_TO_BUILD).
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// Captures the start time and duration of this configured target.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Test specific attributes for this ConfiguredTarget.
	TestAttributes *ConfiguredTestAttributes `protobuf:"bytes,6,opt,name=test_attributes,json=testAttributes,proto3" json:"test_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for configured target level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	// Summary of aggregate coverage across all test actions for this target in
	// this configuration. The server populates this for you in the
	// post-processing phase.
	CoverageSummaries    []*LanguageCoverageSummary `protobuf:"bytes,9,rep,name=coverage_summaries,json=coverageSummaries,proto3" json:"coverage_summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ConfiguredTarget) Reset()         { *m = ConfiguredTarget{} }
func (m *ConfiguredTarget) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTarget) ProtoMessage()    {}
func (*ConfiguredTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_configured_target_48a4f73ce345d3e8, []int{0}
}
func (m *ConfiguredTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTarget.Unmarshal(m, b)
}
func (m *ConfiguredTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTarget.Marshal(b, m, deterministic)
}
func (dst *ConfiguredTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTarget.Merge(dst, src)
}
func (m *ConfiguredTarget) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTarget.Size(m)
}
func (m *ConfiguredTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTarget.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTarget proto.InternalMessageInfo

func (m *ConfiguredTarget) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfiguredTarget) GetId() *ConfiguredTarget_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfiguredTarget) GetStatusAttributes() *StatusAttributes {
	if m != nil {
		return m.StatusAttributes
	}
	return nil
}

func (m *ConfiguredTarget) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *ConfiguredTarget) GetTestAttributes() *ConfiguredTestAttributes {
	if m != nil {
		return m.TestAttributes
	}
	return nil
}

func (m *ConfiguredTarget) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ConfiguredTarget) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *ConfiguredTarget) GetCoverageSummaries() []*LanguageCoverageSummary {
	if m != nil {
		return m.CoverageSummaries
	}
	return nil
}

// The resource ID components that identify the ConfiguredTarget.
type ConfiguredTarget_Id struct {
	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Target ID.
	TargetId string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// The Configuration ID.
	ConfigurationId      string   `protobuf:"bytes,3,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfiguredTarget_Id) Reset()         { *m = ConfiguredTarget_Id{} }
func (m *ConfiguredTarget_Id) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTarget_Id) ProtoMessage()    {}
func (*ConfiguredTarget_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_configured_target_48a4f73ce345d3e8, []int{0, 0}
}
func (m *ConfiguredTarget_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTarget_Id.Unmarshal(m, b)
}
func (m *ConfiguredTarget_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTarget_Id.Marshal(b, m, deterministic)
}
func (dst *ConfiguredTarget_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTarget_Id.Merge(dst, src)
}
func (m *ConfiguredTarget_Id) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTarget_Id.Size(m)
}
func (m *ConfiguredTarget_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTarget_Id.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTarget_Id proto.InternalMessageInfo

func (m *ConfiguredTarget_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *ConfiguredTarget_Id) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *ConfiguredTarget_Id) GetConfigurationId() string {
	if m != nil {
		return m.ConfigurationId
	}
	return ""
}

// Attributes that apply only to test actions under this configured target.
type ConfiguredTestAttributes struct {
	// Total number of test runs. For example, in bazel this is specified with
	// --runs_per_test. Zero if runs_per_test is not used.
	TotalRunCount int32 `protobuf:"varint,2,opt,name=total_run_count,json=totalRunCount,proto3" json:"total_run_count,omitempty"`
	// Total number of test shards. Zero if shard count was not specified.
	TotalShardCount int32 `protobuf:"varint,3,opt,name=total_shard_count,json=totalShardCount,proto3" json:"total_shard_count,omitempty"`
	// How long test is allowed to run.
	TimeoutDuration      *duration.Duration `protobuf:"bytes,5,opt,name=timeout_duration,json=timeoutDuration,proto3" json:"timeout_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ConfiguredTestAttributes) Reset()         { *m = ConfiguredTestAttributes{} }
func (m *ConfiguredTestAttributes) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTestAttributes) ProtoMessage()    {}
func (*ConfiguredTestAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_configured_target_48a4f73ce345d3e8, []int{1}
}
func (m *ConfiguredTestAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTestAttributes.Unmarshal(m, b)
}
func (m *ConfiguredTestAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTestAttributes.Marshal(b, m, deterministic)
}
func (dst *ConfiguredTestAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTestAttributes.Merge(dst, src)
}
func (m *ConfiguredTestAttributes) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTestAttributes.Size(m)
}
func (m *ConfiguredTestAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTestAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTestAttributes proto.InternalMessageInfo

func (m *ConfiguredTestAttributes) GetTotalRunCount() int32 {
	if m != nil {
		return m.TotalRunCount
	}
	return 0
}

func (m *ConfiguredTestAttributes) GetTotalShardCount() int32 {
	if m != nil {
		return m.TotalShardCount
	}
	return 0
}

func (m *ConfiguredTestAttributes) GetTimeoutDuration() *duration.Duration {
	if m != nil {
		return m.TimeoutDuration
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfiguredTarget)(nil), "google.devtools.resultstore.v2.ConfiguredTarget")
	proto.RegisterType((*ConfiguredTarget_Id)(nil), "google.devtools.resultstore.v2.ConfiguredTarget.Id")
	proto.RegisterType((*ConfiguredTestAttributes)(nil), "google.devtools.resultstore.v2.ConfiguredTestAttributes")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/configured_target.proto", fileDescriptor_configured_target_48a4f73ce345d3e8)
}

var fileDescriptor_configured_target_48a4f73ce345d3e8 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0x69, 0xbb, 0x5d, 0xb7, 0x59, 0xd7, 0xb6, 0xb9, 0x1a, 0x2b, 0x48, 0xa9, 0xb2, 0x74,
	0x15, 0x66, 0xa4, 0x8b, 0x7f, 0x50, 0x10, 0xb4, 0x8b, 0x38, 0xe0, 0xc5, 0x32, 0xdd, 0x2b, 0x41,
	0x86, 0xb4, 0x49, 0x63, 0x60, 0x26, 0xe9, 0x26, 0x27, 0x85, 0x7d, 0x2f, 0x9f, 0xc3, 0x67, 0x92,
	0x26, 0x99, 0xda, 0x16, 0xd6, 0xc1, 0xbb, 0xe6, 0xeb, 0xf7, 0xfd, 0x72, 0x92, 0x9c, 0x33, 0xe8,
	0x0d, 0x57, 0x8a, 0x17, 0x2c, 0xa1, 0x6c, 0x0d, 0x4a, 0x15, 0x26, 0xd1, 0xcc, 0xd8, 0x02, 0x0c,
	0x28, 0xcd, 0x92, 0xf5, 0x24, 0x59, 0x28, 0xb9, 0x14, 0xdc, 0x6a, 0x46, 0x73, 0x20, 0x9a, 0x33,
	0x88, 0x57, 0x5a, 0x81, 0xc2, 0x4f, 0x7d, 0x2e, 0xae, 0x72, 0xf1, 0x4e, 0x2e, 0x5e, 0x4f, 0x06,
	0x2f, 0x6b, 0xb9, 0x65, 0xa9, 0xa4, 0x87, 0x0d, 0x5e, 0xd7, 0x9a, 0xd7, 0x4c, 0x13, 0xce, 0x72,
	0x63, 0xcb, 0x92, 0xe8, 0xbb, 0x10, 0xbb, 0xa8, 0x89, 0x2d, 0x45, 0xc1, 0x82, 0x35, 0x94, 0x9b,
	0xb8, 0xd5, 0xdc, 0x2e, 0x13, 0x6a, 0x35, 0x01, 0x51, 0x55, 0x30, 0xfa, 0xdd, 0x46, 0xbd, 0xe9,
	0xf6, 0xa8, 0x37, 0xee, 0xa4, 0x18, 0xa3, 0x23, 0x49, 0x4a, 0x16, 0x35, 0x86, 0x8d, 0x71, 0x27,
	0x73, 0xbf, 0xf1, 0x14, 0x35, 0x05, 0x8d, 0x9a, 0xc3, 0xc6, 0xf8, 0x74, 0x72, 0x19, 0xff, 0xfb,
	0x12, 0xe2, 0x43, 0x62, 0x9c, 0xd2, 0xac, 0x29, 0x28, 0xfe, 0x81, 0xfa, 0x06, 0x08, 0x58, 0x93,
	0x13, 0x00, 0x2d, 0xe6, 0x16, 0x98, 0x89, 0x5a, 0x8e, 0xf9, 0xaa, 0x8e, 0x39, 0x73, 0xc1, 0x4f,
	0xdb, 0x5c, 0xd6, 0x33, 0x07, 0x0a, 0xfe, 0x88, 0x8e, 0x41, 0x94, 0x42, 0xf2, 0xe8, 0xc8, 0x31,
	0xcf, 0xeb, 0x98, 0x37, 0xce, 0x9d, 0x85, 0x14, 0x26, 0xa8, 0x0b, 0xcc, 0xc0, 0x6e, 0x71, 0xc7,
	0x0e, 0xf4, 0xee, 0x3f, 0x0e, 0xcc, 0x0c, 0xec, 0x14, 0xf9, 0x08, 0xf6, 0xd6, 0xf8, 0x2b, 0x42,
	0x2b, 0xad, 0x56, 0x4c, 0x83, 0x60, 0x26, 0x7a, 0x30, 0x6c, 0x8d, 0x4f, 0x27, 0xe3, 0x3a, 0xfa,
	0xb5, 0x4f, 0xdc, 0x65, 0x3b, 0x59, 0xfc, 0x1e, 0xb5, 0x37, 0xef, 0x6c, 0xa2, 0x13, 0x07, 0x79,
	0x5e, 0x07, 0xf9, 0x22, 0x0a, 0x96, 0xf9, 0x08, 0x5e, 0x22, 0x7c, 0xd0, 0x5a, 0x9b, 0x6a, 0x3a,
	0x0e, 0xf4, 0xb6, 0x0e, 0xf4, 0x8d, 0x48, 0x6e, 0x09, 0x67, 0xd3, 0x40, 0x98, 0xf9, 0xde, 0xcc,
	0xfa, 0x8b, 0x3d, 0x41, 0x30, 0x33, 0xb8, 0x45, 0xcd, 0x94, 0xe2, 0x67, 0xe8, 0x4c, 0xc8, 0xb5,
	0x5a, 0xb8, 0xbe, 0xcb, 0x05, 0x0d, 0x7d, 0xf5, 0xf0, 0xaf, 0x98, 0x52, 0xfc, 0x04, 0x75, 0xfc,
	0x9c, 0xe5, 0xa1, 0xcd, 0x3a, 0xd9, 0x89, 0x17, 0x52, 0x8a, 0x2f, 0x50, 0xaf, 0x9a, 0xc7, 0x2d,
	0xa4, 0xe5, 0x3c, 0xdd, 0x3d, 0x3d, 0xa5, 0xa3, 0x5f, 0x0d, 0x14, 0xdd, 0xf7, 0x1a, 0xf8, 0x1c,
	0x75, 0x41, 0x01, 0x29, 0x72, 0x6d, 0x65, 0xbe, 0x50, 0x56, 0x82, 0xdb, 0xaa, 0x9d, 0x9d, 0x39,
	0x39, 0xb3, 0x72, 0xba, 0x11, 0xf1, 0x0b, 0xd4, 0xf7, 0x3e, 0xf3, 0x93, 0x68, 0x1a, 0x9c, 0x2d,
	0xe7, 0xf4, 0x80, 0xd9, 0x46, 0xf7, 0xde, 0x2b, 0xd4, 0x03, 0x51, 0x32, 0x65, 0x21, 0xaf, 0x66,
	0x2b, 0x6a, 0xbb, 0xae, 0x79, 0x5c, 0xdd, 0x64, 0x35, 0x7c, 0xf1, 0x55, 0x30, 0x64, 0xdd, 0x10,
	0xa9, 0x84, 0xcf, 0xb7, 0x68, 0xb4, 0x50, 0x65, 0xcd, 0xd5, 0x5f, 0x37, 0xbe, 0xa7, 0xc1, 0xc1,
	0x55, 0x41, 0x24, 0x8f, 0x95, 0xe6, 0x09, 0x67, 0xd2, 0x6d, 0x90, 0xf8, 0xbf, 0xc8, 0x4a, 0x98,
	0xfb, 0xbe, 0x0c, 0x1f, 0x76, 0x96, 0xf3, 0x63, 0x97, 0xba, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0xb1, 0xb2, 0xc2, 0x0a, 0x05, 0x00, 0x00,
}