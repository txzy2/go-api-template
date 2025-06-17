package user

type UserInput struct {
	Id string `json:"id"`
}

type UserOutput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// IUserService определяет контракт для сервиса пользователя
type IUserService interface {
	GetUserByID(UserInput) (UserOutput, error)
	// CreateUser(UserInput) (UserOutput, error)
}
