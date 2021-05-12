package main

import (
	"goframework/api"
	"goframework/lib"
	"goframework/util"
)

func main() {
	//初始化配置
	lib.LoadConfig(".")

	//初始化数据库
	lib.InitDB()

	//验证中文
	util.ValidateInit()


	x := api.NewServer(lib.DB)
	x.Run(lib.Config.ServerAddress)
}
