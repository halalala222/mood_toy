package main

import (
	"fmt"
	"lwh.com/database"
	"lwh.com/rpackage"
	"lwh.com/setting"
)

func main()  {
	setting.InitSetting(&setting.Configone)
	database.Link()
	r :=rpackage.Newrouter()
	r.Run(":1225")
	fmt.Println(setting.Configone.Jwt.Secret)
}