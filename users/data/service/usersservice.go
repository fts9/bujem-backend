package service

import (
	"bujem/users/data/access"
	"bujem/users/data/access/impl/postgres"
	"bujem/users/model"
	"log"
)

var usersAccess access.Users

func Initialize(dbms string) {
	log.Printf("Initializing users data service for DBMS: %s", dbms)

	switch dbms {
	case "postgres":
		usersAccess = postgres.UsersAccessPostgres{}
	default:
		log.Fatalf("No data access configured for DBMS: %s", dbms)
	}
}

func Create(user *model.User) error {
	return usersAccess.Create(user)
}

func Update(user *model.User) error {
	return usersAccess.Update(user)
}

func GetById(id int) (model.User, error) {
	return usersAccess.FindById(id)
}

func DeleteById(id int) error {
	return usersAccess.DeleteById(id)
}
