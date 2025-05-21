package model

type Info struct {
	Id       string `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
