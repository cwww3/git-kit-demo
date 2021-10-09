package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

//声明一个中间件类型
type Middleware func(endpoint.Endpoint) endpoint.Endpoint

// 日志中间件 在开始和结束打印日志
func LogMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "call start")
			defer logger.Log("msg", "call end")
			return next(ctx, request)
		}
	}
}
