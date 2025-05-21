package routes

import (
	"GOTASK/handler"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RoutersSetup(router *gin.Engine, db *sqlx.DB) {
	router.POST("/analyze", handler.Authorize(), handler.AnalyzeText(db))
	router.POST("/signup", handler.Signup(db))
	router.POST("/login", handler.Login(db))
	router.POST("/auth", handler.Authorize())
	router.POST("/ref", handler.Refresh(db))
	router.Run(":8080")
}
