package service

import "fmt"

type IService interface {
	Hello(name string) (string, error)
}

type Service struct {
	Name string
}

func (s Service) Hello(name string) (string, error) {
	return fmt.Sprintf("name:%v", name), nil
}

// 定于业务结构
type Request struct {
	Name string
}

type Response struct {
	Reply string
	Err   error
}
