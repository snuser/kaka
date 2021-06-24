package invoker

type Output struct {
	error error
	Data  interface{}
	ErrorMsg string
	ErrorCode int
}

func (o *Output) SetError(e error)  {
	o.error = e
	o.ErrorMsg = e.Error()
}

type Invoker interface {
	Invoke(inv *Invocation) *Output
}

type Invocation struct {
	MethodName  string
	ServiceName string
	Input       []byte
}
