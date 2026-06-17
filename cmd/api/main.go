package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	router.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "1.0.0",
		})
	})

	router.GET("/environment", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"environment": "local",
		})
	})

	router.Run(":8081")
}