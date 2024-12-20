package main

import (
	"github.com/Triatmono/Golang-FP-CRUD-Gin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
