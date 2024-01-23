package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/rifansyah/go-crud/app/helpers"
	"github.com/rifansyah/go-crud/app/http/requests"
	"github.com/rifansyah/go-crud/app/initializers"
	"github.com/rifansyah/go-crud/app/models"
)

type PostsController struct{}

func NewPostsController() *PostsController {
	return &PostsController{}
}

func (ctl *PostsController) Create(c *gin.Context) {
	var postData requests.PostsRequest
	if err := c.ShouldBind(&postData); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err, "Failed to read body")
		return
	}

	post := models.Post{Title: postData.Title, Body: postData.Body}
	if result := initializers.DB.Create(&post); result.Error != nil {
		response.ResponseError(c, http.StatusInternalServerError, result.Error, "Failed to create the post data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (ctl *PostsController) Index(c *gin.Context) {
	var posts []models.Post
	if result := initializers.DB.Find(&posts); result.Error != nil {
		response.ResponseError(c, http.StatusInternalServerError, result.Error, "Failed to get all post data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (ctl *PostsController) GetPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if result := initializers.DB.First(&post, id); result.Error != nil {
		response.ResponseError(c, http.StatusBadRequest, result.Error, "Post data not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (ctl *PostsController) Update(c *gin.Context) {
	id := c.Param("id")

	var postData requests.PostsRequest
	if err := c.ShouldBind(&postData); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err, "Failed to get the body data")
		return
	}

	var post models.Post
	if result := initializers.DB.First(&post, id); result.Error != nil {
		response.ResponseError(c, http.StatusBadRequest, result.Error, "Post data not found")
		return
	}

	if result := initializers.DB.Model(&post).Updates(postData); result.Error != nil {
		response.ResponseError(c, http.StatusInternalServerError, result.Error, "Failed to update the post data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (ctl *PostsController) Detele(c *gin.Context) {
	id := c.Param("id")

	if result := initializers.DB.Delete(&models.Post{}, id); result.Error != nil {
		response.ResponseError(c, http.StatusBadRequest, result.Error, "Failed to remove data or post data not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
