package services

import (
	"database/sql"
	"fmt"
)

type (
	IUserService interface {
		GetUserByID(UserInput) (UserOutput, error)
		// CreateUser(UserInput) (UserOutput, error)
	}

	UserOutput struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	UserInput struct {
		Id string `json:"id"`
	}

	UserService struct {
		db *sql.DB
	}
)

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// GetUserByID возвращает пользователя по ID или ошибку, если пользователь не найден.
func (dbs *UserService) GetUserByID(input UserInput) (UserOutput, error) {
	var user UserOutput
	const query = "SELECT id, name, age, email, created_at, updated_at FROM users WHERE id = $1"

	err := dbs.db.QueryRow(query, input.Id).
		Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return UserOutput{}, sql.ErrNoRows
		}

		return UserOutput{}, err
	}

	return user, nil
}

func (s *UserService) CreateUser(name string) (int, error) {
	fmt.Printf("Создан пользователь: %s\n", name)
	return 2, nil // Возвращаем пример ID
}
