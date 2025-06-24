package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"makerble-crudapi/models"
	"makerble-crudapi/config"
)

func LoginHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser models.User
	if err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate a simple token with role included
	token := "simple-token-" + user.Username + "-" + dbUser.Role
	c.JSON(http.StatusOK, gin.H{"token": token, "role": dbUser.Role})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		// Expect token format: simple-token-username-role
		parts := strings.Split(token, "-")
		if len(parts) != 4 || parts[0] != "simple" || parts[1] != "token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		username := parts[2]
		role := parts[3]

		// Verify user exists
		var dbUser models.User
		if err := config.DB.Where("username = ?", username).First(&dbUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		// Attach info to context
		c.Set("username", username)
		c.Set("role", role)
		c.Next()
	}
}
