package config

import (
	"IndustrialInternet/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitConnect()  {
	dsn:="root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local"
	DB,_=gorm.Open("mysql",dsn)
	//if err != nil {
	//	panic(err)
	//}
	// 强制限制表明为自己定义的模型名单数形式
	DB.SingularTable(true)

	//student.CreateTable()
	//AutoMigrate()
}

func CloseDB() {
	_ = DB.Close()
}
func AutoMigrate() {
	//DB.AutoMigrate(&student.Student{},&student.Class{},&student.IDCard{},&student.Teacher{})
	DB.AutoMigrate(&model.User{},&model.Role{},&model.Permission{})
	// 通过结构体创建关联关系以及插入数据
	/*i:=student.IDCard{
		Num: 123456,
	}
	s:=student.Student{
		StudentName: "jiandan",
		IDCard: i,					// 关联学生和学生Idcard信息
	}
	t:=student.Teacher{
		TeacherName: "老师傅了",
		Students: []student.Student{s},  // 多对多关系创建方式,任意一方创建
	}
	c:=student.Class{
		ClassName: "jiandan的班级",
		Students: []student.Student{s},  //把学生放进班级里
	}
	_ = DB.Debug().Create(&c).Error // 创建班级
	_ = DB.Debug().Create(&t).Error // 创建老师
	*/
}