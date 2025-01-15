package handlers

import (
	"books-service/pkg/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		query := `INSERT INTO books (title, author, created_at) VALUES ($1, $2, NOW()) RETURNING id`
		err := db.QueryRow(query, book.Title, book.Author).Scan(&book.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": book.ID})
	}
}

func GetBookByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book

		query := `SELECT id, title, author, created_at FROM books WHERE id = $1`
		err := db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.CreatedAt)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}
