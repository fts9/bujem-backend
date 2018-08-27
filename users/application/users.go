package application

import (
	data "bujem/users/data/access"
	"bujem/users/model"
)

func GetUser(ID int) (model.User, error) {
	return data.FindUser(ID)
}
