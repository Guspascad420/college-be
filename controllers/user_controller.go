package controllers

import (
	"college-be/database"
	"college-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	record := database.Db.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"meta": models.Meta{true, "success"}, "data": user})
}

func UpdateUserData(c *gin.Context) {
	var user models.User
	var newUserData models.User
	nim, err := ExtractNIM(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&newUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.Db.Preload("Major").Where("nim = ?", nim).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}

	update := database.Db.Model(&user).Updates(newUserData)
	if update.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": update.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"meta": models.Meta{true, "user data successfully updated"}, "data": user})
}

func GetUserProfile(c *gin.Context) {
	var user models.User
	nim, err := ExtractNIM(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	record := database.Db.Preload("Major").Where("nim = ?", nim).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"meta": models.Meta{true, "success"},
		"mahasiswa": models.UserProfileResponse{user.ID, user.Name, user.NIM,
			user.Angkatan, user.Major.Name},
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	record := database.Db.Preload("Major").Find(&users)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "success"}, "mahasiswa": users})
}

func GetUserByNIM(c *gin.Context) {
	nim := c.Param("nim")
	var user models.User

	record := database.Db.Preload("Major").Preload("Courses").Where("nim = ?", nim).Find(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "success"}, "mahasiswa": user})
}
