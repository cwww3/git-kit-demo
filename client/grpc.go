package main

import (
	"context"
	"fmt"
	"go-kit-demo/transport/pb"
	"google.golang.org/grpc"
)

func main() {
	serviceAddress := "127.0.0.1:8080"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	helloClient := pb.NewHelloClient(conn)
	hello, _ := helloClient.Hello(context.Background(), &pb.Request{Name: "cw"})
	fmt.Println(hello.GetReply(), hello.GetErr())
}
