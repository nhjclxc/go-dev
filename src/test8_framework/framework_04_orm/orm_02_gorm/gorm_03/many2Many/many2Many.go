package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_03/config"
	"log"
)


type Student struct {
	Id      uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name    string    `gorm:"column:name" json:"name"`

	// gorm:"many2many:<中间表名>;joinForeignKey: 当前 struct 的字段 ;joinReferences:关联 struct 的字段"
	Courses []Course `gorm:"many2many:student_courses;joinForeignKey:StudentId;joinReferences:CourseId"`
}

type Course struct {
	Id    uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title string    `gorm:"column:title" json:"title"`

	Students []Student `gorm:"many2many:student_courses;joinForeignKey:CourseId;joinReferences:StudentId" json:"students"` // 反向多对多
}

type StudentCourse struct {
	StudentId uint64 `gorm:"column:student_id"`
	CourseId uint64 `gorm:"column:course_id"`
}

func (Student) TableName() string {
	return "students"
}

func (Course) TableName() string {
	return "courses"
}

func GetStudentWithCourses(db *gorm.DB, studentID uint64) (*Student, error) {
	var student Student
	err := db.Preload("Courses").First(&student, "id = ?", studentID).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetCourseWithStudents(db *gorm.DB, courseID uint64) (*Course, error) {
	var course Course
	err := db.Preload("Students").First(&course, "id = ?", courseID).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}


func main() {

	studentID := uint64(1)
	student, err := GetStudentWithCourses(config.DB, studentID)
	if err != nil {
		log.Fatal("查询学生失败:", err)
	}

	fmt.Printf("学生: %s\n选修课程:\n", student.Name)
	for _, course := range student.Courses {
		fmt.Printf("- %s\n", course.Title)
	}


	fmt.Println()
	fmt.Println()

	courseID := uint64(2)
	course, err := GetCourseWithStudents(config.DB, courseID)
	if err != nil {
		log.Fatal("查询失败:", err)
	}

	fmt.Printf("课程：%s\n选修学生：\n", course.Title)
	for _, student := range course.Students {
		fmt.Printf("- %s\n", student.Name)
	}
}
