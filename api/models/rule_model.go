package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rule struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" example:"634f787edc90bf2b7c868153"`
	Order    int                `json:"order,omitempty" example:"1"`
	Question string             `json:"question" validate:"required" example:"What time is it?"`
	Answers  []Answer           `json:"answers" validate:"required"`
}

type Answer struct {
	Order int    `json:"order,omitempty" example:"1"`
	Text  string `json:"text" validate:"required" example:"12 o'clock"`
}
