package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin_group()

	r.Run("127.0.0.1:9001")
}

func gin_default() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return router
}

// 分组路由
func gin_group() *gin.Engine {
	router := gin.New()

	router_v1 := router.Group("v1")
	router_v1.GET("/ping", ping_v1)

	router_v2 := router.Group("v2")
	router_v2.GET("/ping", ping_v2)

	return router
}

func ping_v1(c *gin.Context) {
	resp := gin.H{
		"message": "pong",
		"version": "v1",
	}
	c.JSON(200, resp)
}

func ping_v2(c *gin.Context) {
	resp := gin.H{
		"message": "pong",
		"version": "v2",
	}
	c.JSON(200, resp)
}
