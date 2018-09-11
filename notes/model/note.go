package model

import (
	"time"
)

// Note is a model representing a single user entered note
type Note struct {
	ID          int64     `json:"id"`
	NoteType    string    `json:"noteType"`
	NoteContent string    `json:"noteContent"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	OwnerUserID int       `json:"ownerUserId"`
}
