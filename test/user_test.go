package test

import (
	"kaka/services"
	"testing"
)

func TestGetUserById(t *testing.T) {
	service := services.UserService{}
	input := &services.GetUserInput{Id: 3}
	user,err := service.GetUser(input)
	if err !=nil {
		panic("error")
	}
	if user.IsEmpty(){
		panic("error")
	}
	if user.UserName != "lan2"{
		panic("error")
	}


}
