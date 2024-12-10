package controllers

import (
	"github.com/Triatmono/Golang-FP-CRUD-Gin/initializers"
	"github.com/Triatmono/Golang-FP-CRUD-Gin/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context){

	var body struct {
		Title string 
		Body string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context){

	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context){

	//Get the id 
	id := c.Param("id")

	//Get Data 
	var body struct {
		Title string 
		Body string
	}

	c.Bind(&body)

	//Find the id to update
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context){

	//Get the id 
	id := c.Param("id")

	//Delete the Post with id
	initializers.DB.Delete(&models.Post{}, id)

	//Status Response
	c.Status(200)

}