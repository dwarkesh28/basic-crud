package controllers

import (
	"go-crud/dto"
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)


	//  Create a post
	post := models.PostData{Title: body.Title, Body: body.Body }
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	finalData := dto.ToDto(post)

	// return it
	c.JSON(200, gin.H{
		"post": finalData,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.PostData
	var final []*dto.GetPostData

	// initializers.DB.Select("title", "body", "created_at").Find(&posts)
	initializers.DB.Find(&posts)

	for _, post := range posts {
		final = append(final, dto.ToDto(post))
	}

	// return it
	c.JSON(200, gin.H{
		"posts": final,
	})
}

func PostById(c *gin.Context) {
	// Get if from url
	id := c.Param("id")

	// Get the posts
	var post models.PostData

	initializers.DB.First(&post, id)

	// return it
	c.JSON(200, gin.H{
		"post": dto.ToDto(post),
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	var post models.PostData

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.PostData{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": dto.ToDto(post),
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.PostData{}, id)
	c.Status(200)
}
