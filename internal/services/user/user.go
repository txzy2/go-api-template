package user

import (
	"fmt"
)

type UserService struct {
	// Здесь могут быть поля для зависимостей, например, подключение к БД
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserByID(input UserInput) (UserOutput, error) {
	if input.Id == "2" {
		return UserOutput{
			Id:    input.Id,
			Name:  "Anton",
			Age:   24,
			Email: "kamaeff2@gmail.com",
		}, nil
	}
	return UserOutput{}, fmt.Errorf("пользователь с ID %s не найден", input.Id)
}

func (s *UserService) CreateUser(name string) (int, error) {
	fmt.Printf("Создан пользователь: %s\n", name)
	return 2, nil // Возвращаем пример ID
}
