package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex
var a string


func f1()  {
	a ="hello world"
	l.Unlock()
}

func main()  {
	go f1()
	l.Lock()
	fmt.Println(a)
}