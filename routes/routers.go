package routes

import (
	analyzer "GOTASK/handler/Analyzer"
	authorization "GOTASK/handler/Authorization"
	login "GOTASK/handler/Login"
	refreshtoken "GOTASK/handler/RefreshToken"
	signup "GOTASK/handler/Signup"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RoutersSetup(router *gin.Engine, db *sqlx.DB, userHandler *signup.UserHandler, loginHandler *login.LoginHandler) {
	router.POST("/analyze", authorization.Authorize(), analyzer.AnalyzeText(db))
	router.POST("/signup", userHandler.Signup)
	router.POST("/login",loginHandler.Login)
	router.POST("/auth", authorization.Authorize())
	router.POST("/ref", refreshtoken.Refresh(db))
	router.Run(":8080")
}
