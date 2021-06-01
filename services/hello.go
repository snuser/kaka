package services

import "fmt"

type HelloService struct {
	Service
}

func (HelloService) ServiceName() string {
	return "hello"
}

func (HelloService) Hello(input *Input) (*Output, error) {
	return &Output{Msg: "hello called"}, nil
}

func (HelloService) ShutDown() string {
	fmt.Println("shutdown helloService")
	return "shutdown helloService"
}
