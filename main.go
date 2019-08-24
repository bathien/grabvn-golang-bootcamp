package main

import (
	"math/rand"
	"net/http"

	"github.com/DataDog/datadog-go/statsd"
	log "github.com/sirupsen/logrus"
)

func main() {
	var client, err = statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	client.Namespace = "server."
	// Count two events
	// Create our server
	logger := log.New()
	server := Server{
		logger: logger,
		datadog: client,
	}

	// Start the server
	server.ListenAndServe()


}

// Server represents our server.
type Server struct {
	logger *log.Logger
	datadog *statsd.Client

}

// ListenAndServe starts the server
func (s *Server) ListenAndServe() {
	s.logger.Info("echo server is starting on port 8080...")
	http.HandleFunc("/", s.echo)
	http.ListenAndServe(":8080", nil)
}

// Echo echos back the request as a response
func (s *Server) echo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	// 30% chance of failure
	if rand.Intn(100) < 30 {
		err := s.datadog.Count("server_broken", 1, []string{"error"}, 1)
		if err != nil {
			log.Fatal(err)
		}
		s.logger.Error("a chaos monkey broke your server")
		writer.WriteHeader(500)
		writer.Write([]byte("a chaos monkey broke your server"))
		return
	}
	err := s.datadog.Count("server_happy", 1, []string{"info"}, 1)
	if err != nil {
		log.Fatal(err)
	}
	// Happy pat
	writer.WriteHeader(200)
	request.Write(writer)
}
