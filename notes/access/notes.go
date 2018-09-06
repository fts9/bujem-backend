package access

import "bujem/notes/model"

// Notes models the IO functions required by the service
type Notes interface {
	Create(note *model.Note) error
	Update(note *model.Note) error
	FindByID(id int64) (model.Note, error)
	DeleteByID(id int64) error
}
