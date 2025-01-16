# Books Service

## Описание
    Сервис для управления книгами с использованием Gin и PostgreSQL.

## Инструкция по запуску и использованию

    0. Подготовительный этап:
        0.1. Если у вас нет docker'а установите docker
    1. Клонирование проекта: 
        - git clone https://github.com/stnqq/book_service_go.git
    2. Переход в папку проекта: 
        - cd book_service_go
    3. Сборка и запуск контейнера:
        - docker-compose up --build
    4. Проверка сервиса: 
    - Добавление книги (POST /books)
        - Принимает JSON с полями:

            {
                "title": "Название книги",
                "author": "Автор книги"
            }

        - Возвращает ID созданной книги и статус `201 Created`.

        Пример запроса (через Postman):
            POST http://localhost:8080/books
                {
                    "Title": "New Test Book",
                    "Author": "Doe John"
                }
            Ответ на запрос:
                201 Created
                {
                    "id": 3
                }

    - Получение информации о книге по ID (GET /books/:id)
        - Возвращает JSON с информацией о книге:
            {
                "id": 1,
                "title": "Название книги",
                "author": "Автор книги",
                "created_at": "Дата создания"
            }
            
            - Если книга не найдена, возвращает статус `404 Not Found`.

        Пример запроса (через Postman):
            GET http://localhost:8080/books/2

            Ответ на запрос:
                200 OK
                {
                    "id": 1,
                    "title": "Go Programming",
                    "author": "John Doe",
                    "created_at": "2025-01-15T23:04:31.261519Z"
                }
    5. Завершение проекта:
        - Если вы в терминале находитесь в контейнере, 
            то используйте сочетание клавиш ctrl + c
        - Чтобы удалить контейнер, выполните команду:
            - docker-compose down

## Docker Compose
    - Сервис и база данных PostgreSQL запускаются через `docker-compose`.
    - Сервис доступен на порту `8080`.

## Эндпоинты
### - Добавление книги (POST /books)
    - Принимает JSON с полями:

       {
         "title": "Название книги",
         "author": "Автор книги"
       }

    - Возвращает ID созданной книги и статус `201 Created`.

    Пример запроса (через Postman):

    POST http://localhost:8080/books
        {
            "Title": "New Test Book",
            "Author": "Doe John"
        }
    Ответ на запрос:
        201 Created
        {
            "id": 3
        }

### - Получение информации о книге по ID (GET /books/:id)
    - Возвращает JSON с информацией о книге:
       {
         "id": 1,
         "title": "Название книги",
         "author": "Автор книги",
         "created_at": "Дата создания"
       }
       
     - Если книга не найдена, возвращает статус `404 Not Found`.

    Пример запроса (через Postman):

    GET http://localhost:8080/books/2

    Ответ на запрос:
        200 OK
        {
            "id": 1,
            "title": "Go Programming",
            "author": "John Doe",
            "created_at": "2025-01-15T23:04:31.261519Z"
        }

## Интеграция с PostgreSQL
    - Подключение к базе данных через переменные окружения.
    - Хранение данных о книгах в базе данных.
    - Создана таблица `books` с полями:
        - `id`
        - `title`
        - `author`
        - `created_at`
    
# *Опциональные трeбования*
## Middleware для логирования запросов
    - Логирует метод, URL, статус код и время выполнения каждого запроса.
    - Пример лога:
     ```
     [Middleware] POST /books 201 26.345ms
     ```

## Миграция для базы данных
    - Использован инструмент [pressly/goose]
        - (https://github.com/pressly/goose).
    - Создание файла для миграции:
        - goose -dir pkg/db/migrations create books_table sql
    - Создание таблицы `books` с полями:

        id (UUID или SERIAL, Primary Key),
        title (VARCHAR, Not Null),
        author (VARCHAR, Not Null),
        created_at (TIMESTAMP, Not Null, default NOW()).

    - Реализована миграция для создания таблицы `books`:
        - goose -dir pkg/db/migrations postgres "postgresql://user:password@localhost:5432/books_db?sslmode=disable" up

## Multi-stage Dockerfile
    - Образ с минимальным размером.
    - Первый этап: компиляция Go-приложения с использованием базового образа Go.
    - Второй этап: запуск приложения на облегчённом образе Alpine.

## Используемые инструменты
1. **Go (Golang) version: 1.23.4**  — основной язык программирования.
2. **Gin version: 1.10.0** — веб-фреймворк для обработки HTTP-запросов.
3. **PostgreSQL version: 15** — база данных для хранения данных о книгах.
4. **Docker и Docker Compose version: 27.4.0 and 2.31.0** — для контейнеризации и управления сервисами.
5. **pressly/goose version: 3.24.1** — инструмент для выполнения миграций.
6. **Alpine Linux version: 3.21.2** — минималистичный образ для финального этапа сборки Docker.