package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "grab-bootcamp/feedback/pb"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB  parameters section
	// DBHost is host of database
	DBHost string
	// DBUser is username to connect to database
	DBUser string
	// DBPassword password to connect to database
	DBPassword string
	// DBSchema is schema of database
	DBName string
}

var db *gorm.DB

// RunServer runs gRPC server and HTTP gateway
func RunServer(ctx context.Context, feedbackAPI pb.BookingFeedbackServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	pb.RegisterBookingFeedbackServiceServer(server, feedbackAPI)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	reflection.Register(server)

	return server.Serve(listen)
}

func main() {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DBName, "db-name", "", "Database name")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		log.Fatal("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
	var err error

	db, err = gorm.Open("postgres", "host="+cfg.DBHost+" sslmode=disable port=5432 user="+cfg.DBUser+" dbname="+cfg.DBName+" password="+cfg.DBPassword)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	db.LogMode(true)

	defer db.Close()

	err = db.AutoMigrate(BookingFeedback{}).Error
	if err != nil {
		log.Fatal("Error migrate")
	}

	feedBackAPI := Test(db)

	err = RunServer(ctx, feedBackAPI, cfg.GRPCPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
