package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	Db *gorm.DB
}

// ฟังก์ชันสร้าง instance ของ SubjectRepository และส่งคืนกลับไป
func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{Db: db}
}

// ฟังก์ชันดึงข้อมูลวิชาทั้งหมดจากฐานข้อมูลและส่งกลับไปให้ผู้ใช้งานผ่าน JSON response
func (r *SubjectRepository) GetSubjects(c *gin.Context) {
	var subjects []Subject
	r.Db.Find(&subjects)
	c.JSON(200, subjects)
}

// ฟังก์ชันเพิ่มข้อมูลวิชาลงในฐานข้อมูลและส่งกลับไปให้ผู้ใช้งานผ่าน JSON response
func (r *SubjectRepository) PostSubject(c *gin.Context) {
	var newSubject Subject
	c.BindJSON(&newSubject)
	r.Db.Create(&newSubject)
	c.JSON(200, newSubject)
}

// ฟังก์ชันค้นหาวิชาจากฐานข้อมูล โดยใช้ id เป็นเงื่อนไข
func (r *SubjectRepository) GetSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.First(&subject, id)
	c.JSON(200, subject)
}

// ฟังก์ชันอัปเดตข้อมูลวิชาลงในฐานข้อมูลและส่งกลับไปให้ผู้ใช้งานผ่าน JSON response
func (r *SubjectRepository) UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.First(&subject, id)
	c.BindJSON(&subject)
	r.Db.Save(&subject)
	c.JSON(200, subject)
}

// ฟังก์ชันลบข้อมูลวิชาออกจากฐานข้อมูลและส่งกลับไปให้ผู้ใช้งานผ่าน JSON response
func (r *SubjectRepository) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.Delete(&subject, id)
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}
