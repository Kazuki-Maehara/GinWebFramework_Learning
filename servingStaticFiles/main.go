package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("/assets", "../")
	router.StaticFS("/more_static", http.Dir("../"))
	router.StaticFile("/image.jpg", "../image.jpg")
	router.StaticFileFS("/more_image.jpg", "image.jpg", http.Dir("../"))

	router.LoadHTMLGlob("./templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", gin.H{
			"title": "Test image!",
		})
	})
	router.Run(":8080")
}
