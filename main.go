package main

import (
	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	localEndpoint "go-kit-demo/endpoint"
	"go-kit-demo/middleware"
	"go-kit-demo/service"
	"go-kit-demo/transport"
	"net/http"
	"os"
)

func main() {
	// 业务接口服务
	s := service.Service{}
	// 创建业务服务
	helloEndpoint := localEndpoint.NewHelloEndpoint(s)
	// 添加了log中间件的服务
	logger := log.NewLogfmtLogger(os.Stderr)
	helloEndpoint = middleware.LogMiddleware(logger)(helloEndpoint)
	// 使用 kit 创建 handler
	// 传入 业务服务 以及 定义的 加密解密方法
	helloServer := httpTransport.NewServer(helloEndpoint, transport.HelloDecodeRequest, transport.HelloEncodeResponse)
	http.ListenAndServe(":9090", helloServer)
}
