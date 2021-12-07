package apipackage

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"lwh.com/database"
	"lwh.com/jwt"
	"lwh.com/models"
	"net/http"
)

//Login 定义一个登录的api判断用户输入是否正确，然后返回给用户一个token
func Login(c *gin.Context)  {
	var user = models.User{}
	var user1 = models.User{}
	if err := c.ShouldBind(&user);err == nil {
		//从数据库中查找信息
		database.Db.Where("userid = ?",user.Userid).First(&user1)
		//对于存在数据库中的密码进行解码
		err1 := bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password))
		if err1 != nil {
			c.JSON(http.StatusOK,gin.H{
				"message" : "用户id或者密码错误",
			})
		}else {
			//生成一个token
			token,err2 := jwt.GenToken(user.Userid)
			if err2 == nil {
				c.JSON(http.StatusOK,gin.H{
					"token" :token,
					"user" : gin.H{
						"userid" : user.Userid,
						"username" : user.Username,
					},
				})
			}
		}
	}
}


//Register 定义一个注册的api，对于用户注册需要判断用户的
func Register(c *gin.Context)  {
	var user = models.User{}
	if err :=c.ShouldBind(&user);err == nil{
		//对用户传过来的密码进行一个加密
		hash, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err1 != nil {
			log.Panic(err1.Error())
		}
		encodePW := string(hash)
		//把加密过后的东西赋值给password然后存入数据库中
		user.Password = encodePW
		//查找表中userid这一列
		database.Db.Select("userid").Find(&models.Users)
		//遍历切片中的结构体，如果中有用户名相同的话，那么就返回一个message该用户ID已经被注册
		for _,user1 := range models.Users{
			if user.Userid == user1.Userid {
				c.JSON(http.StatusOK,gin.H{
					"message" : "该用户ID已经被注册",
				})
			}
		}
		//如果没有的话就成功注册
		database.Db.Model(&models.User{}).Create(&user)
	} else{
		c.JSON(http.StatusOK,gin.H{
			"message" : err.Error(),
		})
	}
}

//CORS 一个跨域的中间件对于用户的请求头上添加东西
func CORS() func(c *gin.Context) {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		}
	}
}

