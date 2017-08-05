// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/library/book_service.proto

/*
Package library is a generated protocol buffer package.

Package library exposes a list of books
over a gRPC API.

It is generated from these files:
	proto/library/book_service.proto

It has these top-level messages:
	Publisher
	Book
	GetBookRequest
	QueryBooksRequest
*/
package library

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

// BookType describes the different types
// a book in the library can be.
type BookType int32

const (
	// Hardcover is a book with a hard back.
	BookType_HARDCOVER BookType = 0
	// Paperback is a book with a soft back.
	BookType_PAPERBACK BookType = 1
	// Audiobook is an audio recording of the book.
	BookType_AUDIOBOOK BookType = 2
)

var BookType_name = map[int32]string{
	0: "HARDCOVER",
	1: "PAPERBACK",
	2: "AUDIOBOOK",
}
var BookType_value = map[string]int32{
	"HARDCOVER": 0,
	"PAPERBACK": 1,
	"AUDIOBOOK": 2,
}

func (x BookType) String() string {
	return proto.EnumName(BookType_name, int32(x))
}
func (BookType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Publisher describes a Book Publisher.
type Publisher struct {
	// Name is the name of the Publisher.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Publisher) Reset()                    { *m = Publisher{} }
func (m *Publisher) String() string            { return proto.CompactTextString(m) }
func (*Publisher) ProtoMessage()               {}
func (*Publisher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Publisher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Book represents a book in the library.
type Book struct {
	// Isbn is the ISBN number of the book.
	Isbn int64 `protobuf:"varint,1,opt,name=isbn" json:"isbn,omitempty"`
	// Title is the title of the book.
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// Author is the author of the book.
	Author string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	// BookType is the type of the book.
	BookType BookType `protobuf:"varint,4,opt,name=book_type,json=bookType,enum=library.BookType" json:"book_type,omitempty"`
	// PublishingMethod is the publishing method
	// used for this Book.
	//
	// Types that are valid to be assigned to PublishingMethod:
	//	*Book_SelfPublished
	//	*Book_Publisher
	PublishingMethod isBook_PublishingMethod `protobuf_oneof:"publishing_method"`
	// PublicationDate is the time of publication of the book.
	PublicationDate *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=publication_date,json=publicationDate" json:"publication_date,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isBook_PublishingMethod interface {
	isBook_PublishingMethod()
}

type Book_SelfPublished struct {
	SelfPublished bool `protobuf:"varint,5,opt,name=self_published,json=selfPublished,oneof"`
}
type Book_Publisher struct {
	Publisher *Publisher `protobuf:"bytes,6,opt,name=publisher,oneof"`
}

func (*Book_SelfPublished) isBook_PublishingMethod() {}
func (*Book_Publisher) isBook_PublishingMethod()     {}

func (m *Book) GetPublishingMethod() isBook_PublishingMethod {
	if m != nil {
		return m.PublishingMethod
	}
	return nil
}

func (m *Book) GetIsbn() int64 {
	if m != nil {
		return m.Isbn
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetBookType() BookType {
	if m != nil {
		return m.BookType
	}
	return BookType_HARDCOVER
}

func (m *Book) GetSelfPublished() bool {
	if x, ok := m.GetPublishingMethod().(*Book_SelfPublished); ok {
		return x.SelfPublished
	}
	return false
}

func (m *Book) GetPublisher() *Publisher {
	if x, ok := m.GetPublishingMethod().(*Book_Publisher); ok {
		return x.Publisher
	}
	return nil
}

func (m *Book) GetPublicationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.PublicationDate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Book) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Book_OneofMarshaler, _Book_OneofUnmarshaler, _Book_OneofSizer, []interface{}{
		(*Book_SelfPublished)(nil),
		(*Book_Publisher)(nil),
	}
}

func _Book_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Book)
	// publishing_method
	switch x := m.PublishingMethod.(type) {
	case *Book_SelfPublished:
		t := uint64(0)
		if x.SelfPublished {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Book_Publisher:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Publisher); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Book.PublishingMethod has unexpected type %T", x)
	}
	return nil
}

func _Book_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Book)
	switch tag {
	case 5: // publishing_method.self_published
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.PublishingMethod = &Book_SelfPublished{x != 0}
		return true, err
	case 6: // publishing_method.publisher
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Publisher)
		err := b.DecodeMessage(msg)
		m.PublishingMethod = &Book_Publisher{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Book_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Book)
	// publishing_method
	switch x := m.PublishingMethod.(type) {
	case *Book_SelfPublished:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *Book_Publisher:
		s := proto.Size(x.Publisher)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// GetBookRequest is the input to the GetBook method.
type GetBookRequest struct {
	// Isbn is the ISBN with which
	// to match against the ISBN of a book in the library.
	Isbn int64 `protobuf:"varint,1,opt,name=isbn" json:"isbn,omitempty"`
}

func (m *GetBookRequest) Reset()                    { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()               {}
func (*GetBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetBookRequest) GetIsbn() int64 {
	if m != nil {
		return m.Isbn
	}
	return 0
}

// QueryBooksRequest is the input to the QueryBooks method.
type QueryBooksRequest struct {
	// AuthorPrefix is the prefix with which
	// to match against the author of a book in the library.
	AuthorPrefix string `protobuf:"bytes,1,opt,name=author_prefix,json=authorPrefix" json:"author_prefix,omitempty"`
}

func (m *QueryBooksRequest) Reset()                    { *m = QueryBooksRequest{} }
func (m *QueryBooksRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryBooksRequest) ProtoMessage()               {}
func (*QueryBooksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryBooksRequest) GetAuthorPrefix() string {
	if m != nil {
		return m.AuthorPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*Publisher)(nil), "library.Publisher")
	proto.RegisterType((*Book)(nil), "library.Book")
	proto.RegisterType((*GetBookRequest)(nil), "library.GetBookRequest")
	proto.RegisterType((*QueryBooksRequest)(nil), "library.QueryBooksRequest")
	proto.RegisterEnum("library.BookType", BookType_name, BookType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	// GetBook returns a Book from the library
	// that matches the ISBN provided, if found.
	// Otherwise it returns a NotFound error.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// QueryBooks returns all Books whos author
	// matches the author prefix provided, as a stream
	// of Books.
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpc.CallOption) (BookService_QueryBooksClient, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/library.BookService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpc.CallOption) (BookService_QueryBooksClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BookService_serviceDesc.Streams[0], c.cc, "/library.BookService/QueryBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceQueryBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_QueryBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceQueryBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceQueryBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BookService service

type BookServiceServer interface {
	// GetBook returns a Book from the library
	// that matches the ISBN provided, if found.
	// Otherwise it returns a NotFound error.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// QueryBooks returns all Books whos author
	// matches the author prefix provided, as a stream
	// of Books.
	QueryBooks(*QueryBooksRequest, BookService_QueryBooksServer) error
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_QueryBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryBooksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).QueryBooks(m, &bookServiceQueryBooksServer{stream})
}

type BookService_QueryBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceQueryBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceQueryBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "library.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryBooks",
			Handler:       _BookService_QueryBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/library/book_service.proto",
}

func init() { proto.RegisterFile("proto/library/book_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x8f, 0xd2, 0x40,
	0x14, 0xa6, 0x2c, 0x0b, 0xf4, 0xad, 0x20, 0x8c, 0x46, 0x1b, 0x2e, 0xdb, 0x54, 0x13, 0x89, 0x87,
	0x62, 0xd8, 0x83, 0x26, 0x9e, 0x60, 0x21, 0x62, 0xf6, 0x00, 0x8e, 0xab, 0xd7, 0xa6, 0x5d, 0x1e,
	0x30, 0xd9, 0xd2, 0xa9, 0xd3, 0xa9, 0xb1, 0x27, 0xff, 0x26, 0xff, 0x43, 0x33, 0xd3, 0x69, 0xd7,
	0x0d, 0xde, 0xe6, 0x7d, 0x3f, 0xa6, 0xdf, 0xfb, 0xa6, 0xe0, 0xa6, 0x82, 0x4b, 0x3e, 0x89, 0x59,
	0x24, 0x42, 0x51, 0x4c, 0x22, 0xce, 0xef, 0x83, 0x0c, 0xc5, 0x4f, 0x76, 0x87, 0xbe, 0xa6, 0x48,
	0xc7, 0x70, 0xa3, 0xcb, 0x3d, 0xe7, 0xfb, 0x18, 0x27, 0x1a, 0x8e, 0xf2, 0xdd, 0x44, 0xb2, 0x23,
	0x66, 0x32, 0x3c, 0xa6, 0xa5, 0xd2, 0xbb, 0x04, 0x7b, 0x93, 0x47, 0x31, 0xcb, 0x0e, 0x28, 0x08,
	0x81, 0x56, 0x12, 0x1e, 0xd1, 0xb1, 0x5c, 0x6b, 0x6c, 0x53, 0x7d, 0xf6, 0xfe, 0x34, 0xa1, 0x35,
	0xe7, 0xfc, 0x5e, 0x91, 0x2c, 0x8b, 0x12, 0x4d, 0x9e, 0x51, 0x7d, 0x26, 0xcf, 0xe1, 0x5c, 0x32,
	0x19, 0xa3, 0xd3, 0xd4, 0x8e, 0x72, 0x20, 0x2f, 0xa0, 0x1d, 0xe6, 0xf2, 0xc0, 0x85, 0x73, 0xa6,
	0x61, 0x33, 0x11, 0x1f, 0x6c, 0x9d, 0x55, 0x16, 0x29, 0x3a, 0x2d, 0xd7, 0x1a, 0xf7, 0xa7, 0x43,
	0xdf, 0x24, 0xf5, 0xd5, 0x37, 0x6e, 0x8b, 0x14, 0x69, 0x37, 0x32, 0x27, 0xf2, 0x06, 0xfa, 0x19,
	0xc6, 0xbb, 0x20, 0x35, 0x01, 0xb7, 0xce, 0xb9, 0x6b, 0x8d, 0xbb, 0xab, 0x06, 0xed, 0x29, 0xbc,
	0xca, 0xbd, 0x25, 0x53, 0xb0, 0x2b, 0x8d, 0x70, 0xda, 0xae, 0x35, 0xbe, 0x98, 0x92, 0xfa, 0xe2,
	0x7a, 0xbd, 0x55, 0x83, 0x3e, 0xc8, 0xc8, 0x12, 0x06, 0x7a, 0xb8, 0x0b, 0x25, 0xe3, 0x49, 0xb0,
	0x0d, 0x25, 0x3a, 0x1d, 0x6d, 0x1d, 0xf9, 0x65, 0x69, 0x7e, 0x55, 0x9a, 0x7f, 0x5b, 0x95, 0x46,
	0x9f, 0xfe, 0xe3, 0x59, 0x84, 0x12, 0xe7, 0xcf, 0x60, 0x68, 0xee, 0x64, 0xc9, 0x3e, 0x38, 0xa2,
	0x3c, 0xf0, 0xad, 0xf7, 0x1a, 0xfa, 0x9f, 0x50, 0xaa, 0x8d, 0x28, 0xfe, 0xc8, 0x31, 0x93, 0xff,
	0x2b, 0xcf, 0xfb, 0x00, 0xc3, 0x2f, 0x39, 0x8a, 0x42, 0xe9, 0xb2, 0x4a, 0xf8, 0x0a, 0x7a, 0x65,
	0x5b, 0x41, 0x2a, 0x70, 0xc7, 0x7e, 0x99, 0xb7, 0x78, 0x52, 0x82, 0x1b, 0x8d, 0xbd, 0x7d, 0x0f,
	0xdd, 0xaa, 0x2e, 0xd2, 0x03, 0x7b, 0x35, 0xa3, 0x8b, 0xeb, 0xf5, 0xf7, 0x25, 0x1d, 0x34, 0xd4,
	0xb8, 0x99, 0x6d, 0x96, 0x74, 0x3e, 0xbb, 0xbe, 0x19, 0x58, 0x6a, 0x9c, 0x7d, 0x5b, 0x7c, 0x5e,
	0xcf, 0xd7, 0xeb, 0x9b, 0x41, 0x73, 0xfa, 0x1b, 0x2e, 0x94, 0xf1, 0x6b, 0xf9, 0xb3, 0x90, 0x2b,
	0xe8, 0x98, 0x9c, 0xe4, 0x65, 0xdd, 0xd7, 0xe3, 0xe4, 0xa3, 0xde, 0xa3, 0x17, 0xf2, 0x1a, 0xe4,
	0x23, 0xc0, 0x43, 0x6c, 0x32, 0xaa, 0xe9, 0x93, 0x5d, 0x4e, 0xac, 0xef, 0xac, 0xa8, 0xad, 0x3b,
	0xbd, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xeb, 0x95, 0x76, 0xc3, 0x02, 0x00, 0x00,
}
