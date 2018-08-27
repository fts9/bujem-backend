package model

import "time"

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Created  time.Time
	Modified time.Time
}
