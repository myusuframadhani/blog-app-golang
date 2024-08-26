package controllers

import (
	"blog-app/config"
	"blog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateComment handles creating a new comment
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetComments handles retrieving all comments
func GetComments(c *gin.Context) {
	var comments []models.Comment
	result := config.DB.Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetComment handles retrieving a single comment by ID
func GetComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	result := config.DB.First(&comment, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// UpdateComment handles updating a comment
func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Model(&comment).Where("id = ?", id).Updates(comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment handles deleting a comment
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Comment{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
