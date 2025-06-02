package repository

import (
	"GOTASK/model"

	"golang.org/x/crypto/bcrypt"
)
func HashedPassword(User model.Info) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (r *UserRepo) GetId(email string) (string, error) {
	var id string
	query := `SELECT id FROM signup WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserRepo) GetUser(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM signup WHERE email = $1)`
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserRepo) GetUserData(email string) (string, string, error) {
	var id string
	var dbpass string
	query := `SELECT id,password FROM signup WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&id, &dbpass)
	if err != nil {
		return "", "", err
	}
	return id, dbpass, nil
}

func (r *UserRepo) CreateUser(user model.Info) error {
	hashed, err := HashedPassword(user)
	if err != nil {
		return err
	}
	query := `INSERT INTO signup(username,email,password) VALUES($1,$2,$3) RETURNING id `
	_, err = r.db.Exec(query, user.Username, user.Email, hashed)
	return err
}
