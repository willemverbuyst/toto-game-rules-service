package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rule struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Question string             `json:"question,omitempty" validate:"required"`
	Answers  []Answer           `json:"answers,omitempty" validate:"required"`
}

type Answer struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Text string             `json:"text,omitempty" validate:"required"`
}
