package service

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"go-kit-demo/transport/pb"
)

type HelloServer struct {
	HelloHandle grpc.Handler
}

func (s *HelloServer) Hello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	_, result, err := s.HelloHandle.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return result.(*pb.Response), nil
}
