package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 解决跨域（开发环境用）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// 示例：视频流接口
	r.GET("/api/video", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"url": "https://vjs.zencdn.net/v/oceans.mp4", //实际视频URL
		})
	})

	// 示例：获取直播间信息
	r.GET("/api/live-info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "测试直播间",
			"views": 1024,
		})
	})

	r.Run(":8080") // 启动服务
}
