package search

import (
	"context"
	pb "jikkaem/internal/proto/search"
)

type SearchServer struct {
	pb.UnimplementedSearchServer
}

func (s *SearchServer) SearchBar(ctx context.Context, id *pb.Text) (*pb.FancamList, error) {
	return nil, nil
}
