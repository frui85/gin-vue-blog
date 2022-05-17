package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	//声明 server 段配置变量
	AppModel string
	HttpPort string

	//声明 dabase 段配置变量
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPrefix   string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径！", err)
	}
	LoadServer(cfg)
	LoadDatabase(cfg)
}

func LoadServer(f *ini.File) {
	AppModel = f.Section("server").Key("app_model").MustString("debug")
	HttpPort = f.Section("server").Key("http_port").MustString("8080")
}

func LoadDatabase(f *ini.File) {
	Db = f.Section("database").Key("db").MustString("mysql")
	DbHost = f.Section("database").Key("db_host").MustString("localhost")
	DbPort = f.Section("database").Key("db_port").MustString("3306")
	DbUser = f.Section("database").Key("db_user").MustString("root")
	DbPassword = f.Section("database").Key("db_password").MustString("root")
	DbName = f.Section("database").Key("db_name").MustString("gin-vue-blog")
	DbPrefix = f.Section("database").Key("db_prefix").String()
}
