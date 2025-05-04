package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theguitarvity/go-book-api/internal/adapter/in/rest"
	"github.com/theguitarvity/go-book-api/internal/adapter/out/mongo"
	"github.com/theguitarvity/go-book-api/internal/application/service"
	"github.com/theguitarvity/go-book-api/internal/infrastructure/database"
)

func main() {
	bookCollection := database.Connect()

	repo := mongo.NewBookMongoRepository(bookCollection)

	bookService := service.NewBookService(repo)

	bookHandler := rest.NewHandler(bookService)

	engine := gin.Default()

	bookHandler.RegisterRoutes(engine)

	err := engine.Run(":8080")
	if err != nil {
		return
	}

}
