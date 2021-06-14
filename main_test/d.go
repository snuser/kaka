package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func serverApp()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, serverApp")
	})
	s := http.Server{
		Addr: "0.0.0.0:8071",
		Handler: mux,
	}
	s.ListenAndServe()
	if err := http.ListenAndServe("0.0.0.0:8071", mux);err != nil{
		log.Fatalln(err)
	}
}

func serverDebug()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, serverDebug")
	})
	time.AfterFunc(time.Second*3, func() {
		log.Fatalln("server Debug Fatal")
	})
	if err := http.ListenAndServe("0.0.0.0:8072", mux);err != nil{
		log.Fatalln(err)
	}
}
func main()  {
	go serverDebug()
	go serverApp()
	select {

	}
}