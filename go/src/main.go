package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", greeting)
	r.Run(":" + os.Getenv("PORT"))
}

func greeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!",
	})
}
