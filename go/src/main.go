package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", greeting)               // 挨拶
	r.GET("/users", getUsers)          // user の一覧を表示
	r.GET("/users/:id", getUserByID)   //指定した id の user を表示
	r.POST("/users", addNewUser)       //user を追加
	r.PUT("/users/:id", updateUser)    //指定した id の user を更新
	r.DELETE("/users/:id", deleteUser) //指定した id の user を削除

	r.Run(":" + os.Getenv("PORT"))
}
