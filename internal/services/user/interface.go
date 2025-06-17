package user

type UserInput struct {
	Id string
}

type UserOutput struct {
	Id   string
	Name string
	Age  int
}

// IUserService определяет контракт для сервиса пользователя
type IUserService interface {
	GetUserByID(UserInput) (UserOutput, error)
	CreateUser(name string) (int, error)
}
