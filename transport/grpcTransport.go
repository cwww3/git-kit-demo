package transport

import (
	"context"
	"go-kit-demo/service"
	"go-kit-demo/transport/pb"
)



func HelloGrpcDecodeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Request)
	return service.Request{Name: req.GetName()}, nil
}

func HelloGrpcEncodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(service.Response)
	var err string
	if resp.Err != nil {
		err = resp.Err.Error()
	}
	return &pb.Response{
		Reply: resp.Reply,
		Err:   err,
	}, nil
}
