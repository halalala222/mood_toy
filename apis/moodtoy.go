package apis

import (
	"github.com/gin-gonic/gin"
	"lwh.com/database"
	"lwh.com/models"
	"net/http"
	"strings"
)

//GetAllmoodtoy 定义一个获取所有娃娃部件的api
func GetAllmoodtoy(c *gin.Context)  {
	database.Db.Select("init_url").Find(&models.Init)
	database.Db.Select("clothes_url").Find(&models.Clothes)
	database.Db.Select("eyes_url").Find(&models.Eyes)
	database.Db.Select("hair_url").Find(&models.Hairs)
	database.Db.Select("eyebrow_url").Find(&models.Eyebrows)
	database.Db.Select("mouth_url").Find(&models.Mouthes)
	var init []string
	var clohes []string
	var eyes []string
	var hairs []string
	var eyebrows []string
	var mouthes []string
	for _,value := range models.Init{
		init = append(init, value.InitUrl)
	}
	for _,value := range models.Clothes{
		clohes = append(clohes, value.ClothesUrl)
	}
	for _,value := range models.Eyes{
		eyes = append(eyes, value.EyesUrl)
	}
	for _,value := range models.Eyebrows{
		eyebrows = append(eyebrows,value.EyebrowUrl)
	}
	for _,value := range models.Hairs{
		hairs = append(hairs, value.HairUrl)
	}
	for _,value := range models.Mouthes{
		mouthes = append(mouthes, value.MouthUrl)
	}
	c.JSON(http.StatusOK,gin.H{
		"init" : init,
		"clothes" : clohes,
		"eyes" : eyes,
		"eyebrow" : eyebrows,
		"hair" : hairs,
		"mouth" : mouthes,
	})
}
//PostMoodToy 对于这个的实现moodtoy部件的上传
func PostMoodToy(c *gin.Context)  {
	var moodtoy models.Moodtoy
	var user = models.User{}
	var diary = models.Diary{}
	var color = models.Color{}
	var teyebrow = models.ToyEyebrow{}
	//从请求头中的上下文中获取user_id，通过token的获取的
	userid,_ :=c.Get("user_id")
	//对于前端传来的进行一个存储
	if err :=c.ShouldBind(&moodtoy);err == nil{
		//对传来的地址进行一个切片，然后有一个中括号，通过中括号分隔,在进行切获得那个颜色的中文
		eyebrow := strings.SplitN(moodtoy.Eyebrow,"（",2)
		eyebrowcolor := strings.SplitN(eyebrow[1],".",2)
		//进行查找如果没有就新建一个
		database.Db.Model(&moodtoy).Where(moodtoy).FirstOrCreate(&moodtoy)
		//update对用户的mood_id进行一个赋值
		database.Db.Model(&models.User{}).Where("user_id=?",userid).Update("moodtoy_id",moodtoy.ID)
		database.Db.Model(&models.User{}).Where("user_id=?",userid.(string)).First(&user)
		//把mood娃娃的id赋值给user的moodtoyid
		database.Db.Model(&models.Moodtoy{}).Where("id=?",user.MoodtoyID).Updates(&moodtoy)
		diary.UserRefer = user.Model.ID
		//通过颜色查找对应的颜色数字
		database.Db.Model(&models.Color{}).Where("chinese = ?",eyebrowcolor[0]).First(&color)
		//通过颜色的number然后还有路径查找eyebrow中对应的feelingstatus
		database.Db.Where("number=? AND eyebrow_url=?",color.Number,moodtoy.Eyebrow).First(&teyebrow)
		//对用户初始的feeling值为100
		if user.Feeling == 100 {
			user.Feeling = teyebrow.FeelingStatus * 5
			database.Db.Model(&models.User{}).Where("user_id=?",userid).Update("feeling",user.Feeling)
		}
		c.JSON(http.StatusOK,gin.H{
			"init" : "/初始.png",
			"clohes" : moodtoy.Clothes,
			"eyes" : moodtoy.Eyes,
			"hair" :moodtoy.Hair,
			"eyebrow" : moodtoy.Eyebrow,
			"mouth" : moodtoy.Mouth,
		})
	}else {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "传入数据错误",
		})
	}
}

//GetMoodtoy 定义一个用户获取娃娃部件的api
func GetMoodtoy(c *gin.Context)  {
	//解析token的user_id
	var user models.User
	var moodtoy models.Moodtoy
	var eyebrow  models.ToyEyebrow
	var  feelingstatus uint
	userid,_ := c.Get("user_id")
	//通过user_id查找user
	database.Db.Model(&models.User{}).Where("user_id=?",userid).Find(&user)
	//获取的feeling值，然后求出feelingsataus
	feelingstatus = user.Feeling/5
	//获取用户对应一开始对应的moodtoy娃娃
	database.Db.Where("id=?",user.MoodtoyID).Find(&moodtoy)
	//获取moodtoy娃娃对应的eyebrows
	moodtoy.ID = 0
	database.Db.Model(&models.ToyEyebrow{}).Where("eyebrow_url=?",moodtoy.Eyebrow).First(&eyebrow)
	//通过颜色还有用户的feelingstatus获取新的eyebrows
	database.Db.Model(&models.ToyEyebrow{}).Where("number=? and feeling_status=?",eyebrow.Number,feelingstatus).First(&eyebrow)
	moodtoy.Eyebrow = eyebrow.EyebrowUrl
	//判断用户更新的moodtoy模型是否有，没有就新建一个有话就获取
	database.Db.Model(&models.Moodtoy{}).Where(moodtoy).FirstOrCreate(&moodtoy)
	//更新用户的moodtoyid
	database.Db.Model(&models.User{}).Where("user_id=?",userid).Update("moodtoy_id",moodtoy.ID)
	c.JSON(http.StatusOK,gin.H{
		"init" : "/初始.png",
		"clohes" : moodtoy.Clothes,
		"eyes" : moodtoy.Eyes,
		"hair" :moodtoy.Hair,
		"eyebrow" : moodtoy.Eyebrow,
		"mouth" : moodtoy.Mouth,
	})
}
