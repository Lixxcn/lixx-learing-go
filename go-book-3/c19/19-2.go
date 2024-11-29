// 代码清单19-2 自定义监听端口
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		ctx.String(200, "gin的第一个请求")
	})

	r.Run(":9080")
}
