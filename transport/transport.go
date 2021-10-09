package transport

import (
	"context"
	"encoding/json"
	"errors"
	localEndpoint "go-kit-demo/endpoint"
	"net/http"
)

func HelloDecodeRequest(ctx context.Context,request *http.Request) (interface{},error) {
	name := request.URL.Query().Get("name")
	if name == "" {
		return nil,errors.New("参数为空")
	}
	return localEndpoint.HelloRequest{Name: name},nil
}
func HelloEncodeResponse(ctx context.Context,w http.ResponseWriter,response interface{}) error {
	w.Header().Set("Content-type","application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
