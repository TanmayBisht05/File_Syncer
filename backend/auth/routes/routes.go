package routes

import (
	"File_Syncer/auth/handlers"
	"File_Syncer/auth/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", handlers.Logout)
		auth.GET("/check", middleware.JWTAuthMiddleware(), handlers.CheckAuth)
	}
}
