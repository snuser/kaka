package services


type Input struct {
	Name string
}
type Output struct {
	Msg string
}

type Service interface {
	ServiceName() string
	ShutDown() string
}




