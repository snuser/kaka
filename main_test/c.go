package main

import (
	"fmt"
	"time"
)

func search() (string, error) {
	time.Sleep(time.Second * 1)
	return "search result", nil
}

type searchReuslt struct {
	record string
	err    error
}

func main() {
	ch := make(chan searchReuslt,1)
	done := make(chan int,1)
	go func() {
		record, err := search()
		ch <- searchReuslt{record: record, err: err}
	}()

	select {
	case result := <-ch:
		if result.err != nil{
			done <- 1
			fmt.Println(result.err)
			break
		}
		fmt.Printf("receiver %s", result.record)
	case <- done:
		fmt.Printf("done")
	}

}
