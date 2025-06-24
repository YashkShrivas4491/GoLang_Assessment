package main

import (
	"makerble-crudapi/config"
	"makerble-crudapi/routes"
    "os"
	"log"
	"github.com/gin-gonic/gin"

)

func main() {
	// Initialize database
	config.InitDB()

	

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Makerble CRUD API")
	})

	// Setup routes
	routes.SetupRoutes(r)

	// r.Run(":6000")

	// Get port from environment variable or default to 6000
	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	// Run server
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}