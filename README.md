# Go API

Простое REST API на Go с использованием Gin, PostgreSQL и Docker.

## Требования

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Опционально) [Make](https://www.gnu.org/software/make/) для локального запуска миграций

## Быстрый старт

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/txzy2/go-api-template.git
   cd go-api
   ```

2. **Создайте файл окружения `.env`**  
   Пример содержимого:
   ```
   DB_NAME=goapi
   DB_USER=postgres
   DB_PASS=password
   DATABASE_URL=postgres://postgres:password@db:5432/goapi?sslmode=disable
   ```

3. **Запустите проект:**
   ```bash
   docker-compose up --build
   ```
   - Приложение будет доступно на [http://localhost:8080](http://localhost:8080)
   - База данных PostgreSQL будет доступна на порту 5432

4. **Выполните миграции:**
   В новом терминале выполните:
   ```bash
   docker exec -it go_app make migrate-up
   ```
   Это применит миграции из папки `migrations/` к базе данных.

## Структура проекта

- `internal/` — бизнес-логика и сервисы
- `config/` — настройка роутов и инициализация приложения
- `pkg/database/` — подключение к базе данных
- `migrations/` — SQL-файлы для миграций схемы БД
- `Dockerfile`, `docker-compose.yml` — контейнеризация и оркестрация

## Основные команды

- **Сборка и запуск контейнеров:**
  ```bash
  docker-compose up --build
  ```
- **Остановка контейнеров:**
  ```bash
  docker-compose down
  ```
- **Применить миграции:**
  ```bash
  docker exec -it go_app make migrate-up
  ```
- **Откатить последнюю миграцию:**
  ```bash
  docker exec -it go_app make migrate-down
  ```

## Пример API

- `GET /api/user/:id` — получить пользователя по ID

## Примечания

- Для разработки используется hot-reload через [air](https://github.com/air-verse/air).
- Все данные БД сохраняются в Docker volume `postgres_data` и не теряются при перезапуске контейнеров (если не использовать `down -v`).
