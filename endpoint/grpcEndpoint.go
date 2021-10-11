package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-demo/service"
)

func NewGRPCHelloEndpoint(s service.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(service.Request)
		// 业务逻辑
		v, err := s.Hello(r.Name)
		return service.Response{Reply: v}, err
	}
}
