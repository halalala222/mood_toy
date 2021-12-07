package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	Mysql Mysql
	Jwt Jwt
}
var Configone = config{}

type Mysql struct {
	Username  string `yaml:"username"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Password string `yaml:"password"`
}
type Jwt struct {
	Secret string `yaml:"mysecret"`
}
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


