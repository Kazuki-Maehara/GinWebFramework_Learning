package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set a lowe memory limit for multipart forms (default is 32MiB)
	router.MaxMultipartMemory = 8 << 19 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific destination.
		dst := "./" + "savedFile.txt"
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' updeaded!", file.Filename))
	})
	router.Run(":8080")
}
