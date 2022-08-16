package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v\n", ids, names)
	})
	router.Run(":8080")
}

// curl -X POST --data 'names[first]=thinkerou&names[second]=tianou' \
// 'http://localhost:8080/post?ids[aa]=1234&ids[b]=xxx'
