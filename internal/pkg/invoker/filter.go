package invoker

import (
	"fmt"
	"os"
	"syscall"
)

type FilterInvoker struct {
	Invoker
	Filter []Filter
}
type Filter func(inv *Invocation)

func (f *FilterInvoker) Invoke(inv *Invocation) (output *Output) {
	for _, flt := range f.Filter {
		flt(inv)
	}
	return f.Invoker.Invoke(inv)
}

func LogFilter(inv *Invocation) {
	f, _ := os.OpenFile("/tmp/1111", syscall.O_RDWR|syscall.O_CREAT|syscall.O_APPEND, 0777)
	defer func() {
		f.Close()
	}()
	logStr := fmt.Sprintf("logFilter ==== serivceName:%s, MethodName: %s \n", inv.ServiceName, inv.MethodName)
	f.Write([]byte(logStr))
}
