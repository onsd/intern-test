package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", greeting)             // 挨拶
	r.GET("/users", getUsers)        // user の一覧を表示
	r.GET("/users/:id", getUserByID) //指定した id の user を表示
	r.POST("/users", addNewUser)     //user を追加

	r.Run(":" + os.Getenv("PORT"))
}
