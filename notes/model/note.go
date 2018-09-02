package model

import "time"

// Note is a model representing a single user entered note
type Note struct {
	ID          int64
	NoteType    string
	NoteContent string
	Created     time.Time
	Modified    time.Time
	OwnerUserID int
}
