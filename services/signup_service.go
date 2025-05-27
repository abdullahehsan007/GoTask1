package services

import (
	"GOTASK/api/repository"
	"GOTASK/model"
	"errors"
)

type UserService interface {
	SignUp(user model.Info) error
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
func (s *userService) SignUp(user model.Info) error {
	exists, err := s.repo.GetUser(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}
	return s.repo.CreateUser(user)
}
