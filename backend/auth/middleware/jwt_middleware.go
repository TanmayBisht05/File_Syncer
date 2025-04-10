package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecret") // Must match the key used in utils.GenerateJWT

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from HTTP-only cookie
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token not found in cookies"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username, ok := claims["username"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}
			c.Set("userId", username)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			c.Abort()
			return
		}

		c.Next()
	}
}
