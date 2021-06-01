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

func (HttpInvoker) Invoke(inv *Invocation) (re []byte, e error) {
	defer func() {
		err := recover()
		if err != nil {
			errMsg := fmt.Sprintf("panic recover : %+v", err)
			re = []byte(errMsg)
			e = errors.New("panic error"+errMsg)
		}
	}()
	serviceName := inv.ServiceName
	methodName := inv.MethodName
	data := inv.Input
	service, err := services.GetService(serviceName)
	if err != nil {
		return nil, errors.Wrap(err, "执行invoke出错")
	}
	val := reflect.ValueOf(service)
	method := val.MethodByName(methodName)
	inType := method.Type().In(0)
	in := reflect.New(inType.Elem())
	err = json.Unmarshal(data, in.Interface())
	if err != nil {
		return nil, errors.Wrap(err, "json解析出错")
	}
	res := method.Call([]reflect.Value{in})
	output, err := json.Marshal(res[0].Interface())
	if err != nil {
		return nil, err
	}
	return output, e
}
