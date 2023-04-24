// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Artist struct {
	ID              string     `json:"id"`
	StageName       string     `json:"stageName"`
	FullName        string     `json:"fullName"`
	KoreanName      string     `json:"koreanName"`
	KoreanStageName string     `json:"koreanStageName"`
	Dob             *time.Time `json:"dob,omitempty"`
	Group           *string    `json:"group,omitempty"`
	Country         string     `json:"country"`
	Height          *int       `json:"height,omitempty"`
	Weight          *float64   `json:"weight,omitempty"`
	Birthplace      string     `json:"birthplace"`
	Gender          Gender     `json:"gender"`
	Instagram       *string    `json:"instagram,omitempty"`
}

type Fancam struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	PublishedAt   time.Time      `json:"publishedAt"`
	ChannelID     string         `json:"channelID"`
	ChannelTitle  string         `json:"channelTitle"`
	RootThumbnail string         `json:"rootThumbnail"`
	RecordDate    *time.Time     `json:"recordDate,omitempty"`
	SuggestedTags *SuggestedTags `json:"suggestedTags"`
}

type LatestFancamsInput struct {
	MaxResults int `json:"maxResults"`
}

type ListIDs struct {
	Ids []string `json:"ids"`
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SingleID struct {
	ID string `json:"id"`
}

type SuggestedTags struct {
	EnArtist []string `json:"enArtist"`
	EnGroup  []string `json:"enGroup"`
	EnSong   []string `json:"enSong"`
	KrArtist []string `json:"krArtist"`
	KrGroup  []string `json:"krGroup"`
	KrSong   []string `json:"krSong"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}