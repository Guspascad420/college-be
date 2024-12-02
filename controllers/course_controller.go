package controllers

import (
	"college-be/database"
	"college-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCourses(c *gin.Context) {
	var courses []models.Course

	record := database.Db.Find(&courses)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "success"}, "matakuliah": courses})
}

func AddCourse(c *gin.Context) {
	var user models.User
	nim := c.Param("nim")
	courseId, _ := strconv.Atoi(c.Param("courseId"))

	currentUserNim, err := ExtractNIM(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	if currentUserNim != nim {
		c.JSON(http.StatusBadRequest, gin.H{"meta": models.Meta{false, "Anda tidak memiliki akses " +
			"untuk menambah mata kuliah"}})
		c.Abort()
		return
	}
	currentUserRecord := database.Db.Where("nim = ?", currentUserNim).First(&user)
	if currentUserRecord.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, currentUserRecord.Error.Error()}})
		c.Abort()
		return
	}
	insertCourseErr := database.Db.Model(&user).Association("Courses").Append(&models.Course{ID: uint(courseId)})
	if insertCourseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, insertCourseErr.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"meta": models.Meta{true, "mata kuliah berhasil ditambah"}})
}

func RemoveCourse(c *gin.Context) {
	var user models.User
	nim := c.Param("nim")
	courseId, _ := strconv.Atoi(c.Param("courseId"))

	currentUserNim, err := ExtractNIM(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	if currentUserNim != nim {
		c.JSON(http.StatusBadRequest, gin.H{"meta": models.Meta{false, "Anda tidak memiliki akses " +
			"untuk menambah mata kuliah"}})
		c.Abort()
		return
	}
	currentUserRecord := database.Db.Where("nim = ?", currentUserNim).First(&user)
	if currentUserRecord.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, currentUserRecord.Error.Error()}})
		c.Abort()
		return
	}

	removeCourseErr := database.Db.Model(&user).Association("Courses").Delete(&models.Course{ID: uint(courseId)})
	if removeCourseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, removeCourseErr.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"meta": models.Meta{true, "mata kuliah berhasil dihapus"}})
}
