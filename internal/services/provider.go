package services

import "gorm.io/gorm"

// Provider содержит все сервисы приложения
type Provider struct {
	UserService     IUserService
	IncidentService IIncidentService
	// Здесь можно добавить другие сервисы по мере их создания
	// например:
	// AuthService IAuthService
	// OrderService IOrderService
	DB *gorm.DB
}

// NewProvider создает новый экземпляр провайдера со всеми сервисами
func NewProvider(db *gorm.DB) *Provider {
	return &Provider{
		UserService:     NewUserService(db),
		IncidentService: NewIncidentService(db),
		// Инициализация других сервисов
		// AuthService: NewAuthService(),
		// OrderService: NewOrderService(),
		DB: db,
	}
}
