package handler

import (
	"fmt"
	"os"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_Key"))
var refreshSecretKey = []byte(os.Getenv("JWT_REFRESH_KEY"))

func BearerToken(header string) string {
	if strings.HasPrefix(header, "Bearer") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return ""
}
func ParseToken(token string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
func TokenType(token *jwt.Token, Type string) bool {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokentype, ok := claims["type"].(string); ok && tokentype == Type {
			return true
		}
	}
	return false
}

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := BearerToken(ctx.GetHeader("Authorization"))
		if token == "" {
			ctx.String(http.StatusUnauthorized, "Missing authorization header")
			ctx.Abort()
			return
		}

		token2, err := ParseToken(token, secretKey)
		if TokenType(token2, "access") {
			fmt.Println("You are Authorized")
			ctx.Next()
			return
		}

		token2, err = ParseToken(token, refreshSecretKey)
		if err == nil {
			if TokenType(token2, "refresh") {
				ctx.String(http.StatusOK, "This is Refresh token")
				ctx.Abort()
				return
			}
		} else if strings.Contains(err.Error(), "expired") {
			ctx.String(http.StatusUnauthorized, "Token expired")
			ctx.Abort()
			return
		}

		ctx.String(http.StatusUnauthorized, "Invalid Token")
		ctx.Abort()

	}
}
