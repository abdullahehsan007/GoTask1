package login

import (
	"GOTASK/api/repository"
	"GOTASK/model"
	"GOTASK/services"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type implService struct {
	repo repository.UserRepository
	auth services.AuthService
}

func NewAuthService(repo repository.UserRepository, auth services.AuthService) AuthService {
	return &implService{repo: repo,
		auth: auth}
}

func (s *implService) Login(email, password string) (string, error) {
	id, err := s.auth.Authenticator(model.User{Email: email,
		Password: password})
	if err != nil {
		return "", err
	}
	return id, nil
}
