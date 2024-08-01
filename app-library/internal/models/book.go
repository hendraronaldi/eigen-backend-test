package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID         primitive.ObjectID `json:"-" bson:"_id"`
	BookIDs    []string           `json:"book_ids" bson:"book_ids"`
	Code       string             `json:"code" bson:"code"`
	Title      string             `json:"title" bson:"title"`
	Author     string             `json:"author" bson:"author"`
	IsBorrowed bool               `json:"-" bson:"is_borrowed"`
	BorrowedAt *time.Time         `json:"-" bson:"borrowed_at"`
	Stock      int                `json:"stock" bson:"stock"`
}
