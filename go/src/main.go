package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", greeting)      // æ¨æ¶
	r.GET("/users", getUsers) // user ã®ä¸è¦§ãè¡¨ç¤º

	r.Run(":" + os.Getenv("PORT"))
}
