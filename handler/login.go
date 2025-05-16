package handler

import (
	"GOTASK/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type persons struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func Login(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var entry persons

		if err := ctx.ShouldBindJSON(&entry); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
			return
		}
		var checkpass string
		var name string

		err := db.QueryRow(
			`SELECT username, password FROM signup WHERE email = $1 `,
			entry.Email).Scan(&name, &checkpass)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid User"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(checkpass), []byte(entry.Password)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Password"})
			return
		}
		tokenString,refresh,err := services.CreateToken(entry.Email)
		if err != nil {
			fmt.Errorf("No username found")
		}

		// if entry.Password != checkpass{
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		// 	return
		// }
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Login Successful",
			"user":    name,
			"Token":   tokenString,
			"Refresh Token": refresh,
		})
	}
}
