package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func Server() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "test")
	})
	s := http.Server{Addr: "0.0.0.0:8082", Handler: mux}
	time.AfterFunc(time.Second*1, func() {
		fmt.Println("time trigger")
		s.Shutdown(nil)
	})
	return s.ListenAndServe()
}
func main()  {
	var g errgroup.Group
	g.Go(func() error {
		return Server()
	})

	if err := g.Wait(); err != nil {
		fmt.Println("err wait error")
	}

}
