package main

import (
	"fmt"
	"sync"
	"time"
)

var limit = make(chan int, 3)
var index = 0
var stop = make(chan int, 1)

func f() {
	fmt.Println(index)
	fmt.Println(time.Now())
}
var wg sync.WaitGroup
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				<-limit
				wg.Done()
			}()
			limit <- 1
			index++
			f()
			time.Sleep(time.Second * 5)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
