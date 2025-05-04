package contract

import (
	"context"
	"github.com/theguitarvity/go-book-api/internal/domain/entities"
)

type BookService interface {
	CreateBook(ctx context.Context, book *entities.Book)
}
