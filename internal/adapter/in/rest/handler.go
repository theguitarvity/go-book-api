package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/theguitarvity/go-book-api/internal/application/service"
	"github.com/theguitarvity/go-book-api/internal/domain/entities"
	"net/http"
)

type Handler struct {
	Service *service.BookService
}

func NewHandler(s *service.BookService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) RegisterRoutes(engine *gin.Engine) {
	engine.POST("/books", h.CreateBook)
	engine.GET("/books", h.GetBooks)
	engine.GET("/books/:id", h.GetBook)
	engine.PUT("/books/:id", h.UpdateBook)
	engine.DELETE("/books/:id", h.UpdateBook)

}

func (h *Handler) CreateBook(ginContext *gin.Context) {
	var book entities.Book
	if err := ginContext.ShouldBindJSON(&book); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateBook(ginContext.Request.Context(), &book)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "book created"})
}

func (h *Handler) GetBooks(ginContext *gin.Context) {
	books, err := h.Service.GetAllBooks(ginContext)

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, books)
}

func (h *Handler) GetBook(ginContext *gin.Context) {
	id := ginContext.Param("id")
	book, err := h.Service.GetBookById(ginContext, id)

	if err != nil {
		ginContext.JSON(http.StatusNotFound, gin.H{"error": "not Found"})
		return
	}

	ginContext.JSON(http.StatusOK, book)
}

func (h *Handler) UpdateBook(ginContext *gin.Context) {
	id := ginContext.Param("id")
	var book entities.Book
	if err := ginContext.ShouldBindJSON(&book); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateBook(ginContext, id, &book); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *Handler) DeleteBook(ginContext *gin.Context) {
	id := ginContext.Param("id")

	if err := h.Service.DeleteBook(ginContext, id); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
