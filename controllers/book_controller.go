package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/theguitarvity/go-book-api/database"
	"github.com/theguitarvity/go-book-api/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func GetBooks(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.BookCollection.Find(ctx, bson.M{})

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on fetching books"})

		return
	}
	defer cursor.Close(ctx)

	var books []entities.Book

	if err := cursor.All(ctx, &books); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on parsing books"})
		return
	}
	ginCtx.JSON(http.StatusOK, books)

}

func CreateBook(ginCtx *gin.Context) {
	var book entities.Book

	if err := ginCtx.ShouldBindJSON(&book); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create book"})
		return
	}
	book.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := database.BookCollection.InsertOne(ctx, book)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})

		return
	}

	ginCtx.JSON(http.StatusCreated, book)
}
func UpdateBook(ginCtx *gin.Context) {
	idParam := ginCtx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}
	var book entities.Book
	if err := ginCtx.ShouldBindJSON(&book); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":       book.Title,
			"author":      book.Author,
			"description": book.Description,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := database.BookCollection.UpdateByID(ctx, id, update)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on updating book"})
	}
	if res.MatchedCount == 0 {
		ginCtx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
	ginCtx.JSON(http.StatusOK, book)
}
func DeleteBook(ginCtx *gin.Context) {
	idParam := ginCtx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid book id"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := database.BookCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil || res.DeletedCount == 0 {
		ginCtx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
	ginCtx.Status(http.StatusNoContent)
}
