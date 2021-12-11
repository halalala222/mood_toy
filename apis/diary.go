package apis

import (
	"github.com/gin-gonic/gin"
	"lwh.com/database"
	"lwh.com/models"
	"net/http"
)

//PostDiary 定义给一个上传diary的api
func PostDiary(c *gin.Context)  {
	var user models.User
	var diary models.Diary
	userid,_ := c.Get("user_id")
	//进行一个参数绑定
	if err := c.ShouldBind(&diary);err == nil{
		//对diary中的userid进行一个赋值
		database.Db.Model(&models.User{}).Where("user_id=?",userid).First(&user)
		diary.UserRefer = user.Model.ID
		//通过用户传来的数据，进行一个轻日记的判断，如果是good就+1，反之则-1
		if diary.Feeling == "good"{
			user.Feeling += 1
		}else {
			user.Feeling -= 1
		}
		//进行一个对于用户的更新
		database.Db.Model(&models.User{}).Where("user_id",userid).Updates(&user)
		database.Db.Model(&models.Diary{}).Where(diary).FirstOrCreate(&diary)
		c.JSON(http.StatusOK,gin.H{
			"title" : diary.Title,
			"time" : diary.Time,
			"text_content" : diary.TextContext,
			"feeling" : diary.Feeling,
		})
	}else {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "传入参数可能有问题",
		})
	}
}

func GetDiary(c *gin.Context)  {
	var user models.User
	var txt []string
	var time []string
	userid,_ := c.Get("user_id")
	database.Db.Model(&models.User{}).Where("user_id=?",userid).First(&user)
	err := database.Db.Model(&user).Association("Diaries").Find(&models.Diaries)
	if err != nil {
		c.JSON(http.StatusHTTPVersionNotSupported,gin.H{
			"message" : err,
		})
	}
	for _,value := range models.Diaries{
		txt = append(txt,value.TextContext)
		time = append(time,value.Time)
	}
	c.JSON(http.StatusOK,gin.H{
		//返回一个文本内容
		"time" : time,
		"textcontent" : txt,
	})
}


