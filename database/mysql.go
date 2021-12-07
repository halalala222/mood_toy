package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"lwh.com/models"
	"lwh.com/setting"
	"strings"
)

//Db 一个指针，用来存放数据库类型？（不懂）
var Db *gorm.DB
//Link 定义一个映射一个结构体表格，然后链接数据库的函数
func Link()  {
	var err error
	path := strings.Join([]string{setting.Configone.Mysql.Username, ":", setting.Configone.Mysql.Password, "@tcp(",setting.Configone.Mysql.Host, ":", setting.Configone.Mysql.Port, ")/","user", "?charset=utf8&parseTime=true&loc=Local"},"")
	//链接数据库
	Db , err = gorm.Open(mysql.Open(path),&gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	//创建一个表格
	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic(err.Error())
	}

}
