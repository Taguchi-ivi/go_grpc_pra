package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"
	"os"
	"time"

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
	// callDownload(client)
	callUpload(client)

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

func callUpload(client pb.FileServiceClient) {
	filename := "sports.txt"
	path := os.Getenv("STORAGE_PATH") + "/" + filename

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalf("Failed to upload: %v", err)
	}

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read chunk: %v", err)
		}

		req := &pb.UploadRequest{
			Data: buf[:n],
		}
		sendErr := stream.Send(req)
		if sendErr != nil {
			log.Fatalf("Failed to send chunk: %v", sendErr)
		}

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	log.Printf("Upload response: %v", res.GetSize())
}
