package routes

import (

	// analyzer "GOTASK/handler/Analyzer"
	// authorization "GOTASK/handler/Authorization"

	"GOTASK/handler"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RoutersSetup(router *gin.Engine, db *sqlx.DB) {
	handlers := handler.NewUserHandler(db)
	router.POST("/analyze", handlers.Authorize(), handler.AnalyzeText(db))
	router.POST("/signup", handlers.Signup())
	router.POST("/login", handlers.Login())
	router.POST("/auth", handlers.Authorize())
	router.POST("/ref", handlers.Refresh())
	router.Run(":8080")
}
