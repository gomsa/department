// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/department/department.proto

package department

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Department struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string   `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Manages              string   `protobuf:"bytes,5,opt,name=manages,proto3" json:"manages,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{0}
}

func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (m *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(m, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Department) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Department) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Department) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Department) GetManages() string {
	if m != nil {
		return m.Manages
	}
	return ""
}

func (m *Department) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Department) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string   `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{1}
}

func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListQuery) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	Department           *Department   `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	Departments          []*Department `protobuf:"bytes,2,rep,name=departments,proto3" json:"departments,omitempty"`
	Total                int64         `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Valid                bool          `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetDepartment() *Department {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *Response) GetDepartments() []*Department {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Department)(nil), "department.Department")
	proto.RegisterType((*ListQuery)(nil), "department.ListQuery")
	proto.RegisterType((*Request)(nil), "department.Request")
	proto.RegisterType((*Response)(nil), "department.Response")
}

func init() { proto.RegisterFile("proto/department/department.proto", fileDescriptor_a231da9bf7b7101c) }

var fileDescriptor_a231da9bf7b7101c = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xfd, 0xf2, 0xd3, 0xb4, 0xb9, 0x81, 0x6f, 0x31, 0xf4, 0x2b, 0xc3, 0x07, 0x42, 0xcd, 0xca,
	0x55, 0x85, 0x8a, 0xd2, 0x8d, 0x8b, 0xd6, 0x82, 0x1b, 0x17, 0x3a, 0xe0, 0x5a, 0x46, 0x73, 0xa9,
	0x81, 0xfc, 0x99, 0x4c, 0x85, 0x3e, 0x80, 0xef, 0xe2, 0x13, 0xf8, 0x4c, 0x3e, 0x86, 0xcc, 0x4d,
	0xd2, 0xcc, 0xa2, 0x16, 0xda, 0xdd, 0x3d, 0xe7, 0xe4, 0x70, 0xef, 0xb9, 0x77, 0x02, 0xa7, 0x45,
	0x99, 0xab, 0xfc, 0x3c, 0xc2, 0x42, 0x96, 0x2a, 0xc5, 0x4c, 0x19, 0xe5, 0x84, 0x34, 0x06, 0x1d,
	0x13, 0x7e, 0x59, 0x00, 0xcb, 0x2d, 0x64, 0x7f, 0xc1, 0x8e, 0x23, 0x6e, 0x8d, 0xad, 0x33, 0x5f,
	0xd8, 0x71, 0xc4, 0x46, 0xe0, 0x15, 0xb2, 0xc4, 0x4c, 0x71, 0x9b, 0xb8, 0x06, 0x31, 0x06, 0x6e,
	0x26, 0x53, 0xe4, 0x0e, 0xb1, 0x54, 0xb3, 0x21, 0xf4, 0x8a, 0xd7, 0x3c, 0x43, 0xee, 0x12, 0x59,
	0x03, 0xc6, 0xa1, 0x9f, 0xca, 0x4c, 0xae, 0xb0, 0xe2, 0x3d, 0xe2, 0x5b, 0xc8, 0x4e, 0x00, 0x5e,
	0x4a, 0x94, 0x0a, 0xa3, 0x27, 0xa9, 0xb8, 0x47, 0xa2, 0xdf, 0x30, 0x73, 0xa5, 0xe5, 0x75, 0x11,
	0xb5, 0x72, 0xbf, 0x96, 0x1b, 0x66, 0xae, 0xc2, 0x0f, 0x0b, 0xfc, 0xbb, 0xb8, 0x52, 0x0f, 0x6b,
	0x2c, 0x37, 0xba, 0x77, 0x12, 0xa7, 0xb1, 0xa2, 0xd1, 0x1d, 0x51, 0x03, 0x3d, 0x65, 0x21, 0x57,
	0x48, 0xb3, 0x3b, 0x82, 0x6a, 0xcd, 0x55, 0x79, 0xa9, 0xda, 0xc9, 0x75, 0xdd, 0xa4, 0x76, 0x77,
	0xa4, 0xee, 0xed, 0x4c, 0xed, 0x75, 0xa9, 0x43, 0x1f, 0xfa, 0x02, 0xdf, 0xd6, 0x58, 0xa9, 0xf0,
	0xd3, 0x82, 0x81, 0xc0, 0xaa, 0xc8, 0xb3, 0x0a, 0xd9, 0x15, 0x18, 0x6b, 0xa6, 0xb1, 0x82, 0xe9,
	0x68, 0x62, 0xdc, 0xa2, 0xdb, 0xba, 0x30, 0xbe, 0x64, 0x33, 0x08, 0x3a, 0x54, 0x71, 0x7b, 0xec,
	0xec, 0x31, 0x9a, 0x9f, 0xea, 0x1d, 0xa8, 0x5c, 0xc9, 0x84, 0xa2, 0x39, 0xa2, 0x06, 0x9a, 0x7d,
	0x97, 0x49, 0x13, 0x6f, 0x20, 0x6a, 0x30, 0xfd, 0xb6, 0x21, 0x58, 0x1a, 0xde, 0x6b, 0x08, 0x16,
	0xa8, 0x16, 0x9b, 0xfb, 0x3a, 0xe8, 0x2f, 0xfd, 0xfe, 0x0f, 0x4d, 0xbe, 0x8d, 0x1a, 0xfe, 0x61,
	0x97, 0xe0, 0xea, 0x5b, 0xb0, 0x7f, 0xa6, 0xbe, 0xbd, 0xce, 0x1e, 0x9b, 0x73, 0x8b, 0x87, 0x77,
	0x9b, 0x81, 0x77, 0x43, 0xcf, 0xe4, 0x18, 0xe7, 0x23, 0xbd, 0xa0, 0x63, 0x9c, 0x4b, 0x4c, 0xf0,
	0x70, 0xe7, 0xb3, 0x47, 0x3f, 0xdd, 0xc5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x88, 0x04,
	0x60, 0x99, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Departments service

type DepartmentsClient interface {
	// 根据父级 id 获取相关部门
	BetByParent(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error)
	// 获取部门列表
	List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取部门
	Get(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error)
	// 创建部门
	Create(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error)
	// 更新部门
	Update(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error)
	// 删除部门
	Delete(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error)
}

type departmentsClient struct {
	c           client.Client
	serviceName string
}

func NewDepartmentsClient(serviceName string, c client.Client) DepartmentsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "department"
	}
	return &departmentsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *departmentsClient) BetByParent(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.BetByParent", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Get(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Create(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Update(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Delete(ctx context.Context, in *Department, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Departments service

type DepartmentsHandler interface {
	// 根据父级 id 获取相关部门
	BetByParent(context.Context, *Department, *Response) error
	// 获取部门列表
	List(context.Context, *ListQuery, *Response) error
	// 根据 唯一 获取部门
	Get(context.Context, *Department, *Response) error
	// 创建部门
	Create(context.Context, *Department, *Response) error
	// 更新部门
	Update(context.Context, *Department, *Response) error
	// 删除部门
	Delete(context.Context, *Department, *Response) error
}

func RegisterDepartmentsHandler(s server.Server, hdlr DepartmentsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Departments{hdlr}, opts...))
}

type Departments struct {
	DepartmentsHandler
}

func (h *Departments) BetByParent(ctx context.Context, in *Department, out *Response) error {
	return h.DepartmentsHandler.BetByParent(ctx, in, out)
}

func (h *Departments) List(ctx context.Context, in *ListQuery, out *Response) error {
	return h.DepartmentsHandler.List(ctx, in, out)
}

func (h *Departments) Get(ctx context.Context, in *Department, out *Response) error {
	return h.DepartmentsHandler.Get(ctx, in, out)
}

func (h *Departments) Create(ctx context.Context, in *Department, out *Response) error {
	return h.DepartmentsHandler.Create(ctx, in, out)
}

func (h *Departments) Update(ctx context.Context, in *Department, out *Response) error {
	return h.DepartmentsHandler.Update(ctx, in, out)
}

func (h *Departments) Delete(ctx context.Context, in *Department, out *Response) error {
	return h.DepartmentsHandler.Delete(ctx, in, out)
}