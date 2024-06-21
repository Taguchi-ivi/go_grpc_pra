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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	// callUpload(client)
	// callUploadAndNotifyProgress(client)

}

func callListFiles(client pb.FileServiceClient) {
	md := metadata.New(map[string]string{
		"authorization": "Bearer test-token",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &pb.ListFilesRequest{}
	res, err := client.ListFiles(ctx, req)
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}

	fmt.Println(res.GetFilename())
}

func callDownload(client pb.FileServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.DownloadRequest{
		// 存在しないファイル名を指定して、gRPCのエラーを確認する
		Filename: "name.txt",
	}
	stream, err := client.Download(ctx, req)
	if err != nil {
		log.Fatalf("Failed to download: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			resErr, ok := status.FromError(err)
			if ok {
				if resErr.Code() == codes.NotFound {
					log.Fatalf("error code %v, error message %v", resErr.Code(), resErr.Message())
				} else if resErr.Code() == codes.DeadlineExceeded {
					log.Fatalln("deadline exceeded")
				} else {
					log.Fatalln("unknown grpc error")
				}
			} else {
				log.Fatalf("Failed to receive chunk: %v", err)
			}
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

func callUploadAndNotifyProgress(client pb.FileServiceClient) {
	filename := "sports.txt"
	path := os.Getenv("STORAGE_PATH") + "/" + filename

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := client.UploadAndNotifyProgress(context.Background())
	if err != nil {
		log.Fatalf("Failed to upload: %v", err)
	}

	// request
	buf := make([]byte, 5)
	go func() {
		for {
			n, err := file.Read(buf)
			if n == 0 || err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to read chunk: %v", err)
			}

			req := &pb.UploadAndNotifyProgressRequest{
				Data: buf[:n],
			}
			sendErr := stream.Send(req)
			if sendErr != nil {
				log.Fatalf("Failed to send chunk: %v", sendErr)
			}
			time.Sleep(1 * time.Second)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Failed to close stream: %v", err)
		}
	}()

	// response
	ch := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive response: %v", err)
			}

			log.Printf("Received response: %v", res.GetMsg())
		}
		close(ch)
	}()
	<-ch
}
