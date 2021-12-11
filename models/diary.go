package models

import "gorm.io/gorm"

var Diaries []Diary
//Diary 一个用户轻日记的一个结构体
type Diary struct {
	Title string `json:"title"`
	TextContext string `json:"text_content"`
	Time string `json:"time"`
	Feeling string `json:"feeling"`
	UserRefer uint//UseRefer 用来对应user中的id
	gorm.Model
}
