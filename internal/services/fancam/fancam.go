package fancam

import (
	"context"
	"jikkaem/internal/model"
	"jikkaem/internal/mongodb"
	pb "jikkaem/internal/proto/fancam"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	database string = "fancam"
	uri      string = "mongodb://localhost:10001"
)

type FancamServer struct {
	pb.UnimplementedFancamServer
}

func (s *FancamServer) GetFancam(ctx context.Context, input *pb.GetFancamRequest) (*pb.FancamObject, error) {
	coll, err := s.getColl("fancams")
	if err != nil {
		return nil, err
	}

	var result model.Fancam
	hex, err := primitive.ObjectIDFromHex(input.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	// Convert suggestedTags into grpc model
	suggestedTags := &pb.SuggestedTags{
		EnArtist: result.SuggestedTags.EnArtist,
		EnGroup:  result.SuggestedTags.EnGroup,
		EnSong:   result.SuggestedTags.EnSong,
		KrArtist: result.SuggestedTags.KrArtist,
		KrGroup:  result.SuggestedTags.KrGroup,
		KrSong:   result.SuggestedTags.KrSong,
	}

	return &pb.FancamObject{
		Id:            result.ID,
		Title:         result.Title,
		Description:   result.Description,
		PublishedAt:   timestamppb.New(result.PublishedAt),
		ChannelId:     result.ChannelID,
		ChannelTitle:  result.ChannelTitle,
		RootThumbnail: result.RootThumbnail,
		RecordDate:    timestamppb.New(result.RecordDate),
		SuggestedTags: suggestedTags,
	}, nil
}

func (s *FancamServer) GetFancams(ctx context.Context, input *pb.GetFancamsRequest) (*pb.FancamList, error) {
	// Process input, convert to primitive.ObjectID
	ids := input.GetIds()
	if lenIds := len(ids); lenIds > 50 {
		err := status.Errorf(codes.FailedPrecondition, "Number of IDs exceeded limit of 50, received %d", lenIds)
		return nil, err
	}
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			err := status.Error(codes.InvalidArgument, "Provided string is not a hex value")
			return nil, err
		}
		objectIds = append(objectIds, objectId)
	}

	coll, err := s.getColl("fancams")
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objectIds}}}}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []model.Fancam
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	fancamList := &pb.FancamList{
		Fancams: []*pb.FancamObject{},
	}

	for _, result := range results {
		// Convert suggestedTags into grpc model
		suggestedTags := &pb.SuggestedTags{
			EnArtist: result.SuggestedTags.EnArtist,
			EnGroup:  result.SuggestedTags.EnGroup,
			EnSong:   result.SuggestedTags.EnSong,
			KrArtist: result.SuggestedTags.KrArtist,
			KrGroup:  result.SuggestedTags.KrGroup,
			KrSong:   result.SuggestedTags.KrSong,
		}

		tmp := &pb.FancamObject{
			Id:            result.ID,
			Title:         result.Title,
			Description:   result.Description,
			PublishedAt:   timestamppb.New(result.PublishedAt),
			ChannelId:     result.ChannelID,
			ChannelTitle:  result.ChannelTitle,
			RootThumbnail: result.RootThumbnail,
			RecordDate:    timestamppb.New(result.RecordDate),
			SuggestedTags: suggestedTags,
		}
		fancamList.Fancams = append(fancamList.Fancams, tmp)
	}

	return fancamList, err
}

func (s *FancamServer) GetFancamsLatest(ctx context.Context, input *pb.GetFancamsLatestRequest) (*pb.FancamList, error) {
	// Validate max_results input
	maxResults := input.GetMaxResults()
	if maxResults > 50 {
		err := status.Errorf(codes.FailedPrecondition, "maxResults received is greater than 50, received %d", maxResults)
		return nil, err
	}

	// Get collection
	coll, err := s.getColl("fancams")
	if err != nil {
		return nil, err
	}

	// Build query
	// No filter
	filter := bson.D{}
	// Sorts entries from newest to oldest
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})
	// Sets limit on how many docs to fetch
	opts = opts.SetLimit(int64(maxResults))
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	// Get results
	var results []model.Fancam
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	// Map results into gRPC model
	fancamList := &pb.FancamList{
		Fancams: []*pb.FancamObject{},
	}
	for _, result := range results {
		// Convert suggestedTags into grpc model
		suggestedTags := &pb.SuggestedTags{
			EnArtist: result.SuggestedTags.EnArtist,
			EnGroup:  result.SuggestedTags.EnGroup,
			EnSong:   result.SuggestedTags.EnSong,
			KrArtist: result.SuggestedTags.KrArtist,
			KrGroup:  result.SuggestedTags.KrGroup,
			KrSong:   result.SuggestedTags.KrSong,
		}

		mappedFancam := &pb.FancamObject{
			Id:            result.ID,
			Title:         result.Title,
			Description:   result.Description,
			PublishedAt:   timestamppb.New(result.PublishedAt),
			ChannelId:     result.ChannelID,
			ChannelTitle:  result.ChannelTitle,
			RootThumbnail: result.RootThumbnail,
			RecordDate:    timestamppb.New(result.RecordDate),
			SuggestedTags: suggestedTags,
		}
		fancamList.Fancams = append(fancamList.Fancams, mappedFancam)
	}

	return fancamList, err
}

func (s *FancamServer) CreateFancams(ctx context.Context, input *pb.FancamList) (*emptypb.Empty, error) {
	// Convert grpc fancamlist into mongodb model
	inputFancams := input.GetFancams()
	fancams := []model.Fancam{}

	for _, fancam := range inputFancams {
		mappedTag := model.SuggestedTags{
			EnArtist: fancam.SuggestedTags.EnArtist,
			EnGroup:  fancam.SuggestedTags.EnGroup,
			EnSong:   fancam.SuggestedTags.EnSong,
			KrArtist: fancam.SuggestedTags.KrArtist,
			KrGroup:  fancam.SuggestedTags.KrGroup,
			KrSong:   fancam.SuggestedTags.KrSong,
		}

		mappedFancam := model.Fancam{
			ID:            fancam.Id,
			Title:         fancam.GetTitle(),
			Description:   fancam.GetDescription(),
			PublishedAt:   fancam.PublishedAt.AsTime(),
			ChannelID:     fancam.GetChannelId(),
			ChannelTitle:  fancam.GetChannelTitle(),
			RootThumbnail: fancam.GetRootThumbnail(),
			RecordDate:    fancam.RecordDate.AsTime(),
			SuggestedTags: mappedTag,
		}
		fancams = append(fancams, mappedFancam)
	}

	// Insert list into database
	coll, err := s.getColl("fancams")
	if err != nil {
		return nil, err
	}

	for _, fancam := range fancams {
		filter := bson.D{{Key: "_id", Value: fancam.ID}}
		update := bson.D{{Key: "$set", Value: fancam}}
		opts := options.Update().SetUpsert(true)
		_, err = coll.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *FancamServer) DeleteFancam(ctx context.Context, id *pb.DeleteFancamRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *FancamServer) getColl(collName string) (*mongo.Collection, error) {
	client, err := mongodb.GetMongoClient(uri)
	if err != nil {
		return nil, err
	}
	return client.Database(database).Collection(collName), nil
}
