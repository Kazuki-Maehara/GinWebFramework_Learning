package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "./"+file.Filename)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}

// curl -X POST -F "upload[]=@./test.txt" -F "upload[]=@./test2.txt" -H "Content-Type: multipart/form-lhost:8080/upload
