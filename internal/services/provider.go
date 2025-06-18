package services

import "database/sql"

// Provider содержит все сервисы приложения
type Provider struct {
	UserService IUserService
	// Здесь можно добавить другие сервисы по мере их создания
	// например:
	// AuthService IAuthService
	// OrderService IOrderService
}

// NewProvider создает новый экземпляр провайдера со всеми сервисами
func NewProvider(db *sql.DB) *Provider {
	return &Provider{
		UserService: NewUserService(db),
		// Инициализация других сервисов
		// AuthService: NewAuthService(),
		// OrderService: NewOrderService(),
	}
}
