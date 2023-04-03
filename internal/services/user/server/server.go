package server

import (
	"context"

	svc "jikkaem/internal/services/user/service"
	"jikkaem/internal/shared/model"
	pb "jikkaem/internal/shared/proto/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) GetUserByID(ctx context.Context, id *pb.ID) (*pb.UserObjectWithID, error) {
	result, err := svc.GetUserById(id.Id)
	if err != nil {
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

	result, err := svc.CreateUser(mappedUser)
	if err != nil {
		return nil, err
	}

	return &pb.UserObjectWithID{
		Id:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, id *pb.ID) (*pb.UserObjectWithID, error) {
	res, err := svc.DeleteUser(id.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserObjectWithID{
		Id:    res.ID.Hex(),
		Name:  res.Name,
		Email: res.Email,
	}, nil

}
