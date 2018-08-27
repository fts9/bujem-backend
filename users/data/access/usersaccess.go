package access

import "bujem/users/model"

type Users interface {
	Create(user *model.User) error
	Update(user *model.User) error
	FindById(id int) (model.User, error)
	DeleteById(id int) error
}
