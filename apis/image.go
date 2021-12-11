package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"lwh.com/database"
	"lwh.com/models"
	"net/http"
	"path"
	"strings"
)

//PostPicture 定义一个上传单个文件(用户的头像)的api
func PostPicture(c *gin.Context)  {
	var user models.User
	var userimg models.UserImg
	uerid,_ := c.Get("user_id")
	//获取FormFile中的key为file的文件
	file,err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "上传图片错误",
		})
		return
	}
	userimg.ImgUrl = file.Filename//因为我把文件都存在一个固定的文件夹里面，那么查找
	//就仅仅只需要查找存在文件夹中的文件名字就可以了，
	database.Db.Model(&models.User{}).Where("user_id=?",uerid).First(&user)
	userimg.UserRefer2 = user.Model.ID
	database.Db.Model(&models.UserImg{}).Create(&userimg)
	//用来验证检查图片的后缀名
	ext := strings.ToLower(path.Ext(file.Filename))
	if ext != ".jpg" && ext != ".png" {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "只支持jpg/png图片上传",
		})
	}
	//设置存储的路径
	dst := fmt.Sprintf("D:/picture/%s", file.Filename)
	//保存文件到指定的目录
	err = c.SaveUploadedFile(file,dst)
	if err != nil{
		log.Panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

//GetPicture 定义一个获取用户头像的api
func GetPicture(c *gin.Context)  {
	useid,_ := c.Get("user_id")
	var user models.User
	var userimage = models.UserImg{}
	database.Db.Model(&models.User{}).Where("user_id=?",useid).First(&user)
	err :=database.Db.Model(&user).Association("UserImg").Find(&userimage)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"message" : err,
		})
	}
	dst := fmt.Sprintf("D:/picture/%s",userimage.ImgUrl)
	c.File(dst)
}
