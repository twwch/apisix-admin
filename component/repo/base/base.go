package base

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseIRepository interface {
	FindOne(ctx context.Context, filter bson.M, result interface{}, opts ...*options.FindOneOptions) error
	Find(ctx context.Context, filter bson.M, results interface{}, opts ...*options.FindOptions) error
	Count(ctx context.Context, filter bson.M, opts ...*options.CountOptions) (int64, error)
	DeleteOne(ctx context.Context, filter bson.M) error
	DeleteMany(ctx context.Context, filter bson.M) error
	UpdateOne(ctx context.Context, filter bson.M, replacement bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Upsert(ctx context.Context, filter bson.M, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter bson.M, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}