package utility

import (
	"bujem/notes/access"
	"bujem/notes/access/impl/postgres"
	"log"
)

// GetDataAccess returns the notes data access implementation for the given DBMS
func GetDataAccess(dbms string) access.Notes {
	log.Printf("Retrieving data access for DBMS: %s", dbms)
	var notesAccess access.Notes

	switch dbms {
	case "postgres":
		notesAccess = postgres.NotesPostgres{}
	default:
		log.Fatalf("Unable to locate data access for DBMS: %s", dbms)
	}

	return notesAccess
}
