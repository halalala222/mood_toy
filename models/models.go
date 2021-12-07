package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var  Users []User

//User 一个用户数据结构，里面有着Id，Name(name是用户使用的时候的一个昵称)，Username(username是用来登录注册的)
//如果用户没有更改昵称的话，name就等于username？？？？？
type User struct {
	Id uint `json:"id"`
	Userid string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	gorm.Model
}

//Moodtoy 一个mood娃娃的状态
type Moodtoy struct {
	Id uint `json:"id"`
	Sex string `json:"sex"`
	Feeling  uint `json:"feeling"`
	Hair uint `json:"hair"`
 	Eyes  uint  `json:"eyes"`
	Face  uint `json:"face"`
	Clothes uint `json:"clothes"`
	Trousers uint `json:"trousers"`
}

//Lists 一个用户轻日记的一个结构体
type Lists struct {
	List string `json:"list"`
	Time string `json:"time"`
	Feeling string `json:"feeling"`
}

//Myclaims 定义一个我的自己字段
type Myclaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

