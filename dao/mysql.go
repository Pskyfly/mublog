package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	//注意不能加:,否则会生成一个局部变量而没有初始化全局变量
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8")
	//DB, err = gorm.Open("mysql", "root:123456@tcp(192.168.62.186:3306)/bubble?charset=utf8")
	if err != nil {
		return
	}
	return DB.DB().Ping() //如果有错返回错误信息，否则返回nil
}
