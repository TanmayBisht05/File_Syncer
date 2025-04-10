package handlers

import (
	"File_Syncer/auth/models"
	"File_Syncer/auth/utils"
	"File_Syncer/db"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitAuthHandler() {
	userCollection = db.Client.Database("file_syncer_db").Collection("auth_users")
}

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Check if email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	// Insert user
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Generate JWT and set it in HTTP-only cookie
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 3600*24, "/", "", false, true) // Secure=true in prod with HTTPS
	c.JSON(http.StatusOK, gin.H{"username": user.Username, "email": user.Email})
}

func Login(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dbUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&dbUser)
	if err != nil || !utils.CheckPasswordHash(req.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
		return
	}

	token, err := utils.GenerateJWT(dbUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set token as HTTP-only cookie
	c.SetCookie("token", token, 3600*24, "/", "", false, true) // Secure=true in prod
	c.JSON(http.StatusOK, gin.H{"username": dbUser.Username, "email": dbUser.Email})
}

func Logout(c *gin.Context) {
	// Invalidate the token by expiring the cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func CheckAuth(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": userID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
	})
}

// func CheckAuth(c *gin.Context) {
// 	userID, exists := c.Get("userId")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"username": userID})
// }
