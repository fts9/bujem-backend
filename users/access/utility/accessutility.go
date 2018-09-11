package utility

import (
	"bujem/users/access"
	"bujem/users/access/impl/postgres"
	"log"
)

// GetDataAccess returns the users data access implementation for the given DBMS
func GetDataAccess(dbms string) access.Users {
	log.Printf("Retrieving data access for DBMS: %s", dbms)
	var notesAccess access.Users

	switch dbms {
	case "postgres":
		notesAccess = postgres.UsersAccessPostgres{}
	default:
		log.Fatalf("Unable to locate data access for DBMS: %s", dbms)
	}

	return notesAccess
}
