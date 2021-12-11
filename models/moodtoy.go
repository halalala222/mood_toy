package models

import "gorm.io/gorm"

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

