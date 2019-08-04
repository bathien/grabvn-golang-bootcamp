// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedback.proto

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

// Booking feedback of customer
type BookingFeedback struct {
	// Unique integer identifier of the feedback
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// content of the feedback
	Feedback string `protobuf:"bytes,2,opt,name=feedback,proto3" json:"feedback,omitempty"`
	// booking code reference to the feedback
	BookingCodeId        string   `protobuf:"bytes,3,opt,name=bookingCodeId,proto3" json:"bookingCodeId,omitempty"`
	PassengerId          string   `protobuf:"bytes,4,opt,name=passengerId,proto3" json:"passengerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingFeedback) Reset()         { *m = BookingFeedback{} }
func (m *BookingFeedback) String() string { return proto.CompactTextString(m) }
func (*BookingFeedback) ProtoMessage()    {}
func (*BookingFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{0}
}

func (m *BookingFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingFeedback.Unmarshal(m, b)
}
func (m *BookingFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingFeedback.Marshal(b, m, deterministic)
}
func (m *BookingFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingFeedback.Merge(m, src)
}
func (m *BookingFeedback) XXX_Size() int {
	return xxx_messageInfo_BookingFeedback.Size(m)
}
func (m *BookingFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_BookingFeedback proto.InternalMessageInfo

func (m *BookingFeedback) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BookingFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

func (m *BookingFeedback) GetBookingCodeId() string {
	if m != nil {
		return m.BookingCodeId
	}
	return ""
}

func (m *BookingFeedback) GetPassengerId() string {
	if m != nil {
		return m.PassengerId
	}
	return ""
}

// Request data to create new feedback
type CreateRequest struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Feedback entity to add
	BookingFeedback      *BookingFeedback `protobuf:"bytes,2,opt,name=bookingFeedback,proto3" json:"bookingFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetBookingFeedback() *BookingFeedback {
	if m != nil {
		return m.BookingFeedback
	}
	return nil
}

// Contains data of created feedback
type CreateResponse struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created feedback
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains feedbacks of passenger ID
type GetFeedbacksByPassengerIdResponse struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Feedbacks by passenger ID
	BookingFeedback      []*BookingFeedback `protobuf:"bytes,2,rep,name=bookingFeedback,proto3" json:"bookingFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetFeedbacksByPassengerIdResponse) Reset()         { *m = GetFeedbacksByPassengerIdResponse{} }
func (m *GetFeedbacksByPassengerIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeedbacksByPassengerIdResponse) ProtoMessage()    {}
func (*GetFeedbacksByPassengerIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{3}
}

func (m *GetFeedbacksByPassengerIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbacksByPassengerIdResponse.Unmarshal(m, b)
}
func (m *GetFeedbacksByPassengerIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbacksByPassengerIdResponse.Marshal(b, m, deterministic)
}
func (m *GetFeedbacksByPassengerIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbacksByPassengerIdResponse.Merge(m, src)
}
func (m *GetFeedbacksByPassengerIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeedbacksByPassengerIdResponse.Size(m)
}
func (m *GetFeedbacksByPassengerIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbacksByPassengerIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbacksByPassengerIdResponse proto.InternalMessageInfo

func (m *GetFeedbacksByPassengerIdResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetFeedbacksByPassengerIdResponse) GetBookingFeedback() []*BookingFeedback {
	if m != nil {
		return m.BookingFeedback
	}
	return nil
}

// Request data to read feedbacks by passenger ID
type GetFeedbacksByPassengerIdRequest struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// passenger ID
	PassengerID          string   `protobuf:"bytes,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbacksByPassengerIdRequest) Reset()         { *m = GetFeedbacksByPassengerIdRequest{} }
func (m *GetFeedbacksByPassengerIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedbacksByPassengerIdRequest) ProtoMessage()    {}
func (*GetFeedbacksByPassengerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{4}
}

func (m *GetFeedbacksByPassengerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbacksByPassengerIdRequest.Unmarshal(m, b)
}
func (m *GetFeedbacksByPassengerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbacksByPassengerIdRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedbacksByPassengerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbacksByPassengerIdRequest.Merge(m, src)
}
func (m *GetFeedbacksByPassengerIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedbacksByPassengerIdRequest.Size(m)
}
func (m *GetFeedbacksByPassengerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbacksByPassengerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbacksByPassengerIdRequest proto.InternalMessageInfo

func (m *GetFeedbacksByPassengerIdRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetFeedbacksByPassengerIdRequest) GetPassengerID() string {
	if m != nil {
		return m.PassengerID
	}
	return ""
}

// Contains feedback data specified in by booking code request
type GetFeedbackByBookingCodeRequest struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// booking code
	BookingCodeId        string   `protobuf:"bytes,2,opt,name=bookingCodeId,proto3" json:"bookingCodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbackByBookingCodeRequest) Reset()         { *m = GetFeedbackByBookingCodeRequest{} }
func (m *GetFeedbackByBookingCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackByBookingCodeRequest) ProtoMessage()    {}
func (*GetFeedbackByBookingCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{5}
}

func (m *GetFeedbackByBookingCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackByBookingCodeRequest.Unmarshal(m, b)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackByBookingCodeRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackByBookingCodeRequest.Merge(m, src)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackByBookingCodeRequest.Size(m)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackByBookingCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackByBookingCodeRequest proto.InternalMessageInfo

func (m *GetFeedbackByBookingCodeRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetFeedbackByBookingCodeRequest) GetBookingCodeId() string {
	if m != nil {
		return m.BookingCodeId
	}
	return ""
}

// Request data to read feedback by booking code
type GetFeedbackByBookingCodeResponse struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Feedback entity read by booking code
	BookingFeedback      *BookingFeedback `protobuf:"bytes,2,opt,name=bookingFeedback,proto3" json:"bookingFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetFeedbackByBookingCodeResponse) Reset()         { *m = GetFeedbackByBookingCodeResponse{} }
func (m *GetFeedbackByBookingCodeResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackByBookingCodeResponse) ProtoMessage()    {}
func (*GetFeedbackByBookingCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{6}
}

func (m *GetFeedbackByBookingCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackByBookingCodeResponse.Unmarshal(m, b)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackByBookingCodeResponse.Marshal(b, m, deterministic)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackByBookingCodeResponse.Merge(m, src)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackByBookingCodeResponse.Size(m)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackByBookingCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackByBookingCodeResponse proto.InternalMessageInfo

func (m *GetFeedbackByBookingCodeResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetFeedbackByBookingCodeResponse) GetBookingFeedback() *BookingFeedback {
	if m != nil {
		return m.BookingFeedback
	}
	return nil
}

// Request data to delete feedback
type DeleteRequest struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the feedback to delete
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains status of delete operation
type DeleteResponse struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed deleted
	// Equals 1 in case of succesfull delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// Request data to read all feedback
type GetAllRequest struct {
	// API versioning
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{9}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func (m *GetAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all feedbacks
type GetAllResponse struct {
	// API versioning
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// List of all feedbacks
	BookingFeedbacks     []*BookingFeedback `protobuf:"bytes,2,rep,name=bookingFeedbacks,proto3" json:"bookingFeedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{10}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetAllResponse) GetBookingFeedbacks() []*BookingFeedback {
	if m != nil {
		return m.BookingFeedbacks
	}
	return nil
}

func init() {
	proto.RegisterType((*BookingFeedback)(nil), "pb.BookingFeedback")
	proto.RegisterType((*CreateRequest)(nil), "pb.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "pb.CreateResponse")
	proto.RegisterType((*GetFeedbacksByPassengerIdResponse)(nil), "pb.GetFeedbacksByPassengerIdResponse")
	proto.RegisterType((*GetFeedbacksByPassengerIdRequest)(nil), "pb.GetFeedbacksByPassengerIdRequest")
	proto.RegisterType((*GetFeedbackByBookingCodeRequest)(nil), "pb.GetFeedbackByBookingCodeRequest")
	proto.RegisterType((*GetFeedbackByBookingCodeResponse)(nil), "pb.GetFeedbackByBookingCodeResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pb.DeleteResponse")
	proto.RegisterType((*GetAllRequest)(nil), "pb.GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "pb.GetAllResponse")
}

func init() { proto.RegisterFile("feedback.proto", fileDescriptor_7b189e8c8330c03e) }

var fileDescriptor_7b189e8c8330c03e = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x0d, 0x60, 0xd0, 0xde, 0x06, 0x5a, 0xc7, 0xc4, 0x10, 0x36, 0xd2, 0xb1, 0x26, 0x5d, 0xb5,
	0xb1, 0x6e, 0x35, 0x46, 0xda, 0xd8, 0x74, 0x67, 0x30, 0x31, 0x71, 0x61, 0x14, 0xca, 0xb5, 0x21,
	0x6d, 0x0a, 0x32, 0x68, 0xd2, 0x0f, 0xf0, 0x6b, 0xdf, 0x4f, 0xbc, 0x94, 0x61, 0x5e, 0x99, 0x29,
	0xd0, 0xbe, 0x1d, 0x73, 0x39, 0xf7, 0x9e, 0x73, 0xe7, 0x1c, 0x00, 0xfb, 0x37, 0x62, 0x1c, 0x85,
	0x9b, 0xdd, 0x34, 0xcb, 0xd3, 0x22, 0x25, 0x7a, 0x16, 0xd1, 0xff, 0x1a, 0x0c, 0xfc, 0x34, 0xdd,
	0x25, 0x87, 0xed, 0xe7, 0xea, 0x2d, 0xb1, 0x41, 0x4f, 0x62, 0x47, 0xf3, 0xb4, 0x89, 0x11, 0xe8,
	0x49, 0x4c, 0x5c, 0x78, 0x26, 0x3a, 0x1d, 0xdd, 0xd3, 0x26, 0xbd, 0xe0, 0xe1, 0x4c, 0xc6, 0x60,
	0x45, 0xbc, 0x7d, 0x91, 0xc6, 0xb8, 0x8e, 0x1d, 0xa3, 0x04, 0xc8, 0x45, 0xe2, 0x41, 0x3f, 0x0b,
	0x19, 0xc3, 0xc3, 0x16, 0xf3, 0x75, 0xec, 0x3c, 0x29, 0x31, 0xf5, 0x12, 0xfd, 0x05, 0xd6, 0x22,
	0xc7, 0xb0, 0xc0, 0x00, 0xff, 0xfc, 0x45, 0x56, 0x90, 0x21, 0x18, 0x61, 0x96, 0x94, 0x2a, 0x7a,
	0xc1, 0xe9, 0x91, 0x7c, 0x80, 0x41, 0x24, 0x2b, 0x2d, 0xd5, 0xf4, 0xe7, 0x2f, 0xa6, 0x59, 0x34,
	0x55, 0x96, 0x08, 0x54, 0x2c, 0x9d, 0x83, 0x2d, 0x18, 0x58, 0x96, 0x1e, 0x18, 0x36, 0x50, 0xf0,
	0xcd, 0x75, 0xb1, 0x39, 0x2d, 0x60, 0xb4, 0xc2, 0x42, 0x8c, 0x60, 0xfe, 0xf1, 0xcb, 0x59, 0x72,
	0xc7, 0x98, 0x46, 0xa5, 0xc6, 0xcd, 0x4a, 0xbf, 0x81, 0xd7, 0xc1, 0xda, 0x76, 0x3d, 0xd2, 0x1d,
	0x2f, 0x2b, 0xa3, 0xea, 0x25, 0xfa, 0x1d, 0x5e, 0xd5, 0xe6, 0xfa, 0x47, 0xff, 0xec, 0x51, 0xfb,
	0xd8, 0x0b, 0x83, 0xf5, 0x06, 0x83, 0x29, 0x93, 0x24, 0x2b, 0xa3, 0x1f, 0x77, 0x4f, 0xb7, 0x3b,
	0xfa, 0x16, 0xac, 0x25, 0xee, 0xb1, 0x2b, 0x33, 0xaa, 0xa1, 0xef, 0xc1, 0x16, 0x2d, 0xad, 0xaa,
	0x1c, 0x78, 0x1a, 0x97, 0x18, 0xd1, 0x28, 0x8e, 0x74, 0x04, 0xd6, 0x0a, 0x8b, 0x4f, 0xfb, 0x7d,
	0x2b, 0x21, 0xdd, 0x80, 0x2d, 0x20, 0xad, 0x04, 0x1f, 0x61, 0xa8, 0xac, 0xc2, 0xba, 0xf2, 0x71,
	0x01, 0x9e, 0xdf, 0xe9, 0xf0, 0x52, 0x41, 0x7d, 0xc5, 0xfc, 0x5f, 0xb2, 0x41, 0x32, 0x03, 0x93,
	0xa7, 0x9c, 0x3c, 0x3f, 0xcd, 0x92, 0xbe, 0x29, 0x97, 0xd4, 0x4b, 0x95, 0xbc, 0x9f, 0x30, 0x5c,
	0x61, 0x51, 0xcf, 0xd8, 0x92, 0x8c, 0x4f, 0xb8, 0x6b, 0x11, 0x74, 0xdf, 0x5c, 0x41, 0x55, 0x04,
	0x3f, 0x2a, 0x82, 0x5a, 0x24, 0xc8, 0x6b, 0xa5, 0xb5, 0x29, 0x8b, 0xee, 0xb8, 0x1b, 0x54, 0x8d,
	0x9f, 0x81, 0xc9, 0x1d, 0xe5, 0x0b, 0x4b, 0x81, 0xe0, 0x0b, 0x2b, 0x86, 0xcf, 0xc0, 0xe4, 0x0e,
	0xf1, 0x06, 0xc9, 0x50, 0xde, 0x20, 0x1b, 0x18, 0x99, 0xe5, 0xdf, 0xf2, 0xdd, 0x7d, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0xcb, 0x47, 0x9e, 0x3f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookingFeedbackServiceClient is the client API for BookingFeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookingFeedbackServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetByPassengerID(ctx context.Context, in *GetFeedbacksByPassengerIdRequest, opts ...grpc.CallOption) (*GetFeedbacksByPassengerIdResponse, error)
	GetByBookingCode(ctx context.Context, in *GetFeedbackByBookingCodeRequest, opts ...grpc.CallOption) (*GetFeedbackByBookingCodeResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type bookingFeedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookingFeedbackServiceClient(cc *grpc.ClientConn) BookingFeedbackServiceClient {
	return &bookingFeedbackServiceClient{cc}
}

func (c *bookingFeedbackServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingFeedbackService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingFeedbackServiceClient) GetByPassengerID(ctx context.Context, in *GetFeedbacksByPassengerIdRequest, opts ...grpc.CallOption) (*GetFeedbacksByPassengerIdResponse, error) {
	out := new(GetFeedbacksByPassengerIdResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingFeedbackService/GetByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingFeedbackServiceClient) GetByBookingCode(ctx context.Context, in *GetFeedbackByBookingCodeRequest, opts ...grpc.CallOption) (*GetFeedbackByBookingCodeResponse, error) {
	out := new(GetFeedbackByBookingCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingFeedbackService/GetByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingFeedbackServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingFeedbackService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingFeedbackServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingFeedbackService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingFeedbackServiceServer is the server API for BookingFeedbackService service.
type BookingFeedbackServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetByPassengerID(context.Context, *GetFeedbacksByPassengerIdRequest) (*GetFeedbacksByPassengerIdResponse, error)
	GetByBookingCode(context.Context, *GetFeedbackByBookingCodeRequest) (*GetFeedbackByBookingCodeResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
}

// UnimplementedBookingFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookingFeedbackServiceServer struct {
}

func (*UnimplementedBookingFeedbackServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBookingFeedbackServiceServer) GetByPassengerID(ctx context.Context, req *GetFeedbacksByPassengerIdRequest) (*GetFeedbacksByPassengerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPassengerID not implemented")
}
func (*UnimplementedBookingFeedbackServiceServer) GetByBookingCode(ctx context.Context, req *GetFeedbackByBookingCodeRequest) (*GetFeedbackByBookingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByBookingCode not implemented")
}
func (*UnimplementedBookingFeedbackServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedBookingFeedbackServiceServer) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterBookingFeedbackServiceServer(s *grpc.Server, srv BookingFeedbackServiceServer) {
	s.RegisterService(&_BookingFeedbackService_serviceDesc, srv)
}

func _BookingFeedbackService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingFeedbackServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingFeedbackService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingFeedbackServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingFeedbackService_GetByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbacksByPassengerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingFeedbackServiceServer).GetByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingFeedbackService/GetByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingFeedbackServiceServer).GetByPassengerID(ctx, req.(*GetFeedbacksByPassengerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingFeedbackService_GetByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbackByBookingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingFeedbackServiceServer).GetByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingFeedbackService/GetByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingFeedbackServiceServer).GetByBookingCode(ctx, req.(*GetFeedbackByBookingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingFeedbackService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingFeedbackServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingFeedbackService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingFeedbackServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingFeedbackService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingFeedbackServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingFeedbackService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingFeedbackServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookingFeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BookingFeedbackService",
	HandlerType: (*BookingFeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookingFeedbackService_Create_Handler,
		},
		{
			MethodName: "GetByPassengerID",
			Handler:    _BookingFeedbackService_GetByPassengerID_Handler,
		},
		{
			MethodName: "GetByBookingCode",
			Handler:    _BookingFeedbackService_GetByBookingCode_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookingFeedbackService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BookingFeedbackService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback.proto",
}
