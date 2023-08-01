package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello, gin")
	})
	router.Run("192.168.1.33:9999")
}
