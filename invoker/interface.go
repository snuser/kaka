package invoker

type Invoker interface {
	Invoke(inv *Invocation) ([]byte, error)
}

type Invocation struct {
	MethodName  string
	ServiceName string
	Input       []byte
}
