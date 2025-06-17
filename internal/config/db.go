package config

// DBConfig конфигурация базы данных
type DBConfig struct {
	Host string
	Name string
	User string
	Pass string
	Port string
}

// LoadDBConfig загружает конфигурацию БД
func LoadDBConfig() DBConfig {
	return DBConfig{
		Host: getEnv("DB_HOST", "localhost"),
		Name: getEnv("DB_NAME", "mydb"),
		User: getEnv("DB_USER", "user"),
		Pass: getEnv("DB_PASS", "pass"),
		Port: getEnv("DB_PORT", "5432"),
	}
}
