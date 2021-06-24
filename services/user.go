package services

import (
	"errors"
	"kaka/biz"
	"kaka/dao"
	"kaka/internal/pkg/services_manage"
)

type UserService struct {
	services_manage.Service
}

func (UserService) ServiceName() string {
	return "user"
}

func (UserService) ShutDown() string {
	return "shutdown userService"
}

type GetUserInput struct {
	Id int
}

func (UserService) GetUser(input *GetUserInput) (*biz.User, error) {
	id := input.Id
	user, err := dao.GetUserDao().GetUserById(id)
	if err != nil {
		return nil, err
	}

	if user.IsEmpty() {
		return user, errors.New("user empty")
	}
	return user, nil
}
