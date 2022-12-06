package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/sample", func(c *gin.Context) {
		c.String(http.StatusOK, "Sample API")
	})

	r.Run(":8080")
}
