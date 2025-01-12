package repository

import (
	"context"
	"errors"
	"time"

	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/mongodb"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IFooRepository interface {
	GetAll() ([]*model.Foo, error)
	GetById(id string) (*model.Foo, error)
	Create(foo *model.Foo) error
	Update(foo *model.Foo) error
}

type FooRepository struct {
	db           *mongodb.MongoDB
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewFooRepository(db *mongodb.MongoDB, cfg *config.Config) *FooRepository {
	return &FooRepository{
		db:           db,
		collection:   db.Client.Database("fiber").Collection("foo"),
		queryTimeout: cfg.Mongo.QueryTimeout,
	}
}

func (f *FooRepository) GetAll() ([]*model.Foo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), f.queryTimeout)
	defer cancel()

	cursor, err := f.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var foos []*model.Foo
	if err = cursor.All(ctx, &foos); err != nil {
		return nil, err
	}
	return foos, nil
}

func (f *FooRepository) GetById(id string) (*model.Foo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), f.queryTimeout)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var foo model.Foo
	err = f.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &foo, nil
}

func (f *FooRepository) Create(foo *model.Foo) error {
	ctx, cancel := context.WithTimeout(context.Background(), f.queryTimeout)
	defer cancel()

	_, err := f.collection.InsertOne(ctx, foo)
	return err
}

func (f *FooRepository) Update(foo *model.Foo) error {
	ctx, cancel := context.WithTimeout(context.Background(), f.queryTimeout)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(foo.ID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":        foo.Name,
			"description": foo.Description,
		},
	}

	result, err := f.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		update,
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("not found")
	}

	if result.ModifiedCount == 0 {
		return errors.New("cannot update")
	}

	return nil
}
