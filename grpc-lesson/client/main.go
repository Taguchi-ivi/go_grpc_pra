package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	callListFiles(client)

}

func callListFiles(client pb.FileServiceClient) {
	req := &pb.ListFilesRequest{}
	res, err := client.ListFiles(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}

	fmt.Println(res.GetFilename())
}
