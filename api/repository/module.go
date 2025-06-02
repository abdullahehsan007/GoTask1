package repository

import (
	"GOTASK/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUser(email string) (bool, error)
	CreateUser(user model.Info) error
	GetId(email string) (string, error)
	GetUserData(email string) (string, string, error)
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}
