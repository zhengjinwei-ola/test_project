package main

import (
	"test_project/app/server"
	"test_project/util/config"
)

func init() {
	// 初始化配置
	config.InitConfig()
	// 初始化数据库
	//mysql.InitDB()
}

func main() {
	server.Run()
}
