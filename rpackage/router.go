package rpackage

import (
	"github.com/gin-gonic/gin"
	"lwh.com/apis"
	"lwh.com/middleware"
)

func Newrouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	baseRoute := r.Group("/api")
	ruser := baseRoute.Group("/")
	{
		ruser.POST("/login", apis.Login)
		ruser.POST("/register", apis.Register)
		ruser.DELETE("/user", middleware.JWTAuthMiddleware(), apis.DeleteUser)
		ruser.PUT("/user", middleware.JWTAuthMiddleware(), apis.UpdateUser)
	}
	rmood := baseRoute.Group("/moodtoy")
	{
		rmood.POST("", apis.PostMoodToy)
		rmood.GET("/all", apis.GetAllmoodtoy)
		rmood.GET("", apis.GetMoodtoy)
		//这个就是访问这个/static就是可以获取./moodtoy中的文件
		rmood.Static("/static","./moodtoy")
			
	}
	rdiary := baseRoute.Group("/diary",middleware.JWTAuthMiddleware())
	{
		rdiary.POST("", apis.PostDiary)
		rdiary.GET("", apis.GetDiary)
	}
	rImage := baseRoute.Group("/picture",middleware.JWTAuthMiddleware())
	{
		rImage.POST("", apis.PostPicture)
		rImage.GET("", apis.GetPicture)

	}
	rrobort := baseRoute.Group("/robot")
	rrobort.POST("",apis.GetXXG)
	return r
}
