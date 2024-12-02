package controllers

import (
	"college-be/database"
	"college-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMajors(c *gin.Context) {
	var majors []models.Major

	record := database.Db.Find(&majors)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "success"}, "data": majors})
}
