package handler

import (
	"context"
	"fmt"

	hellopb "github.com/Taguchi-ivi/go_grpc_pra/basics/pkg/grpc"
)

type myServer struct {
	hellopb.UnimplementedHelloServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}
