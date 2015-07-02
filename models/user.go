package models

import "time"

// User Object
type User struct {
	ID        string    `gorethink:"id, omitempty" json:"id"`
	Username  string    `gorethink:"username" json:"username"`
	Password  string    `gorethink:"password" json:"-"`
	Name      string    `gorethink:"name" json:"name"`
	FID       string    `gorethink:"fid" json:"fid"`
	UpdatedAt time.Time `gorethink:"created_at" json:"updated_at"`
	CreatedAt time.Time `gorethink:"updated_at" json:"created_at"`
}
