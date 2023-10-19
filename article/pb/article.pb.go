// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article/article.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Article struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author               string   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{0}
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

func (m *Article) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

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

type ArticleInput struct {
	Author               string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleInput) Reset()         { *m = ArticleInput{} }
func (m *ArticleInput) String() string { return proto.CompactTextString(m) }
func (*ArticleInput) ProtoMessage()    {}
func (*ArticleInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{1}
}

func (m *ArticleInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleInput.Unmarshal(m, b)
}
func (m *ArticleInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleInput.Marshal(b, m, deterministic)
}
func (m *ArticleInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleInput.Merge(m, src)
}
func (m *ArticleInput) XXX_Size() int {
	return xxx_messageInfo_ArticleInput.Size(m)
}
func (m *ArticleInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleInput.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleInput proto.InternalMessageInfo

func (m *ArticleInput) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ArticleInput) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleInput) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateArticleRequest struct {
	ArticleInput         *ArticleInput `protobuf:"bytes,1,opt,name=articleInput,proto3" json:"articleInput,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateArticleRequest) Reset()         { *m = CreateArticleRequest{} }
func (m *CreateArticleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArticleRequest) ProtoMessage()    {}
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{2}
}

func (m *CreateArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArticleRequest.Unmarshal(m, b)
}
func (m *CreateArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArticleRequest.Marshal(b, m, deterministic)
}
func (m *CreateArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArticleRequest.Merge(m, src)
}
func (m *CreateArticleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArticleRequest.Size(m)
}
func (m *CreateArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArticleRequest proto.InternalMessageInfo

func (m *CreateArticleRequest) GetArticleInput() *ArticleInput {
	if m != nil {
		return m.ArticleInput
	}
	return nil
}

type CreateArticleResponse struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateArticleResponse) Reset()         { *m = CreateArticleResponse{} }
func (m *CreateArticleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateArticleResponse) ProtoMessage()    {}
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{3}
}

func (m *CreateArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArticleResponse.Unmarshal(m, b)
}
func (m *CreateArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArticleResponse.Marshal(b, m, deterministic)
}
func (m *CreateArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArticleResponse.Merge(m, src)
}
func (m *CreateArticleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateArticleResponse.Size(m)
}
func (m *CreateArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArticleResponse proto.InternalMessageInfo

func (m *CreateArticleResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type ReadArticleRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadArticleRequest) Reset()         { *m = ReadArticleRequest{} }
func (m *ReadArticleRequest) String() string { return proto.CompactTextString(m) }
func (*ReadArticleRequest) ProtoMessage()    {}
func (*ReadArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{4}
}

func (m *ReadArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadArticleRequest.Unmarshal(m, b)
}
func (m *ReadArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadArticleRequest.Marshal(b, m, deterministic)
}
func (m *ReadArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadArticleRequest.Merge(m, src)
}
func (m *ReadArticleRequest) XXX_Size() int {
	return xxx_messageInfo_ReadArticleRequest.Size(m)
}
func (m *ReadArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadArticleRequest proto.InternalMessageInfo

func (m *ReadArticleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadArticleResponse struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadArticleResponse) Reset()         { *m = ReadArticleResponse{} }
func (m *ReadArticleResponse) String() string { return proto.CompactTextString(m) }
func (*ReadArticleResponse) ProtoMessage()    {}
func (*ReadArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{5}
}

func (m *ReadArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadArticleResponse.Unmarshal(m, b)
}
func (m *ReadArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadArticleResponse.Marshal(b, m, deterministic)
}
func (m *ReadArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadArticleResponse.Merge(m, src)
}
func (m *ReadArticleResponse) XXX_Size() int {
	return xxx_messageInfo_ReadArticleResponse.Size(m)
}
func (m *ReadArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadArticleResponse proto.InternalMessageInfo

func (m *ReadArticleResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type UpdateArticleRequest struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ArticleInput         *ArticleInput `protobuf:"bytes,2,opt,name=articleInput,proto3" json:"articleInput,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateArticleRequest) Reset()         { *m = UpdateArticleRequest{} }
func (m *UpdateArticleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateArticleRequest) ProtoMessage()    {}
func (*UpdateArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{6}
}

func (m *UpdateArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateArticleRequest.Unmarshal(m, b)
}
func (m *UpdateArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateArticleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateArticleRequest.Merge(m, src)
}
func (m *UpdateArticleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateArticleRequest.Size(m)
}
func (m *UpdateArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateArticleRequest proto.InternalMessageInfo

func (m *UpdateArticleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateArticleRequest) GetArticleInput() *ArticleInput {
	if m != nil {
		return m.ArticleInput
	}
	return nil
}

type UpdateArticleResponse struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateArticleResponse) Reset()         { *m = UpdateArticleResponse{} }
func (m *UpdateArticleResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateArticleResponse) ProtoMessage()    {}
func (*UpdateArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{7}
}

func (m *UpdateArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateArticleResponse.Unmarshal(m, b)
}
func (m *UpdateArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateArticleResponse.Marshal(b, m, deterministic)
}
func (m *UpdateArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateArticleResponse.Merge(m, src)
}
func (m *UpdateArticleResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateArticleResponse.Size(m)
}
func (m *UpdateArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateArticleResponse proto.InternalMessageInfo

func (m *UpdateArticleResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type DeleteArticleRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArticleRequest) Reset()         { *m = DeleteArticleRequest{} }
func (m *DeleteArticleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArticleRequest) ProtoMessage()    {}
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{8}
}

func (m *DeleteArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArticleRequest.Unmarshal(m, b)
}
func (m *DeleteArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArticleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArticleRequest.Merge(m, src)
}
func (m *DeleteArticleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteArticleRequest.Size(m)
}
func (m *DeleteArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArticleRequest proto.InternalMessageInfo

func (m *DeleteArticleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteArticleResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArticleResponse) Reset()         { *m = DeleteArticleResponse{} }
func (m *DeleteArticleResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteArticleResponse) ProtoMessage()    {}
func (*DeleteArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{9}
}

func (m *DeleteArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArticleResponse.Unmarshal(m, b)
}
func (m *DeleteArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArticleResponse.Marshal(b, m, deterministic)
}
func (m *DeleteArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArticleResponse.Merge(m, src)
}
func (m *DeleteArticleResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteArticleResponse.Size(m)
}
func (m *DeleteArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArticleResponse proto.InternalMessageInfo

func (m *DeleteArticleResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListArticleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArticleRequest) Reset()         { *m = ListArticleRequest{} }
func (m *ListArticleRequest) String() string { return proto.CompactTextString(m) }
func (*ListArticleRequest) ProtoMessage()    {}
func (*ListArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{10}
}

func (m *ListArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArticleRequest.Unmarshal(m, b)
}
func (m *ListArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArticleRequest.Marshal(b, m, deterministic)
}
func (m *ListArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArticleRequest.Merge(m, src)
}
func (m *ListArticleRequest) XXX_Size() int {
	return xxx_messageInfo_ListArticleRequest.Size(m)
}
func (m *ListArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArticleRequest proto.InternalMessageInfo

type ListArticleResponse struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArticleResponse) Reset()         { *m = ListArticleResponse{} }
func (m *ListArticleResponse) String() string { return proto.CompactTextString(m) }
func (*ListArticleResponse) ProtoMessage()    {}
func (*ListArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12321969c539ac1b, []int{11}
}

func (m *ListArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArticleResponse.Unmarshal(m, b)
}
func (m *ListArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArticleResponse.Marshal(b, m, deterministic)
}
func (m *ListArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArticleResponse.Merge(m, src)
}
func (m *ListArticleResponse) XXX_Size() int {
	return xxx_messageInfo_ListArticleResponse.Size(m)
}
func (m *ListArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArticleResponse proto.InternalMessageInfo

func (m *ListArticleResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

func init() {
	proto.RegisterType((*Article)(nil), "article.Article")
	proto.RegisterType((*ArticleInput)(nil), "article.ArticleInput")
	proto.RegisterType((*CreateArticleRequest)(nil), "article.CreateArticleRequest")
	proto.RegisterType((*CreateArticleResponse)(nil), "article.CreateArticleResponse")
	proto.RegisterType((*ReadArticleRequest)(nil), "article.ReadArticleRequest")
	proto.RegisterType((*ReadArticleResponse)(nil), "article.ReadArticleResponse")
	proto.RegisterType((*UpdateArticleRequest)(nil), "article.UpdateArticleRequest")
	proto.RegisterType((*UpdateArticleResponse)(nil), "article.UpdateArticleResponse")
	proto.RegisterType((*DeleteArticleRequest)(nil), "article.DeleteArticleRequest")
	proto.RegisterType((*DeleteArticleResponse)(nil), "article.DeleteArticleResponse")
	proto.RegisterType((*ListArticleRequest)(nil), "article.ListArticleRequest")
	proto.RegisterType((*ListArticleResponse)(nil), "article.ListArticleResponse")
}

func init() { proto.RegisterFile("article/article.proto", fileDescriptor_12321969c539ac1b) }

var fileDescriptor_12321969c539ac1b = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x89, 0xb6, 0x38, 0xfd, 0x40, 0xc6, 0x44, 0x42, 0xfd, 0xa0, 0x04, 0xd1, 0xe2, 0xa1,
	0x4a, 0x3d, 0x79, 0xac, 0xf5, 0xa0, 0x22, 0x05, 0x23, 0x7a, 0xf0, 0xb6, 0x6d, 0x06, 0x0c, 0x94,
	0x24, 0x26, 0x5b, 0x7f, 0x8a, 0xbf, 0x57, 0x4c, 0x37, 0x65, 0x77, 0xbb, 0x45, 0xda, 0x53, 0x3b,
	0xb3, 0x93, 0xf7, 0xde, 0xbe, 0x37, 0x2c, 0x78, 0x2c, 0xe7, 0xf1, 0x74, 0x46, 0x57, 0xe2, 0xb7,
	0x9f, 0xe5, 0x29, 0x4f, 0xb1, 0x2e, 0xca, 0x80, 0x41, 0x7d, 0xb8, 0xf8, 0x8b, 0x6d, 0xb0, 0xe3,
	0xc8, 0xb7, 0xba, 0x56, 0xcf, 0x09, 0xed, 0x38, 0xc2, 0x43, 0xa8, 0xb1, 0x39, 0xff, 0x4c, 0x73,
	0xdf, 0xee, 0x5a, 0xbd, 0xbd, 0x50, 0x54, 0xe8, 0xc2, 0x2e, 0x8f, 0xf9, 0x8c, 0x7c, 0xa7, 0x6c,
	0x2f, 0x0a, 0xf4, 0xa1, 0x3e, 0x4d, 0x13, 0x4e, 0x09, 0xf7, 0x77, 0xca, 0x7e, 0x55, 0x06, 0xef,
	0xd0, 0x14, 0x14, 0x8f, 0x49, 0x36, 0xe7, 0x12, 0xae, 0x65, 0xc6, 0xb5, 0xd7, 0xe0, 0x3a, 0x2a,
	0xee, 0x0b, 0xb8, 0xa3, 0x9c, 0x18, 0x27, 0x81, 0x1e, 0xd2, 0xd7, 0x9c, 0x0a, 0x8e, 0xb7, 0xd0,
	0x64, 0x12, 0x5f, 0xc9, 0xd2, 0x18, 0x78, 0xfd, 0xca, 0x01, 0x59, 0x4c, 0xa8, 0x8c, 0x06, 0x23,
	0xf0, 0x34, 0xc8, 0x22, 0x4b, 0x93, 0x82, 0xf0, 0x12, 0x2a, 0xc7, 0x04, 0xdc, 0xbe, 0x0e, 0x17,
	0x2e, 0x2d, 0x3d, 0x03, 0x0c, 0x89, 0x45, 0x9a, 0x2a, 0xcd, 0xdd, 0x60, 0x08, 0x07, 0xca, 0xd4,
	0x16, 0x44, 0x0c, 0xdc, 0xb7, 0x2c, 0x5a, 0x35, 0x40, 0x0f, 0x52, 0x37, 0xc4, 0xde, 0xc8, 0x10,
	0x8d, 0x62, 0x0b, 0x9d, 0xe7, 0xe0, 0xde, 0xd3, 0x8c, 0xfe, 0xd3, 0x19, 0x5c, 0x80, 0xa7, 0xcd,
	0x09, 0x32, 0x7d, 0xd0, 0x05, 0x7c, 0x8e, 0x0b, 0xae, 0xc2, 0xfd, 0x39, 0xaa, 0x74, 0x37, 0x57,
	0x3a, 0xf8, 0x71, 0xa0, 0x2d, 0x9a, 0xaf, 0x94, 0x7f, 0xc7, 0x53, 0xc2, 0x31, 0xb4, 0x94, 0x95,
	0xc0, 0x93, 0xe5, 0xe7, 0xa6, 0xed, 0xeb, 0x9c, 0xae, 0x3b, 0x16, 0x72, 0x1e, 0xa0, 0x21, 0xe5,
	0x8e, 0x47, 0xcb, 0xf1, 0xd5, 0x9d, 0xe9, 0x1c, 0x9b, 0x0f, 0x05, 0xd2, 0x18, 0x5a, 0x4a, 0x36,
	0x92, 0x32, 0xd3, 0x5a, 0x48, 0xca, 0xcc, 0x91, 0x8e, 0xa1, 0xa5, 0xd8, 0x2f, 0xe1, 0x99, 0xe2,
	0x93, 0xf0, 0xcc, 0xa9, 0x3d, 0x41, 0x43, 0xca, 0x43, 0xba, 0xe9, 0x6a, 0x76, 0xd2, 0x4d, 0x0d,
	0x11, 0x5e, 0x5b, 0x77, 0xcd, 0x0f, 0xa8, 0x1e, 0xb2, 0x6c, 0x32, 0xa9, 0x95, 0x8f, 0xd8, 0xcd,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x50, 0x07, 0xb6, 0xdd, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleServiceClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	ReadArticle(ctx context.Context, in *ReadArticleRequest, opts ...grpc.CallOption) (*ReadArticleResponse, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
	ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (ArticleService_ListArticleClient, error)
}

type articleServiceClient struct {
	cc *grpc.ClientConn
}

func NewArticleServiceClient(cc *grpc.ClientConn) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ReadArticle(ctx context.Context, in *ReadArticleRequest, opts ...grpc.CallOption) (*ReadArticleResponse, error) {
	out := new(ReadArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/ReadArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error) {
	out := new(UpdateArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	out := new(DeleteArticleResponse)
	err := c.cc.Invoke(ctx, "/article.ArticleService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (ArticleService_ListArticleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArticleService_serviceDesc.Streams[0], "/article.ArticleService/ListArticle", opts...)
	if err != nil {
		return nil, err
	}
	x := &articleServiceListArticleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArticleService_ListArticleClient interface {
	Recv() (*ListArticleResponse, error)
	grpc.ClientStream
}

type articleServiceListArticleClient struct {
	grpc.ClientStream
}

func (x *articleServiceListArticleClient) Recv() (*ListArticleResponse, error) {
	m := new(ListArticleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArticleServiceServer is the server API for ArticleService service.
type ArticleServiceServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	ReadArticle(context.Context, *ReadArticleRequest) (*ReadArticleResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
	ListArticle(*ListArticleRequest, ArticleService_ListArticleServer) error
}

// UnimplementedArticleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (*UnimplementedArticleServiceServer) CreateArticle(ctx context.Context, req *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (*UnimplementedArticleServiceServer) ReadArticle(ctx context.Context, req *ReadArticleRequest) (*ReadArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadArticle not implemented")
}
func (*UnimplementedArticleServiceServer) UpdateArticle(ctx context.Context, req *UpdateArticleRequest) (*UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (*UnimplementedArticleServiceServer) DeleteArticle(ctx context.Context, req *DeleteArticleRequest) (*DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (*UnimplementedArticleServiceServer) ListArticle(req *ListArticleRequest, srv ArticleService_ListArticleServer) error {
	return status.Errorf(codes.Unimplemented, "method ListArticle not implemented")
}

func RegisterArticleServiceServer(s *grpc.Server, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ReadArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ReadArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/ReadArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ReadArticle(ctx, req.(*ReadArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ListArticle_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListArticleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArticleServiceServer).ListArticle(m, &articleServiceListArticleServer{stream})
}

type ArticleService_ListArticleServer interface {
	Send(*ListArticleResponse) error
	grpc.ServerStream
}

type articleServiceListArticleServer struct {
	grpc.ServerStream
}

func (x *articleServiceListArticleServer) Send(m *ListArticleResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "ReadArticle",
			Handler:    _ArticleService_ReadArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticleService_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticleService_DeleteArticle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListArticle",
			Handler:       _ArticleService_ListArticle_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "article/article.proto",
}
