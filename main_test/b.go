package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Printf("receiver value %d",val)
	}()
	select {
	case <-time.After(time.Second*5):
		fmt.Printf("select done")
	}
}
