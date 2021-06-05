package services

import (
	"errors"
	"kaka/dao"
	"kaka/model"
)

type UserService struct {
	Service
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

func (UserService) GetUser(input *GetUserInput) (*model.User, error) {
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
