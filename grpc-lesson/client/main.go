package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
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
	// callListFiles(client)
	callDownload(client)

}

func callListFiles(client pb.FileServiceClient) {
	req := &pb.ListFilesRequest{}
	res, err := client.ListFiles(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}

	fmt.Println(res.GetFilename())
}

func callDownload(client pb.FileServiceClient) {
	req := &pb.DownloadRequest{
		Filename: "name.txt",
	}
	stream, err := client.Download(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to download: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive chunk: %v", err)
		}

		log.Println(res.GetData())
		log.Println(string(res.GetData()))
	}
}
