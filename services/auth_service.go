package services

import (
	api "GOTASK/api/repository"
	"GOTASK/model"
	"errors"
	"regexp"

	"github.com/jmoiron/sqlx"
)

func IsValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}

func RegisterUser(db *sqlx.DB, user model.Info) error {
	exists, err := api.GetUser(db, user.Username, user.Email)
	if err == nil && exists {
		return errors.New("user already exists")
	}

	return api.CreateUser(db, user)
}
