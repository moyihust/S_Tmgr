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
	Cpno    *string `gorm:"foreignKey:Cno"`
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

	dsn := "moyi:Lgt1441339168@tcp(obmtkk8o4cuc0ig0-mi.cn-hangzhou.oceanbase.aliyuncs.com:3306)/S_T?charset=utf8mb4&parseTime=True&loc=Local"

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
	r.GET("/course", func(c *gin.Context) {
		var courses []Course
		db.Find(&courses)
		c.JSON(200, courses)
	})
	r.GET("/sc", func(c *gin.Context) {
		type Result struct {
			Sno   string
			Sname string
			Cno   string
			Cname string
			Grade int
		}
		var results []Result
		db.Raw(`SELECT s.Sno, s.Sname, c.Cno, c.Cname, s2.Grade 
				FROM Student s, SC s2, Course c 
				WHERE s.Sno = s2.Sno AND s2.Cno = c.Cno`).Scan(&results)
		c.JSON(200, results)
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
	r.POST("/inputCourse", func(c *gin.Context) {
		var course Course
		if err := c.ShouldBind(&course); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if course.Cpno != nil && *course.Cpno == "" {
			course.Cpno = nil
		}
		db.Create(&course)
		c.JSON(200, course)
	})
	r.POST("/inputSC", func(c *gin.Context) {
		var sc SC
		if err := c.ShouldBind(&sc); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&sc)
		c.JSON(200, sc)
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
	r.PUT("/updateCourse/:cno", func(c *gin.Context) {
		var course Course
		if err := c.ShouldBind(&course); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		cno := c.Param("cno")
		db.Model(&Course{}).Where("cno = ?", cno).Updates(course)
		c.JSON(200, course)
	})
	r.PUT("/updateSC/:sno/:cno", func(c *gin.Context) {
		var sc SC
		if err := c.ShouldBind(&sc); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		sno := c.Param("sno")
		cno := c.Param("cno")
		db.Model(&SC{}).Where("sno = ? AND cno = ?", sno, cno).Updates(sc)
		c.JSON(200, sc)
	})
	r.DELETE("/deleteStudent/:sno", func(c *gin.Context) {
		sno := c.Param("sno")
		db.Where("sno = ?", sno).Delete(&Student{})
		c.JSON(200, gin.H{"sno": sno})
	})
	r.DELETE("/deleteCourse/:cno", func(c *gin.Context) {
		cno := c.Param("cno")
		db.Where("cno = ?", cno).Delete(&Course{})
		c.JSON(200, gin.H{"cno": cno})
	})
	r.DELETE("/deleteSC/:sno/:cno", func(c *gin.Context) {
		sno := c.Param("sno")
		cno := c.Param("cno")
		db.Where("sno = ? AND cno = ?", sno, cno).Delete(&SC{})
		c.JSON(200, gin.H{"sno": sno, "cno": cno})
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
	r.GET("/studentStats/:sdept", func(c *gin.Context) {
		sdept := c.Param("sdept")
		type Stats struct {
			Sdept        string
			AverageGrade float64
			BestGrade    int
			WorstGrade   int
			Excellent    int
			Fail         int
		}
		var stats []Stats
		db.Raw(`
			SELECT 
				s.Sdept,
				AVG(s2.Grade) as AverageGrade,
				MAX(s2.Grade) as BestGrade,
				MIN(s2.Grade) as WorstGrade,
				SUM(CASE WHEN s2.Grade >= 90 THEN 1 ELSE 0 END) as Excellent,
				SUM(CASE WHEN s2.Grade < 60 THEN 1 ELSE 0 END) as Fail
			FROM Student s
			JOIN SC s2 ON s.Sno = s2.Sno
			WHERE s.Sdept = ?
			GROUP BY s.Sdept
		`, sdept).Scan(&stats)
		c.JSON(200, stats)
	})
	r.GET("/sdeptlist", func(c *gin.Context) {
		type Sdept struct {
			Sdept string
		}
		var sdepts []Sdept
		db.Raw(`SELECT DISTINCT Sdept FROM Student`).Scan(&sdepts)
		c.JSON(200, sdepts)
	})
	r.GET("/search/:sdept", func(c *gin.Context) {
		sdept := c.Param("sdept")
		type Result struct {
			Sno   string
			Sname string
			Cno   string
			Cname string
			Grade int
		}
		var results []Result
		db.Raw(`SELECT s.Sno, s.Sname, c.Cno, c.Cname, s2.Grade 
				FROM Student s, SC s2, Course c 
				WHERE s.Sno = s2.Sno AND s2.Cno = c.Cno AND s.Sdept = ?
				ORDER BY s2.Grade DESC`, sdept).Scan(&results)
		c.JSON(200, results)
	})
	r.Run(":1145")
}
