package main

import (
	"GOTASK/api/repository"
	"GOTASK/database"
	login "GOTASK/handler/Login"
	signup "GOTASK/handler/Signup"
	"GOTASK/routes"
	"GOTASK/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Database()
	//for signup (dependencies)
	userRepo := repository.NewUserRepository(db)
	userService := signup.NewUserService(userRepo)
	userHandler := signup.NewUserHandler(userService)

	//for login
	authService := services.NewAuthService(userRepo)
	loginservice := login.NewAuthService(userRepo, authService)
	loginHandler := login.NewLoginHandler(loginservice)

	defer db.Close()
	router := gin.Default()
	routes.RoutersSetup(router, db, userHandler, loginHandler)

}
