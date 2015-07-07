// Package conf provides all stuffs related to configuration
package conf

import (
	r "github.com/dancannon/gorethink"
)

var (
	// Session provides access to RethinkDB
	Session *r.Session
)
