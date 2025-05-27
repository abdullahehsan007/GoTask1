package login

import (
	"GOTASK/model"
	"GOTASK/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service services.AuthService
}

func NewLoginHandler(service services.AuthService) *LoginHandler {
	return &LoginHandler{service: service}
}

func (h *LoginHandler) Login(ctx *gin.Context) {
	var entry model.User

	if err := ctx.ShouldBindJSON(&entry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
		return
	}

	id, err := h.service.Login(entry.Email, entry.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokenString, refresh, err := services.CreateToken(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Login Successful",
		"Token":         tokenString,
		"Refresh Token": refresh,
	})
}
