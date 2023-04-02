package service

import (
	"context"
	"fmt"
	"jikkaem/internal/shared/models"
	"jikkaem/internal/shared/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserById(id string) (*models.User, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}
	var result models.User
	filter := bson.D{{Key: "_id", Value: id}}
	if err = userColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found.")
		}
		log.Fatalf("could not get user from db")
	}
	return &result, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	userWithID := &models.User{
		ID:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
	}
	_, err = userColl.InsertOne(context.TODO(), userWithID)
	if err != nil {
		log.Fatalf("could not insert user into db")
	}
	return userWithID, nil
}
