package controller

import (
	"fmt"
	"goWebApp/database/users"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type LoginRequest struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := users.GetUser(loginRequest.Username)
	fmt.Println("login user : ", loginRequest.Username)
	fmt.Println("login password : ", loginRequest.Password)
	fmt.Println("get user : ", user)

	if loginRequest.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Erreur d'authentification"})
		return
	}

	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("username", user.Username)
	session.Set("email", user.Email)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User Sign In successfully"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User Sign out successfully"})
}

func Test(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	email := session.Get("email")
	username := session.Get("username")

	c.JSON(200, gin.H{
		"message":  "Testing Success",
		"id":       id,
		"username": username,
		"email":    email,
	})
}
