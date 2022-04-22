package main

import (
	"goWebApp/controller"
	"goWebApp/sessionMiddleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Logout)
	auth := r.Group("/auth")
	auth.Use(sessionMiddleware.Authentication())
	{
		auth.GET("/test", controller.Test)
	}

	r.Run()
}
