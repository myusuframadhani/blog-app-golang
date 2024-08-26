package controllers

import (
	"blog-app/config"
	"blog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost handles creating a new post
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPosts handles retrieving all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := config.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost handles retrieving a single post by ID
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost handles updating a post
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Model(&post).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost handles deleting a post
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
