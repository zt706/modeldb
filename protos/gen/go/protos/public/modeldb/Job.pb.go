// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/public/modeldb/Job.proto

package modeldb

import (
	context "context"
	fmt "fmt"
	common "github.com/VertaAI/modeldb/protos/gen/go/protos/public/common"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type JobStatusEnum_JobStatus int32

const (
	JobStatusEnum_NOT_STARTED JobStatusEnum_JobStatus = 0
	JobStatusEnum_IN_PROGRESS JobStatusEnum_JobStatus = 1
	JobStatusEnum_COMPLETED   JobStatusEnum_JobStatus = 2
)

var JobStatusEnum_JobStatus_name = map[int32]string{
	0: "NOT_STARTED",
	1: "IN_PROGRESS",
	2: "COMPLETED",
}

var JobStatusEnum_JobStatus_value = map[string]int32{
	"NOT_STARTED": 0,
	"IN_PROGRESS": 1,
	"COMPLETED":   2,
}

func (x JobStatusEnum_JobStatus) String() string {
	return proto.EnumName(JobStatusEnum_JobStatus_name, int32(x))
}

func (JobStatusEnum_JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{0, 0}
}

type JobTypeEnum_JobType int32

const (
	JobTypeEnum_KUBERNETES_JOB JobTypeEnum_JobType = 0
)

var JobTypeEnum_JobType_name = map[int32]string{
	0: "KUBERNETES_JOB",
}

var JobTypeEnum_JobType_value = map[string]int32{
	"KUBERNETES_JOB": 0,
}

func (x JobTypeEnum_JobType) String() string {
	return proto.EnumName(JobTypeEnum_JobType_name, int32(x))
}

func (JobTypeEnum_JobType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{1, 0}
}

type JobStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobStatusEnum) Reset()         { *m = JobStatusEnum{} }
func (m *JobStatusEnum) String() string { return proto.CompactTextString(m) }
func (*JobStatusEnum) ProtoMessage()    {}
func (*JobStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{0}
}

func (m *JobStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatusEnum.Unmarshal(m, b)
}
func (m *JobStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatusEnum.Marshal(b, m, deterministic)
}
func (m *JobStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatusEnum.Merge(m, src)
}
func (m *JobStatusEnum) XXX_Size() int {
	return xxx_messageInfo_JobStatusEnum.Size(m)
}
func (m *JobStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatusEnum proto.InternalMessageInfo

type JobTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobTypeEnum) Reset()         { *m = JobTypeEnum{} }
func (m *JobTypeEnum) String() string { return proto.CompactTextString(m) }
func (*JobTypeEnum) ProtoMessage()    {}
func (*JobTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{1}
}

func (m *JobTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobTypeEnum.Unmarshal(m, b)
}
func (m *JobTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobTypeEnum.Marshal(b, m, deterministic)
}
func (m *JobTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobTypeEnum.Merge(m, src)
}
func (m *JobTypeEnum) XXX_Size() int {
	return xxx_messageInfo_JobTypeEnum.Size(m)
}
func (m *JobTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_JobTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_JobTypeEnum proto.InternalMessageInfo

type Job struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	StartTime            string                  `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string                  `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Metadata             []*common.KeyValue      `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty"`
	JobStatus            JobStatusEnum_JobStatus `protobuf:"varint,6,opt,name=job_status,json=jobStatus,proto3,enum=ai.verta.modeldb.JobStatusEnum_JobStatus" json:"job_status,omitempty"`
	JobType              JobTypeEnum_JobType     `protobuf:"varint,7,opt,name=job_type,json=jobType,proto3,enum=ai.verta.modeldb.JobTypeEnum_JobType" json:"job_type,omitempty"`
	Owner                string                  `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{2}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Job) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Job) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Job) GetMetadata() []*common.KeyValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Job) GetJobStatus() JobStatusEnum_JobStatus {
	if m != nil {
		return m.JobStatus
	}
	return JobStatusEnum_NOT_STARTED
}

func (m *Job) GetJobType() JobTypeEnum_JobType {
	if m != nil {
		return m.JobType
	}
	return JobTypeEnum_KUBERNETES_JOB
}

func (m *Job) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type CreateJob struct {
	Description          string                  `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	StartTime            string                  `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string                  `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Metadata             []*common.KeyValue      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty"`
	JobStatus            JobStatusEnum_JobStatus `protobuf:"varint,5,opt,name=job_status,json=jobStatus,proto3,enum=ai.verta.modeldb.JobStatusEnum_JobStatus" json:"job_status,omitempty"`
	JobType              JobTypeEnum_JobType     `protobuf:"varint,6,opt,name=job_type,json=jobType,proto3,enum=ai.verta.modeldb.JobTypeEnum_JobType" json:"job_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateJob) Reset()         { *m = CreateJob{} }
func (m *CreateJob) String() string { return proto.CompactTextString(m) }
func (*CreateJob) ProtoMessage()    {}
func (*CreateJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{3}
}

func (m *CreateJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJob.Unmarshal(m, b)
}
func (m *CreateJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJob.Marshal(b, m, deterministic)
}
func (m *CreateJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJob.Merge(m, src)
}
func (m *CreateJob) XXX_Size() int {
	return xxx_messageInfo_CreateJob.Size(m)
}
func (m *CreateJob) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJob.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJob proto.InternalMessageInfo

func (m *CreateJob) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateJob) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *CreateJob) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *CreateJob) GetMetadata() []*common.KeyValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateJob) GetJobStatus() JobStatusEnum_JobStatus {
	if m != nil {
		return m.JobStatus
	}
	return JobStatusEnum_NOT_STARTED
}

func (m *CreateJob) GetJobType() JobTypeEnum_JobType {
	if m != nil {
		return m.JobType
	}
	return JobTypeEnum_KUBERNETES_JOB
}

type CreateJob_Response struct {
	Job                  *Job     `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateJob_Response) Reset()         { *m = CreateJob_Response{} }
func (m *CreateJob_Response) String() string { return proto.CompactTextString(m) }
func (*CreateJob_Response) ProtoMessage()    {}
func (*CreateJob_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{3, 0}
}

func (m *CreateJob_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJob_Response.Unmarshal(m, b)
}
func (m *CreateJob_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJob_Response.Marshal(b, m, deterministic)
}
func (m *CreateJob_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJob_Response.Merge(m, src)
}
func (m *CreateJob_Response) XXX_Size() int {
	return xxx_messageInfo_CreateJob_Response.Size(m)
}
func (m *CreateJob_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJob_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJob_Response proto.InternalMessageInfo

func (m *CreateJob_Response) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type UpdateJob struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EndTime              string                  `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	JobStatus            JobStatusEnum_JobStatus `protobuf:"varint,2,opt,name=job_status,json=jobStatus,proto3,enum=ai.verta.modeldb.JobStatusEnum_JobStatus" json:"job_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateJob) Reset()         { *m = UpdateJob{} }
func (m *UpdateJob) String() string { return proto.CompactTextString(m) }
func (*UpdateJob) ProtoMessage()    {}
func (*UpdateJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{4}
}

func (m *UpdateJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJob.Unmarshal(m, b)
}
func (m *UpdateJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJob.Marshal(b, m, deterministic)
}
func (m *UpdateJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJob.Merge(m, src)
}
func (m *UpdateJob) XXX_Size() int {
	return xxx_messageInfo_UpdateJob.Size(m)
}
func (m *UpdateJob) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJob.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJob proto.InternalMessageInfo

func (m *UpdateJob) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateJob) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *UpdateJob) GetJobStatus() JobStatusEnum_JobStatus {
	if m != nil {
		return m.JobStatus
	}
	return JobStatusEnum_NOT_STARTED
}

type UpdateJob_Response struct {
	Job                  *Job     `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateJob_Response) Reset()         { *m = UpdateJob_Response{} }
func (m *UpdateJob_Response) String() string { return proto.CompactTextString(m) }
func (*UpdateJob_Response) ProtoMessage()    {}
func (*UpdateJob_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{4, 0}
}

func (m *UpdateJob_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJob_Response.Unmarshal(m, b)
}
func (m *UpdateJob_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJob_Response.Marshal(b, m, deterministic)
}
func (m *UpdateJob_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJob_Response.Merge(m, src)
}
func (m *UpdateJob_Response) XXX_Size() int {
	return xxx_messageInfo_UpdateJob_Response.Size(m)
}
func (m *UpdateJob_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJob_Response.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJob_Response proto.InternalMessageInfo

func (m *UpdateJob_Response) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type DeleteJob struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteJob) Reset()         { *m = DeleteJob{} }
func (m *DeleteJob) String() string { return proto.CompactTextString(m) }
func (*DeleteJob) ProtoMessage()    {}
func (*DeleteJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{5}
}

func (m *DeleteJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteJob.Unmarshal(m, b)
}
func (m *DeleteJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteJob.Marshal(b, m, deterministic)
}
func (m *DeleteJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteJob.Merge(m, src)
}
func (m *DeleteJob) XXX_Size() int {
	return xxx_messageInfo_DeleteJob.Size(m)
}
func (m *DeleteJob) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteJob.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteJob proto.InternalMessageInfo

func (m *DeleteJob) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteJob_Response struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteJob_Response) Reset()         { *m = DeleteJob_Response{} }
func (m *DeleteJob_Response) String() string { return proto.CompactTextString(m) }
func (*DeleteJob_Response) ProtoMessage()    {}
func (*DeleteJob_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{5, 0}
}

func (m *DeleteJob_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteJob_Response.Unmarshal(m, b)
}
func (m *DeleteJob_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteJob_Response.Marshal(b, m, deterministic)
}
func (m *DeleteJob_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteJob_Response.Merge(m, src)
}
func (m *DeleteJob_Response) XXX_Size() int {
	return xxx_messageInfo_DeleteJob_Response.Size(m)
}
func (m *DeleteJob_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteJob_Response.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteJob_Response proto.InternalMessageInfo

func (m *DeleteJob_Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type GetJob struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJob) Reset()         { *m = GetJob{} }
func (m *GetJob) String() string { return proto.CompactTextString(m) }
func (*GetJob) ProtoMessage()    {}
func (*GetJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{6}
}

func (m *GetJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJob.Unmarshal(m, b)
}
func (m *GetJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJob.Marshal(b, m, deterministic)
}
func (m *GetJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJob.Merge(m, src)
}
func (m *GetJob) XXX_Size() int {
	return xxx_messageInfo_GetJob.Size(m)
}
func (m *GetJob) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJob.DiscardUnknown(m)
}

var xxx_messageInfo_GetJob proto.InternalMessageInfo

func (m *GetJob) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetJob_Response struct {
	Job                  *Job     `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJob_Response) Reset()         { *m = GetJob_Response{} }
func (m *GetJob_Response) String() string { return proto.CompactTextString(m) }
func (*GetJob_Response) ProtoMessage()    {}
func (*GetJob_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcdcababea00ad0, []int{6, 0}
}

func (m *GetJob_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJob_Response.Unmarshal(m, b)
}
func (m *GetJob_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJob_Response.Marshal(b, m, deterministic)
}
func (m *GetJob_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJob_Response.Merge(m, src)
}
func (m *GetJob_Response) XXX_Size() int {
	return xxx_messageInfo_GetJob_Response.Size(m)
}
func (m *GetJob_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJob_Response.DiscardUnknown(m)
}

var xxx_messageInfo_GetJob_Response proto.InternalMessageInfo

func (m *GetJob_Response) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterEnum("ai.verta.modeldb.JobStatusEnum_JobStatus", JobStatusEnum_JobStatus_name, JobStatusEnum_JobStatus_value)
	proto.RegisterEnum("ai.verta.modeldb.JobTypeEnum_JobType", JobTypeEnum_JobType_name, JobTypeEnum_JobType_value)
	proto.RegisterType((*JobStatusEnum)(nil), "ai.verta.modeldb.JobStatusEnum")
	proto.RegisterType((*JobTypeEnum)(nil), "ai.verta.modeldb.JobTypeEnum")
	proto.RegisterType((*Job)(nil), "ai.verta.modeldb.Job")
	proto.RegisterType((*CreateJob)(nil), "ai.verta.modeldb.CreateJob")
	proto.RegisterType((*CreateJob_Response)(nil), "ai.verta.modeldb.CreateJob.Response")
	proto.RegisterType((*UpdateJob)(nil), "ai.verta.modeldb.UpdateJob")
	proto.RegisterType((*UpdateJob_Response)(nil), "ai.verta.modeldb.UpdateJob.Response")
	proto.RegisterType((*DeleteJob)(nil), "ai.verta.modeldb.DeleteJob")
	proto.RegisterType((*DeleteJob_Response)(nil), "ai.verta.modeldb.DeleteJob.Response")
	proto.RegisterType((*GetJob)(nil), "ai.verta.modeldb.GetJob")
	proto.RegisterType((*GetJob_Response)(nil), "ai.verta.modeldb.GetJob.Response")
}

func init() {
	proto.RegisterFile("protos/public/modeldb/Job.proto", fileDescriptor_7dcdcababea00ad0)
}

var fileDescriptor_7dcdcababea00ad0 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6b, 0xa7, 0x4d, 0xe2, 0x89, 0x1a, 0xc2, 0x02, 0x95, 0x1b, 0x5a, 0x11, 0x2c, 0x10,
	0x01, 0x21, 0x5b, 0xa4, 0xe2, 0x82, 0x10, 0xd0, 0xb4, 0x56, 0x69, 0x4a, 0x9b, 0xca, 0x49, 0x7b,
	0x40, 0x42, 0x91, 0x1d, 0xaf, 0xc2, 0x46, 0xb1, 0xd7, 0xb2, 0xd7, 0x45, 0xbd, 0xf2, 0x0a, 0x3c,
	0x08, 0x27, 0x5e, 0x82, 0x2b, 0x0f, 0xc0, 0x85, 0x07, 0x41, 0x5e, 0x3b, 0x9b, 0xfe, 0x71, 0x83,
	0x14, 0x71, 0x4a, 0x76, 0xbe, 0xdd, 0x6f, 0x66, 0x7e, 0x3b, 0xc9, 0xc2, 0x83, 0x20, 0xa4, 0x8c,
	0x46, 0x46, 0x10, 0x3b, 0x13, 0x32, 0x34, 0x3c, 0xea, 0xe2, 0x89, 0xeb, 0x18, 0x1d, 0xea, 0xe8,
	0x5c, 0x41, 0x35, 0x9b, 0xe8, 0x67, 0x38, 0x64, 0xb6, 0x9e, 0x69, 0xf5, 0xe6, 0xe5, 0x23, 0x43,
	0xea, 0x79, 0xd4, 0x37, 0x76, 0xf8, 0x47, 0x0f, 0x87, 0x67, 0x64, 0x88, 0xd3, 0xb3, 0xf5, 0x8d,
	0x11, 0xa5, 0xa3, 0x09, 0x36, 0xec, 0x80, 0x18, 0xb6, 0xef, 0x53, 0x66, 0x33, 0x42, 0xfd, 0x28,
	0x55, 0xb5, 0x43, 0x58, 0xed, 0x50, 0xa7, 0xc7, 0x6c, 0x16, 0x47, 0xa6, 0x1f, 0x7b, 0xda, 0x6b,
	0x50, 0x44, 0x00, 0xdd, 0x82, 0xca, 0x51, 0xb7, 0x3f, 0xe8, 0xf5, 0xb7, 0xad, 0xbe, 0xb9, 0x5b,
	0x5b, 0x4a, 0x02, 0xfb, 0x47, 0x83, 0x63, 0xab, 0xbb, 0x67, 0x99, 0xbd, 0x5e, 0x4d, 0x42, 0xab,
	0xa0, 0xec, 0x74, 0x0f, 0x8f, 0x3f, 0x98, 0x89, 0x2e, 0x6b, 0xcf, 0xa1, 0xd2, 0xa1, 0x4e, 0xff,
	0x3c, 0xc0, 0xdc, 0x6c, 0x13, 0x4a, 0xd9, 0x12, 0x21, 0xa8, 0x1e, 0x9c, 0xb4, 0x4d, 0xeb, 0xc8,
	0xec, 0x9b, 0xbd, 0x41, 0xa7, 0xdb, 0xae, 0x2d, 0x69, 0x3f, 0x65, 0x28, 0x74, 0xa8, 0x83, 0xaa,
	0x20, 0x13, 0x57, 0x95, 0x1a, 0x52, 0x53, 0xb1, 0x64, 0xe2, 0xa2, 0x06, 0x54, 0x5c, 0x1c, 0x0d,
	0x43, 0x12, 0x24, 0xa5, 0xaa, 0x32, 0x17, 0x2e, 0x86, 0xd0, 0x26, 0x40, 0xc4, 0xec, 0x90, 0x0d,
	0x18, 0xf1, 0xb0, 0x5a, 0xe0, 0x1b, 0x14, 0x1e, 0xe9, 0x13, 0x0f, 0xa3, 0x75, 0x28, 0x63, 0xdf,
	0x4d, 0xc5, 0x65, 0x2e, 0x96, 0xb0, 0xef, 0x72, 0xe9, 0x25, 0x94, 0x3d, 0xcc, 0x6c, 0xd7, 0x66,
	0xb6, 0xba, 0xd2, 0x28, 0x34, 0x2b, 0xad, 0x75, 0x5d, 0xd0, 0x4d, 0x31, 0xea, 0x07, 0xf8, 0xfc,
	0xd4, 0x9e, 0xc4, 0xd8, 0x12, 0x5b, 0xd1, 0x7b, 0x80, 0x31, 0x75, 0x06, 0x11, 0xe7, 0xa2, 0x16,
	0x1b, 0x52, 0xb3, 0xda, 0x7a, 0xaa, 0x5f, 0xbd, 0x16, 0xfd, 0x12, 0xcb, 0xd9, 0xca, 0x52, 0xc6,
	0x82, 0xe9, 0x3b, 0x28, 0x27, 0x4e, 0xec, 0x3c, 0xc0, 0x6a, 0x89, 0xfb, 0x3c, 0xce, 0xf5, 0x99,
	0x42, 0x9c, 0x7e, 0xb7, 0x4a, 0xe3, 0x0c, 0xe5, 0x5d, 0x58, 0xa1, 0x5f, 0x7c, 0x1c, 0xaa, 0x65,
	0xde, 0x5a, 0xba, 0xd0, 0x7e, 0xcb, 0xa0, 0xec, 0x84, 0xd8, 0x66, 0x38, 0x41, 0x7a, 0x05, 0xa1,
	0xf4, 0x2f, 0x84, 0xf2, 0x3c, 0x84, 0x85, 0x9b, 0x11, 0x2e, 0x2f, 0x8a, 0x70, 0xe5, 0x3f, 0x21,
	0x2c, 0x2e, 0x82, 0xb0, 0xbe, 0x05, 0x65, 0x0b, 0x47, 0x01, 0xf5, 0x23, 0x8c, 0x9e, 0x40, 0x61,
	0x4c, 0x1d, 0x8e, 0xa8, 0xd2, 0xba, 0x97, 0x6b, 0x64, 0x25, 0x3b, 0xb4, 0x1f, 0x12, 0x28, 0x27,
	0x81, 0x9b, 0x11, 0xbe, 0x3a, 0xb4, 0x73, 0x80, 0x5d, 0xee, 0x5c, 0x5e, 0xbc, 0xf3, 0xc5, 0xea,
	0x7e, 0x0b, 0xca, 0x2e, 0x9e, 0xe0, 0xdc, 0xb2, 0xeb, 0xda, 0x05, 0xc7, 0x35, 0x28, 0x66, 0x35,
	0x26, 0x7a, 0xd9, 0xca, 0x56, 0xda, 0x21, 0x14, 0xf7, 0x30, 0xcb, 0x3b, 0xbd, 0x48, 0x3d, 0xad,
	0xef, 0x05, 0x80, 0xa4, 0xbb, 0xf4, 0x6f, 0x0a, 0x4d, 0x40, 0x19, 0x8a, 0xb9, 0xbd, 0x7f, 0xfd,
	0x9c, 0x18, 0xea, 0xfa, 0xa3, 0x39, 0xa2, 0x3e, 0xad, 0x43, 0xdb, 0xf8, 0xfa, 0xeb, 0xcf, 0x37,
	0x79, 0x4d, 0xbb, 0x6d, 0x9c, 0xbd, 0x30, 0xc6, 0xd4, 0x31, 0x84, 0xfb, 0x2b, 0xe9, 0x19, 0xfa,
	0x04, 0xc5, 0x51, 0xda, 0x8b, 0x7a, 0xdd, 0x2d, 0xed, 0xb2, 0xfe, 0xf0, 0x26, 0x65, 0x96, 0x64,
	0x8d, 0x27, 0xa9, 0xa1, 0xea, 0x34, 0x49, 0x66, 0x4a, 0x40, 0x89, 0xc5, 0x88, 0xe4, 0x34, 0x23,
	0xe6, 0x27, 0xaf, 0x19, 0x21, 0xce, 0xf2, 0xac, 0xf3, 0x3c, 0x77, 0x90, 0x68, 0x66, 0xe6, 0x4e,
	0x40, 0x71, 0xc5, 0xb5, 0xe6, 0xa4, 0x12, 0x77, 0x9e, 0x97, 0x4a, 0x88, 0x73, 0x52, 0x09, 0xf7,
	0x76, 0xfb, 0x58, 0xfa, 0xf8, 0x66, 0x44, 0xd8, 0xe7, 0xd8, 0x49, 0x7e, 0xe0, 0xc6, 0x69, 0xe2,
	0xb5, 0xbd, 0x2f, 0x5e, 0xaa, 0xec, 0x31, 0x1a, 0x61, 0xdf, 0x18, 0x51, 0x23, 0xf7, 0x35, 0x73,
	0x8a, 0x3c, 0xbc, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xab, 0x97, 0xd0, 0xed, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	CreateJob(ctx context.Context, in *CreateJob, opts ...grpc.CallOption) (*CreateJob_Response, error)
	GetJob(ctx context.Context, in *GetJob, opts ...grpc.CallOption) (*GetJob_Response, error)
	UpdateJob(ctx context.Context, in *UpdateJob, opts ...grpc.CallOption) (*UpdateJob_Response, error)
	DeleteJob(ctx context.Context, in *DeleteJob, opts ...grpc.CallOption) (*DeleteJob_Response, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) CreateJob(ctx context.Context, in *CreateJob, opts ...grpc.CallOption) (*CreateJob_Response, error) {
	out := new(CreateJob_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.JobService/createJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *GetJob, opts ...grpc.CallOption) (*GetJob_Response, error) {
	out := new(GetJob_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.JobService/getJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) UpdateJob(ctx context.Context, in *UpdateJob, opts ...grpc.CallOption) (*UpdateJob_Response, error) {
	out := new(UpdateJob_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.JobService/updateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) DeleteJob(ctx context.Context, in *DeleteJob, opts ...grpc.CallOption) (*DeleteJob_Response, error) {
	out := new(DeleteJob_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.JobService/deleteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	CreateJob(context.Context, *CreateJob) (*CreateJob_Response, error)
	GetJob(context.Context, *GetJob) (*GetJob_Response, error)
	UpdateJob(context.Context, *UpdateJob) (*UpdateJob_Response, error)
	DeleteJob(context.Context, *DeleteJob) (*DeleteJob_Response, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) CreateJob(ctx context.Context, req *CreateJob) (*CreateJob_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (*UnimplementedJobServiceServer) GetJob(ctx context.Context, req *GetJob) (*GetJob_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (*UnimplementedJobServiceServer) UpdateJob(ctx context.Context, req *UpdateJob) (*UpdateJob_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJob not implemented")
}
func (*UnimplementedJobServiceServer) DeleteJob(ctx context.Context, req *DeleteJob) (*DeleteJob_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.JobService/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CreateJob(ctx, req.(*CreateJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.JobService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*GetJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_UpdateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).UpdateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.JobService/UpdateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).UpdateJob(ctx, req.(*UpdateJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.JobService/DeleteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).DeleteJob(ctx, req.(*DeleteJob))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.verta.modeldb.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createJob",
			Handler:    _JobService_CreateJob_Handler,
		},
		{
			MethodName: "getJob",
			Handler:    _JobService_GetJob_Handler,
		},
		{
			MethodName: "updateJob",
			Handler:    _JobService_UpdateJob_Handler,
		},
		{
			MethodName: "deleteJob",
			Handler:    _JobService_DeleteJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/public/modeldb/Job.proto",
}
