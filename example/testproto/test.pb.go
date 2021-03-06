// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testproto/test.proto

package testproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetArticlesReq struct {
	// @inject_tag: form:"title"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" form:"title"`
	// @inject_tag: form:"page"
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	// @inject_tag: form:"page_size"
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorId             int32    `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty" form:"author_id" uri:"author_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArticlesReq) Reset()         { *m = GetArticlesReq{} }
func (m *GetArticlesReq) String() string { return proto.CompactTextString(m) }
func (*GetArticlesReq) ProtoMessage()    {}
func (*GetArticlesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111ee603e695d6e, []int{0}
}

func (m *GetArticlesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticlesReq.Unmarshal(m, b)
}
func (m *GetArticlesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticlesReq.Marshal(b, m, deterministic)
}
func (m *GetArticlesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticlesReq.Merge(m, src)
}
func (m *GetArticlesReq) XXX_Size() int {
	return xxx_messageInfo_GetArticlesReq.Size(m)
}
func (m *GetArticlesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticlesReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticlesReq proto.InternalMessageInfo

func (m *GetArticlesReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetArticlesReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetArticlesReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetArticlesReq) GetAuthorId() int32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type GetArticlesResp struct {
	Total                int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Articles             []*Article `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetArticlesResp) Reset()         { *m = GetArticlesResp{} }
func (m *GetArticlesResp) String() string { return proto.CompactTextString(m) }
func (*GetArticlesResp) ProtoMessage()    {}
func (*GetArticlesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111ee603e695d6e, []int{1}
}

func (m *GetArticlesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticlesResp.Unmarshal(m, b)
}
func (m *GetArticlesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticlesResp.Marshal(b, m, deterministic)
}
func (m *GetArticlesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticlesResp.Merge(m, src)
}
func (m *GetArticlesResp) XXX_Size() int {
	return xxx_messageInfo_GetArticlesResp.Size(m)
}
func (m *GetArticlesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticlesResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticlesResp proto.InternalMessageInfo

func (m *GetArticlesResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetArticlesResp) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorId             int32    `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty" form:"author_id" uri:"author_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111ee603e695d6e, []int{2}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetAuthorId() int32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111ee603e695d6e, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetArticlesReq)(nil), "testproto.GetArticlesReq")
	proto.RegisterType((*GetArticlesResp)(nil), "testproto.GetArticlesResp")
	proto.RegisterType((*Article)(nil), "testproto.Article")
	proto.RegisterType((*Empty)(nil), "testproto.Empty")
}

func init() { proto.RegisterFile("testproto/test.proto", fileDescriptor_a111ee603e695d6e) }

var fileDescriptor_a111ee603e695d6e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0x4d, 0x6f, 0xda, 0x40,
	0x14, 0x94, 0x31, 0x14, 0xbc, 0xf4, 0x4b, 0x2b, 0x0e, 0xae, 0x5b, 0xa9, 0xe0, 0x4b, 0xb9, 0xe0,
	0x15, 0xf4, 0xd0, 0x43, 0x4f, 0xa5, 0x89, 0xa2, 0x5c, 0x8d, 0x94, 0x48, 0x1c, 0x82, 0x16, 0xf3,
	0xb4, 0x5e, 0xc9, 0xf6, 0x3a, 0xf6, 0x33, 0x4a, 0x88, 0xb8, 0xe4, 0x2f, 0xe4, 0xa7, 0xe5, 0x2f,
	0xe4, 0x98, 0x1f, 0x11, 0xb1, 0x06, 0x0b, 0xa2, 0xa0, 0x9c, 0xf6, 0xbd, 0x99, 0xd1, 0x9b, 0x99,
	0x25, 0x1d, 0x84, 0x1c, 0xd3, 0x4c, 0xa1, 0x62, 0x9b, 0xc9, 0xd3, 0x23, 0xb5, 0x2a, 0xd4, 0xf9,
	0x21, 0x94, 0x12, 0x11, 0x30, 0x9e, 0x4a, 0xc6, 0x93, 0x44, 0x21, 0x47, 0xa9, 0x92, 0xbc, 0x14,
	0xba, 0x48, 0x3e, 0x9f, 0x01, 0xfe, 0xcb, 0x50, 0x06, 0x11, 0xe4, 0x3e, 0x5c, 0xd3, 0x0e, 0x69,
	0xa0, 0xc4, 0x08, 0x6c, 0xa3, 0x6b, 0xf4, 0x2d, 0xbf, 0x5c, 0x28, 0x25, 0xf5, 0x94, 0x0b, 0xb0,
	0x6b, 0x5d, 0xa3, 0xdf, 0xf0, 0xf5, 0x4c, 0xbf, 0x13, 0x6b, 0xf3, 0xce, 0x72, 0xb9, 0x02, 0xdb,
	0xd4, 0x44, 0x6b, 0x03, 0x4c, 0xe4, 0x4a, 0x93, 0xbc, 0xc0, 0x50, 0x65, 0x33, 0xb9, 0xb0, 0xeb,
	0x25, 0x59, 0x02, 0xe7, 0x0b, 0xf7, 0x92, 0x7c, 0x39, 0x70, 0xcd, 0x53, 0x6d, 0xab, 0x90, 0x47,
	0xda, 0xd6, 0xf4, 0xcb, 0x85, 0x7a, 0xa4, 0xc5, 0xb7, 0x2a, 0xbb, 0xd6, 0x35, 0xfb, 0xed, 0x11,
	0xf5, 0xaa, 0x6a, 0xde, 0xf6, 0x80, 0x5f, 0x69, 0xdc, 0x0b, 0xd2, 0xdc, 0x82, 0x47, 0x7a, 0xd8,
	0xa4, 0x19, 0xa8, 0x04, 0x21, 0x41, 0x5d, 0xc5, 0xf2, 0x77, 0xeb, 0x61, 0x60, 0xf3, 0x55, 0xe0,
	0x26, 0x69, 0x9c, 0xc6, 0x29, 0xde, 0x8e, 0x9e, 0x0d, 0xd2, 0x1e, 0x47, 0x4a, 0x4c, 0x20, 0x5b,
	0xca, 0x00, 0xe8, 0x9a, 0xb4, 0xf7, 0x9a, 0xd0, 0x6f, 0x7b, 0xe9, 0x0e, 0xff, 0xd5, 0x71, 0x8e,
	0x51, 0x79, 0xea, 0xfe, 0xb9, 0x7f, 0x7c, 0x7a, 0xa8, 0x0d, 0xe9, 0x47, 0xb6, 0x1c, 0xb2, 0x5d,
	0x99, 0x69, 0x8f, 0xfe, 0xd4, 0xbb, 0x0e, 0xc1, 0xee, 0xaa, 0x74, 0xeb, 0x4a, 0x42, 0xaf, 0xc8,
	0xa7, 0xff, 0x19, 0x70, 0x84, 0x5d, 0xeb, 0x37, 0xbe, 0xc7, 0xf9, 0xba, 0x87, 0xe9, 0x16, 0xee,
	0x2f, 0xed, 0xd7, 0x73, 0xdf, 0xbb, 0x3f, 0x3e, 0x99, 0x8e, 0x85, 0xc4, 0xb0, 0x98, 0x7b, 0x81,
	0x8a, 0x59, 0xac, 0xc2, 0x42, 0xe6, 0xa1, 0x2a, 0x98, 0x3e, 0x16, 0x0c, 0x04, 0x24, 0x03, 0xa1,
	0x06, 0x42, 0x26, 0x0c, 0x6e, 0x78, 0x9c, 0x46, 0xc0, 0x2a, 0xab, 0xbf, 0xd5, 0x34, 0xff, 0xa0,
	0x9f, 0xdf, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x25, 0x6f, 0x42, 0xac, 0x02, 0x00, 0x00,
}
