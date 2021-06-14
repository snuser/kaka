package apps

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var amOnce sync.Once

type AppManager struct {
	Apps   []*App
	Cancel context.CancelFunc
	Ctx    context.Context
}

type App struct {
	Name       string
	Server     *http.Server
	ShutdownFn func()
}

var appManager *AppManager

func NewAppManager() *AppManager {
	amOnce.Do(func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		appManager = &AppManager{}
		appManager.Ctx = ctx
		appManager.Cancel = cancel
	})
	return appManager
}

func (am *AppManager) Add(addr string, handler http.Handler, name string, shutdownFn func()) {
	httpServer := &http.Server{Addr: addr, Handler: handler}
	app := &App{
		Name:       name,
		Server:     httpServer,
		ShutdownFn: shutdownFn,
	}
	appManager.Apps = append(appManager.Apps, app)
}

func (am *AppManager) Start() {
	group, errCtx := errgroup.WithContext(am.Ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	group.Go(am.innerStart)
	group.Go(func() error {
		for {
			select {
			case <-signalChan:
				am.Cancel()
			case <-errCtx.Done():
				return am.Shutdown()
			}
		}
	})
	if err := group.Wait(); err != nil {
		fmt.Println("apps error", err)
	}
	fmt.Println("all apps done")
}

//使用 sync.WaitGroup 启动多个app
func (am *AppManager) innerStart() error {
	var wg sync.WaitGroup
	wg.Add(len(am.Apps))
	appErr := make([]error, 0)
	for _, app := range am.Apps {
		goApp := app
		go func() {
			defer func() {
				wg.Done()
			}()
			fmt.Printf("app: %s started, listen:%s\n", goApp.Name, goApp.Server.Addr)
			err := goApp.Server.ListenAndServe()
			fmt.Printf("app: %s, canceled, Err:%s \n", goApp.Name, err)
			appErr = append(appErr, err)
		}()
	}
	wg.Wait()
	for _, err := range appErr {
		if err != nil {
			return err
		}
	}
	return nil
}

//停止所有的APP
func (am *AppManager) Shutdown() error {
	var swg sync.WaitGroup
	swg.Add(len(am.Apps))
	shutdownErrs := make([]error, 0)
	for _, app := range am.Apps {
		goApp := app
		go func() {
			defer swg.Done()
			goApp.ShutdownFn()
			shutdownErrs = append(shutdownErrs, goApp.Server.Shutdown(am.Ctx))
		}()
	}
	swg.Wait()
	for _, err := range shutdownErrs {
		if err != nil {
			return err
		}
	}
	return nil
}
