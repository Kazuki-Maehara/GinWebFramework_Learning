package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("templates/camera.html")
	router.GET("/camera", func(c *gin.Context) {
		c.HTML(http.StatusOK, "camera.html", gin.H{
			"title": "Main website",
			"para":  "This is a paragraph!",
		})
	})

	router.Run(":8080")
}
