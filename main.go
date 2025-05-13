package main

import (
	"GOTASK/database"
	"GOTASK/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Database()
	router := gin.Default()
	router.POST("/analyze", handler.AnalyzeText(db))
	router.POST("/signup", handler.Signup(db))
	router.POST("/login", handler.Login(db))
	
	router.Run(":8080")
}
