// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/proto/v1/scans.proto

package analytics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId                 string                        `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	ProjectId                string                        `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Deleted                  bool                          `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	TimeZone                 string                        `protobuf:"bytes,5,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
	CreatedAt                int64                         `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt                int64                         `protobuf:"varint,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UpdatedAt                int64                         `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Source                   string                        `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	Origin                   string                        `protobuf:"bytes,10,opt,name=origin,proto3" json:"origin,omitempty"`
	Status                   string                        `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	EngineExecutionSummaries []*ScanEngineExecutionSummary `protobuf:"bytes,12,rep,name=engine_execution_summaries,json=engineExecutionSummaries,proto3" json:"engine_execution_summaries,omitempty"`
	ConcurrentScans          uint32                        `protobuf:"varint,13,opt,name=concurrent_scans,json=concurrentScans,proto3" json:"concurrent_scans,omitempty"`
	Tags                     []string                      `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	Initiator                string                        `protobuf:"bytes,15,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Branch                   string                        `protobuf:"bytes,16,opt,name=branch,proto3" json:"branch,omitempty"`
	ExecutionTime            uint64                        `protobuf:"varint,17,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	PreScan                  *PreScan                      `protobuf:"bytes,18,opt,name=preScan,proto3" json:"preScan,omitempty"`
	Engines                  []string                      `protobuf:"bytes,19,rep,name=engines,proto3" json:"engines,omitempty"`
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{0}
}

func (x *ScanRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScanRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ScanRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ScanRequest) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *ScanRequest) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *ScanRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ScanRequest) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *ScanRequest) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ScanRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ScanRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ScanRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ScanRequest) GetEngineExecutionSummaries() []*ScanEngineExecutionSummary {
	if x != nil {
		return x.EngineExecutionSummaries
	}
	return nil
}

func (x *ScanRequest) GetConcurrentScans() uint32 {
	if x != nil {
		return x.ConcurrentScans
	}
	return 0
}

func (x *ScanRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ScanRequest) GetInitiator() string {
	if x != nil {
		return x.Initiator
	}
	return ""
}

func (x *ScanRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *ScanRequest) GetExecutionTime() uint64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *ScanRequest) GetPreScan() *PreScan {
	if x != nil {
		return x.PreScan
	}
	return nil
}

func (x *ScanRequest) GetEngines() []string {
	if x != nil {
		return x.Engines
	}
	return nil
}

type ScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ScanResponse) Reset() {
	*x = ScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResponse) ProtoMessage() {}

func (x *ScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResponse.ProtoReflect.Descriptor instead.
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{1}
}

func (x *ScanResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type ScanEngineExecutionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engine        string               `protobuf:"bytes,1,opt,name=engine,proto3" json:"engine,omitempty"`
	Loc           uint64               `protobuf:"varint,2,opt,name=loc,proto3" json:"loc,omitempty"`
	ScanErrorCode string               `protobuf:"bytes,3,opt,name=scan_error_code,json=scanErrorCode,proto3" json:"scan_error_code,omitempty"`
	Sources       []string             `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	ScanType      string               `protobuf:"bytes,5,opt,name=scan_type,json=scanType,proto3" json:"scan_type,omitempty"`
	ExecutionTime uint64               `protobuf:"varint,6,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	EngineStatus  string               `protobuf:"bytes,7,opt,name=engine_status,json=engineStatus,proto3" json:"engine_status,omitempty"`
	Severities    []*ScanVulnerability `protobuf:"bytes,8,rep,name=severities,proto3" json:"severities,omitempty"`
	States        []*ScanVulnerability `protobuf:"bytes,9,rep,name=states,proto3" json:"states,omitempty"`
	Statuses      []*ScanVulnerability `protobuf:"bytes,10,rep,name=statuses,proto3" json:"statuses,omitempty"`
}

func (x *ScanEngineExecutionSummary) Reset() {
	*x = ScanEngineExecutionSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanEngineExecutionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanEngineExecutionSummary) ProtoMessage() {}

func (x *ScanEngineExecutionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanEngineExecutionSummary.ProtoReflect.Descriptor instead.
func (*ScanEngineExecutionSummary) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{2}
}

func (x *ScanEngineExecutionSummary) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *ScanEngineExecutionSummary) GetLoc() uint64 {
	if x != nil {
		return x.Loc
	}
	return 0
}

func (x *ScanEngineExecutionSummary) GetScanErrorCode() string {
	if x != nil {
		return x.ScanErrorCode
	}
	return ""
}

func (x *ScanEngineExecutionSummary) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *ScanEngineExecutionSummary) GetScanType() string {
	if x != nil {
		return x.ScanType
	}
	return ""
}

func (x *ScanEngineExecutionSummary) GetExecutionTime() uint64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *ScanEngineExecutionSummary) GetEngineStatus() string {
	if x != nil {
		return x.EngineStatus
	}
	return ""
}

func (x *ScanEngineExecutionSummary) GetSeverities() []*ScanVulnerability {
	if x != nil {
		return x.Severities
	}
	return nil
}

func (x *ScanEngineExecutionSummary) GetStates() []*ScanVulnerability {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *ScanEngineExecutionSummary) GetStatuses() []*ScanVulnerability {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type ScanVulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Counter int64  `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *ScanVulnerability) Reset() {
	*x = ScanVulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanVulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanVulnerability) ProtoMessage() {}

func (x *ScanVulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanVulnerability.ProtoReflect.Descriptor instead.
func (*ScanVulnerability) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{3}
}

func (x *ScanVulnerability) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScanVulnerability) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type ScanStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ScanCount int64  `protobuf:"varint,2,opt,name=scanCount,proto3" json:"scanCount,omitempty"`
}

func (x *ScanStatus) Reset() {
	*x = ScanStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanStatus) ProtoMessage() {}

func (x *ScanStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanStatus.ProtoReflect.Descriptor instead.
func (*ScanStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{4}
}

func (x *ScanStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ScanStatus) GetScanCount() int64 {
	if x != nil {
		return x.ScanCount
	}
	return 0
}

type ScanIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId  string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	EventTime int64  `protobuf:"varint,3,opt,name=eventTime,proto3" json:"eventTime,omitempty"`
}

func (x *ScanIdentifier) Reset() {
	*x = ScanIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanIdentifier) ProtoMessage() {}

func (x *ScanIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanIdentifier.ProtoReflect.Descriptor instead.
func (*ScanIdentifier) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{5}
}

func (x *ScanIdentifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScanIdentifier) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ScanIdentifier) GetEventTime() int64 {
	if x != nil {
		return x.EventTime
	}
	return 0
}

type PreScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode     string `protobuf:"bytes,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ExecutionTime uint64 `protobuf:"varint,2,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	Status        string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PreScan) Reset() {
	*x = PreScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_scans_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreScan) ProtoMessage() {}

func (x *PreScan) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_scans_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreScan.ProtoReflect.Descriptor instead.
func (*PreScan) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_scans_proto_rawDescGZIP(), []int{6}
}

func (x *PreScan) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *PreScan) GetExecutionTime() uint64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *PreScan) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_proto_v1_scans_proto protoreflect.FileDescriptor

var file_api_proto_v1_scans_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x63, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x22, 0xfd, 0x04, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x63, 0x0a, 0x1a,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x63, 0x61,
	0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x18, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x07, 0x70, 0x72, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x07, 0x70, 0x72, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x03,
	0x0a, 0x1a, 0x53, 0x63, 0x61, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x63, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x61,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6e,
	0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22,
	0x41, 0x0a, 0x11, 0x53, 0x63, 0x61, 0x6e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x22, 0x42, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x63, 0x61,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x0e, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xbb, 0x01, 0x0a,
	0x05, 0x53, 0x63, 0x61, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x1a, 0x17, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f,
	0x3b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_v1_scans_proto_rawDescOnce sync.Once
	file_api_proto_v1_scans_proto_rawDescData = file_api_proto_v1_scans_proto_rawDesc
)

func file_api_proto_v1_scans_proto_rawDescGZIP() []byte {
	file_api_proto_v1_scans_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_scans_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_scans_proto_rawDescData)
	})
	return file_api_proto_v1_scans_proto_rawDescData
}

var file_api_proto_v1_scans_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_v1_scans_proto_goTypes = []interface{}{
	(*ScanRequest)(nil),                // 0: analytics.ScanRequest
	(*ScanResponse)(nil),               // 1: analytics.ScanResponse
	(*ScanEngineExecutionSummary)(nil), // 2: analytics.ScanEngineExecutionSummary
	(*ScanVulnerability)(nil),          // 3: analytics.ScanVulnerability
	(*ScanStatus)(nil),                 // 4: analytics.ScanStatus
	(*ScanIdentifier)(nil),             // 5: analytics.ScanIdentifier
	(*PreScan)(nil),                    // 6: analytics.PreScan
}
var file_api_proto_v1_scans_proto_depIdxs = []int32{
	2, // 0: analytics.ScanRequest.engine_execution_summaries:type_name -> analytics.ScanEngineExecutionSummary
	6, // 1: analytics.ScanRequest.preScan:type_name -> analytics.PreScan
	3, // 2: analytics.ScanEngineExecutionSummary.severities:type_name -> analytics.ScanVulnerability
	3, // 3: analytics.ScanEngineExecutionSummary.states:type_name -> analytics.ScanVulnerability
	3, // 4: analytics.ScanEngineExecutionSummary.statuses:type_name -> analytics.ScanVulnerability
	0, // 5: analytics.Scans.create:input_type -> analytics.ScanRequest
	0, // 6: analytics.Scans.update:input_type -> analytics.ScanRequest
	5, // 7: analytics.Scans.delete:input_type -> analytics.ScanIdentifier
	1, // 8: analytics.Scans.create:output_type -> analytics.ScanResponse
	1, // 9: analytics.Scans.update:output_type -> analytics.ScanResponse
	1, // 10: analytics.Scans.delete:output_type -> analytics.ScanResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_v1_scans_proto_init() }
func file_api_proto_v1_scans_proto_init() {
	if File_api_proto_v1_scans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_scans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_scans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_scans_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanEngineExecutionSummary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_scans_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanVulnerability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_scans_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_scans_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_scans_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreScan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_scans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_scans_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_scans_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_scans_proto_msgTypes,
	}.Build()
	File_api_proto_v1_scans_proto = out.File
	file_api_proto_v1_scans_proto_rawDesc = nil
	file_api_proto_v1_scans_proto_goTypes = nil
	file_api_proto_v1_scans_proto_depIdxs = nil
}
