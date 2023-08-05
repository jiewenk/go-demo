package main

import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context) {
	ctx.HTML(200, "index/index.html", "这是Index界面")
}

func User(ctx *gin.Context) {
	ctx.HTML(200, "user/user.html", gin.H{
		"content": "这是User界面,Content",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/**/*")
	// 静态资源
	router.Static("/static", "static")
	router.GET("/", Index)
	router.GET("/user", User)
	router.Run(":9999")
}
