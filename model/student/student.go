package student

import (
	"IndustrialInternet/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Class struct {
	gorm.Model
	ClassName string `json:"class_name"`

	// 一个班级有多个学生   学生通过ClassID知道自己在哪个班级    一对多关系
	Students  []Student   // 班级里面有多个学生
}
type Student struct {
	gorm.Model
	StudentName string  `json:"student_name"`
	ClassID uint		`json:"class_id"`// 告诉学生属于哪个班级

	// 每个学生有一张IDCard  IDCard 通过 StudentID知道是哪个学生的   一对一关系
	IDCard IDCard

	// 多对多  TeacherID uint
	Teachers []Teacher `gorm:"many2many:student_teacher"` // 一个学生有多个老师
}
type IDCard struct {
	gorm.Model
	StudentID uint  	// 标识属于哪个学生的卡
	Num int			`json:"num"`
}
type Teacher struct {
	gorm.Model
	TeacherName string	`json:"teacher_name"`

	// 多对多		StudentID uint
	Students []Student `gorm:"many2many:student_teacher"`
}

func CreateTable() {
	//config.DB.AutoMigrate(&Student{},&Class{},&IDCard{},&Teacher{})
}

/*
请求格式范例
{
    "student_name":"xiaomi",
    "class_id":1,
    "IDCard":{
        "num":55555
    },
    "Teachers":[{
      "teacher_name":"老师一号"
    },{
      "teacher_name":"老师一号"
    }]
}
*/
// 创建学生表
func CreateStudent() gin.HandlerFunc{
	return func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		config.DB.Create(&student)
	}
}
//获取学生表
func GetStudent() gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		// 关联查询之前先进行预加载 ,否则关联数据不能正确返回
		config.DB.Preload("Teachers").Preload("IDCard").First(&student,"id = ?",id)
		c.JSON(http.StatusOK,gin.H{
			"student":student,
		})
	}
}
// 获取班级
func GetClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("ID")
		var class Class
		config.DB.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").First(&class,"id = ?",id)
		c.JSON(http.StatusOK,gin.H{
			"class":class,
		})
	}
}
