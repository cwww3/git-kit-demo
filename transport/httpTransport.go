package transport

import (
	"context"
	"encoding/json"
	"errors"
	"go-kit-demo/service"
	"net/http"
)

func HelloHttpDecodeRequest(ctx context.Context,request *http.Request) (interface{},error) {
	name := request.URL.Query().Get("name")
	if name == "" {
		return nil,errors.New("参数为空")
	}
	return service.Request{Name: name},nil
}
func HelloHttpEncodeResponse(ctx context.Context,w http.ResponseWriter,response interface{}) error {
	w.Header().Set("Content-type","application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
