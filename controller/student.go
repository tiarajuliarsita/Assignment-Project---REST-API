package controller

import (
	"assigment_project_rest_api/database"
	"assigment_project_rest_api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateStudent(c *gin.Context) {
	db := database.GetDB()
	newStudent := models.Student{}

	err := c.ShouldBindJSON(&newStudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	student := models.Student{
		Name:   newStudent.Name,
		Age:    newStudent.Age,
		Scores: newStudent.Scores,
	}

	err = db.Create(&student).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "succes created student",
		"status":  true,
		"student": student,
	})

}

func GetAllStudent(c *gin.Context) {
	db := database.GetDB()
	students := []models.Student{}
	err := db.Preload("Scores").Find(&students).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

func UpdateStudentByID(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	updatedStudent := models.Student{}

	err := c.ShouldBindJSON(&updatedStudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existStudent models.Student
	err = db.First(&existStudent, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	existStudent.Name = updatedStudent.Name
	existStudent.Age = updatedStudent.Age
	existStudent.Scores = updatedStudent.Scores

	err = db.Save(&existStudent).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated successfully",
		"status":  true,
		"student": existStudent,
	})
}

func DeleteStudentByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	db := database.GetDB()
	var student models.Student
	var score models.Score

	err := db.Where("student_id= ?", idInt).Delete(&score).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = db.Where("id = ?", idInt).Delete(&student).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "secces deleted student",
	})
}
