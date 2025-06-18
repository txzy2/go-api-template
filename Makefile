DB_URL=$(DATABASE_URL)

# Применить все миграции вверх
migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

# Откатить одну миграцию вниз
migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

# Откатить все миграции
migrate-drop:
	migrate -path ./migrations -database "$(DB_URL)" drop -f

# Создать новую миграцию (пример: make migrate-new name=create_table)
migrate-new:
	migrate create -ext sql -dir ./migrations -seq $(name)
