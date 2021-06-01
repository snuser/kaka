package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan string , 0)
	ch2 := make(chan string, 0 )

	go func() {
		time.Sleep(time.Second)
		ch1 <- "msg from ch1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "msg from ch2"
	}()
	for i:=0; i<2; i++ {
		select {
			case msg := <- ch1:
				fmt.Println(msg)
			case msg := <- ch2:
				fmt.Println(msg)
		}
	}
}