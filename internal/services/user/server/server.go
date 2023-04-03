package server

import (
	"context"

	svc "jikkaem/internal/services/user/service"
	"jikkaem/internal/shared/model"
	pb "jikkaem/internal/shared/proto/user"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) GetUserById(ctx context.Context, id *pb.ID) (*pb.UserObjectWithID, error) {
	result, err := svc.GetUserById(id.GetId())
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
