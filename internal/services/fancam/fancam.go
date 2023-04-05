package fancam

import (
	"context"
	pb "jikkaem/internal/proto/fancam"
)

type FancamServer struct {
	pb.UnimplementedFancamServer
}

func (s *FancamServer) GetFancamByID(ctx context.Context, id *pb.ID) (*pb.FancamObjectWithID, error) {
	return nil, nil
}

func (s *FancamServer) CreateFancams(ctx context.Context, fancam *pb.FancamList) (*pb.FancamListWithID, error) {
	return nil, nil
}

func (s *FancamServer) DeleteFancam(ctx context.Context, id *pb.ID) (*pb.FancamObjectWithID, error) {
	return nil, nil
}
