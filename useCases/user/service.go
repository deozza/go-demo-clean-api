package user

import "api-test/entity"


type Service struct {
	repo Repository
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(email, password, username string) (entity.ID, error) {
	e, err := entity.NewUser(email, password, username)
	if err != nil {
		return e.ID, err
	}

	return e.ID, err
}