package models

import "time"

// Session Object
type Session struct {
	ID          string     `gorethink:"id, omitempty" json:"id"`
	Name        string     `gorethink:"name" json:"name"`
	Description string     `gorethink:"description" json:"description"`
	Location    string     `gorethink:"location" json:"location"`
	Questions   []Question `json:"questions"`
	UpdatedAt   time.Time  `gorethink:"created_at" json:"updated_at"`
	CreatedAt   time.Time  `gorethink:"updated_at" json:"created_at"`
}
