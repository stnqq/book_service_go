package main

import (
	"books-service/pkg/db"
	"books-service/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Инициализация маршрутов
	r := gin.Default()
	r.Use(handlers.LoggingMiddleware())

	r.POST("/books", handlers.CreateBook(database))
	r.GET("/books/:id", handlers.GetBookByID(database))

	// Запуск сервиса
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
