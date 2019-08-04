package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "grab-bootcamp/feedback/pb"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBookingFeedbackServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)
	passengerID := "Passenger " + pfx

	// Call Create by passenger 1
	req1 := pb.CreateRequest{
		Api: apiVersion,
		BookingFeedback: &pb.BookingFeedback{
			Feedback:      "feedback (" + pfx + ")",
			BookingCodeId: pfx,
			PassengerId:   passengerID,
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	// Call Create by passenger 1
	req2 := pb.CreateRequest{
		Api: apiVersion,
		BookingFeedback: &pb.BookingFeedback{
			Feedback:      "feedback 2 (" + pfx + ")",
			BookingCodeId: "code 2",
			PassengerId:   passengerID,
		},
	}
	res2, err := c.Create(ctx, &req2)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res2)

	// Get by passengerId
	req3 := pb.GetFeedbacksByPassengerIdRequest{
		Api:         apiVersion,
		PassengerID: passengerID,
	}
	res3, err := c.GetByPassengerID(ctx, &req3)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res3)

	// Get by passengerId
	req6 := pb.GetFeedbackByBookingCodeRequest{
		Api:           apiVersion,
		BookingCodeId: "code 2",
	}
	res6, err := c.GetByBookingCode(ctx, &req6)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res6)

	// Call GetAll
	req4 := pb.GetAllRequest{
		Api: apiVersion,
	}
	res4, err := c.GetAll(ctx, &req4)
	if err != nil {
		log.Fatalf("GetAll failed: %v", err)
	}
	log.Printf("GetAll result: <%+v>\n\n", res4)
	id := res1.Id
	// Delete
	req5 := pb.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	res5, err := c.Delete(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)
}
