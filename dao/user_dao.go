package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"kaka/biz"
	"sync"
)
var userDaoOnce sync.Once
var userDao *UserDao

type User struct {
	Id         int
	UserName   string
	Email      string
	PasswdHash string
}

func (u *User) ToBizUser() *biz.User {
	return &biz.User{
		Id:         u.Id,
		UserName:   u.UserName,
		Email:      u.Email,
	}
}

type UserDao struct {
	DB *sql.DB
}

func GetUserDao() *UserDao{
	userDaoOnce.Do(func() {
		userDao = &UserDao{DB: GetDB()}
	})
	return userDao
}

//这里如果判断error如果是errNoRows的话 应该如何处理
//从业务角度出发不应该返回error， 如果数据为空， 这里可以返回
//1. 返回 nil  调用方需要判断 返回值是否为nil, 但函数的签名返回的*User
//2. 返回空的User对象， 在User对象设置IsEmpty函数 用来判断是否为空数据
//后面如何处理 调用方来决定

func (ud *UserDao) GetUserById(id int) (*biz.User, error) {
	user := &User{}
	sqlStr := "select * from user where id = ?"
	err := ud.DB.QueryRow(sqlStr, id).Scan(&user.Id, &user.UserName, &user.Email, &user.PasswdHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.ToBizUser(), nil
		}
		return nil, errors.Wrap(err, "GetUserById error occurred")
	}
	return user.ToBizUser(), nil
}

