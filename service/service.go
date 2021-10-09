package service

import "fmt"

type IService interface {
	Hello(name string) string
}

type Service struct {
	Name string
}

func (s Service)Hello(name string) string {
	return fmt.Sprintf("name:%v",name)
}
