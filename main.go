package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	t := time.Now()

	r.GET("/sample", func(c *gin.Context) {
		c.String(http.StatusOK, "Sample API")
	})

	r.GET("/time/now", func(c *gin.Context) {
		c.String(http.StatusOK, t.String())
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
