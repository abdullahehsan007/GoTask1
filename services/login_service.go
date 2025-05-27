package services

import (
	"GOTASK/api/repository"
	"GOTASK/model"
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Authenticator(credential model.User) (string, error)
	RegisterUser(user model.Info) error
	Login(email, password string) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func IsValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}
func (r *authService) Login(email, password string) (string, error) {
	// var l AuthService
	// id, err := l.Authenticator(model.User{Email: email,
	// 	Password: password})
	// if err != nil {
	// 	return "", err
	// }
	// return id, nil
	return r.Authenticator(model.User{
		Email:    email,
		Password: password,
	})
}
func (r *authService) Authenticator(credential model.User) (string, error) {
	exist, err := r.repo.GetUser(credential.Email)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", fmt.Errorf("user not found")
	}
	id, pass, err := r.repo.GetUserData(credential.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(credential.Password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}
	return id, err
}
func (r *authService) RegisterUser(user model.Info) error {
	exists, err := r.repo.GetUser(user.Email)
	if err == nil && exists {
		return errors.New("user already exists")
	}

	return r.repo.CreateUser(user)

}
