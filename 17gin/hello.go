package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"blog":   "www.blog.com",
			"wechat": "www.wechat.com",
		})
	})
	r.Run(":8080")
}
