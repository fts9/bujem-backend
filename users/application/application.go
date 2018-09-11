package application

import (
	"bujem/users/access"
	"bujem/users/access/utility"
	"bujem/users/model"
	"log"
)

var usersAccess access.Users

// CreateUser saves a new User to the DB
func CreateUser(user *model.User) (model.User, error) {
	log.Printf("Creating new user record: %s", user.Username)
	initializeAccess()
	err := usersAccess.Create(user)
	if err != nil {
		log.Printf("Failed to create user record for user: %s", user.Username)
		return model.User{}, err
	}
	return usersAccess.FindByID(user.ID)
}

// UpdateUser updates an existing user on the DB
func UpdateUser(user *model.User) (model.User, error) {
	log.Printf("Updating user with ID: %d", user.ID)
	initializeAccess()
	err := usersAccess.Update(user)
	if err != nil {
		log.Printf("Failed to update user record with ID: %d", user.ID)
		return model.User{}, err
	}
	return usersAccess.FindByID(user.ID)
}

// GetUser retrieves an existing user record from the DB
func GetUser(ID int64) (model.User, error) {
	log.Printf("Get user with ID: %d\n", ID)
	initializeAccess()
	user, err := usersAccess.FindByID(ID)
	if err != nil {
		log.Printf("Failed to retrieve user record with ID: %d", user.ID)
	}
	return user, err
}

// DeleteUser deletes an existing user record from the DB
func DeleteUser(ID int64) error {
	log.Printf("Delete user with ID: %d", ID)
	initializeAccess()
	return usersAccess.DeleteByID(ID)
}

func initializeAccess() {
	if usersAccess == nil {
		log.Println("Initializing application service data access")
		usersAccess = utility.GetDataAccess("postgres")
	}
}
