package services

import (
	"errors"
	"sync"
)

var services sync.Map

func AddServices(service Service) {
	services.Store(service.ServiceName(), service)
}

func GetServicesList() []Service {
	serivceList := make([]Service, 0)
	services.Range(func(key, value interface{}) bool {
		service := value.(Service)
		serivceList = append(serivceList, service)
		return true
	})
	return serivceList
}

var ServiceNotFoundError = errors.New("service not found")

func GetService(name string) (Service, error) {
	if service, ok := services.Load(name); ok {
		return service.(Service), nil
	}

	return nil,  ServiceNotFoundError
}
