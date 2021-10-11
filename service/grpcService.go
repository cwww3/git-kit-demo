package service

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"go-kit-demo/transport/pb"
)

type GrpcServer struct {
	HelloHandle grpc.Handler
	pb.UnimplementedHelloServer
}

func (s GrpcServer) Hello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	_, result, err := s.HelloHandle.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return result.(*pb.Response), nil
}
