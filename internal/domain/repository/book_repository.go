package repository

import (
	"context"
	"github.com/theguitarvity/go-book-api/internal/domain/entities"
)

type BookRepository interface {
	Create(ctx context.Context, book *entities.Book) error
	FindAll(ctx context.Context) ([]entities.Book, error)
	FindById(ctx context.Context, id string) (*entities.Book, error)
	Update(ctx context.Context, id string, book *entities.Book) error
	Delete(ctx context.Context, id string) error
}
