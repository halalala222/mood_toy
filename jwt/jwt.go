package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"lwh.com/models"
	"lwh.com/setting"
)

//GenToken 定义一个生成token的函数
//通过用户传入的userid来判断是否是该用户
func GenToken(userid string)(string,error)  {
	c := models.Myclaims{
		UserID: userid,
		StandardClaims: jwt.StandardClaims{
			Issuer :"moodtoy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString([]byte(setting.Configone.Jwt.Secret))
}

//ParseToken 定义一个解析token的方法
func ParseToken(tokenString string) (*models.Myclaims,error) {
	token,err := jwt.ParseWithClaims(tokenString,&models.Myclaims{}, func(token *jwt.Token) (i interface{},err error) {
		return []byte(setting.Configone.Jwt.Secret),nil
	})
	if err != err {
		return nil,err
	}
	//token.Claims.(*models.Claims)一个强制类型转换，转换为这个的一个指针
	if claims,ok :=token.Claims.(*models.Myclaims);ok && token.Valid{
		return claims,nil
	}
	//最后如说token绑定可以但是还是错误，那么就是返回一个nil的models.Claims的指针，然后错误是无效的token
	return nil,errors.New("invalid token")
}
