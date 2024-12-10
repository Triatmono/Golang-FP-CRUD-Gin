package main

import (
	"github.com/Triatmono/Golang-FP-CRUD-Gin/controllers"
	"github.com/Triatmono/Golang-FP-CRUD-Gin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/id", controllers.PostsShow)
	r.GET("/posts/id", controllers.PostsUpdate)
	r.GET("/posts/id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
