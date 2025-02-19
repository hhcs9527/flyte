// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/spark.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SparkApplication_Type int32

const (
	SparkApplication_PYTHON SparkApplication_Type = 0
	SparkApplication_JAVA   SparkApplication_Type = 1
	SparkApplication_SCALA  SparkApplication_Type = 2
	SparkApplication_R      SparkApplication_Type = 3
)

var SparkApplication_Type_name = map[int32]string{
	0: "PYTHON",
	1: "JAVA",
	2: "SCALA",
	3: "R",
}

var SparkApplication_Type_value = map[string]int32{
	"PYTHON": 0,
	"JAVA":   1,
	"SCALA":  2,
	"R":      3,
}

func (x SparkApplication_Type) String() string {
	return proto.EnumName(SparkApplication_Type_name, int32(x))
}

func (SparkApplication_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{0, 0}
}

type SparkApplication struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SparkApplication) Reset()         { *m = SparkApplication{} }
func (m *SparkApplication) String() string { return proto.CompactTextString(m) }
func (*SparkApplication) ProtoMessage()    {}
func (*SparkApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{0}
}

func (m *SparkApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparkApplication.Unmarshal(m, b)
}
func (m *SparkApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparkApplication.Marshal(b, m, deterministic)
}
func (m *SparkApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparkApplication.Merge(m, src)
}
func (m *SparkApplication) XXX_Size() int {
	return xxx_messageInfo_SparkApplication.Size(m)
}
func (m *SparkApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_SparkApplication.DiscardUnknown(m)
}

var xxx_messageInfo_SparkApplication proto.InternalMessageInfo

// Custom Proto for Spark Plugin.
type SparkJob struct {
	ApplicationType     SparkApplication_Type `protobuf:"varint,1,opt,name=applicationType,proto3,enum=flyteidl.plugins.SparkApplication_Type" json:"applicationType,omitempty"`
	MainApplicationFile string                `protobuf:"bytes,2,opt,name=mainApplicationFile,proto3" json:"mainApplicationFile,omitempty"`
	MainClass           string                `protobuf:"bytes,3,opt,name=mainClass,proto3" json:"mainClass,omitempty"`
	SparkConf           map[string]string     `protobuf:"bytes,4,rep,name=sparkConf,proto3" json:"sparkConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HadoopConf          map[string]string     `protobuf:"bytes,5,rep,name=hadoopConf,proto3" json:"hadoopConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExecutorPath        string                `protobuf:"bytes,6,opt,name=executorPath,proto3" json:"executorPath,omitempty"`
	// Databricks job configuration.
	// Config structure can be found here. https://docs.databricks.com/dev-tools/api/2.0/jobs.html#request-structure.
	DatabricksConf *_struct.Struct `protobuf:"bytes,7,opt,name=databricksConf,proto3" json:"databricksConf,omitempty"`
	// Databricks access token. https://docs.databricks.com/dev-tools/api/latest/authentication.html
	// This token can be set in either flytepropeller or flytekit.
	DatabricksToken string `protobuf:"bytes,8,opt,name=databricksToken,proto3" json:"databricksToken,omitempty"`
	// Domain name of your deployment. Use the form <account>.cloud.databricks.com.
	// This instance name can be set in either flytepropeller or flytekit.
	DatabricksInstance   string   `protobuf:"bytes,9,opt,name=databricksInstance,proto3" json:"databricksInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SparkJob) Reset()         { *m = SparkJob{} }
func (m *SparkJob) String() string { return proto.CompactTextString(m) }
func (*SparkJob) ProtoMessage()    {}
func (*SparkJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{1}
}

func (m *SparkJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparkJob.Unmarshal(m, b)
}
func (m *SparkJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparkJob.Marshal(b, m, deterministic)
}
func (m *SparkJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparkJob.Merge(m, src)
}
func (m *SparkJob) XXX_Size() int {
	return xxx_messageInfo_SparkJob.Size(m)
}
func (m *SparkJob) XXX_DiscardUnknown() {
	xxx_messageInfo_SparkJob.DiscardUnknown(m)
}

var xxx_messageInfo_SparkJob proto.InternalMessageInfo

func (m *SparkJob) GetApplicationType() SparkApplication_Type {
	if m != nil {
		return m.ApplicationType
	}
	return SparkApplication_PYTHON
}

func (m *SparkJob) GetMainApplicationFile() string {
	if m != nil {
		return m.MainApplicationFile
	}
	return ""
}

func (m *SparkJob) GetMainClass() string {
	if m != nil {
		return m.MainClass
	}
	return ""
}

func (m *SparkJob) GetSparkConf() map[string]string {
	if m != nil {
		return m.SparkConf
	}
	return nil
}

func (m *SparkJob) GetHadoopConf() map[string]string {
	if m != nil {
		return m.HadoopConf
	}
	return nil
}

func (m *SparkJob) GetExecutorPath() string {
	if m != nil {
		return m.ExecutorPath
	}
	return ""
}

func (m *SparkJob) GetDatabricksConf() *_struct.Struct {
	if m != nil {
		return m.DatabricksConf
	}
	return nil
}

func (m *SparkJob) GetDatabricksToken() string {
	if m != nil {
		return m.DatabricksToken
	}
	return ""
}

func (m *SparkJob) GetDatabricksInstance() string {
	if m != nil {
		return m.DatabricksInstance
	}
	return ""
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.SparkApplication_Type", SparkApplication_Type_name, SparkApplication_Type_value)
	proto.RegisterType((*SparkApplication)(nil), "flyteidl.plugins.SparkApplication")
	proto.RegisterType((*SparkJob)(nil), "flyteidl.plugins.SparkJob")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.plugins.SparkJob.HadoopConfEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.plugins.SparkJob.SparkConfEntry")
}

func init() { proto.RegisterFile("flyteidl/plugins/spark.proto", fileDescriptor_ca8a069b9820144a) }

var fileDescriptor_ca8a069b9820144a = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0xfd, 0xb2, 0xb9, 0x2b, 0x6d, 0xb8, 0x0a, 0x86, 0xd2, 0x87, 0xd2, 0x17, 0xa3, 0xe0,
	0x44, 0xea, 0x83, 0x1f, 0x28, 0x92, 0x2d, 0xea, 0x5a, 0x44, 0xd7, 0xb4, 0x08, 0xfa, 0x36, 0x49,
	0xa7, 0x69, 0xe8, 0xec, 0xcc, 0x90, 0x4c, 0xc4, 0xfc, 0x79, 0x91, 0x4c, 0xec, 0x66, 0x37, 0xac,
	0x82, 0x6f, 0xc9, 0xb9, 0xe7, 0x9c, 0x7b, 0x39, 0x87, 0x81, 0xe9, 0x8e, 0x97, 0x9a, 0xa5, 0x5b,
	0xee, 0x2b, 0x5e, 0x24, 0xa9, 0xc8, 0xfd, 0x5c, 0xd1, 0xec, 0x40, 0x54, 0x26, 0xb5, 0x44, 0xe7,
	0x38, 0x25, 0x7f, 0xa6, 0x93, 0x69, 0x22, 0x65, 0xc2, 0x99, 0x6f, 0xe6, 0x51, 0xb1, 0xf3, 0x73,
	0x9d, 0x15, 0xb1, 0xae, 0xf9, 0xf3, 0x53, 0x70, 0xd6, 0x95, 0x3c, 0x50, 0x8a, 0xa7, 0x31, 0xd5,
	0xa9, 0x14, 0x73, 0x02, 0xbd, 0x4d, 0xa9, 0x18, 0x02, 0x0c, 0xce, 0xbf, 0x6d, 0xce, 0x3e, 0x7f,
	0x72, 0x6e, 0xe1, 0x10, 0x7a, 0xab, 0xe0, 0x6b, 0xe0, 0x58, 0x68, 0x43, 0x7f, 0xbd, 0x0c, 0x3e,
	0x06, 0x4e, 0x07, 0xfb, 0x60, 0x85, 0x4e, 0x77, 0xfe, 0xab, 0x07, 0x43, 0x63, 0xb2, 0x92, 0x11,
	0x7e, 0x81, 0x31, 0x6d, 0xbc, 0x2a, 0x1f, 0xd7, 0x9a, 0x59, 0xde, 0x68, 0xf1, 0x80, 0xb4, 0x4f,
	0x23, 0xed, 0xcd, 0xa4, 0xa2, 0x87, 0x6d, 0x3d, 0x3e, 0x81, 0xbb, 0x17, 0x34, 0x15, 0x57, 0x88,
	0xef, 0x52, 0xce, 0xdc, 0xce, 0xcc, 0xf2, 0xec, 0xf0, 0xa6, 0x11, 0x4e, 0xc1, 0xae, 0xe0, 0x25,
	0xa7, 0x79, 0xee, 0x76, 0x0d, 0xaf, 0x01, 0xf0, 0x3d, 0xd8, 0x26, 0xb2, 0xa5, 0x14, 0x3b, 0xb7,
	0x37, 0xeb, 0x7a, 0x27, 0x8b, 0x87, 0x7f, 0x39, 0x6e, 0x25, 0xa3, 0xfa, 0xa3, 0xe2, 0xbe, 0x15,
	0x3a, 0x2b, 0xc3, 0x46, 0x8b, 0x2b, 0x80, 0x3d, 0xdd, 0x4a, 0xa9, 0x8c, 0x53, 0xdf, 0x38, 0x3d,
	0xfa, 0x87, 0xd3, 0xd9, 0x25, 0xb9, 0xb6, 0xba, 0xa2, 0xc6, 0x39, 0xdc, 0x61, 0x3f, 0x59, 0x5c,
	0x68, 0x99, 0x9d, 0x53, 0xbd, 0x77, 0x07, 0xe6, 0xea, 0x6b, 0x18, 0xbe, 0x81, 0xd1, 0x96, 0x6a,
	0x1a, 0x65, 0x69, 0x7c, 0xc8, 0xcd, 0xce, 0xdb, 0x33, 0xcb, 0x3b, 0x59, 0xdc, 0x27, 0x75, 0xc7,
	0xe4, 0xd8, 0x31, 0x59, 0x9b, 0x8e, 0xc3, 0x16, 0x1d, 0x3d, 0x18, 0x37, 0xc8, 0x46, 0x1e, 0x98,
	0x70, 0x87, 0x66, 0x4f, 0x1b, 0x46, 0x02, 0xd8, 0x40, 0x1f, 0x44, 0xae, 0xa9, 0x88, 0x99, 0x6b,
	0x1b, 0xf2, 0x0d, 0x93, 0xc9, 0x2b, 0x18, 0x5d, 0xcf, 0x09, 0x1d, 0xe8, 0x1e, 0x58, 0x69, 0xca,
	0xb7, 0xc3, 0xea, 0x13, 0xef, 0x41, 0xff, 0x07, 0xe5, 0xc5, 0xb1, 0xb9, 0xfa, 0xe7, 0x65, 0xe7,
	0xb9, 0x35, 0x79, 0x0d, 0xe3, 0x56, 0x36, 0xff, 0x23, 0x3f, 0x7d, 0xf1, 0xfd, 0x59, 0x92, 0xea,
	0x7d, 0x11, 0x91, 0x58, 0x5e, 0xf8, 0x26, 0x7f, 0x99, 0x25, 0xfe, 0xe5, 0x43, 0x49, 0x98, 0xf0,
	0x55, 0xf4, 0x38, 0x91, 0x7e, 0xfb, 0xed, 0x44, 0x03, 0x13, 0xd9, 0xd3, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x02, 0x51, 0x7a, 0xdc, 0x56, 0x03, 0x00, 0x00,
}
