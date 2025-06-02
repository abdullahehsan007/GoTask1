package handler

import (
	"GOTASK/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *userHandler) Signup() gin.HandlerFunc{
	return func(ctx *gin.Context){
	var user model.Info
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if err := h.service.SignUp(user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered!",
		"user":    user.Username,
	})
}
}