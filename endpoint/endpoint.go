package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-demo/service"
)

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Reply string
}

func NewHelloEndpoint(s service.IService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r,ok := request.(HelloRequest)
		if !ok {
			return HelloResponse{},nil
		}
		return HelloResponse{Reply: s.Hello(r.Name)},nil
	}
}
