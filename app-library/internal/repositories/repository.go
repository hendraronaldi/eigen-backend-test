package repositories

import (
	"app-library/internal/app_config"
	"app-library/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface: an abstraction of library repository
type Interface interface {
	FindOneMember(ctx context.Context, oid primitive.ObjectID) (*models.Member, error)
	FindOneBook(ctx context.Context, oid primitive.ObjectID) (*models.Book, error)
	FindAllMembers(ctx context.Context) ([]*models.Member, error)
	FindAllBooks(ctx context.Context) ([]*models.Book, error)
	UpdateMember(ctx context.Context, m *models.Member) error
	UpdateBook(ctx context.Context, b *models.Book) error
}

// Repository: repository of library
type Repository struct {
	db                *mongo.Database
	collectionBooks   string
	collectionMembers string
}

// NewRepository: to initialize repository of user
func NewRepository(pdb *mongo.Database) *Repository {
	config := app_config.Get()
	return &Repository{
		db:                pdb,
		collectionBooks:   config.ENV.COLLECTION_BOOKS,
		collectionMembers: config.ENV.COLLECTION_MEMBERS,
	}
}
