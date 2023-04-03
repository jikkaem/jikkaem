package userservice

import (
	"context"
	model "jikkaem/internal/shared/model"
	"jikkaem/internal/shared/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(id string) (*model.User, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	var result model.User
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = userColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteUser(id string) (*model.User, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	var result *model.User
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = userColl.FindOneAndDelete(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user *model.User) (*model.User, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	_, err = userColl.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
