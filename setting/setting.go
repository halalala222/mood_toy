package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//总的一个配置的结构体
type config struct {
	Mysql Mysql
	Jwt Jwt
	XXG XXG
}
//Configone 声明一个总结构体的变量
var Configone = config{}

//Mysql 一个mysql对应的结构体
type Mysql struct {
	Username  string `yaml:"username"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Password string `yaml:"password"`
}

//Jwt 一个对应配置文件的结构体
type Jwt struct {
	Secret string `yaml:"mysecret"`
}

type XXG struct {
	AccessKeyID string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}
//InitSetting 对配置文件进行读写的一个函数
func InitSetting(con *config)  {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile,con)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}


