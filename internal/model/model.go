package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

type Artist struct {
	ID              primitive.ObjectID `bson:"_id"`
	StageName       string             `bson:"stageName"`
	FullName        string             `bson:"fullName"`
	KoreanName      string             `bson:"koreanName"`
	KoreanStageName string             `bson:"koreanStageName"`
	DOB             time.Time          `bson:"dob,omitempty"`
	Group           string             `bson:"group,omitempty"`
	Country         string             `bson:"country"`
	Height          int8               `bson:"height,omitempty"`
	Weight          int8               `bson:"weight,omitempty"`
	Birthplace      string             `bson:"birthplace"`
	Gender          string             `bson:"gender"`
	Instagram       string             `bson:"instagram,omitempty"`
}

type Fancam struct {
	ID            primitive.ObjectID `bson:"_id"`
	Title         string             `bson:"title"`
	Description   string             `bson:"description"`
	PublishedAt   time.Time          `bson:"publishedAt"`
	ChannelID     string             `bson:"channelId"`
	ChannelTitle  string             `bson:"channelTitle"`
	RootThumbnail string             `bson:"rootThumbnail"`
	RecordDate    time.Time          `bson:"recordDate,omitempty"`
	Artists       []Artist           `bson:"artists"`
	SuggestedTags SuggestedTags      `bson:"suggestedTags,omitempty"`
}

type SuggestedTags struct {
	EnArtist []string `bson:"enArtist,omitempty"`
	EnGroup  []string `bson:"enGroup,omitempty"`
	EnSong   []string `bson:"enSong,omitempty"`
	KrArtist []string `bson:"krArtist,omitempty"`
	KrGroup  []string `bson:"krGroup,omitempty"`
	KrSong   []string `bson:"krSong,omitempty"`
}
