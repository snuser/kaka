package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type hello struct {
	endpoint string
	SayHi    func(name string) (string, error)
}

func (h hello) SayHello(name string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(h.endpoint + "/" + name)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	h := &hello{endpoint: "http://baidu.com"}
	changeFunc(h)
	fmt.Println(h.SayHi("landfall"))
}


func changeFunc(val interface{}) {
	v := reflect.ValueOf(val)
	ele := v.Elem()
	t := ele.Type()
	num := t.NumField()
	for i := 0; i < num; i++ {
		field := t.Field(i)
		fieldValue := ele.Field(i)
		if fieldValue.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				name := args[0].Interface().(string)
				value := fmt.Sprintf("name: %s", name)
				return []reflect.Value{reflect.ValueOf(value), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fieldValue.Set(reflect.MakeFunc(field.Type, fn))
		}
	}
}
