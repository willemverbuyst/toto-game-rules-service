package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rule struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Order    int                `json:"order,omitempty"`
	Question string             `json:"question" validate:"required"`
	Answers  []Answer           `json:"answers" validate:"required"`
}

type Answer struct {
	Order int    `json:"order,omitempty"`
	Text  string `json:"text" validate:"required"`
}
