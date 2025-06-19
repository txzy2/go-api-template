package services

import (
	"fmt"
	"time"

	"github.com/txzy2/simple-api/internal/models"
	"gorm.io/gorm"
)

type (
	// публичные методы
	IUserService interface {
		GetUserByID(UserInput) (UserOutput, error)
		CreateUser(CreateUserInput) (uint, error)
	}

	UserOutput struct {
		Id        uint   `json:"id"`
		Name      string `json:"name"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	UserInput struct {
		Id string `json:"id" binding:"required"`
	}

	CreateUserInput struct {
		Name  string `json:"name" binding:"required"`
		Age   int    `json:"age" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	UserService struct {
		db *gorm.DB
	}
)

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetUserByID возвращает пользователя по ID или ошибку, если пользователь не найден.
func (s *UserService) GetUserByID(input UserInput) (UserOutput, error) {
	var user models.User

	result := s.db.First(&user, "id = ?", input.Id)
	if result.Error != nil {
		return UserOutput{}, result.Error
	}
	// Преобразуем User в UserOutput, если нужно
	return UserOutput{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) CreateUser(newUser CreateUserInput) (uint, error) {
	user := models.User{
		Name:  newUser.Name,
		Age:   newUser.Age,
		Email: newUser.Email,
	}
	result := s.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	fmt.Printf("Создан пользователь: %s\n", user.Name)
	return user.Id, nil // Возвращаем ID созданного пользователя
}
