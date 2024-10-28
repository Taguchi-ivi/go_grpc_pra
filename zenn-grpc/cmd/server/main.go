package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Taguchi-ivi/go_grpc_pra/basics/handler"
	hellopb "github.com/Taguchi-ivi/go_grpc_pra/basics/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	// gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterHelloServiceServer(s, handler.NewMyServer())

	// サーバーリフレクションの設定(gRPCurlなどでリクエストを送るため)
	// grpcurl -plaintext -d '{"name": "hsaki"}' localhost:8080 pb.HelloService.Hello
	reflection.Register(s)

	go func() {
		log.Printf("start server on port %d", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
