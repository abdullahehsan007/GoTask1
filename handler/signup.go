package handler

import (
	"GOTASK/model"
	"GOTASK/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Signup(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var User model.Info
		if err := ctx.ShouldBindJSON(&User); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := services.RegisterUser(db, User)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		// exists, err = api.GetUser(db, User.Username, User.Email)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		// 	return
		// }
		// if exists {
		// 	ctx.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		// 	return
		// }

		// _ = db.QueryRow(
		// 	`INSERT INTO signup(username,email,password) VALUES($1,$2,$3) RETURNING id `, User.Username, User.Email, string(hashedPassword))
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    User.Username,
		})
	}
}
