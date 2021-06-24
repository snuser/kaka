package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kaka/internal/pkg/invoker"
	"kaka/internal/pkg/services_manage"
	"kaka/pkg"
	"kaka/services"
	"net/http"
	"os"
	"sync"
	"time"
)

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

func AppShutdown() {
	fmt.Println("App shutdown start")
	forceShutdownIfNeed()
	ShutdownAppServices()
	fmt.Println("shutdown done")
}

func ShutdownAppServices() {
	var wg sync.WaitGroup
	servicesList := services_manage.GetServicesList()
	wg.Add(len(servicesList))
	for _, service := range servicesList {
		s := service
		go func() {
			defer wg.Done()
			fmt.Println(s.ShutDown())
		}()
	}
	wg.Wait()
}

// GetAppServer 返回地址, handler, APP的名字和卸载服务的回调函数
func GetAppServer() (string, http.Handler, string, func()) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return "0.0.0.0:8081", mux, "APP1", AppShutdown
}

func GetDebugServer() (string, http.Handler, string, func()) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return "0.0.0.0:8082", mux, "Debug", AppShutdown
}

func main() {
	services_manage.AddServices(&services.UserService{})
	appManager := pkg.NewAppManager()
	appManager.Add(GetAppServer())
	appManager.Add(GetDebugServer())
	appManager.Start()
}
