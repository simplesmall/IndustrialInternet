package config

import (
	"IndustrialInternet/model"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
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
// 用户
type UserForm struct {
	model.User
	Roles []int `json:"Roles"`
}
//权限
type PermissionForm struct {
	model.Permission
}
func AutoMigrate() {
	DB.AutoMigrate(&model.User{},&model.Role{},&model.Permission{})
	//填充数据
	db :=DB
	result := db.Find(model.User{}, "1")
	if result.Error != nil {
		var rootInfo UserForm
		var roles []int
		rootInfo.User.Name = "超级管理员"
		rootInfo.User.Username = "root"
		rootInfo.User.Password = "123456"
		rootInfo.User.Email = "root@threejs.com"
		rootInfo.User.Type = "00"
		rootInfo.User.Phone = "13999999999"
		rootInfo.User.Describe = "就是超管"
		rootInfo.Roles = roles

		_, _ = CreateUser(rootInfo)
	}
	var perArr [9]PermissionForm

	perArr[0].Permission = model.Permission{Name: "首页", Parent: "0", Status: "1", Type: "01", Uid: "home", Url: "test", Icon: "test", Describe: "首页"}
	perArr[1].Permission = model.Permission{Name: "系统管理", Parent: "0", Status: "1", Type: "01", Uid: "system", Url: "test", Icon: "test", Describe: "系统管理"}
	perArr[2].Permission = model.Permission{Name: "用户管理", Parent: "2", Status: "1", Type: "02", Uid: "user", Url: "test", Icon: "test", Describe: "用户管理"}
	perArr[3].Permission = model.Permission{Name: "日志管理", Parent: "2", Status: "1", Type: "02", Uid: "log", Url: "test", Icon: "test", Describe: "日志管理"}
	perArr[4].Permission = model.Permission{Name: "模型管理", Parent: "0", Status: "1", Type: "01", Uid: "model", Url: "test", Icon: "test", Describe: "模型管理"}
	perArr[5].Permission = model.Permission{Name: "应用管理", Parent: "0", Status: "1", Type: "01", Uid: "app", Url: "test", Icon: "test", Describe: "应用管理"}
	perArr[6].Permission = model.Permission{Name: "角色管理", Parent: "2", Status: "1", Type: "02", Uid: "role", Url: "test", Icon: "test", Describe: "角色管理"}
	perArr[7].Permission = model.Permission{Name: "分组管理", Parent: "5", Status: "1", Type: "02", Uid: "model_group", Url: "test", Icon: "test", Describe: "分组管理"}
	perArr[8].Permission = model.Permission{Name: "模型列表", Parent: "5", Status: "1", Type: "02", Uid: "model_list", Url: "test", Icon: "test", Describe: "模型列表"}

	perArr[0].Permission.BaseModel.ID = 1
	perArr[1].Permission.BaseModel.ID = 2
	perArr[2].Permission.BaseModel.ID = 3
	perArr[3].Permission.BaseModel.ID = 4
	perArr[4].Permission.BaseModel.ID = 5
	perArr[5].Permission.BaseModel.ID = 6
	perArr[6].Permission.BaseModel.ID = 7
	perArr[7].Permission.BaseModel.ID = 8
	perArr[8].Permission.BaseModel.ID = 9

	for i := 0; i < len(perArr); i++ {
		db.FirstOrCreate(&perArr[i].Permission)
	}
}

func CreateUser(userData UserForm) (uint, error) {

	var UserList []model.User
	//查重
	err := DB.Table("user").Where("username = ?", userData.Username).Find(&UserList).Error
	if err != nil {
		return 0, errors.New("user search fail")
	} else {
		if len(UserList) > 0 {
			return 0, errors.New("user exit")
		}
	}

	originPwd := []byte(userData.User.Password)
	//此方法生成hash值
	hashPwd, _ := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost) //password为string类型
	userData.User.Password = string(hashPwd)
	result := DB.Create(&userData.User)
	id := userData.User.ID

	var roles []model.Role
	DB.Find(&roles, userData.Roles)

	//更新关系
	DB.Model(&userData.User).Association("Role").Replace(roles)

	if result.Error != nil {
		return 0, errors.New("user create fail")
	}
	return id, nil
}