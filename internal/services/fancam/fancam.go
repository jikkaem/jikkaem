package fancam

import (
	"context"
	"jikkaem/internal/model"
	"jikkaem/internal/mongodb"
	pb "jikkaem/internal/proto/fancam"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	database string = "fancam"
	uri      string = "mongodb://localhost:10001"
)

type FancamServer struct {
	pb.UnimplementedFancamServer
}

func (s *FancamServer) GetFancamByID(ctx context.Context, id *pb.ID) (*pb.FancamObject, error) {
	coll, err := s.getColl("fancams")
	if err != nil {
		return nil, err
	}

	var result model.Fancam
	hex, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	if err = coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	// Convert artists list into grpc model
	var artists []*pb.ArtistObject
	for _, artist := range result.Artists {
		mappedArtist := &pb.ArtistObject{
			Id:              artist.ID.Hex(),
			StageName:       artist.StageName,
			FullName:        artist.FullName,
			KoreanName:      artist.KoreanName,
			KoreanStageName: artist.KoreanStageName,
			Dob:             timestamppb.New(artist.DOB),
			Group:           artist.Group,
			Country:         artist.Country,
			Height:          int32(artist.Height),
			Weight:          int32(artist.Weight),
			Birthplace:      artist.Birthplace,
			Gender:          pb.Gender(pb.Gender_value[artist.Gender]),
			Instagram:       artist.Instagram,
		}
		artists = append(artists, mappedArtist)
	}

	// Convert suggestedTags into grpc model
	var suggestedTags []*pb.SuggestedTags
	for _, tag := range result.SuggestedTags {
		mappedTag := &pb.SuggestedTags{
			EnArtist: tag.EnArtist,
			EnGroup:  tag.EnGroup,
			EnSong:   tag.EnSong,
			KrArtist: tag.KrArtist,
			KrGroup:  tag.KrGroup,
			KrSong:   tag.KrSong,
		}
		suggestedTags = append(suggestedTags, mappedTag)
	}

	return &pb.FancamObject{
		Id:            result.ID.Hex(),
		Title:         result.Title,
		YtLink:        result.YtLink,
		RecordDate:    timestamppb.New(result.RecordDate),
		Artists:       artists,
		SuggestedTags: suggestedTags,
	}, nil
}

func (s *FancamServer) CreateFancams(ctx context.Context, fancam *pb.FancamList) (*pb.FancamList, error) {
	return nil, nil
}

func (s *FancamServer) DeleteFancam(ctx context.Context, id *pb.ID) (*pb.FancamObject, error) {
	return nil, nil
}

func (s *FancamServer) getColl(collName string) (*mongo.Collection, error) {
	client, err := mongodb.GetMongoClient(uri)
	if err != nil {
		return nil, err
	}
	return client.Database(database).Collection(collName), nil
}
