// models/StudentRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentRepository struct {
	Db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{Db: db}
}

func (r *StudentRepository) GetStudents(c *gin.Context) {
	var students []Student // Change variable name from "Students" to "students"
	r.Db.Find(&students)
	c.JSON(200, students)
}

func (r *StudentRepository) PostStudent(c *gin.Context) {
	var newStudent Student // Change variable name from "newStudent" to "newStudent"
	c.BindJSON(&newStudent)
	r.Db.Create(&newStudent)
	c.JSON(200, newStudent)
}

func (r *StudentRepository) GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student // Change variable name from "Student" to "student"
	r.Db.First(&student, id)
	c.JSON(200, student)
}

func (r *StudentRepository) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student // Change variable name from "Student" to "student"
	r.Db.First(&student, id)
	c.BindJSON(&student)
	r.Db.Save(&student)
	c.JSON(200, student)
}

func (r *StudentRepository) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student // Change variable name from "Student" to "student"
	r.Db.Delete(&student, id)
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}
