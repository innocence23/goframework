package main

import (
	"goframework/api"
	"goframework/lib"
)

func main() {
	//初始化配置
	lib.LoadConfig(".")

	//初始化数据库
	lib.InitDB()

	x := api.NewServer(lib.DB)
	x.Run(lib.Config.ServerAddress)
}
