package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"lwh.com/models"
	"lwh.com/setting"
	"strings"
)
var err error

//定义一函数就是链接数据库的函数
func  mysqllink(dbname string)  {
	path := strings.Join([]string{setting.Configone.Mysql.Username, ":", setting.Configone.Mysql.Password, "@tcp(",setting.Configone.Mysql.Host, ":", setting.Configone.Mysql.Port, ")/",dbname, "?charset=utf8&parseTime=true&loc=Local"},"")
	//链接数据库
	Db , err = gorm.Open(mysql.Open(path),&gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
}
//Db 一个指针，用来存放数据库类型？（不懂）
var Db *gorm.DB
//Link 定义一个映射一个结构体表格，然后链接数据库的函数
func Link()  {
	//链接数据库
	mysqllink("user")
	//自动跟随结构体在数据库中建立和更新表格
	err = Db.AutoMigrate(&models.User{})
	err = Db.AutoMigrate(&models.ToyHair{})
	err = Db.AutoMigrate(&models.ToyInit{})
	err = Db.AutoMigrate(&models.ToyEyes{})
	err = Db.AutoMigrate(&models.ToyClothes{})
	err = Db.AutoMigrate(&models.ToyMouth{})
	err = Db.AutoMigrate(&models.ToyEyebrow{})
	err = Db.AutoMigrate(&models.Moodtoy{})
	err = Db.AutoMigrate(&models.Color{})
	err = Db.AutoMigrate(&models.Diary{})
	err = Db.AutoMigrate(&models.UserImg{})
	if err != nil {
		log.Panic(err.Error())
	}

}
