package rpackage
import (
	"github.com/gin-gonic/gin"
	"lwh.com/apipackage"
)
func Newrouter()  *gin.Engine{
	r := gin.Default()
	//一个登录的中间件
	r.Use(apipackage.CORS())
	//一个登录的api
	r.POST("/login",apipackage.Login)
	//一个注册的apio
	r.POST("/register",apipackage.Register)
	//一个
	return r
}
