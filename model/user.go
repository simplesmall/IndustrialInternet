package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
// 1. 创建 time.Time 类型的副本 XTime；
type XTime struct {
	time.Time
}

const TimeFormat = "2006-01-02 15:04:05"

//MyTime 自定义时间
type MyTime time.Time

func (t *XTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = XTime{now}
	return err
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
type BaseModel struct {
	ID        uint `gorm:"primary_key,AUTO_INCREMENT"`
	CreatedAt XTime
	UpdatedAt XTime
	DeletedAt *XTime `sql:"index"`
}

// 用户
type User struct {
	BaseModel
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Username string `gorm:"size:100;column(username)" json:"username" form:"username"`
	Password string `gorm:"size:255;column(password)" json:"password" form:"password"`
	Email    string `gorm:"size:200;column(email)" json:"email" form:"email"`
	Status   string `gorm:"size:20;column(status)" json:"status" form:"status"`
	Avatar   string `gorm:"size:200;column(avatar)" json:"avatar" form:"avatar"`
	Type     string `gorm:"size:20;column(type)" json:"type" form:"type"`
	Phone    string `gorm:"size:20;column(phone)" json:"phone" form:"phone"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
	Token    string `gorm:"size:500;column:token" json:"-"`
	Expire   int64  `gorm:"column:expire" json:"-"`
	Role     []Role `gorm:"many2many:user_roles"`
}

//权限
type Permission struct {
	BaseModel
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Parent   string `gorm:"size:10;column(parent)" json:"parent" form:"parent"`
	Status   string `gorm:"size:10;column(status)" json:"status" form:"status"`
	Uid      string `gorm:"size:50;column(uid)" json:"uid" form:"uid"`
	Type     string `gorm:"size:10;column(type)" json:"type" form:"type"`
	Url      string `gorm:"size:500;column(url)" json:"url" form:"url"`
	Icon     string `gorm:"size:500;column(icon)" json:"icon" form:"icon"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
}

//角色
type Role struct {
	BaseModel
	Name       string       `gorm:"size:100;column(name)" json:"name" form:"name"`
	UserId     int          `gorm:"" json:"user_id"`
	Desc       string       `gorm:"size:500;column(desc)" json:"desc" form:"desc"`
	User       User         `gorm:"foreignKey:UserId"`
	Permission []Permission `gorm:"many2many:role_permissions"`
}
