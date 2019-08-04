package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "grab-bootcamp/feedback/pb"
)

type BookingFeedback struct {
	ID            int64  `gorm:"primary_key"`
	Feedback      string `gorm:"not null"`
	BookingCodeID string `gorm:"not null"`
	PassengerID   string `gorm:"not null"`

	CreatedAt time.Time
}

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// BookingFeedbackServiceServer is implementation of pb.BookingFeedbackServiceServer proto interface
type BookingFeedbackServiceServer struct {
	db *gorm.DB
}

// NewBookingFeedbackServiceServer creates BookingFeedback service
func Test(db *gorm.DB) *BookingFeedbackServiceServer {
	return &BookingFeedbackServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *BookingFeedbackServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new BookingFeedback task
func (s *BookingFeedbackServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	feedback := BookingFeedback{
		Feedback:      req.BookingFeedback.Feedback,
		BookingCodeID: req.BookingFeedback.BookingCodeId,
		PassengerID:   req.BookingFeedback.PassengerId,
	}
	log.Printf("Get feedback", feedback)

	dbc := db.Create(&feedback)
	if dbc.Error != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into BookingFeedback-> ")
	}

	return &pb.CreateResponse{
		Api: apiVersion,
		Id:  feedback.ID,
	}, nil
}

// Read BookingFeedback task
func (s *BookingFeedbackServiceServer) GetByPassengerID(ctx context.Context, req *pb.GetFeedbacksByPassengerIdRequest) (*pb.GetFeedbacksByPassengerIdResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	list := []*pb.BookingFeedback{}

	dbList := db.Where("passenger_id = ?", req.PassengerID).Find(&list)

	if err := dbList.Error; err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BookingFeedback-> ")
	}

	return &pb.GetFeedbacksByPassengerIdResponse{
		Api:             apiVersion,
		BookingFeedback: list,
	}, nil

}

// Read BookingFeedback task
func (s *BookingFeedbackServiceServer) GetByBookingCode(ctx context.Context, req *pb.GetFeedbackByBookingCodeRequest) (*pb.GetFeedbackByBookingCodeResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	var feedback pb.BookingFeedback

	dbq := db.Where("booking_code_id = ?", req.BookingCodeId).First(&feedback)

	if err := dbq.Error; err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BookingFeedback-> ")
	}

	return &pb.GetFeedbackByBookingCodeResponse{
		Api:             apiVersion,
		BookingFeedback: &feedback,
	}, nil

}

// Delete BookingFeedback task
func (s *BookingFeedbackServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	dbd := db.Where("id = ?", req.Id).Delete(&BookingFeedback{})

	if dbd.Error != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("BookingFeedback with ID='%d' is not found",
			req.Id))
	}

	return &pb.DeleteResponse{
		Api:     apiVersion,
		Deleted: req.Id,
	}, nil
}

// Get all BookingFeedback
func (s *BookingFeedbackServiceServer) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	list := []*pb.BookingFeedback{}
	dbList := db.Find(&list)

	if dbList.Error != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from BookingFeedback-> ")
	}

	return &pb.GetAllResponse{
		Api:              apiVersion,
		BookingFeedbacks: list,
	}, nil
}
