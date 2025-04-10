package routes

import (
	"File_Syncer/auth/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterProtectedRoutes(r *gin.Engine) {
	protected := r.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/data", func(c *gin.Context) {
			// c.JSON(200, gin.H{"message": "You accessed a protected route"})
			c.JSON(http.StatusOK, gin.H{"message": "You accessed a protected route"})
		})
	}
}
