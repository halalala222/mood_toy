package apis

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
	var password string
	if err := c.ShouldBind(&user);err == nil {
		//从数据库中查找信息
		password = user.Password
		//对于存在数据库中的密码进行解码
		database.Db.Model(&models.User{}).Where("user_id=?",user.UserID).First(&user)
		err1 := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		if err1 != nil {

			c.JSON(http.StatusUnauthorized ,gin.H{
				"message" : "用户id或者密码错误",
			})
		}else {
			//生成一个token
			token,err2 := jwt.GenToken(user.UserID)
			if err2 == nil {
				c.JSON(http.StatusOK,gin.H{
					"token" :token,
					"user" : gin.H{
						"user_id" : user.UserID,
						"username" : user.UserName,
					},
				})
			}
		}
	}else {
		c.JSON(http.StatusUnauthorized,gin.H{
			"message" : "请输入正确",
		})
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
		database.Db.Where("user_id=?",user.UserID).Find(&models.Users)
		//遍历切片中的结构体，如果中有用户名相同的话，那么就返回一个message该用户ID已经被注册
		for _,user1 := range models.Users{
			if user.UserID == user1.UserID {
				c.JSON(http.StatusBadRequest,gin.H{
					"message" : "该用户ID已经被注册",
				})
				return
			}
		}
		database.Db.Model(&models.User{}).Where("user_id=?",user.UserID).Create(&user)
	} else{
		c.JSON(http.StatusOK,gin.H{
			"message" : err.Error(),
		})
	}
}


//DeleteUser 定义一个删除账户的api
func DeleteUser(c *gin.Context)  {
	var user models.User
	userid,_ := c.Get("user_id")
	//根据主键来进行一个删除
	database.Db.Model(&models.User{}).Where("user_id=?",userid).First(&user)
	database.Db.Delete(&models.User{},user.Model.ID)
}

//UpdateUser 定义一个更新用户用户名或者用户密码的一个api
//记得对密码进行一个加密
//对于这个的调试，因为如果我更改了用户的user_id那么就需要重新给一个token
func UpdateUser(c *gin.Context)  {
	var user models.User
	userid,_ := c.Get("user_id")
	if err := c.ShouldBind(&user);err == nil{
		if user.Password != ""{
			hash, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err1 != nil {
				log.Panic(err1.Error())
			}
			encodePW := string(hash)
			//把加密过后的东西赋值给password然后存入数据库中
			user.Password = encodePW
			database.Db.Model(&models.User{}).Where("user_id",userid).Update("password",user.Password)
		}
		database.Db.Model(&models.User{}).Where("user_id",userid).Updates(user)
		c.JSON(http.StatusOK,gin.H{
			"userid" : user.UserID,
			"username" : user.UserName,
		})
	}else {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
	}
}

