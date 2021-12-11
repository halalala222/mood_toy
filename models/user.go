package models

import "gorm.io/gorm"

var  Users []User

//User 一个用户数据结构，里面有着Id，Name(name是用户使用的时候的一个昵称)，Username(username是用来登录注册的)
//一个user有着做个list然后就是has many的关联
type User struct {
	//binding:"required"就是说这个不能传入是空的，如果是空的就会返回一个err
	UserID string `json:"user_id"`//userid作为用户的表示符，因为userid是不能够重复的，但是username是可以重复的？
	UserName string `json:"username"`
	Password string `json:"password"`
	//user has many关联
	Diaries []Diary `gorm:"foreignKey:UserRefer;references:id"`
	//user has one关联
	UserImg UserImg `gorm:"foreignKey:UserRefer2;references:id"`
	Feeling uint `gorm:"default:100"`
	MoodtoyID uint
	gorm.Model
}
