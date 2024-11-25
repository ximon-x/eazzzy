package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	ping := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}

	r.GET("/ping", ping)

	// TODO: Get port from env
	os.Setenv("PORT", "3001")
	r.Run()
}
