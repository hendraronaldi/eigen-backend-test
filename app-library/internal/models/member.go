package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID                 primitive.ObjectID `json:"-" bson:"_id"`
	Code               string             `json:"code" bson:"code"`
	Name               string             `json:"name" bson:"name"`
	BoorowedBooks      []*Book            `json:"-" bson:"borrowed_books"`
	TotalBorrowedBooks int                `json:"total_borrowed_books" bson:"-"`
	PenalizedAt        *time.Time         `json:"-" bson:"penalized_at"`
}
