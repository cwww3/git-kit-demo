package main

import (
	"github.com/go-kit/kit/log"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	httpTransport "github.com/go-kit/kit/transport/http"
	localEndpoint "go-kit-demo/endpoint"
	"go-kit-demo/middleware"
	"go-kit-demo/service"
	"go-kit-demo/transport"
	"go-kit-demo/transport/pb"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

func main() {
	// 业务接口服务
	s := service.Service{}
	// 创建Endpoint
	helloHttpEndpoint := localEndpoint.NewHttpHelloEndpoint(s)
	helloGrpcEndpoint := localEndpoint.NewGRPCHelloEndpoint(s)
	// 添加了log中间件的服务
	logger := log.NewLogfmtLogger(os.Stderr)
	helloHttpEndpoint = middleware.LogMiddleware(logger)(helloHttpEndpoint)
	helloGrpcEndpoint = middleware.LogMiddleware(logger)(helloGrpcEndpoint)
	// 使用 kit 创建 handler
	// 传入 Endpoint 以及 定义的 加密解密方法
	helloHttpServer := httpTransport.NewServer(helloHttpEndpoint, transport.HelloHttpDecodeRequest, transport.HelloHttpEncodeResponse)
	helloGrpcServer := grpcTransport.NewServer(helloGrpcEndpoint, transport.HelloGrpcDecodeRequest, transport.HelloGrpcEncodeResponse)

	// Grpc服务
	go func() {
		hs := service.HelloServer{
			HelloHandle: helloGrpcServer,
		}
		ls, _ := net.Listen("tcp", ":8080")
		gs := grpc.NewServer()
		pb.RegisterHelloServer(gs, hs)
		gs.Serve(ls)
	}()
	// Http服务
	http.ListenAndServe(":9090", helloHttpServer)
}
