package main

import (
	"github.com/gin-gonic/gin"
	helpers "github.com/horlabyc/golang-jwt-project/helpers"
	"github.com/horlabyc/golang-jwt-project/routes"
)

func main() {
	port := helpers.LoadEnv("PORT")
	if port == "" {
		port = "3000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
