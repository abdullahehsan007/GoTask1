package handler

import (
	"GOTASK/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler interface {
	Login() gin.HandlerFunc
	Signup() gin.HandlerFunc
	Authorize() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}

type userHandler struct {
	db      *sqlx.DB
	service services.AuthService
}

func NewUserHandler(db *sqlx.DB) Handler {
	return &userHandler{
		db:      db,
		service: services.Newauthservice(db),
	}
}

