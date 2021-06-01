package main

import "fmt"

type base struct {
	name string
}

func (b base) GetName() string {
	return "name:" + b.name
}

type User struct {
	base
	name string
}

func (u *User) GetName() string {
	return "name user:" + u.name
}

func main(){
	u := &User{name: "user"}
	fmt.Println(u.GetName())
}