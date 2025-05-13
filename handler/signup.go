package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Info struct {
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func Signup(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user Info
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
			return
		}
		var exists bool
		err := db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM signup WHERE username = $1 OR email = $2)`,
			user.Username, user.Email).Scan(&exists)
		if err != nil || exists {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
			return
		}
		_, err = db.Exec(
			`INSERT INTO signup(username,email,password) VALUES($1,$2,$3)`, user.Username, user.Email, user.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user.Username,
		})
	}
}
