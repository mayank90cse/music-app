package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//swagger:response track
type Track struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ImageURI string             `json:"imageURI,omitempty" bson:"imageURI,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Artists  []string           `json:"artists,omitempty" bson:"artists,omitempty"`
	ISRC     string             `json:"isrc,omitempty" bson:"isrc,omitempty"`
}

type ErrorDetails struct {
	Code       string `json:"code,omitempty"`
	Desc       string `json:"description,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

type MusicData struct {
	TrackDetails TrackDetails `json:"tracks,omitempty" bson:"tracks,omitempty"`
}

type TrackDetails struct {
	Items []Item `json:"items,omitempty" bson:"items,omitempty"`
}

type Item struct {
	Name        string      `json:"name,omitempty" bson:"name,omitempty"`
	Album       Album       `json:"album,omitempty" bson:"album,omitempty"`
	Artists     []Artist    `json:"artists,omitempty" bson:"artists,omitempty"`
	Popularity  int         `json:"popularity,omitempty" bson:"popularity,omitempty"`
	ExternalIDs ExternalIDs `json:"external_ids,omitempty" bson:"external_ids,omitempty"`
}

type Album struct {
	Type    string   `json:"album_type,omitempty" bson:"album_type,omitempty"`
	Images  []Image  `json:"images,omitempty" bson:"images,omitempty"`
	Artists []Artist `json:"artists,omitempty" bson:"artists,omitempty"`
}

type Artist struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
	URI  string `json:"uri,omitempty" bson:"uri,omitempty"`
}

type Image struct {
	Height float64 `json:"height,omitempty" bson:"height,omitempty"`
	Width  float64 `json:"width,omitempty" bson:"width,omitempty"`
	URL    string  `json:"url,omitempty" bson:"url,omitempty"`
}

type ExternalIDs struct {
	ISRC string `json:"isrc,omitempty" bson:"isrc,omitempty"`
}
