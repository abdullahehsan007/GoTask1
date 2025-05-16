package handler

import (
	"GOTASK/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Refresh(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken := ctx.PostForm("r_token")
		if refreshToken == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.String(http.StatusUnauthorized, "Refresh Token Required")
			return
		}

		email, err := services.VerifyRefreshToken(refreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
			return
		}
		newAccessToken,_,err:= services.CreateToken(email)
        if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Could not generate new token"})
			return
		}
		
	     ctx.JSON(http.StatusCreated, gin.H{
			"Message": "Access Token Generated",
			"Access Token": newAccessToken,
		})

	}
}
