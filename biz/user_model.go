package biz


type User struct {
	Id         int
	UserName   string
	Email      string
	PasswdHash string
}

func (u *User) IsEmpty() bool {
	return u.Id == 0
}