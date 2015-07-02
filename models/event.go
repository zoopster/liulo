package models

import "time"

// Event Object
type Event struct {
	ID        string    `gorethink:"id, omitempty" json:"id"`
	Name      string    `gorethink:"name" json:"name"`
	Slug      string    `gorethink: "slug" json:"slug"`
	Location  string    `gorethink:"location" json:"location"`
	Restrict  bool      `gorethink:"restrict" json:"restrict"`
	Private   bool      `gorethink:"private" json:"private"`
	Single    bool      `gorethink:"single" json:"single"`
	Creator   string    `gorethink:"creator" json:"creator"`
	From      time.Time `gorethink:"from" json:"from"`
	To        time.Time `gorethink:"to" json:"to"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Sessions  []Session `json:"sessions"`
}
