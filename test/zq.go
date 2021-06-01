package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	Ach := make(chan struct{}, 1)
	Bch := make(chan struct{}, 1)
	Cch := make(chan struct{}, 1)

	Ach <- struct{}{}
	go func() {
		count := 0
		for count < 100 {
			select {
			case <-Ach:
				fmt.Println("A")
				count++
				Bch <- struct{}{}
			}
		}
		wg.Done()
	}()

	go func() {
		count := 0
		for count < 100 {
			select {
			case <-Bch:
				fmt.Println("B")
				count++
				Cch <- struct{}{}
			}
		}
		wg.Done()
	}()

	go func() {
		count := 0
		for count < 100 {
			select {
			case <-Cch:
				fmt.Println("C")
				count++
				Ach <- struct{}{}
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
