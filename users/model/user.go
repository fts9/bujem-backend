package model

import "time"

// User is an abstraction of a single Users DB record
type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}
