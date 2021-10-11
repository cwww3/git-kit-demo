package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-demo/service"
)

func NewHttpHelloEndpoint(s service.IService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(service.Request)
		v, err := s.Hello(r.Name)
		return service.Response{Reply: v}, err
	}
}
