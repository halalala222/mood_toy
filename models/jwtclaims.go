package models

import "github.com/dgrijalva/jwt-go"

//Myclaims 定义一个我的自己字段
//添加了一个userid的字段
type Myclaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

