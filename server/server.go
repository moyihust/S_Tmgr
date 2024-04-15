package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Sno         string `gorm:"primaryKey"`
	Sname       string `gorm:"unique"`
	Ssex        string
	Sage        int16
	Sdept       string
	Scholarship string
}

func (Student) TableName() string {
	return "Student"
}

type Course struct {
	Cno     string `gorm:"primaryKey"`
	Cname   string
	Cpno    string `gorm:"foreignKey:Cno"`
	Ccredit int16
}

func (Course) TableName() string {
	return "Course"
}

type SC struct {
	Sno   string `gorm:"primaryKey"`
	Cno   string `gorm:"primaryKey"`
	Grade int16
}

func (SC) TableName() string {
	return "Sc"
}

// // 将字符串转换为指定的编码格式
// func transformString(str string, encoder transform.Transformer) string {
// 	result, _, err := transform.String(encoder, str)
// 	if err != nil {
// 		return str
// 	}
// 	return result
// }

func main() {

	// 初始化 logrus
	// logrus.SetFormatter(&logrus.TextFormatter{})
	// logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetOutput(os.Stdout)

	dsn := "root@tcp(127.0.0.1:2881)/S_T?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	fmt.Println("连接数据库成功")
	r := gin.Default()
	// 添加 CORS 中间件
	r.Use(cors.Default())
	r.GET("/student", func(c *gin.Context) {
		var students []Student
		db.Find(&students)
		// fmt.Println(students)
		c.JSON(200, students)
	})
	r.POST("/inputStudent", func(c *gin.Context) {
		var student Student
		if err := c.ShouldBind(&student); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&student)
		c.JSON(200, student)
	})
	r.PUT("/updateStudent/:sno", func(c *gin.Context) {
		var student Student
		if err := c.ShouldBind(&student); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		sno := c.Param("sno")
		db.Model(&Student{}).Where("sno = ?", sno).Updates(student)
		c.JSON(200, student)
	})
	r.DELETE("/deleteStudent/:sno", func(c *gin.Context) {
		sno := c.Param("sno")
		db.Where("sno = ?", sno).Delete(&Student{})
		c.JSON(200, gin.H{"sno": sno})
	})
	r.GET("/studentCourses/:sno", func(c *gin.Context) {
		sno := c.Param("sno")
		type Result struct {
			Sno         string
			Sname       string
			Ssex        string
			Sage        int16
			Sdept       string
			Scholarship string
			Cno         string
			Cname       string
			Grade       int16
		}
		var results []Result
		db.Raw(`SELECT Student.Sno, Student.Sname, Student.Ssex, Student.Sage, Student.Sdept, Student.Scholarship, Course.Cno, Course.Cname, SC.Grade 
		FROM Student
		LEFT JOIN SC ON Student.Sno = SC.Sno
		LEFT JOIN Course ON SC.Cno = Course.Cno
		WHERE Student.Sno = ?`, sno).Scan(&results)
		c.JSON(200, results)
	})
	r.Run(":1145")
}
