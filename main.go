package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theguitarvity/go-book-api/controllers"
	"github.com/theguitarvity/go-book-api/database"
)

func main() {
	database.Connect()

	engine := gin.Default()

	api := engine.Group("/api")
	{
		api.GET("/books", controllers.GetBooks)
		api.POST("/books", controllers.CreateBook)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)
	}

	err := engine.Run(":8080")
	if err != nil {
		return
	}

}
