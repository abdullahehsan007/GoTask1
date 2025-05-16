package handler

import (
	"os"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

var secretKey = []byte(os.Getenv("JWT_Key"))
var refreshSecretKey = []byte(os.Getenv("JWT_REFRESH_KEY"))

func Authorize(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer") {
			ctx.String(http.StatusUnauthorized, "Missing authorization header")
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		token2, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err == nil && token2.Valid {
			if claims, ok := token2.Claims.(jwt.MapClaims); ok && claims["type"] == "access" {
				ctx.String(http.StatusOK, "You are authorized")
				return
			}
		}
		token2, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return refreshSecretKey, nil
		})
		if err == nil && token2.Valid {
			if claims, ok := token2.Claims.(jwt.MapClaims); ok && claims["type"] == "refresh" {
				ctx.String(http.StatusOK, "This is Refresh token")
				return
			}
		}
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				ctx.String(http.StatusUnauthorized, "Token expired")
				return
			}
		}
		ctx.String(http.StatusUnauthorized, "Invalid Token")

	}
}
