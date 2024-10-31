package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieEntity struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName   string             `json:"movieName,omitempty"`
	Thumblaine  string             `json:"thumblaine,omitempty"`
	Description string             `json:"description,omitempty"`
	Rating      string             `json:"rating,omitempty"`
	DownloadUrl string             `json:"downloadUrl"`
	WatchUrl    string             `json:"watchUrl"`
}
