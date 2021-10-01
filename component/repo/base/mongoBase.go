package base

import (
	"context"
	"github.com/twwch/gin-sdk/mongoBase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoBase struct {
	*mongoBase.Base
	softDelete     bool
	deletedAtField string
	updatedAtField string
	createdAtField string
}

func NewBaseModel(host, db, collection string) *mongoBase.Base {
	return mongoBase.NewBase(host, db, collection)
}

func (m *MongoBase) DeletedAtField() string {
	return m.deletedAtField
}

func (m *MongoBase) UpdatedAtField() string {
	return m.updatedAtField
}

func (m *MongoBase) CreatedAtField() string {
	return m.createdAtField
}

func NewMongoBase(host, db, collection string) *MongoBase {
	return &MongoBase{
		Base:           NewBaseModel(host, db, collection),
		updatedAtField: "updated_at",
		createdAtField: "created_at",
		deletedAtField: "deleted_at",
	}
}

func (m *MongoBase) SetSoftDelete(isOpen bool) *MongoBase {
	m.softDelete = isOpen
	return m
}

func (m *MongoBase) FindOne(ctx context.Context, filter bson.M, result interface{}, opts ...*options.FindOneOptions) error {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}
	}
	return m.Base.FindOne(ctx, filter, result, opts...)
}

func (m *MongoBase) Find(ctx context.Context, filter bson.M, results interface{}, opts ...*options.FindOptions) error {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}
	}
	return m.Base.Find(ctx, filter, results, opts...)
}

func (m *MongoBase) Count(ctx context.Context, filter bson.M, opts ...*options.CountOptions) (int64, error) {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}
	}
	return m.Base.Count(ctx, filter, opts...)
}

func (m *MongoBase) DeleteOne(ctx context.Context, filter bson.M) error {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}

		_, err := m.Base.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				m.deletedAtField: time.Now().Unix(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	}

	_, err := m.Base.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoBase) DeleteMany(ctx context.Context, filter bson.M) error {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}

		_, err := m.Base.UpdateMany(ctx, filter, bson.M{
			"$set": bson.M{
				m.deletedAtField: time.Now().Unix(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	}

	_, err := m.Base.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoBase) UpdateOne(ctx context.Context, filter bson.M, replacement bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}
		if _, ok := replacement["$set"]; ok {
			now := time.Now().Unix()

			setter, err := m.convertToDoc(ctx, replacement["$set"])
			if err != nil {
				return nil, err
			}
			setter[m.updatedAtField] = now
			replacement["$set"] = setter
		}
	}
	return m.Base.UpdateOne(ctx, filter, replacement, opts...)
}

func (m *MongoBase) Upsert(ctx context.Context, filter bson.M, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}
	}
	return m.Base.Upsert(ctx, filter, replacement, opts...)
}

func (m *MongoBase) UpdateMany(ctx context.Context, filter bson.M, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.softDelete {
		filter[m.deletedAtField] = bson.M{
			"$eq": 0,
		}
	}
	return m.Base.UpdateMany(ctx, filter, replacement, opts...)
}

func (m *MongoBase) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	doc, err := m.convertToDoc(ctx, document)
	if err != nil {
		return nil, err
	}

	if m.softDelete {
		now := time.Now().Unix()
		doc[m.createdAtField] = now
		doc[m.updatedAtField] = now
		doc[m.deletedAtField] = int64(0)
	}

	return m.Base.InsertOne(ctx, doc, opts...)
}

func (m *MongoBase) convertToDoc(ctx context.Context, document interface{}) (primitive.M, error) {
	var doc primitive.M

	roleFeatureRaw, err := bson.Marshal(document)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(roleFeatureRaw, &doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
