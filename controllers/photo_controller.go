package controllers

import (
	"PROJECT_BTPN/database"
	"PROJECT_BTPN/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userId")
	photo := models.Photo{Title: input.Title, Caption: input.Caption, PhotoUrl: input.PhotoUrl, UserID: userID.(uint)}
	database.DB.Create(&photo)

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	database.DB.Find(&photos)

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := database.DB.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	userID, _ := c.Get("userId")
	if photo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&photo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	if err := database.DB.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	userID, _ := c.Get("userId")
	if photo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	database.DB.Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
