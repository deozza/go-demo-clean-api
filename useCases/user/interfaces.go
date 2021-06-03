package user

import "api-test/entity"

type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateUser(email, password, username string) (entity.ID, error)
}