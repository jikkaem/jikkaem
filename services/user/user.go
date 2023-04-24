package user

import (
	"context"

	"jikkaem/model"
	"jikkaem/mongodb"
	pb "jikkaem/proto/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database string = "user"
	uri      string = "mongodb://localhost:10000"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) GetUserByID(ctx context.Context, id *pb.ID) (*pb.UserObject, error) {
	coll, err := s.getColl("users")
	if err != nil {
		return nil, err
	}

	var result model.User
	hex, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return &pb.UserObject{
		Id:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, user *pb.UserObject) (*pb.UserObject, error) {
	mappedUser := &model.User{
		ID:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
	}

	coll, err := s.getColl("users")
	if err != nil {
		return nil, err
	}

	_, err = coll.InsertOne(context.TODO(), mappedUser)
	if err != nil {
		return nil, err
	}

	return &pb.UserObject{
		Id:    mappedUser.ID.Hex(),
		Name:  mappedUser.Name,
		Email: mappedUser.Email,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, id *pb.ID) (*pb.UserObject, error) {
	coll, err := s.getColl("users")
	if err != nil {
		return nil, err
	}

	var result *model.User
	hex, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = coll.FindOneAndDelete(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return &pb.UserObject{
		Id:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil

}

func (s *UserServer) getColl(collName string) (*mongo.Collection, error) {
	client, err := mongodb.GetMongoClient(uri)
	if err != nil {
		return nil, err
	}
	return client.Database(database).Collection(collName), nil
}
