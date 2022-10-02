package models

type Rule struct {
	Id       int      `json:"id,omitempty"`
	Question string   `json:"question,omitempty" validate:"required"`
	Answers  []Answer `json:"answers,omitempty" validate:"required"`
}

type Answer struct {
	Id   int    `json:"id,omitempty"`
	Text string `json:"text,omitempty" validate:"required"`
}
