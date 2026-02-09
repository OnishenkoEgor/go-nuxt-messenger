package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sso/domain/user"
)

const (
	dbName         string = "sso"
	collectionName string = "users"
)

type UserRepository struct {
	ctx        *context.Context
	collection *mongo.Collection
}

func NewUserRepository(ctx *context.Context, dbClient *mongo.Client) UserRepository {
	collection := dbClient.Database(dbName).Collection(collectionName)

	return UserRepository{
		ctx:        ctx,
		collection: collection,
	}
}

func (repo UserRepository) Create(user *user.User) error {
	_, err := repo.collection.InsertOne(*repo.ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) GetById(id string) (*user.User, error) {
	var user user.User

	filter := bson.D{{"uuid", id}}
	err := repo.collection.FindOne(*repo.ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) Get() ([]*user.User, error) {
	cursor, err := repo.collection.Find(*repo.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var users []*user.User
	err = cursor.All(*repo.ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo UserRepository) Update(user *user.User) error {
	return nil
}

func (repo UserRepository) Delete(id string) error {
	filter := bson.D{{"uuid", id}}
	fmt.Println(id)
	result, err := repo.collection.DeleteOne(*repo.ctx, filter)
	fmt.Println(result)
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		fmt.Println(result.DeletedCount)
		return errors.New("incorrect removed rows count")
	}

	return nil
}
