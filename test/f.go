package main

import (
	"os"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	f, _ := os.OpenFile("/tmp/2", syscall.O_RDWR|syscall.O_CREAT|syscall.O_APPEND, 0777)
	defer f.Close()
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			f.Write([]byte("\nmsg:" + strconv.Itoa(i)))
			wg.Done()
		}(i)
	}
	wg.Wait()

}
