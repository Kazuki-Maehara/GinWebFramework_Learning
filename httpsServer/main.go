package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Println("Index")
	c.HTML(http.StatusOK, "ssl_test.html", gin.H{})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", Index)
	r.RunTLS(":8080", "./certificates/server.crt", "./certificates/server.key")
}
