package invoker

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"kaka/services"
	"reflect"
)

type HttpInvoker struct {
}

func (HttpInvoker) Invoke(inv *Invocation) (output *Output) {
	defer func() {
		err := recover()
		if err != nil {
			errMsg := fmt.Sprintf("Invoke panic recover : %+v", err)
			fmt.Println(errMsg)
			output.SetError(errors.New(errMsg))
		}
	}()
	output = &Output{}
	serviceName := inv.ServiceName
	methodName := inv.MethodName
	data := inv.Input
	service, err := services.GetService(serviceName)
	if err != nil {
		output.SetError(errors.Wrap(err, "执行invoker.GetService报错"))
		return output
	}
	val := reflect.ValueOf(service)
	method := val.MethodByName(methodName)
	inType := method.Type().In(0)
	in := reflect.New(inType.Elem())
	err = json.Unmarshal(data, in.Interface())
	if err != nil {
		output.SetError(errors.Wrap(err, "解析resp数据出错"))
		return output
	}
	res := method.Call([]reflect.Value{in})
	serviceRes := res[0].Interface()
	if err, ok := res[1].Interface().(error); ok {
		output.SetError(err)
	}
	output.Data = serviceRes
	return output
}
