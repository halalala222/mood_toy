package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var  Users []User

//User 一个用户数据结构，里面有着Id，Name(name是用户使用的时候的一个昵称)，Username(username是用来登录注册的)
//一个user有着做个list然后就是has many的关联
type User struct {
	//binding:"required"就是说这个不能传入是空的，如果是空的就会返回一个err
	UserID string `json:"user_id"`//userid作为用户的表示符，因为userid是不能够重复的，但是username是可以重复的？
	UserName string `json:"user_name"`
	Password string `json:"password"`
	//user has many关联
	Diaries []Diary `gorm:"foreignKey:UserRefer;references:id"`
	//user has one关联
	UserImg UserImg `gorm:"foreignKey:UserRefer2;references:id"`
	Feeling uint `gorm:"default:100"`
	MoodtoyID uint
	gorm.Model
}

//Moodtoy 一个mood娃娃的状态,一个moodtoy有着多个user
//就是moodtoyhas many的关联
type Moodtoy struct {
	Hair string `json:"hair"`
	Eyebrow string `json:"eyebrow"`
 	Eyes string  `json:"eyes"`
	Mouth string `json:"mouth"`
	Clothes string `json:"clothes"`
	gorm.Model
}

var Diaries []Diary
//Diary 一个用户轻日记的一个结构体
type Diary struct {
	TextContext string `json:"text_context"`
	Time string `json:"time"`
	Feeling string `json:"feeling"`
	UserRefer uint//UseRefer 用来对应user中的id
	gorm.Model
}

//Color 一个color中的表格对应中文与number对应
type Color struct {
	Chinese string
	Number uint
}

//UserImg 定义一个存储用户的头像路径以及用户的一个关联数据库
//这个是需要入库的
type UserImg struct {
	UserRefer2 uint
	ImgUrl string
	gorm.Model
}
//Myclaims 定义一个我的自己字段
//添加了一个userid的字段
type Myclaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var Hairs []ToyHair
type ToyHair struct {
	HairUrl string
}

var Eyebrows []ToyEyebrow
//ToyEyebrow 加入一个心情的等级
type ToyEyebrow struct {
	EyebrowUrl string
	FeelingStatus uint
	Number uint
}

var Mouthes  []ToyMouth
type ToyMouth struct {
	MouthUrl string
}

var Clothes  []ToyClothes
type ToyClothes struct {
	ClothesUrl string
}

var Eyes []ToyEyes
type ToyEyes struct {
	EyesUrl string
}

var Init  []ToyInit
type ToyInit struct {
	InitUrl string
}

type Utterance struct {
	Utt string `json:"utt"`
}

type Txt struct {
	Text string `json:"text"`
}