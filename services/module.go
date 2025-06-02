package services

import (
	"GOTASK/api/repository"
	"GOTASK/model"

	"github.com/jmoiron/sqlx"
)

type AuthService interface {
	Authenticator(credential model.User) (string, error)
	RegisterUser(user model.Info) error
	Login(email, password string) (string, error)
	SignUp(user model.Info) error
}

type authservice struct {
	repo repository.UserRepository
}

func Newauthservice(db *sqlx.DB) AuthService {
	return &authservice{
		repo: repository.NewUserRepo(db),
	}
}
