package api

import (
	"GOTASK/model"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(User model.Info) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	return string(hashed), err
}

func GetUser(db *sqlx.DB, username, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM signup WHERE username = $1 OR email = $2)`
	err := db.QueryRow(query, username, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CreateUser(db *sqlx.DB, user model.Info) error {
	hashed, err := HashedPassword(user)
	if err != nil {
		return err
	}
    query:= `INSERT INTO signup(username,email,password) VALUES($1,$2,$3) RETURNING id `
	_,err = db.Exec(query, user.Username, user.Email, hashed)
	return err
}
