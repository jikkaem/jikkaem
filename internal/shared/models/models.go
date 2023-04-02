package models

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
	StageName       string    `bson:"stageName"`
	FullName        string    `bson:"fullName"`
	KoreanName      string    `bson:"koreanName"`
	KoreanStageName string    `bson:"koreanStageName"`
	DOB             time.Time `bson:"dob,omitempty"`
	Group           string    `bson:"group,omitempty"`
	Country         string    `bson:"country"`
	Height          int8      `bson:"height,omitempty"`
	Weight          int8      `bson:"weight,omitempty"`
	Birthplace      string    `bson:"birthplace"`
	Gender          string    `bson:"gender"`
	Instagram       string    `bson:"instagram,omitempty"`
}

type ArtistWithID struct {
	ID primitive.ObjectID `bson:"_id"`
	*Artist
}

type Fancam struct {
	Title   string   `bson:"title"`
	YtLink  string   `bson:"ytLink"`
	Artists []Artist `bson:"artists"`
}

type FancamWithID struct {
	ID primitive.ObjectID `bson:"_id"`
	*Fancam
}
