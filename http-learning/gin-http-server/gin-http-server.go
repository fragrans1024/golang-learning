package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin_default()

	r.Run("127.0.0.1:9001")
}

func gin_default() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return router
}
