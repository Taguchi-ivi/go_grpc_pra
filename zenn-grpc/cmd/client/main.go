package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	hellopb "github.com/Taguchi-ivi/go_grpc_pra/basics/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  hellopb.HelloServiceClient
)

func main() {
	fmt.Println("start gRPC client.")

	// 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("connection error:", err)
		return
	}
	defer conn.Close()

	// gRPCクライアントを生成
	client = hellopb.NewHelloServiceClient(conn)

	for {
		fmt.Println("1: send request")
		fmt.Println("2: exit")
		fmt.Println("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()
		case "2":
			fmt.Println("bye.")
			goto M
		}
	}

M:
}

func Hello() {
	fmt.Println("please enter your name >")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(res.GetMessage())
	}
}
