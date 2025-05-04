package service

import (
	"context"
	"github.com/theguitarvity/go-book-api/internal/domain/entities"
	"github.com/theguitarvity/go-book-api/internal/domain/repository"
)

type BookService struct {
	Repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, book *entities.Book) error {
	return s.Repo.Create(ctx, book)
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]entities.Book, error) {
	return s.Repo.FindAll(ctx)
}

func (s *BookService) GetBookById(ctx context.Context, id string) (*entities.Book, error) {
	return s.Repo.FindById(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, id string, book *entities.Book) error {
	return s.UpdateBook(ctx, id, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	return s.DeleteBook(ctx, id)
}
