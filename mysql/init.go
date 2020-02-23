package mysql

import (
	"golibrary/config"
	"golibrary/mysql/db"
)

func init(){
	c,_ := config.LoadDefaultConfig()
	print(c)
	//todo:读取配置后转化为mysql配置对象
	myConfig := &db.Config{}
	print(myConfig)
	//todo:通过mysqlConfig 初始化mysql客户端
}