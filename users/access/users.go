package access

import "bujem/users/model"

// Users provides a common interface for users data access
type Users interface {
	Create(user *model.User) error
	Update(user *model.User) error
	FindByID(ID int64) (model.User, error)
	DeleteByID(ID int64) error
}
