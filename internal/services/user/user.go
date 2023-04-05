package user

import (
	"context"

	"jikkaem/internal/model"
	"jikkaem/internal/mongodb"
	pb "jikkaem/internal/proto/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) GetUserByID(ctx context.Context, id *pb.ID) (*pb.UserObjectWithID, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	var result model.User
	hex, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = userColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return &pb.UserObjectWithID{
		Id:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, user *pb.UserObject) (*pb.UserObjectWithID, error) {
	mappedUser := &model.User{
		ID:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
	}

	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	_, err = userColl.InsertOne(context.TODO(), mappedUser)
	if err != nil {
		return nil, err
	}

	return &pb.UserObjectWithID{
		Id:    mappedUser.ID.Hex(),
		Name:  mappedUser.Name,
		Email: mappedUser.Email,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, id *pb.ID) (*pb.UserObjectWithID, error) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		return nil, err
	}

	var result *model.User
	hex, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = userColl.FindOneAndDelete(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return &pb.UserObjectWithID{
		Id:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil

}
