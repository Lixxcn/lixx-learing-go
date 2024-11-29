// 代码清单19-1 第一个Gin框架示例程序
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		ctx.String(200, "gin的第一个请求")
	})

	r.Run()
}
