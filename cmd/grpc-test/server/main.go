package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/kderosha/test/cmd/grpc-test/proto/test"
	"google.golang.org/grpc"
)

type testServer struct {
	test.UnimplementedTesterServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen")
	grpcServer := grpc.NewServer()
	fmt.Println("registering new server")
	test.RegisterTesterServer(grpcServer, &testServer{})
	fmt.Println("Serving")
	grpcServer.Serve(lis)
}

func (ts *testServer) SayHello(request *test.Test, returnStream test.Tester_SayHelloServer) error {
	ctx, cancelFunc := context.WithTimeout(returnStream.Context(), time.Second*10)
	defer cancelFunc()
	// Return
	var x int64 = 0
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			returnStream.Send(&test.Test{Message: x})
		}
		x++
	}
}
