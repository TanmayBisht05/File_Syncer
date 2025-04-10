package httpserver

import (
	"log"
	"time"

	"File_Syncer/auth/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartHTTPServer(port string) {
	router := gin.Default()

	// ✅ CORS middleware for frontend at http://localhost:5173
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ✅ Register authentication routes
	routes.RegisterAuthRoutes(router)

	// ✅ Register protected routes
	routes.RegisterProtectedRoutes(router)

	// ✅ Simple test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	addr := ":" + port
	log.Println("🌐 HTTP server running at port", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}
