package application

import (
	"bujem/users/data/service"
	"bujem/users/model"
	"log"
)

func Initialize() {
	service.Initialize("postgres")
}

func CreateUser(user *model.User) error {
	log.Printf("Creating user: %s", user.Username)
	return service.Create(user)
}

func UpdateUser(user *model.User) error {
	log.Printf("Updating user with ID: %d", user.ID)
	return service.Update(user)
}

func GetUser(ID int) (model.User, error) {
	log.Printf("Get user with ID: %d\n", ID)
	return service.GetById(ID)
}

func DeleteUser(id int) error {
	log.Printf("Delete user with ID: %d", id)
	return service.DeleteById(id)
}
