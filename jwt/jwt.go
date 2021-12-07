package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lwh.com/models"
	"lwh.com/setting"
	"net/http"
	"strings"
)

//GenToken 定义一个生成token的函数
func GenToken(username string)(string,error)  {
	c := models.Myclaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer :"lifiverings",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString([]byte(setting.Configone.Jwt.Secret))
}

//ParseToken 定义一个解析token的方法
func ParseToken(tokenString string) (*models.Myclaims,error) {
	token,err := jwt.ParseWithClaims(tokenString,&models.Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return []string{setting.Configone.Jwt.Secret},nil
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

func authHandler(c *gin.Context)  {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"message" : "无效参数",
		})
		return
	}
	//这里需要从数据库中提取东西出来
	url := "halalala222"
	//应该是进行一个加密，对用户传入的密码进行一个加密
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message" : err,
		})
	}
	//类型转换为字符串
	encodePW := string(hash)
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(url))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message" : "输入的密码错误",

		})
	}
}

//JWTAuthMiddleware 定义一个中间件用来判断前端传来的token格式是否正确以及前端的token是否是对俄
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//这个是获取请求头中的内容
		authHandler := c.Request.Header.Get("Authorization")
		if authHandler == "" {
			c.JSON(http.StatusOK,gin.H{
				"message" : "请求头中auth为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHandler," ",2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK,gin.H{
				"message" : "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		//对请求头中的tokensrting 进行解码
		mc,err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"message" : "无效的token",

			})
			return
		}
		c.Set("username",mc.Username)
		c.Next()
	}
}
