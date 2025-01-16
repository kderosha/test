package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kderosha/test/cmd/grpc-test/proto/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type server struct {
	c test.TesterClient
}

func main() {
	client, err := grpc.Dial("grpc:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	testerClient := test.NewTesterClient(client)
	s := &server{
		c: testerClient,
	}
	router := mux.NewRouter()
	router.HandleFunc("/", s.StreamHandler)
	router.HandleFunc("/grpc", s.GrpcStreamHandler)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Minute,
		ReadTimeout:  15 * time.Minute,
	}
	fmt.Println("Starting server")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}

// Stream a json back with 1 million numbers
func (s *server) StreamHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "{\"numbers\": [")
	fmt.Fprintf(response, "%d", 1)
	for i := 2; i < 1000000000; i++ {
		fmt.Fprintf(response, ",%d", i)
	}
	fmt.Fprintf(response, "]}")
	return
}

// A handler that writes numbers to the response after retreiving them from the grpc service
func (s *server) GrpcStreamHandler(response http.ResponseWriter, request *http.Request) {
	slog.Info("Starting grpc handler")
	test, err := s.c.SayHello(context.Background(), &test.Test{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, err.Error())
		return
	}
	response.WriteHeader(http.StatusOK)
	fmt.Println("Wrote header... reading from stream")
	fmt.Fprintf(response, "{\"numbers\": [")
	fmt.Fprintf(response, "%d", -1)
	for {
		message, err := test.Recv()
		if err == io.EOF {
			fmt.Println("Done reading")
			break
		}
		fmt.Fprintf(response, ",%d", message.Message)
	}
	fmt.Println("Client done reading")
	fmt.Fprintf(response, "]}")
	return
}
