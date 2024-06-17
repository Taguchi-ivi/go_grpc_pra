package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles was invoked")

	dir := "/Users/taguchikazuki/dev/study/go/udemy-go-grpc/protobuf-lesson/grpc-lesson/storage"
	paths, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesResponse{
		Filename: filenames,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
