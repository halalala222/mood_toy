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
		mc,err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"message" : err,

			})
			c.Abort()
			return
		}
		c.Set("user_id",mc.UserID)
		c.Next()
	}
}
