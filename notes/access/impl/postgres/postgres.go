package postgres

import (
	"bujem/notes/model"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // DB driver so is imported anonymously
)

// NotesPostgres implements data access for notes on a PostgreSQL database
type NotesPostgres struct{}

// Create creates a new note record, updating the provided model with the record ID
func (dao NotesPostgres) Create(note *model.Note) error {
	db := getConnection()
	defer db.Close()

	var id int64
	// PostgreSQL doesn't support retrieving the last inserted ID from metadata so QueryRow and Scan act as a work around for this
	err := db.QueryRow("insert into notes (note_type, note_content, created, modified, owner_user_id) VALUES ($1, $2, now()::timestamp, now()::timestamp, $3) RETURNING id", note.NoteType, note.NoteContent, note.OwnerUserID).Scan(&id)
	if err != nil {
		return err
	}
	note.ID = id
	return nil
}

// Update updates an existing note record
func (dao NotesPostgres) Update(note *model.Note) error {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec("update notes set note_type=$1, note_content=$2, modified=now()::timestamp where id=$3", note.NoteType, note.NoteContent, note.ID)
	if err != nil {
		return err
	}
	return err
}

// FindByID locates and returns an existing note record using its database ID value
func (dao NotesPostgres) FindByID(id int64) (model.Note, error) {
	db := getConnection()
	defer db.Close()

	results, err := db.Query("select id, note_type, note_content, created, modified, owner_user_id from notes where id=$1", id)
	if err != nil {
		return model.Note{}, err
	}
	defer results.Close()

	if !results.Next() {
		return model.Note{}, fmt.Errorf("Note record with ID %d not found", id)
	}
	var note model.Note
	err = results.Scan(&note.ID, &note.NoteType, &note.NoteContent, &note.Created, &note.Modified, &note.OwnerUserID)
	if err != nil {
		return model.Note{}, err
	}
	if results.Err() != nil {
		return model.Note{}, results.Err()
	}
	return note, nil
}

// DeleteByID deletes an existing note record using its database ID value
func (dao NotesPostgres) DeleteByID(id int64) error {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec("delete from notes where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/bujem?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
