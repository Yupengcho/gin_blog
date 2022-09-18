package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	Dbuser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err !=nil{
		fmt.Println("读取错误请检查")
	}
	LoadServer(file)
	LoadData(file)
}



func LoadServer(file *ini.File){
	AppMode=file.Section("sever").Key("AppMode").MustString("debug")
	HttpPort=file.Section("sever").Key("HttpPort").MustString("3000")
}

func LoadData(file *ini.File){
	Db         =file.Section("database").Key("Db").MustString("debug")
	DbHost     =file.Section("database").Key("DbHost").MustString("localhost")
	DbPort 	   =file.Section("database").Key("DbPort").MustString("3306")
	Dbuser 	   =file.Section("database").Key("Dbuser").MustString("ginblog")
	DbPassWord =file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName     =file.Section("database").Key("DbName").MustString("ginblog")
}