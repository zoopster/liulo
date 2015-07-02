package models

import "time"

// Question Object
type Question struct {
	ID        string    `gorethink:"id, omitempty" json:"id"`
	Content   string    `gorethink:"content" json:"content"`
	Up        int       `gorethink:"up" json:"up"`
	Creator   string    `gorethink:"creator" json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Approved  bool      `gorethink:"approved" json:"approved"`
	Users     []User    `json:"users"`
}
