package main

import (
	"lwh.com/database"
	"lwh.com/rpackage"
	"lwh.com/setting"
)

func main()  {
	//配置文件的读取
	setting.InitSetting(&setting.Configone)
	//链接数据库
	database.Link()
	//新建路由
	r :=rpackage.Newrouter()
	r.Run(":1225")
}