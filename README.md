# Books Service

## Описание
Сервис для управления книгами с использованием Gin и PostgreSQL.

## Установка
1. Установите Docker и Docker Compose.
2. Выполните `docker-compose up --build`.

## Эндпоинты
- POST `/books`
- GET `/books/:id`

## Миграция
Скачал goose
Проверил версию goose
Создал файл для миграции: goose -dir pkg/db/migrations create new_user_table sql
Накатил миграцию: goose -dir pkg/db/migrations postgres "postgresql://user:password@localhost:5432/books_db?sslmode=disable" up