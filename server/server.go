package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/transform"
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

// 将字符串转换为指定的编码格式
func transformString(str string, encoder transform.Transformer) string {
	result, _, err := transform.String(encoder, str)
	if err != nil {
		return str
	}
	return result
}

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

	r := gin.Default()
	// 添加 CORS 中间件
	r.Use(cors.Default())
	r.GET("/student", func(c *gin.Context) {
		var students []Student
		db.Find(&students)
		c.JSON(200, students)
	})

	r.Run(":1145")
}
