package apipackage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"lwh.com/database"
	"lwh.com/jwt"
	"lwh.com/models"
	"net/http"
	"path"
	"strings"
)

//Login 定义一个登录的api判断用户输入是否正确，然后返回给用户一个token
func Login(c *gin.Context)  {
	var user = models.User{}
	var user1 = models.User{}
	if err := c.ShouldBind(&user);err == nil {
		//从数据库中查找信息
		database.Db.Where("user_id = ?",user.UserID).First(&user1)
		//对于存在数据库中的密码进行解码
		err1 := bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password))
		if err1 != nil {
			c.JSON(http.StatusOK,gin.H{
				"message" : "用户id或者密码错误",
			})
		}else {
			//生成一个token
			token,err2 := jwt.GenToken(user.UserID)
			if err2 == nil {
				c.JSON(http.StatusOK,gin.H{
					"token" :token,
					"user" : gin.H{
						"userid" : user.UserID,
						"username" : user.UserName,
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
		database.Db.Select("user_id").Find(&models.Users)
		//遍历切片中的结构体，如果中有用户名相同的话，那么就返回一个message该用户ID已经被注册
		for _,user1 := range models.Users{
			if user.UserID == user1.UserID {
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
		"hairs" : hairs,
		"mouthe" : mouthes,
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
			"hairs" :moodtoy.Hair,
			"eyebrow" : moodtoy.Eyebrow,
			"mouth" : moodtoy.Mouth,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
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
		"hairs" :moodtoy.Hair,
		"eyebrow" : moodtoy.Eyebrow,
		"mouth" : moodtoy.Mouth,
	})
}

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
			"time" : diary.Time,
			"text_content" : diary.TextContext,
			"feeling" : diary.Feeling,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"message" : "传入参数可以有问题",
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
		c.JSON(http.StatusOK,gin.H{
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
		c.JSON(http.StatusOK,gin.H{
			"message" : "传入的数据可能有误",
		})
	}
}

//PostPicture 定义一个上传单个文件的api
func PostPicture(c *gin.Context)  {
	var user models.User
	uerid,_ := c.Get("user_id")
	//获取FormFile中的key为file的文件
	file,err := c.FormFile("file")
	user.UserImg.ImgUrl = file.Filename//因为我把文件都存在一个固定的文件夹里面，那么查找
	//就仅仅只需要查找存在文件夹中的文件名字就可以了，
	database.Db.Model(&models.User{}).Where("user_id=?",uerid).Updates(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	//用来验证检查图片的后缀名
	ext := strings.ToLower(path.Ext(file.Filename))
	if ext != ".jpg" && ext != ".png" {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : "只支持jpg/png图片上传",
		})
		//defer os.Exit(2)
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

