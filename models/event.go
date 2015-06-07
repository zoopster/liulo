package models

import "time"

type Event struct {
	Id        string
	Name      string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
