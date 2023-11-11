package main

import (
	"context"
	"fmt"
	"io"

	"github.com/kderosha/test/cmd/grpc-test/proto/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := grpc.Dial("localhost:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	testerClient := test.NewTesterClient(client)
	test, err := testerClient.SayHello(context.Background(), &test.Test{})
	if err != nil {
		panic(err)
	}
	for {
		message, err := test.Recv()
		if err == io.EOF {
			fmt.Println("Done reading")
			break
		}
		fmt.Println(message.Message)
	}
	fmt.Println("Client done reading")

}
