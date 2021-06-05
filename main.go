package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kaka/invoker"
	"kaka/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var FLock sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(data))
	serviceName := r.Header.Get("sparrow-service")
	methodName := r.Header.Get("sparrow-service-method")
	inv := &invoker.Invocation{
		MethodName:  methodName,
		ServiceName: serviceName,
		Input:       data,
	}
	filterInvoker := invoker.FilterInvoker{
		Invoker: invoker.HttpInvoker{},
		Filter:  []invoker.Filter{invoker.LogFilter},
	}
	output := filterInvoker.Invoke(inv)
	respData, _ := json.Marshal(output)
	fmt.Fprintf(w, "%s", string(respData))
}

func forceShutdownIfNeed() {
	time.AfterFunc(time.Minute, func() {
		os.Exit(1)
	})
}
func shutdown() {
	servicesList := services.GetServicesList()
	for _, service := range servicesList {
		go func() {
			service.ShutDown()
		}()
	}
}
func listenSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGKILL)
	select {
	case <-signals:
		fmt.Println("收到信号")
		forceShutdownIfNeed()
		shutdown()
		os.Exit(0)
	}
}
func main() {
	services.AddServices(&services.UserService{})
	go func() {
		listenSignal()
	}()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
