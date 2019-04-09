package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", greeting)      // 挨拶
	r.GET("/users", getUsers) // user の一覧を表示

	fmt.Println("OS.ENV:")
	fmt.Println(os.Environ())
	r.Run(":" + os.Getenv("PORT"))
}
