package biz

import dao "kaka/dao"

type User struct {
	Id       int
	UserName string
	Email    string
}

func (u *User) IsEmpty() bool {
	return u.Id == 0
}

func GetUserById(id int) *User {
	userDao := dao.GetUserDao()
	user, _ := userDao.GetUserById(id)
	return user
}
