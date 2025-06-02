package services

import (
	"GOTASK/model"
	"errors"
)

func (s *authservice) SignUp(user model.Info) error {
	exists, err := s.repo.GetUser(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}
	return s.repo.CreateUser(user)
}
