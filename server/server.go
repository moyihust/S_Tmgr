package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	gorm.Model
	Sno         string `gorm:"primary_key"`
	Sname       string `gorm:"unique"`
	Ssex        string
	Sage        int
	Sdept       string
	Scholarship string
	Courses     []Course `gorm:"many2many:student_courses;"`
}

type Course struct {
	gorm.Model
	Cno      string `gorm:"primary_key"`
	Cname    string
	Cpno     string
	Ccredit  int
	Students []Student `gorm:"many2many:student_courses;"`
}

type SC struct {
	gorm.Model
	Sno   string `gorm:"primary_key"`
	Cno   string `gorm:"primary_key"`
	Grade int
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:2881)/S_T?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Student{}, &Course{}, &SC{})

	// Create a new student
	db.Create(&Student{Sno: "200215121", Sname: "张三", Sage: 20, Ssex: "男", Sdept: "CS", Scholarship: "是"})

	// Update student's information
	db.Model(&Student{}).Where("sno = ?", "200215121").Update("Sname", "李四")

	// Add a new course
	db.Create(&Course{Cno: "1", Cname: "计算机科学", Cpno: "", Ccredit: 4})

	// Update course information
	db.Model(&Course{}).Where("cno = ?", "1").Update("Cname", "计算机科学与技术")

	// Delete a course that no student has chosen
	db.Where("cno NOT IN (?)", db.Table("sc").Select("cno")).Delete(&Course{})

	// Enter a student's grade
	db.Create(&SC{Sno: "200215121", Cno: "1", Grade: 95})

	// Update a student's grade
	db.Model(&SC{}).Where("sno = ? AND cno = ?", "200215121", "1").Update("Grade", 96)

	// Calculate the average, best, worst grades, excellent rate, and number of failures of a student
	var avgGrade, bestGrade, worstGrade, excellentRate, numberOfFailures int
	db.Model(&SC{}).Where("sno = ?", "200215121").Select("AVG(grade) as avg_grade, MAX(grade) as best_grade, MIN(grade) as worst_grade").Scan(&avgGrade, &bestGrade, &worstGrade)
	db.Model(&SC{}).Where("sno = ? AND grade >= 90", "200215121").Count(&excellentRate)
	db.Model(&SC{}).Where("sno = ? AND grade < 60", "200215121").Count(&numberOfFailures)

	// Rank students' grades by department
	var students []Student
	db.Preload("Courses").Order("grade desc").Find(&students)

	// Input a student number, display the student's basic information and course selection information
	var student Student
	db.Preload("Courses").Where("sno = ?", "200215121").First(&student)

	fmt.Println(student)
}
