package main

import (
	"github.com/aqibcs/go-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Use the booksRouter for book-related routes
	booksRouter := router.Group("/books")
	{
		booksRouter.GET("", handlers.GetBooks)
		booksRouter.GET("/:id", handlers.BookByID)
		booksRouter.POST("", handlers.CreateBook)
		booksRouter.PATCH("/checkout", handlers.CheckoutBook)
		booksRouter.PATCH("/return", handlers.ReturnBook)
	}

	router.Run("localhost:8080")
}
