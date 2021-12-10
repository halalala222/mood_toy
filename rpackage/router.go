package rpackage

import (
	"github.com/gin-gonic/gin"
	"lwh.com/apipackage"
	"lwh.com/jwt"
)
func Newrouter()  *gin.Engine{
	r := gin.Default()
	//路由分组分成一个是mood的
	rmood := r.Group("/moodtoy")
	rmood.Use(jwt.JWTAuthMiddleware())
	//路由分组分成一个是List的
	rdiary := r.Group("/diary")
	rdiary.Use(jwt.JWTAuthMiddleware())
	//一个登录的中间件
	r.Use(apipackage.CORS())
	//一个登录的api
	r.POST("/login",apipackage.Login)
	//一个注册的apio
	r.POST("/register",apipackage.Register)
	//一个删除用户的api
	r.DELETE("/user",jwt.JWTAuthMiddleware(),apipackage.DeleteUser)
	//一个更新用户的api
	r.PUT("/user",jwt.JWTAuthMiddleware(),apipackage.UpdateUser)
	//一个用户上传用户图片的api
	r.POST("/picture",jwt.JWTAuthMiddleware(),apipackage.PostPicture)
	//一个mood娃娃的接口集合
	rmood.POST("/",apipackage.PostMoodToy)
	rmood.GET("/all",apipackage.GetAllmoodtoy)
	rmood.GET("/",apipackage.GetMoodtoy)
	//rmood.PUT("/")
	//一个list的接口集合
	rdiary.POST("/",apipackage.PostDiary)
	rdiary.GET("/",apipackage.GetDiary)
	//rdiary.PUT("/")
	r.GET("/robot",apipackage.GetBot)
	return r
}
