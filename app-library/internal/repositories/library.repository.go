package repositories

import (
	"app-library/internal/models"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) FindOneBook(ctx context.Context, oid primitive.ObjectID) (*models.Book, error) {
	res := r.db.Collection(r.collectionBooks).FindOne(ctx, bson.M{"_id": oid, "is_borrowed": bson.M{"$ne": true}})

	data := models.Book{}
	err := res.Decode(&data)
	if err != nil {
		return nil, errors.Wrap(err, "Decode FindOneBook")
	}
	return &data, nil
}

func (r *Repository) FindOneMember(ctx context.Context, oid primitive.ObjectID) (*models.Member, error) {
	res := r.db.Collection(r.collectionMembers).FindOne(ctx, bson.M{"_id": oid})

	data := models.Member{}
	err := res.Decode(&data)
	if err != nil {
		return nil, errors.Wrap(err, "Decode FindOneMember")
	}
	return &data, nil
}

func (r *Repository) FindAllMembers(ctx context.Context) ([]*models.Member, error) {
	res, err := r.db.Collection(r.collectionMembers).Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "FindAllMembers")
	}

	datas := []*models.Member{}
	for res.Next(ctx) {
		data := models.Member{}
		res.Decode(&data)
		if err != nil {
			return nil, errors.Wrap(err, "Decode FindAllMembers")
		}
		datas = append(datas, &data)
	}
	return datas, nil
}

func (r *Repository) FindAllBooks(ctx context.Context) ([]*models.Book, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"is_borrowed": bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"code":   "$code",
					"title":  "$title",
					"author": "$author",
				},
				"stock": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":    primitive.NilObjectID,
				"code":   "$_id.code",
				"title":  "$_id.title",
				"author": "$_id.author",
				"stock":  "$stock",
			},
		},
	}

	cursor, err := r.db.Collection(r.collectionBooks).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "FindAllBooks")
	}
	defer cursor.Close(ctx)

	var results []*models.Book
	err = cursor.All(context.Background(), &results)
	if err != nil {
		return nil, errors.Wrap(err, "parse FindAllBooks")
	}
	return results, nil
}

func (r *Repository) UpdateMember(ctx context.Context, m *models.Member) error {
	rs := r.db.Collection(r.collectionMembers).FindOneAndUpdate(ctx, bson.M{
		"_id": m.ID,
	}, bson.M{"$set": m})

	if rs.Err() != nil {
		return errors.Wrap(rs.Err(), "UpdateMember")
	}
	return nil
}

func (r *Repository) UpdateBook(ctx context.Context, b *models.Book) error {
	rs := r.db.Collection(r.collectionBooks).FindOneAndUpdate(ctx, bson.M{
		"_id": b.ID,
	}, bson.M{"$set": bson.M{"is_borrowed": b.IsBorrowed}})

	if rs.Err() != nil {
		return errors.Wrap(rs.Err(), "UpdateBook")
	}
	return nil
}
