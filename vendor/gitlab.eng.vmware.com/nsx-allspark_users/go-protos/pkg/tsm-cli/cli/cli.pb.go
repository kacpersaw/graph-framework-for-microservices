// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/tsm-cli/cli.proto

package cli // import "tsm-cli/cli"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

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

// Severity of the message
type Severity int32

const (
	Severity_INVALID Severity = 0
	Severity_INFO    Severity = 1
	Severity_ERROR   Severity = 2
)

var Severity_name = map[int32]string{
	0: "INVALID",
	1: "INFO",
	2: "ERROR",
}
var Severity_value = map[string]int32{
	"INVALID": 0,
	"INFO":    1,
	"ERROR":   2,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}
func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{0}
}

// Path represents a link to an api object in the cli chain
type Path struct {
	// kind represents a resource type.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// id represents object id.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// path represents a link to an api resource
	Path                 *_struct.Value `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{0}
}
func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (dst *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(dst, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Path) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Path) GetPath() *_struct.Value {
	if m != nil {
		return m.Path
	}
	return nil
}

// Yaml represents the output
type File struct {
	File                 []byte   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{1}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

// Response of an API request
type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Severity             Severity `protobuf:"varint,2,opt,name=severity,proto3,enum=cli.Severity" json:"severity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetSeverity() Severity {
	if m != nil {
		return m.Severity
	}
	return Severity_INVALID
}

// Get query for the resource identified by the path
type GetRequest struct {
	Path                 *Path    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

// Get response to the query
type GetResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	File                 *File     `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{4}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GetResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

// List request for the resource identified by the path
type ListRequest struct {
	Path                 *Path    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{5}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

// List response for the resourse which can be one or many
type ListResponse struct {
	Response             []*Path  `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{6}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetResponse() []*Path {
	if m != nil {
		return m.Response
	}
	return nil
}

// Create or Update request for the resource
type UpsertRequest struct {
	File                 *File    `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertRequest) Reset()         { *m = UpsertRequest{} }
func (m *UpsertRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertRequest) ProtoMessage()    {}
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{7}
}
func (m *UpsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertRequest.Unmarshal(m, b)
}
func (m *UpsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertRequest.Marshal(b, m, deterministic)
}
func (dst *UpsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertRequest.Merge(dst, src)
}
func (m *UpsertRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertRequest.Size(m)
}
func (m *UpsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertRequest proto.InternalMessageInfo

func (m *UpsertRequest) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

// Upsert response for the resource
type UpsertResponse struct {
	Response             []*Response `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpsertResponse) Reset()         { *m = UpsertResponse{} }
func (m *UpsertResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertResponse) ProtoMessage()    {}
func (*UpsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{8}
}
func (m *UpsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertResponse.Unmarshal(m, b)
}
func (m *UpsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertResponse.Marshal(b, m, deterministic)
}
func (dst *UpsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertResponse.Merge(dst, src)
}
func (m *UpsertResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertResponse.Size(m)
}
func (m *UpsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertResponse proto.InternalMessageInfo

func (m *UpsertResponse) GetResponse() []*Response {
	if m != nil {
		return m.Response
	}
	return nil
}

// Delete request for the resource.
type DeleteRequest struct {
	// Types that are valid to be assigned to Request:
	//	*DeleteRequest_File
	//	*DeleteRequest_Path
	Request              isDeleteRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{9}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

type isDeleteRequest_Request interface {
	isDeleteRequest_Request()
}

type DeleteRequest_File struct {
	File *File `protobuf:"bytes,1,opt,name=file,proto3,oneof"`
}

type DeleteRequest_Path struct {
	Path *Path `protobuf:"bytes,2,opt,name=path,proto3,oneof"`
}

func (*DeleteRequest_File) isDeleteRequest_Request() {}

func (*DeleteRequest_Path) isDeleteRequest_Request() {}

func (m *DeleteRequest) GetRequest() isDeleteRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *DeleteRequest) GetFile() *File {
	if x, ok := m.GetRequest().(*DeleteRequest_File); ok {
		return x.File
	}
	return nil
}

func (m *DeleteRequest) GetPath() *Path {
	if x, ok := m.GetRequest().(*DeleteRequest_Path); ok {
		return x.Path
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeleteRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeleteRequest_OneofMarshaler, _DeleteRequest_OneofUnmarshaler, _DeleteRequest_OneofSizer, []interface{}{
		(*DeleteRequest_File)(nil),
		(*DeleteRequest_Path)(nil),
	}
}

func _DeleteRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeleteRequest)
	// request
	switch x := m.Request.(type) {
	case *DeleteRequest_File:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case *DeleteRequest_Path:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Path); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeleteRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _DeleteRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeleteRequest)
	switch tag {
	case 1: // request.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(File)
		err := b.DecodeMessage(msg)
		m.Request = &DeleteRequest_File{msg}
		return true, err
	case 2: // request.path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Path)
		err := b.DecodeMessage(msg)
		m.Request = &DeleteRequest_Path{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeleteRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeleteRequest)
	// request
	switch x := m.Request.(type) {
	case *DeleteRequest_File:
		s := proto.Size(x.File)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DeleteRequest_Path:
		s := proto.Size(x.Path)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Delete response for the resource
type DeleteResponse struct {
	Response             []*Response `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_d2908b4d0ceb18c8, []int{10}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetResponse() []*Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Path)(nil), "cli.Path")
	proto.RegisterType((*File)(nil), "cli.File")
	proto.RegisterType((*Response)(nil), "cli.Response")
	proto.RegisterType((*GetRequest)(nil), "cli.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "cli.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "cli.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "cli.ListResponse")
	proto.RegisterType((*UpsertRequest)(nil), "cli.UpsertRequest")
	proto.RegisterType((*UpsertResponse)(nil), "cli.UpsertResponse")
	proto.RegisterType((*DeleteRequest)(nil), "cli.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "cli.DeleteResponse")
	proto.RegisterEnum("cli.Severity", Severity_name, Severity_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CliClient is the client API for Cli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CliClient interface {
	// CRUD methods to query the api resources
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type cliClient struct {
	cc *grpc.ClientConn
}

func NewCliClient(cc *grpc.ClientConn) CliClient {
	return &cliClient{cc}
}

func (c *cliClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/cli.Cli/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/cli.Cli/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertResponse, error) {
	out := new(UpsertResponse)
	err := c.cc.Invoke(ctx, "/cli.Cli/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/cli.Cli/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CliServer is the server API for Cli service.
type CliServer interface {
	// CRUD methods to query the api resources
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Upsert(context.Context, *UpsertRequest) (*UpsertResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterCliServer(s *grpc.Server, srv CliServer) {
	s.RegisterService(&_Cli_serviceDesc, srv)
}

func _Cli_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.Cli/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.Cli/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.Cli/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.Cli/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cli_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cli.Cli",
	HandlerType: (*CliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Cli_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Cli_List_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _Cli_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Cli_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/tsm-cli/cli.proto",
}

func init() { proto.RegisterFile("protos/tsm-cli/cli.proto", fileDescriptor_cli_d2908b4d0ceb18c8) }

var fileDescriptor_cli_d2908b4d0ceb18c8 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0xcd, 0x7e, 0xd8, 0x24, 0x77, 0x9b, 0x18, 0xaf, 0x20, 0xcb, 0xa2, 0x58, 0x16, 0x84, 0x36,
	0xb6, 0x1b, 0x48, 0xf1, 0xc9, 0x27, 0x6b, 0x6d, 0x0d, 0x94, 0x46, 0x46, 0x8c, 0x20, 0xbe, 0xa4,
	0xc9, 0x6d, 0x3a, 0x38, 0xed, 0xc6, 0x9d, 0x89, 0xe0, 0x5f, 0xf4, 0x57, 0xc9, 0x7c, 0xec, 0x76,
	0x57, 0xaa, 0xd8, 0xb7, 0xd9, 0x73, 0xcf, 0x3d, 0xf7, 0x9c, 0x3b, 0x3b, 0x10, 0xaf, 0x8b, 0x5c,
	0xe5, 0x72, 0xa4, 0xe4, 0xf5, 0xc1, 0x42, 0xf0, 0xd1, 0x42, 0xf0, 0xcc, 0x40, 0x18, 0x2c, 0x04,
	0x4f, 0x9e, 0xae, 0xf2, 0x7c, 0x25, 0x68, 0x64, 0xa0, 0x8b, 0xcd, 0xe5, 0x48, 0xaa, 0x62, 0xb3,
	0x50, 0x96, 0x92, 0xce, 0x20, 0xfc, 0x30, 0x57, 0x57, 0x88, 0x10, 0x7e, 0xe3, 0x37, 0xcb, 0xd8,
	0xdb, 0xf1, 0x76, 0xbb, 0xcc, 0x9c, 0xb1, 0x0f, 0x3e, 0x5f, 0xc6, 0xbe, 0x41, 0x7c, 0xbe, 0xc4,
	0x21, 0x84, 0xeb, 0xb9, 0xba, 0x8a, 0x83, 0x1d, 0x6f, 0x37, 0x1a, 0x3f, 0xc9, 0xac, 0x70, 0x56,
	0x0a, 0x67, 0xb3, 0xb9, 0xd8, 0x10, 0x33, 0x9c, 0x34, 0x81, 0xf0, 0x84, 0x0b, 0xd2, 0xba, 0x97,
	0x5c, 0x90, 0xd1, 0xdd, 0x66, 0xe6, 0x9c, 0x4e, 0xa1, 0xc3, 0x48, 0xae, 0xf3, 0x1b, 0x49, 0x18,
	0x43, 0xfb, 0x9a, 0xa4, 0x9c, 0xaf, 0xc8, 0x8d, 0x2e, 0x3f, 0x71, 0x0f, 0x3a, 0x92, 0x7e, 0x50,
	0xc1, 0xd5, 0x4f, 0xe3, 0xa1, 0x3f, 0xee, 0x65, 0x3a, 0xda, 0x47, 0x07, 0xb2, 0xaa, 0x9c, 0xbe,
	0x04, 0x38, 0x25, 0xc5, 0xe8, 0xfb, 0x86, 0xa4, 0xc2, 0x67, 0xce, 0xa6, 0x67, 0x6c, 0x76, 0x4d,
	0x93, 0xce, 0xe8, 0x9c, 0x7d, 0x86, 0xc8, 0x90, 0x9d, 0x81, 0x3d, 0xe8, 0x14, 0xee, 0xec, 0x3a,
	0xec, 0x98, 0x92, 0xc0, 0xaa, 0xb2, 0x16, 0x36, 0x59, 0xfc, 0x9a, 0xb0, 0x0e, 0xe9, 0x62, 0xed,
	0x43, 0x74, 0xc6, 0xe5, 0xff, 0xda, 0x78, 0x05, 0xdb, 0x96, 0xed, 0xc4, 0x5f, 0x34, 0x7c, 0x04,
	0xcd, 0x96, 0xaa, 0x94, 0x66, 0xd0, 0xfb, 0xb4, 0x96, 0x54, 0xd4, 0xc7, 0x54, 0x0b, 0xbe, 0xc3,
	0xd4, 0x6b, 0xe8, 0x97, 0xfc, 0x3b, 0x03, 0x07, 0xff, 0x08, 0x9c, 0x7e, 0x85, 0xde, 0x31, 0x09,
	0x52, 0x54, 0x0e, 0x7b, 0xfe, 0x97, 0x61, 0xef, 0x5b, 0x76, 0x9c, 0x26, 0x98, 0xd0, 0xfe, 0x1f,
	0xa1, 0x35, 0x41, 0x17, 0x8e, 0xba, 0xd0, 0x2e, 0xac, 0x98, 0xb6, 0x56, 0xaa, 0xdf, 0xdb, 0xda,
	0x70, 0x1f, 0x3a, 0xe5, 0x8f, 0x80, 0x11, 0xb4, 0x27, 0xe7, 0xb3, 0x37, 0x67, 0x93, 0xe3, 0x41,
	0x0b, 0x3b, 0x10, 0x4e, 0xce, 0x4f, 0xa6, 0x03, 0x0f, 0xbb, 0xf0, 0xe0, 0x1d, 0x63, 0x53, 0x36,
	0xf0, 0xc7, 0xbf, 0x3c, 0x08, 0xde, 0x0a, 0x8e, 0x43, 0x08, 0x4e, 0x49, 0xe1, 0x43, 0xa3, 0x7a,
	0xfb, 0xcb, 0x24, 0x83, 0x5b, 0xc0, 0x45, 0x6f, 0xe1, 0x01, 0x84, 0xfa, 0x82, 0xd0, 0xd6, 0x6a,
	0x37, 0x9b, 0x3c, 0xaa, 0x21, 0x15, 0xfd, 0x10, 0xb6, 0xec, 0xa2, 0x11, 0x4d, 0xb9, 0x71, 0x4b,
	0xc9, 0xe3, 0x06, 0x56, 0x6f, 0xb2, 0x2b, 0x70, 0x4d, 0x8d, 0x6d, 0xbb, 0xa6, 0xe6, 0x8e, 0xd2,
	0xd6, 0x51, 0xef, 0x4b, 0x54, 0x7b, 0xea, 0x17, 0x5b, 0xe6, 0xfd, 0x1d, 0xfe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x54, 0xfc, 0xe8, 0x07, 0x04, 0x00, 0x00,
}
