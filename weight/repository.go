package weight

import (
	"context"

	"github.com/ijalalfrz/sirclo-weight-test/entity"
	"github.com/ijalalfrz/sirclo-weight-test/exception"
	"github.com/ijalalfrz/sirclo-weight-test/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository is collection of behaviour weightRepository
type Repository interface {
	InsertOne(ctx context.Context, weight entity.Weight) (err error)
	UpdateOne(ctx context.Context, key int64, weight entity.Weight) (err error)
	FindMany(ctx context.Context, sortBy string, sort int) (bunchOfWeight []entity.Weight, err error)
	FindOne(ctx context.Context, key int64) (weight entity.Weight, err error)
}

type weightRepository struct {
	logger *logrus.Logger
	col    mongodb.Collection
}

// NewWeightRepository is a constructor.
func NewWeightRepository(logger *logrus.Logger, db mongodb.Database) Repository {
	col := db.Collection("weight")
	return &weightRepository{logger, col}
}

func (r weightRepository) InsertOne(ctx context.Context, weight entity.Weight) (err error) {
	_, err = r.col.InsertOne(ctx, weight)
	if err != nil {
		r.logger.Error(err)
		err = exception.ErrInternalServer
		return
	}
	return
}
func (r weightRepository) UpdateOne(ctx context.Context, key int64, weight entity.Weight) (err error) {
	filter := bson.M{
		"date": key,
	}

	updatedData := map[string]interface{}{
		"$set": weight,
	}

	updatedResult, err := r.col.UpdateOne(ctx, filter, updatedData, options.Update().SetUpsert(true))
	if err != nil {
		r.logger.Error(err)
		err = exception.ErrInternalServer
		return
	}

	if updatedResult.MatchedCount < 1 {
		err = exception.ErrNotFound
		return
	}

	return
}
func (r weightRepository) FindMany(ctx context.Context, sortBy string, sort int) (bunchOfWeight []entity.Weight, err error) {
	qSort := map[string]int{
		sortBy: sort,
	}
	opt := options.Find().SetSort(qSort)
	filter := map[interface{}]interface{}{}
	cursor, err := r.col.Find(ctx, filter, opt)
	if err != nil {
		r.logger.Error(err)
		err = exception.ErrInternalServer
		return
	}

	for cursor.Next(ctx) {
		weight := entity.Weight{}
		if err = cursor.Decode(&weight); err != nil {
			err = exception.ErrInternalServer
			return
		}

		bunchOfWeight = append(bunchOfWeight, weight)
	}

	if len(bunchOfWeight) < 1 {
		err = exception.ErrNotFound
		return
	}

	return
}
func (r weightRepository) FindOne(ctx context.Context, key int64) (weight entity.Weight, err error) {
	filter := bson.M{
		"date": key,
	}

	if err = r.col.FindOne(ctx, filter).Decode(&weight); err != nil {
		if err != mongo.ErrNoDocuments {
			r.logger.Error(err)
			err = exception.ErrInternalServer
			return
		}
		err = exception.ErrNotFound
		return
	}

	return
}
