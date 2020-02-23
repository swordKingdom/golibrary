package api

import "time"

type MgoConf struct {
	//Username 用户名称
	Username string
	//Password 数据库密码
	Password string
	//Addrs mongodb主机地址
	Addrs []string
	//Timeout 超时设置
	Timeout time.Duration
	//Database 数据库名称
	Database string
	//MaxIdleTimeMS 最大存活毫秒数
	MaxIdleTimeMS int
	//PoolLimit 链接池大小设置
	PoolLimit int
}

func (c *MgoConf) Cype() *MgoConf {
	cpyeConf := new(MgoConf)
	return cpyeConf
}
