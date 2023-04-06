package fancam

import (
	"context"
	"jikkaem/internal/model"
	"jikkaem/internal/mongodb"
	pb "jikkaem/internal/proto/fancam"

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
		tmp := &pb.FancamObject{
			Id:            result.ID.Hex(),
			Title:         result.Title,
			YtLink:        result.YtLink,
			RecordDate:    timestamppb.New(result.RecordDate),
			Artists:       artists,
			SuggestedTags: suggestedTags,
		}
		fancamList.Fancams = append(fancamList.Fancams, tmp)
	}

	return fancamList, err
}

func (s *FancamServer) GetLatest(ctx context.Context, input *pb.LatestRequest) (*pb.FancamList, error) {
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
		tmp := &pb.FancamObject{
			Id:            result.ID.Hex(),
			Title:         result.Title,
			YtLink:        result.YtLink,
			RecordDate:    timestamppb.New(result.RecordDate),
			Artists:       artists,
			SuggestedTags: suggestedTags,
		}
		fancamList.Fancams = append(fancamList.Fancams, tmp)
	}

	return fancamList, err
}

func (s *FancamServer) CreateFancams(ctx context.Context, input *pb.FancamList) (*emptypb.Empty, error) {
	// Convert grpc fancamlist into mongodb model
	inputFancams := input.GetFancams()
	fancams := make([]interface{}, len(inputFancams))

	// Loop over all entries given
	for _, fancam := range inputFancams {
		// Map artists
		var mappedArtists []model.Artist
		inputArtists := fancam.GetArtists()
		for _, artist := range inputArtists {
			mappedArtist := model.Artist{
				StageName:       artist.GetStageName(),
				FullName:        artist.GetFullName(),
				KoreanName:      artist.GetKoreanName(),
				KoreanStageName: artist.GetKoreanStageName(),
				DOB:             artist.GetDob().AsTime(),
				Group:           artist.GetGroup(),
				Country:         artist.GetCountry(),
				Height:          int8(artist.GetHeight()),
				Weight:          int8(artist.GetWeight()),
				Birthplace:      artist.GetBirthplace(),
				Gender:          artist.GetGender().String(),
				Instagram:       artist.GetInstagram(),
			}
			mappedArtists = append(mappedArtists, mappedArtist)
		}

		// Map suggested tags
		var mappedTags []model.SuggestedTags
		inputTags := fancam.GetSuggestedTags()
		for _, tag := range inputTags {
			mappedTag := model.SuggestedTags{
				EnArtist: tag.GetEnArtist(),
				EnGroup:  tag.GetEnGroup(),
				EnSong:   tag.GetEnSong(),
				KrArtist: tag.GetKrArtist(),
				KrGroup:  tag.GetKrGroup(),
				KrSong:   tag.GetKrSong(),
			}
			mappedTags = append(mappedTags, mappedTag)
		}

		mappedFancam := model.Fancam{
			ID:            primitive.NewObjectID(),
			Title:         fancam.GetTitle(),
			YtLink:        fancam.GetYtLink(),
			RecordDate:    fancam.RecordDate.AsTime(),
			Artists:       mappedArtists,
			SuggestedTags: mappedTags,
		}
		fancams = append(fancams, mappedFancam)
	}

	// Insert list into database
	coll, err := s.getColl("fancams")
	if err != nil {
		return nil, err
	}

	_, err = coll.InsertMany(ctx, fancams)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FancamServer) DeleteFancam(ctx context.Context, id *pb.ID) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *FancamServer) getColl(collName string) (*mongo.Collection, error) {
	client, err := mongodb.GetMongoClient(uri)
	if err != nil {
		return nil, err
	}
	return client.Database(database).Collection(collName), nil
}
