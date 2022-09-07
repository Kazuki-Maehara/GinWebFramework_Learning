package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/local/file", func(c *gin.Context) {
		c.File("./main.go")
	})

	var fs http.FileSystem = http.Dir("./")
	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("main.go", fs)
	})

	router.Run(":8080")
}
