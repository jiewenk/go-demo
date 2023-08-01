package main

import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", Index)
	router.Run(":9999")
}
