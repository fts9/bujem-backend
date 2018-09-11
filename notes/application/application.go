package application

import (
	"bujem/notes/access"
	"bujem/notes/access/utility"
	"bujem/notes/model"
	"log"
)

var notesAccess access.Notes

// Create saves a new Note record to the DB
func Create(note *model.Note) (model.Note, error) {
	log.Printf("Creating new note record for user: %d", note.OwnerUserID)
	initializeAccess()
	err := notesAccess.Create(note)
	if err != nil {
		log.Printf("Failed to create note record for user: %d", note.OwnerUserID)
		return model.Note{}, err
	}
	return notesAccess.FindByID(note.ID)
}

// Update updates an existing Note record on the DB
func Update(note *model.Note) (model.Note, error) {
	log.Printf("Updating note with ID: %d", note.ID)
	initializeAccess()
	err := notesAccess.Update(note)
	if err != nil {
		log.Printf("Failed to update note record with ID: %d", note.ID)
		return model.Note{}, err
	}
	return notesAccess.FindByID(note.ID)
}

// FindByID retrieves a note with the given ID from the DB
func FindByID(id int64) (model.Note, error) {
	log.Printf("Retrieving note with ID: %d", id)
	initializeAccess()
	note, err := notesAccess.FindByID(id)
	if err != nil {
		log.Printf("Failed to retrieve note record with ID: %d", id)
	}
	return note, err
}

// DeleteByID deletes a note record with the given ID from the DB
func DeleteByID(id int64) error {
	log.Printf("Deleting note with ID: %d", id)
	initializeAccess()
	err := notesAccess.DeleteByID(id)
	if err != nil {
		log.Printf("Failed to delete note record with ID: %d", id)
	}
	return err
}

func initializeAccess() {
	if notesAccess == nil {
		log.Println("Initializing application service data access")
		notesAccess = utility.GetDataAccess("postgres")
	}
}
