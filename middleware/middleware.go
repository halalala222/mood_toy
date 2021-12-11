package middleware

import (
	"github.com/gin-gonic/gin"
	"lwh.com/jwt"
	"net/http"
	"strings"
)

//JWTAuthMiddleware 定义一个中间件用来判断前端传来的token格式是否正确以及前端的token是否是对
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//这个是获取请求头中的内容
		authHandler := c.Request.Header.Get("Authorization")
		if authHandler == "" {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "请求头中auth为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHandler," ",2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		//对请求头中的tokensrting 进行解码
		mc,err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"message" : err,

			})
			c.Abort()
			return
		}
		c.Set("user_id",mc.UserID)
		c.Next()
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
