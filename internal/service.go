package internal

type UserService interface {
	GetUser(id uint) string
}

type userService struct{}

func NewUserService() UserService {
	return userService{}
}

func (us userService) GetUser(id uint) string {
	return "hello world"
}
