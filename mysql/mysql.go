package mysql

import "golibrary/mysql/db"

var defaultMysqlClient *db.Client

//GetMysqlClient 获取默认客户端
func GetMysqlClient()*db.Client{
	return defaultMysqlClient
}