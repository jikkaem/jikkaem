package user_service

import (
	"context"
	"fmt"
	"jikkaem/internal/shared/models"
	"jikkaem/internal/shared/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserById(id string) (*models.UserWithID, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}
	var result models.UserWithID
	filter := bson.D{{Key: "_id", Value: id}}
	if err = userColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found.")
		}
		panic(err)
	}
	return &result, nil
}

func CreateUser(user *models.User) (*models.UserWithID, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	userWithID := &models.UserWithID{
		ID:   primitive.NewObjectID(),
		User: user,
	}
	_, err = userColl.InsertOne(context.TODO(), userWithID)
	if err != nil {
		panic(err)
	}
	return userWithID, nil
}
