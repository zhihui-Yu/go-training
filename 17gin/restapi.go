package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	// users := []User{{123, "zhangsan"}, {456, "wangwu"}}
	r := gin.Default()
	// r.GET("/users", func(c *gin.Context) {
	// 	c.JSON(200, users)
	// })

	r.POST("/users", func(c *gin.Context) {})

	r.DELETE("/users", func(c *gin.Context) {})

	r.PUT("/users", func(c *gin.Context) {})

	r.PATCH("/users", func(c *gin.Context) {})

	// 无重定向，url是确定的
	// r.GET("/users/:id", // 使用这种后，/users/xxx 都不能注册了
	// 	func(c *gin.Context) {
	// 		id := c.Param("id")
	// 		c.String(200, "this user id is %s", id)
	// 	})

	r.RedirectTrailingSlash = false // 是否自动重定向

	// 如果/users 没有被注册，会重定向到/users/*id, id = "/"
	r.GET("/users/*id", //会取/users 后面的所有东西当成id，如/users/id/id, id ="/id/id"
		func(c *gin.Context) {
			id := c.Param("id")
			c.String(200, "this user id is %s", id)
		})

	r.GET("/",
		func(c *gin.Context) {
			c.String(200, c.Query("wechat"))
		})

	r.GET("/default",
		func(c *gin.Context) {
			c.String(200, c.DefaultQuery("wechat", "defalu-val"))
		})

	r.Run(":8080")
}
