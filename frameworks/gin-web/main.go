package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// Test Ping
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Hello message
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework. "})
	})

	// Get all books
	engine.GET("api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllBooks())
	})

	// Create new book
	engine.POST("/api/books", func(c *gin.Context) {
		var book Book
		if c.BindJSON(&book) == nil {
			isbn, created := CreateBook(book)
			if created {
				c.Header("Location", "/api/books/"+isbn)
				c.Status(http.StatusCreated)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	// Get book
	engine.GET("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		book, found := GetBook(isbn)
		if found {
			c.JSON(http.StatusOK, book)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	// Update book
	engine.PUT("api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")

		var book Book
		if c.BindJSON(&book) == nil {
			exists := UpdateBook(isbn, book)
			if exists {
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusNotFound)
			}
		}
	})

	// Delete book
	engine.DELETE("api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		DeleteBook(isbn)
		c.Status(http.StatusOK)
	})

	engine.LoadHTMLGlob("./templates/*.html")
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "GOnGIN",
		})
	})

	// Run server on PORT
	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
