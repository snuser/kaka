package services
type UserService struct{
}

func (UserService) ServiceName() string {
	return "user"
}

func (UserService) ShutDown() string {
	return "shutdown userService"
}

func (UserService) GetUser(input *Input) (*Output, error) {
	return &Output{Msg: "getUser called"}, nil
}

