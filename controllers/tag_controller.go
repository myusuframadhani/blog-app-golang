package controllers

import (
	"blog-app/config"
	"blog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTag handles creating a new tag
func CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// GetTags handles retrieving all tags
func GetTags(c *gin.Context) {
	var tags []models.Tag
	result := config.DB.Find(&tags)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tags"})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetTag handles retrieving a single tag by ID
func GetTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag
	result := config.DB.First(&tag, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// UpdateTag handles updating a tag
func UpdateTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Model(&tag).Where("id = ?", id).Updates(tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// DeleteTag handles deleting a tag
func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Tag{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
