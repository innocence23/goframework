package main

import (
	"fmt"
	"goframework/database"
	"goframework/model"

	"github.com/gin-gonic/gin"
)

func main() {

	//初始化数据库
	database.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		user := model.User{}
		fmt.Println(database.DB.First(&user))
		fmt.Println(user)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
