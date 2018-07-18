// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/group_service.proto

package monitoring // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The `ListGroup` request.
type ListGroupsRequest struct {
	// The project whose groups are to be listed. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// An optional filter consisting of a single group name.  The filters limit the
	// groups returned based on their parent-child relationship with the specified
	// group. If no filter is specified, all groups are returned.
	//
	// Types that are valid to be assigned to Filter:
	//	*ListGroupsRequest_ChildrenOfGroup
	//	*ListGroupsRequest_AncestorsOfGroup
	//	*ListGroupsRequest_DescendantsOfGroup
	Filter isListGroupsRequest_Filter `protobuf_oneof:"filter"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken            string   `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsRequest) Reset()         { *m = ListGroupsRequest{} }
func (m *ListGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupsRequest) ProtoMessage()    {}
func (*ListGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{0}
}
func (m *ListGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsRequest.Unmarshal(m, b)
}
func (m *ListGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsRequest.Marshal(b, m, deterministic)
}
func (dst *ListGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsRequest.Merge(dst, src)
}
func (m *ListGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupsRequest.Size(m)
}
func (m *ListGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsRequest proto.InternalMessageInfo

func (m *ListGroupsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListGroupsRequest_Filter interface {
	isListGroupsRequest_Filter()
}

type ListGroupsRequest_ChildrenOfGroup struct {
	ChildrenOfGroup string `protobuf:"bytes,2,opt,name=children_of_group,json=childrenOfGroup,proto3,oneof"`
}

type ListGroupsRequest_AncestorsOfGroup struct {
	AncestorsOfGroup string `protobuf:"bytes,3,opt,name=ancestors_of_group,json=ancestorsOfGroup,proto3,oneof"`
}

type ListGroupsRequest_DescendantsOfGroup struct {
	DescendantsOfGroup string `protobuf:"bytes,4,opt,name=descendants_of_group,json=descendantsOfGroup,proto3,oneof"`
}

func (*ListGroupsRequest_ChildrenOfGroup) isListGroupsRequest_Filter() {}

func (*ListGroupsRequest_AncestorsOfGroup) isListGroupsRequest_Filter() {}

func (*ListGroupsRequest_DescendantsOfGroup) isListGroupsRequest_Filter() {}

func (m *ListGroupsRequest) GetFilter() isListGroupsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListGroupsRequest) GetChildrenOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_ChildrenOfGroup); ok {
		return x.ChildrenOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetAncestorsOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_AncestorsOfGroup); ok {
		return x.AncestorsOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetDescendantsOfGroup() string {
	if x, ok := m.GetFilter().(*ListGroupsRequest_DescendantsOfGroup); ok {
		return x.DescendantsOfGroup
	}
	return ""
}

func (m *ListGroupsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListGroupsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListGroupsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListGroupsRequest_OneofMarshaler, _ListGroupsRequest_OneofUnmarshaler, _ListGroupsRequest_OneofSizer, []interface{}{
		(*ListGroupsRequest_ChildrenOfGroup)(nil),
		(*ListGroupsRequest_AncestorsOfGroup)(nil),
		(*ListGroupsRequest_DescendantsOfGroup)(nil),
	}
}

func _ListGroupsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListGroupsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ListGroupsRequest_ChildrenOfGroup:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ChildrenOfGroup)
	case *ListGroupsRequest_AncestorsOfGroup:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AncestorsOfGroup)
	case *ListGroupsRequest_DescendantsOfGroup:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DescendantsOfGroup)
	case nil:
	default:
		return fmt.Errorf("ListGroupsRequest.Filter has unexpected type %T", x)
	}
	return nil
}

func _ListGroupsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListGroupsRequest)
	switch tag {
	case 2: // filter.children_of_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Filter = &ListGroupsRequest_ChildrenOfGroup{x}
		return true, err
	case 3: // filter.ancestors_of_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Filter = &ListGroupsRequest_AncestorsOfGroup{x}
		return true, err
	case 4: // filter.descendants_of_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Filter = &ListGroupsRequest_DescendantsOfGroup{x}
		return true, err
	default:
		return false, nil
	}
}

func _ListGroupsRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListGroupsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ListGroupsRequest_ChildrenOfGroup:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ChildrenOfGroup)))
		n += len(x.ChildrenOfGroup)
	case *ListGroupsRequest_AncestorsOfGroup:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.AncestorsOfGroup)))
		n += len(x.AncestorsOfGroup)
	case *ListGroupsRequest_DescendantsOfGroup:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.DescendantsOfGroup)))
		n += len(x.DescendantsOfGroup)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The `ListGroups` response.
type ListGroupsResponse struct {
	// The groups that match the specified filters.
	Group []*Group `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupsResponse) Reset()         { *m = ListGroupsResponse{} }
func (m *ListGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupsResponse) ProtoMessage()    {}
func (*ListGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{1}
}
func (m *ListGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupsResponse.Unmarshal(m, b)
}
func (m *ListGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupsResponse.Marshal(b, m, deterministic)
}
func (dst *ListGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupsResponse.Merge(dst, src)
}
func (m *ListGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupsResponse.Size(m)
}
func (m *ListGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupsResponse proto.InternalMessageInfo

func (m *ListGroupsResponse) GetGroup() []*Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ListGroupsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The `GetGroup` request.
type GetGroupRequest struct {
	// The group to retrieve. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupRequest) Reset()         { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()    {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{2}
}
func (m *GetGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupRequest.Unmarshal(m, b)
}
func (m *GetGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupRequest.Marshal(b, m, deterministic)
}
func (dst *GetGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupRequest.Merge(dst, src)
}
func (m *GetGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupRequest.Size(m)
}
func (m *GetGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupRequest proto.InternalMessageInfo

func (m *GetGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The `CreateGroup` request.
type CreateGroupRequest struct {
	// The project in which to create the group. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A group definition. It is an error to define the `name` field because
	// the system assigns the name.
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// If true, validate this request but do not create the group.
	ValidateOnly         bool     `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{3}
}
func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (dst *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(dst, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateGroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *CreateGroupRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// The `UpdateGroup` request.
type UpdateGroupRequest struct {
	// The new definition of the group.  All fields of the existing group,
	// excepting `name`, are replaced with the corresponding fields of this group.
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// If true, validate this request but do not update the existing group.
	ValidateOnly         bool     `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{4}
}
func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(dst, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *UpdateGroupRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// The `DeleteGroup` request. You can only delete a group if it has no children.
type DeleteGroupRequest struct {
	// The group to delete. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{5}
}
func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(dst, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The `ListGroupMembers` request.
type ListGroupMembersRequest struct {
	// The group whose members are listed. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// An optional [list filter](/monitoring/api/learn_more#filtering) describing
	// the members to be returned.  The filter may reference the type, labels, and
	// metadata of monitored resources that comprise the group.
	// For example, to return only resources representing Compute Engine VM
	// instances, use this filter:
	//
	//     resource.type = "gce_instance"
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// An optional time interval for which results should be returned. Only
	// members that were part of the group during the specified interval are
	// included in the response.  If no interval is provided then the group
	// membership over the last minute is returned.
	Interval             *TimeInterval `protobuf:"bytes,6,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListGroupMembersRequest) Reset()         { *m = ListGroupMembersRequest{} }
func (m *ListGroupMembersRequest) String() string { return proto.CompactTextString(m) }
func (*ListGroupMembersRequest) ProtoMessage()    {}
func (*ListGroupMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{6}
}
func (m *ListGroupMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupMembersRequest.Unmarshal(m, b)
}
func (m *ListGroupMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupMembersRequest.Marshal(b, m, deterministic)
}
func (dst *ListGroupMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupMembersRequest.Merge(dst, src)
}
func (m *ListGroupMembersRequest) XXX_Size() int {
	return xxx_messageInfo_ListGroupMembersRequest.Size(m)
}
func (m *ListGroupMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupMembersRequest proto.InternalMessageInfo

func (m *ListGroupMembersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListGroupMembersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListGroupMembersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListGroupMembersRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListGroupMembersRequest) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

// The `ListGroupMembers` response.
type ListGroupMembersResponse struct {
	// A set of monitored resources in the group.
	Members []*monitoredres.MonitoredResource `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	// If there are more results than have been returned, then this field is
	// set to a non-empty value.  To see the additional results, use that value as
	// `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of elements matching this request.
	TotalSize            int32    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGroupMembersResponse) Reset()         { *m = ListGroupMembersResponse{} }
func (m *ListGroupMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ListGroupMembersResponse) ProtoMessage()    {}
func (*ListGroupMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_service_460c0e48a72aa329, []int{7}
}
func (m *ListGroupMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGroupMembersResponse.Unmarshal(m, b)
}
func (m *ListGroupMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGroupMembersResponse.Marshal(b, m, deterministic)
}
func (dst *ListGroupMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGroupMembersResponse.Merge(dst, src)
}
func (m *ListGroupMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ListGroupMembersResponse.Size(m)
}
func (m *ListGroupMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGroupMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGroupMembersResponse proto.InternalMessageInfo

func (m *ListGroupMembersResponse) GetMembers() []*monitoredres.MonitoredResource {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *ListGroupMembersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListGroupMembersResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func init() {
	proto.RegisterType((*ListGroupsRequest)(nil), "google.monitoring.v3.ListGroupsRequest")
	proto.RegisterType((*ListGroupsResponse)(nil), "google.monitoring.v3.ListGroupsResponse")
	proto.RegisterType((*GetGroupRequest)(nil), "google.monitoring.v3.GetGroupRequest")
	proto.RegisterType((*CreateGroupRequest)(nil), "google.monitoring.v3.CreateGroupRequest")
	proto.RegisterType((*UpdateGroupRequest)(nil), "google.monitoring.v3.UpdateGroupRequest")
	proto.RegisterType((*DeleteGroupRequest)(nil), "google.monitoring.v3.DeleteGroupRequest")
	proto.RegisterType((*ListGroupMembersRequest)(nil), "google.monitoring.v3.ListGroupMembersRequest")
	proto.RegisterType((*ListGroupMembersResponse)(nil), "google.monitoring.v3.ListGroupMembersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupServiceClient interface {
	// Lists the existing groups.
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// Gets a single group.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Creates a new group.
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Updates an existing group.
	// You can change any group attributes except `name`.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Deletes an existing group.
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists the monitored resources that are members of a group.
	ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error)
}

type groupServiceClient struct {
	cc *grpc.ClientConn
}

func NewGroupServiceClient(cc *grpc.ClientConn) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error) {
	out := new(ListGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.GroupService/ListGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
type GroupServiceServer interface {
	// Lists the existing groups.
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	// Gets a single group.
	GetGroup(context.Context, *GetGroupRequest) (*Group, error)
	// Creates a new group.
	CreateGroup(context.Context, *CreateGroupRequest) (*Group, error)
	// Updates an existing group.
	// You can change any group attributes except `name`.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error)
	// Deletes an existing group.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*empty.Empty, error)
	// Lists the monitored resources that are members of a group.
	ListGroupMembers(context.Context, *ListGroupMembersRequest) (*ListGroupMembersResponse, error)
}

func RegisterGroupServiceServer(s *grpc.Server, srv GroupServiceServer) {
	s.RegisterService(&_GroupService_serviceDesc, srv)
}

func _GroupService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.GroupService/ListGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroupMembers(ctx, req.(*ListGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGroups",
			Handler:    _GroupService_ListGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupService_GetGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "ListGroupMembers",
			Handler:    _GroupService_ListGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/monitoring/v3/group_service.proto",
}

func init() {
	proto.RegisterFile("google/monitoring/v3/group_service.proto", fileDescriptor_group_service_460c0e48a72aa329)
}

var fileDescriptor_group_service_460c0e48a72aa329 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x7e, 0xdd, 0xa4, 0x69, 0xb2, 0x69, 0xd5, 0x76, 0x55, 0xf5, 0x8d, 0xdc, 0x0f, 0x05, 0xf7,
	0x83, 0xa8, 0x50, 0x5b, 0x24, 0x07, 0x24, 0x10, 0x3d, 0xb4, 0xa0, 0x82, 0x44, 0xd5, 0xca, 0x2d,
	0x3d, 0xa0, 0x4a, 0x91, 0x9b, 0x4c, 0x8c, 0xc1, 0xde, 0x35, 0xf6, 0x26, 0xd0, 0xa2, 0x4a, 0x80,
	0xc4, 0x81, 0x33, 0x37, 0x6e, 0x1c, 0xe1, 0x2f, 0x70, 0xe2, 0xca, 0x95, 0xbf, 0xc0, 0xff, 0x00,
	0x79, 0xbd, 0x9b, 0x38, 0x9f, 0xed, 0x85, 0x5b, 0xb2, 0xf3, 0x8c, 0x9f, 0x67, 0x66, 0x9f, 0x99,
	0x45, 0x25, 0x9b, 0x52, 0xdb, 0x05, 0xc3, 0xa3, 0xc4, 0x61, 0x34, 0x70, 0x88, 0x6d, 0xb4, 0x2a,
	0x86, 0x1d, 0xd0, 0xa6, 0x5f, 0x0d, 0x21, 0x68, 0x39, 0x35, 0xd0, 0xfd, 0x80, 0x32, 0x8a, 0xe7,
	0x62, 0xa4, 0xde, 0x41, 0xea, 0xad, 0x8a, 0xba, 0x28, 0xf2, 0x2d, 0xdf, 0x31, 0x2c, 0x42, 0x28,
	0xb3, 0x98, 0x43, 0x49, 0x18, 0xe7, 0xa8, 0x2b, 0x89, 0xa8, 0xc8, 0x83, 0x7a, 0x35, 0x80, 0x90,
	0x36, 0x03, 0xf9, 0x61, 0xf5, 0xda, 0x40, 0x09, 0x35, 0xea, 0x79, 0x94, 0x08, 0x48, 0x71, 0xb8,
	0x4a, 0x81, 0x58, 0x10, 0x08, 0xfe, 0xef, 0xb4, 0xd9, 0x30, 0xc0, 0xf3, 0xd9, 0x59, 0x1c, 0xd4,
	0xfe, 0x28, 0x68, 0xf6, 0xb1, 0x13, 0xb2, 0xdd, 0x28, 0x21, 0x34, 0xe1, 0x65, 0x13, 0x42, 0x86,
	0x31, 0x4a, 0x13, 0xcb, 0x83, 0xc2, 0x44, 0x51, 0x29, 0xe5, 0x4c, 0xfe, 0x1b, 0xdf, 0x44, 0xb3,
	0xb5, 0x67, 0x8e, 0x5b, 0x0f, 0x80, 0x54, 0x69, 0xa3, 0xca, 0x19, 0x0a, 0x63, 0x11, 0xe0, 0xe1,
	0x7f, 0xe6, 0xb4, 0x0c, 0xed, 0x37, 0xf8, 0x97, 0xb0, 0x8e, 0xb0, 0x45, 0x6a, 0x10, 0x32, 0x1a,
	0x84, 0x1d, 0x78, 0x4a, 0xc0, 0x67, 0xda, 0x31, 0x89, 0x2f, 0xa3, 0xb9, 0x3a, 0x84, 0x35, 0x20,
	0x75, 0x8b, 0xb0, 0x44, 0x46, 0x5a, 0x64, 0xe0, 0x44, 0x54, 0xe6, 0x2c, 0xa0, 0x9c, 0x6f, 0xd9,
	0x50, 0x0d, 0x9d, 0x73, 0x28, 0x8c, 0x17, 0x95, 0xd2, 0xb8, 0x99, 0x8d, 0x0e, 0x0e, 0x9d, 0x73,
	0xc0, 0x4b, 0x08, 0xf1, 0x20, 0xa3, 0x2f, 0x80, 0x14, 0x32, 0xbc, 0x10, 0x0e, 0x3f, 0x8a, 0x0e,
	0xb6, 0xb3, 0x28, 0xd3, 0x70, 0x5c, 0x06, 0x81, 0x46, 0x11, 0x4e, 0x36, 0x20, 0xf4, 0x29, 0x09,
	0x01, 0xdf, 0x42, 0xe3, 0xb1, 0x00, 0xa5, 0x98, 0x2a, 0xe5, 0xcb, 0x0b, 0xfa, 0xa0, 0x2b, 0xd6,
	0x79, 0x92, 0x19, 0x23, 0xf1, 0x3a, 0x9a, 0x26, 0xf0, 0x9a, 0x55, 0x13, 0xb4, 0xbc, 0x3d, 0xe6,
	0x54, 0x74, 0x7c, 0x20, 0xa9, 0xb5, 0x35, 0x34, 0xbd, 0x0b, 0x31, 0x5f, 0x6f, 0xbf, 0x53, 0x9d,
	0x7e, 0x6b, 0x6f, 0x15, 0x84, 0x77, 0x02, 0xb0, 0x18, 0x0c, 0x84, 0xa6, 0x13, 0x57, 0xd3, 0x16,
	0x1b, 0xf1, 0x5d, 0x4d, 0xec, 0x0a, 0x9a, 0x6a, 0x59, 0xae, 0x53, 0xb7, 0x18, 0x54, 0x29, 0x71,
	0xcf, 0x38, 0x75, 0xd6, 0x9c, 0x94, 0x87, 0xfb, 0xc4, 0x3d, 0xd3, 0x5c, 0x84, 0x9f, 0xf8, 0xf5,
	0x5e, 0x05, 0xff, 0x8a, 0xad, 0x84, 0xf0, 0x7d, 0x70, 0x61, 0x48, 0xbd, 0xc9, 0xd6, 0xfc, 0x50,
	0xd0, 0xff, 0xed, 0x3b, 0xdb, 0x03, 0xef, 0x14, 0x82, 0x91, 0xd6, 0xed, 0x32, 0x4a, 0x6a, 0xa4,
	0x51, 0xd2, 0x3d, 0x46, 0xc1, 0xf3, 0xd2, 0x28, 0xdc, 0x61, 0x39, 0x53, 0xfc, 0xc3, 0x5b, 0x28,
	0xeb, 0x10, 0x06, 0x41, 0xcb, 0x72, 0xb9, 0xbb, 0xf2, 0x65, 0x6d, 0x70, 0x23, 0x8e, 0x1c, 0x0f,
	0x1e, 0x09, 0xa4, 0xd9, 0xce, 0xd1, 0x3e, 0x2b, 0xa8, 0xd0, 0x5f, 0x83, 0x70, 0xdf, 0x6d, 0x34,
	0xe1, 0xc5, 0x47, 0xc2, 0x7f, 0x4b, 0xf2, 0xdb, 0x96, 0xef, 0xe8, 0x7b, 0x72, 0x5d, 0x98, 0x62,
	0x5b, 0x98, 0x12, 0x7d, 0x55, 0x0f, 0x46, 0x45, 0x33, 0xca, 0x2c, 0x37, 0xd9, 0x92, 0x1c, 0x3f,
	0x89, 0x7a, 0x52, 0xfe, 0x9e, 0x41, 0x93, 0x5c, 0xd8, 0x61, 0xbc, 0xe7, 0xf0, 0x07, 0x05, 0xa1,
	0xce, 0x94, 0xe0, 0xeb, 0x83, 0x4b, 0xed, 0x5b, 0x24, 0x6a, 0xe9, 0x72, 0x60, 0x5c, 0xb2, 0xb6,
	0xfa, 0xfe, 0xd7, 0xef, 0x4f, 0x63, 0xcb, 0x78, 0x31, 0x5a, 0x5f, 0x6f, 0xa2, 0x6b, 0xbb, 0xe7,
	0x07, 0xf4, 0x39, 0xd4, 0x58, 0x68, 0x6c, 0x5c, 0xc4, 0x0b, 0x2d, 0xc4, 0x2d, 0x94, 0x95, 0xb3,
	0x83, 0xd7, 0x86, 0x18, 0xaf, 0x7b, 0xb6, 0xd4, 0x51, 0xfe, 0xd4, 0xd6, 0x39, 0x6b, 0x11, 0x2f,
	0x0f, 0x62, 0x15, 0xa4, 0xc6, 0xc6, 0x05, 0x7e, 0xa7, 0xa0, 0x7c, 0x62, 0x18, 0xf1, 0x90, 0xba,
	0xfa, 0xe7, 0x75, 0x34, 0xfd, 0x0d, 0x4e, 0xbf, 0xa6, 0x8d, 0x2c, 0xfa, 0x8e, 0x18, 0xa2, 0x8f,
	0x0a, 0xca, 0x27, 0xc6, 0x71, 0x98, 0x86, 0xfe, 0x89, 0x1d, 0xad, 0xa1, 0xc2, 0x35, 0x6c, 0xaa,
	0xab, 0x5c, 0x43, 0xfc, 0x70, 0x0c, 0x6d, 0x84, 0xd4, 0xf2, 0x0a, 0xe5, 0x13, 0xb3, 0x3a, 0x4c,
	0x4a, 0xff, 0x38, 0xab, 0xf3, 0x12, 0x29, 0x5f, 0x23, 0xfd, 0x41, 0xf4, 0x1a, 0xc9, 0x8b, 0xd8,
	0xb8, 0xec, 0x22, 0xbe, 0x28, 0x68, 0xa6, 0x77, 0x6c, 0xf0, 0xe6, 0x25, 0x2e, 0xeb, 0x5e, 0x11,
	0xaa, 0x7e, 0x55, 0xb8, 0xb0, 0xa6, 0xce, 0xb5, 0x95, 0xf0, 0xfa, 0x68, 0x6d, 0x86, 0x18, 0xc2,
	0xed, 0xaf, 0x0a, 0x2a, 0xd4, 0xa8, 0x37, 0x90, 0x65, 0x7b, 0x36, 0x39, 0x57, 0x07, 0x51, 0x13,
	0x0e, 0x94, 0xa7, 0x5b, 0x02, 0x6a, 0x53, 0xd7, 0x22, 0xb6, 0x4e, 0x03, 0xdb, 0xb0, 0x81, 0xf0,
	0x16, 0x19, 0x71, 0xc8, 0xf2, 0x9d, 0xb0, 0xfb, 0x8d, 0xbf, 0xdb, 0xf9, 0xf7, 0x6d, 0x4c, 0xdd,
	0x8d, 0x3f, 0xb0, 0xe3, 0xd2, 0x66, 0x5d, 0x2e, 0x88, 0x88, 0xf1, 0xb8, 0xf2, 0x53, 0x06, 0x4f,
	0x78, 0xf0, 0xa4, 0x13, 0x3c, 0x39, 0xae, 0x9c, 0x66, 0x38, 0x49, 0xe5, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x86, 0x94, 0xf2, 0xde, 0xed, 0x08, 0x00, 0x00,
}
