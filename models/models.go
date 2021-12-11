package models

import (
	"gorm.io/gorm"
)



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
