package handler

import (
	"GOTASK/model"
	"GOTASK/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)


func Login(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var entry model.User
		var checkpass string
		var name string
		var user model.Info

		if err := ctx.ShouldBindJSON(&entry); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
			return
		}

		err := db.QueryRow(`SELECT username, password FROM signup WHERE email = $1 `,
			entry.Email).Scan(&name, &checkpass)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid User"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(checkpass), []byte(entry.Password)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Password"})
			return
		}
		tokenString, refresh, err := services.CreateToken(user.Id)
		if err != nil {
			fmt.Errorf("No username found")
		}

		// if entry.Password != checkpass{
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		// 	return
		// }
		ctx.JSON(http.StatusCreated, gin.H{
			"message":       "Login Successful",
			"user":          name,
			"Token":         tokenString,
			"Refresh Token": refresh,
		})
	}
}
