package mysql

import (
	"golibrary/config"
	"golibrary/mysql/db"
	"sync"
)

var once sync.Once

func initMysql(){
	c := config.GlobalConfig
	print(c)
	//todo:读取配置后转化为mysql配置对象
	myConfig := &db.Config{}
	print(myConfig)
	//todo:通过mysqlConfig 初始化mysql客户端
}

func Init(){
	once.Do(initMysql)
}