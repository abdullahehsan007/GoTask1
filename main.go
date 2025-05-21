package main

import (
	"GOTASK/database"
	"GOTASK/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Database()
	defer db.Close()
	router := gin.Default()
	routes.RoutersSetup(router, db)
	
}
