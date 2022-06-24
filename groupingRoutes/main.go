package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "ping ping")
		})
		v1.GET("/pong", func(c *gin.Context) {
			c.String(http.StatusOK, "pong pong")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "PING PING")
		})
		v2.GET("/pong", func(c *gin.Context) {
			c.String(http.StatusOK, "PONG PONG")
		})
	}

	router.Run(":8080")
}
