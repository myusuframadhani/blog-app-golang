package controllers

import (
	"blog-app/config"
	"blog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory handles creating a new category
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetCategories handles retrieving all categories
func GetCategories(c *gin.Context) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategory handles retrieving a single category by ID
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	result := config.DB.First(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory handles updating a category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Model(&category).Where("id = ?", id).Updates(category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory handles deleting a category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Category{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
