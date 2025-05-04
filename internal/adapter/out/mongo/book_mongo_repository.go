package mongo

import (
	"context"
	"github.com/theguitarvity/go-book-api/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookMongoRepository struct {
	Collection *mongo.Collection
}

func NewBookMongoRepository(col *mongo.Collection) *BookMongoRepository {
	return &BookMongoRepository{Collection: col}
}

func (repository *BookMongoRepository) Create(ctx context.Context, book *entities.Book) error {
	_, err := repository.Collection.InsertOne(ctx, book)
	return err
}

func (repository *BookMongoRepository) FindAll(ctx context.Context) ([]entities.Book, error) {
	cursor, err := repository.Collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []entities.Book

	if err := cursor.All(ctx, &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (repository *BookMongoRepository) FindById(ctx context.Context, id string) (*entities.Book, error) {
	var book entities.Book
	err := repository.Collection.FindOne(ctx, bson.M{"id": id}).Decode(&book)

	return &book, err
}

func (repository *BookMongoRepository) Update(ctx context.Context, id string, book *entities.Book) error {
	_, err := repository.Collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": book})

	return err
}

func (repository *BookMongoRepository) Delete(ctx context.Context, id string) error {
	_, err := repository.Collection.DeleteOne(ctx, bson.M{"id": id})

	return err
}
