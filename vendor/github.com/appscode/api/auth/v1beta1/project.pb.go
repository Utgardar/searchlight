// Code generated by protoc-gen-go.
// source: project.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProjectListRequest struct {
	WithMember bool     `protobuf:"varint,1,opt,name=with_member,json=withMember" json:"with_member,omitempty"`
	Members    []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *ProjectListRequest) Reset()                    { *m = ProjectListRequest{} }
func (m *ProjectListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProjectListRequest) ProtoMessage()               {}
func (*ProjectListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ProjectListRequest) GetWithMember() bool {
	if m != nil {
		return m.WithMember
	}
	return false
}

func (m *ProjectListRequest) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ProjectListResponse struct {
	Projects []*Project `protobuf:"bytes,1,rep,name=projects" json:"projects,omitempty"`
}

func (m *ProjectListResponse) Reset()                    { *m = ProjectListResponse{} }
func (m *ProjectListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectListResponse) ProtoMessage()               {}
func (*ProjectListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ProjectListResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ProjectMemberListRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *ProjectMemberListRequest) Reset()                    { *m = ProjectMemberListRequest{} }
func (m *ProjectMemberListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProjectMemberListRequest) ProtoMessage()               {}
func (*ProjectMemberListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ProjectMemberListRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ProjectMemberListResponse struct {
	Project *Project `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
}

func (m *ProjectMemberListResponse) Reset()                    { *m = ProjectMemberListResponse{} }
func (m *ProjectMemberListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectMemberListResponse) ProtoMessage()               {}
func (*ProjectMemberListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ProjectMemberListResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type Project struct {
	Name             string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phid             string    `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Type             string    `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Status           string    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ViewPolicy       string    `protobuf:"bytes,5,opt,name=view_policy,json=viewPolicy" json:"view_policy,omitempty"`
	EditPolicy       string    `protobuf:"bytes,6,opt,name=edit_policy,json=editPolicy" json:"edit_policy,omitempty"`
	JoinPolicy       string    `protobuf:"bytes,7,opt,name=join_policy,json=joinPolicy" json:"join_policy,omitempty"`
	MembershipLocked bool      `protobuf:"varint,8,opt,name=membership_locked,json=membershipLocked" json:"membership_locked,omitempty"`
	HasSubprojects   bool      `protobuf:"varint,9,opt,name=has_subprojects,json=hasSubprojects" json:"has_subprojects,omitempty"`
	Members          []*Member `protobuf:"bytes,10,rep,name=members" json:"members,omitempty"`
	CreatedAt        int64     `protobuf:"varint,11,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Project) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Project) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Project) GetViewPolicy() string {
	if m != nil {
		return m.ViewPolicy
	}
	return ""
}

func (m *Project) GetEditPolicy() string {
	if m != nil {
		return m.EditPolicy
	}
	return ""
}

func (m *Project) GetJoinPolicy() string {
	if m != nil {
		return m.JoinPolicy
	}
	return ""
}

func (m *Project) GetMembershipLocked() bool {
	if m != nil {
		return m.MembershipLocked
	}
	return false
}

func (m *Project) GetHasSubprojects() bool {
	if m != nil {
		return m.HasSubprojects
	}
	return false
}

func (m *Project) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Project) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Member struct {
	Phid     string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	RealName string `protobuf:"bytes,3,opt,name=real_name,json=realName" json:"real_name,omitempty"`
	IsAdmin  bool   `protobuf:"varint,4,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Member) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Member) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Member) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Member) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func init() {
	proto.RegisterType((*ProjectListRequest)(nil), "appscode.auth.v1beta1.ProjectListRequest")
	proto.RegisterType((*ProjectListResponse)(nil), "appscode.auth.v1beta1.ProjectListResponse")
	proto.RegisterType((*ProjectMemberListRequest)(nil), "appscode.auth.v1beta1.ProjectMemberListRequest")
	proto.RegisterType((*ProjectMemberListResponse)(nil), "appscode.auth.v1beta1.ProjectMemberListResponse")
	proto.RegisterType((*Project)(nil), "appscode.auth.v1beta1.Project")
	proto.RegisterType((*Member)(nil), "appscode.auth.v1beta1.Member")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Projects service

type ProjectsClient interface {
	List(ctx context.Context, in *ProjectListRequest, opts ...grpc.CallOption) (*ProjectListResponse, error)
	Members(ctx context.Context, in *ProjectMemberListRequest, opts ...grpc.CallOption) (*ProjectMemberListResponse, error)
}

type projectsClient struct {
	cc *grpc.ClientConn
}

func NewProjectsClient(cc *grpc.ClientConn) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) List(ctx context.Context, in *ProjectListRequest, opts ...grpc.CallOption) (*ProjectListResponse, error) {
	out := new(ProjectListResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Projects/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Members(ctx context.Context, in *ProjectMemberListRequest, opts ...grpc.CallOption) (*ProjectMemberListResponse, error) {
	out := new(ProjectMemberListResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Projects/Members", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Projects service

type ProjectsServer interface {
	List(context.Context, *ProjectListRequest) (*ProjectListResponse, error)
	Members(context.Context, *ProjectMemberListRequest) (*ProjectMemberListResponse, error)
}

func RegisterProjectsServer(s *grpc.Server, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Projects/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).List(ctx, req.(*ProjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectMemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Projects/Members",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Members(ctx, req.(*ProjectMemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1beta1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Projects_List_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _Projects_Members_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}

func init() { proto.RegisterFile("project.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0xa7, 0xc4, 0xf6, 0x94, 0x8f, 0xb2, 0x08, 0xb4, 0x4d, 0x69, 0x89, 0x7c, 0x21,
	0xa5, 0x28, 0xa6, 0xe5, 0xd0, 0xc2, 0xad, 0x95, 0xb8, 0xb5, 0x10, 0x8c, 0xb8, 0x70, 0xb1, 0x36,
	0xf6, 0xaa, 0xde, 0x92, 0x78, 0x5d, 0xef, 0xba, 0x55, 0x85, 0xb8, 0xf4, 0x15, 0x90, 0x78, 0x11,
	0xc4, 0x93, 0x70, 0xe2, 0xce, 0x99, 0x67, 0x40, 0xfb, 0xe1, 0x26, 0x11, 0xa1, 0xe4, 0xb6, 0xfe,
	0xcd, 0x7f, 0x3c, 0xff, 0x9d, 0x19, 0x1b, 0x6e, 0x95, 0x15, 0x3f, 0xa1, 0xa9, 0xec, 0x97, 0x15,
	0x97, 0x1c, 0xdd, 0x27, 0x65, 0x29, 0x52, 0x9e, 0xd1, 0x3e, 0xa9, 0x65, 0xde, 0x3f, 0xdb, 0x1e,
	0x52, 0x49, 0xb6, 0x3b, 0x0f, 0x8f, 0x39, 0x3f, 0x1e, 0xd1, 0x88, 0x94, 0x2c, 0x22, 0x45, 0xc1,
	0x25, 0x91, 0x8c, 0x17, 0xc2, 0x24, 0x75, 0x36, 0x9a, 0xa4, 0xf9, 0xf1, 0xf0, 0x0d, 0xa0, 0x81,
	0xa9, 0x72, 0xc8, 0x84, 0x8c, 0xe9, 0x69, 0x4d, 0x85, 0x44, 0x8f, 0x60, 0xf9, 0x9c, 0xc9, 0x3c,
	0x19, 0xd3, 0xf1, 0x90, 0x56, 0xd8, 0xe9, 0x3a, 0x3d, 0x3f, 0x06, 0x85, 0x8e, 0x34, 0x41, 0x18,
	0x3c, 0x13, 0x13, 0xd8, 0xed, 0xb6, 0x7a, 0x41, 0xdc, 0x3c, 0x86, 0x6f, 0xe1, 0xde, 0xcc, 0x0b,
	0x45, 0xc9, 0x0b, 0x41, 0xd1, 0x4b, 0xf0, 0xed, 0x6d, 0x04, 0x76, 0xba, 0xad, 0xde, 0xf2, 0xce,
	0x46, 0x7f, 0xee, 0x7d, 0xfa, 0x36, 0x3b, 0xbe, 0xd2, 0x87, 0x4f, 0x01, 0x5b, 0x68, 0xaa, 0x4f,
	0x3b, 0x5d, 0x81, 0x56, 0xcd, 0x32, 0xed, 0x30, 0x88, 0xd5, 0x31, 0x7c, 0x0f, 0xab, 0x73, 0xd4,
	0xd6, 0xc6, 0x1e, 0x78, 0xf6, 0xb5, 0x3a, 0xe5, 0xff, 0x2e, 0x1a, 0x79, 0xf8, 0xdb, 0x05, 0xcf,
	0x42, 0x84, 0x60, 0xa9, 0x20, 0x63, 0x6a, 0xab, 0xea, 0xb3, 0x62, 0x65, 0xce, 0x32, 0xec, 0x1a,
	0xa6, 0xce, 0x8a, 0xc9, 0x8b, 0x92, 0xe2, 0x96, 0x61, 0xea, 0x8c, 0x1e, 0x40, 0x5b, 0x48, 0x22,
	0x6b, 0x81, 0x97, 0x34, 0xb5, 0x4f, 0xaa, 0xe5, 0x67, 0x8c, 0x9e, 0x27, 0x25, 0x1f, 0xb1, 0xf4,
	0x02, 0xdf, 0xd0, 0x41, 0x50, 0x68, 0xa0, 0x89, 0x12, 0xd0, 0x8c, 0xc9, 0x46, 0xd0, 0x36, 0x02,
	0x85, 0x26, 0x82, 0x13, 0xce, 0x8a, 0x46, 0xe0, 0x19, 0x81, 0x42, 0x56, 0xb0, 0x05, 0x77, 0xed,
	0x94, 0x72, 0x56, 0x26, 0x23, 0x9e, 0x7e, 0xa4, 0x19, 0xf6, 0xf5, 0x6c, 0x57, 0x26, 0x81, 0x43,
	0xcd, 0xd1, 0x63, 0xb8, 0x93, 0x13, 0x91, 0x88, 0x7a, 0x78, 0x35, 0xb7, 0x40, 0x4b, 0x6f, 0xe7,
	0x44, 0xbc, 0x9b, 0x50, 0xb4, 0x3b, 0x59, 0x05, 0xd0, 0x83, 0x5d, 0xff, 0x47, 0x4b, 0xcd, 0x38,
	0xae, 0x36, 0x05, 0xad, 0x03, 0xa4, 0x15, 0x25, 0x92, 0x66, 0x09, 0x91, 0x78, 0xb9, 0xeb, 0xf4,
	0x5a, 0x71, 0x60, 0xc9, 0xbe, 0x0c, 0x4f, 0xa1, 0x6d, 0x97, 0xad, 0x69, 0xad, 0x33, 0xd5, 0xda,
	0x35, 0x08, 0x6a, 0x41, 0xab, 0x44, 0xcf, 0xc1, 0xf4, 0xdc, 0x57, 0xe0, 0xb5, 0x9a, 0xc5, 0x1a,
	0x04, 0x15, 0x25, 0x23, 0x13, 0x34, 0xcd, 0xf7, 0x15, 0xd0, 0xc1, 0x55, 0xf0, 0x99, 0x48, 0x48,
	0x36, 0x66, 0x85, 0x1e, 0x81, 0x1f, 0x7b, 0x4c, 0xec, 0xab, 0xc7, 0x9d, 0x9f, 0x2e, 0xf8, 0x83,
	0xe6, 0x5e, 0x5f, 0x1d, 0x58, 0x52, 0xbb, 0x83, 0x36, 0xaf, 0x5f, 0x91, 0xa9, 0x6d, 0xec, 0x3c,
	0x59, 0x44, 0x6a, 0x56, 0x31, 0xdc, 0xbd, 0xfc, 0x8e, 0x5d, 0xdf, 0xb9, 0xfc, 0xf1, 0xeb, 0x8b,
	0xbb, 0x85, 0x36, 0xa3, 0x64, 0xf6, 0x4b, 0xad, 0x65, 0x1e, 0xd9, 0xf4, 0xa8, 0xe9, 0x75, 0x74,
	0x22, 0x78, 0x81, 0xbe, 0x39, 0xe0, 0x1d, 0xd9, 0x1e, 0x46, 0xd7, 0x17, 0xfc, 0xeb, 0x7b, 0xe9,
	0x3c, 0x5b, 0x3c, 0xc1, 0xfa, 0x7c, 0x35, 0xe5, 0xf3, 0x05, 0xda, 0x5d, 0xc8, 0xe7, 0xa7, 0x9a,
	0x65, 0x9f, 0x23, 0x3b, 0x68, 0xed, 0xfa, 0x60, 0x0f, 0xd6, 0x53, 0x3e, 0x9e, 0xaa, 0x5e, 0xb2,
	0x19, 0x07, 0x07, 0x37, 0xad, 0x85, 0x81, 0xfa, 0x2f, 0x0d, 0x9c, 0x0f, 0x9e, 0x0d, 0x0c, 0xdb,
	0xfa, 0x4f, 0xf5, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x3b, 0xa6, 0x8f, 0x0f, 0x05,
	0x00, 0x00,
}
