package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brian20373/gogogrpc/gen/pb/gogogrpc"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		panic(fmt.Errorf("run: %w", err))
	}
}

func run() error {
	listenOn := ":8080"

	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		fmt.Println("fail")
	}

	server := grpc.NewServer()

	gogogrpc.RegisterGreeterServer(server, &GreeterServerHandler{})
	log.Println("Server listen on", listenOn)

	if err := server.Serve(listener); err != nil {
		fmt.Println("Server serve error : ")
	}

	return nil

}

var _ gogogrpc.GreeterServer = &GreeterServerHandler{}

type GreeterServerHandler struct {
	gogogrpc.UnimplementedGreeterServer
}

func (h *GreeterServerHandler) SayHello(ctx context.Context, req *gogogrpc.HelloRequest) (*gogogrpc.HelloReply, error) {

	fmt.Printf("Get %s request.\n", req.Name)

	return &gogogrpc.HelloReply{
		Message: fmt.Sprintln("Hello", req.Name),
	}, nil
}
