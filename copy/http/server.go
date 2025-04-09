package httpserver

import (
	"log"

	"File_Syncer/auth/routes"
	"github.com/gin-gonic/gin"
)

func StartHTTPServer(port string) {
	router := gin.Default()

	// Register authentication routes
	routes.RegisterAuthRoutes(router)

	// Protected routes
	routes.RegisterProtectedRoutes(router)

	// Simple test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	addr := ":" + port
	log.Println("🌐 HTTP server running at port", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}
