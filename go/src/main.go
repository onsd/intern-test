package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", greeting)
	r.GET("/users", getUsers) // user の一覧を表示

	r.Run(":" + os.Getenv("PORT"))
}
