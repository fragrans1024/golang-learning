package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("prefix", prefix)
	}

	router_v1 := router.Group("v1")
	router_v1.Use(gin.Logger(), gin.Recovery())
	router_v1.Use(middleware_v1)
	router_v1.GET("/ping", ping_v1)
	router_v1.POST("/info", info_v1)

	router_v2 := router.Group("v2")
	router_v2.GET("/ping", ping_v2)

	return router
}

func middleware_v1(c *gin.Context) {
	log.Println("begin middle ware v1")
	c.Next()
	log.Println("end middle ware v1")
}

func ping_v1(c *gin.Context) {
	log.Println("process ping v1")
	resp := gin.H{
		"message": "pong",
		"version": "v1",
	}
	c.JSON(200, resp)
}

type Info struct {
	Name string `json:"name" binding:"required,prefix"`
	Age  uint32 `json:"age" binding:"required,min=18"`
}

// curl -X POST  http://127.0.0.1:9001/v1/info -H "Content-Type: application/json" -d '{"name", "abc", "age": 13}'
func info_v1(c *gin.Context) {
	var info Info
	if err := c.ShouldBindBodyWithJSON(&info); err != nil {
		log.Printf("err = %v\n", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%+v", info)

	c.JSON(204, nil)
}

var prefix validator.Func = func(field validator.FieldLevel) bool {
	str, ok := field.Field().Interface().(string)
	if !ok {
		return false
	}
	if strings.HasPrefix(str, "abc") {
		return true
	} else {
		log.Printf("%s is not prefix with abc", str)
		return false
	}
}

func ping_v2(c *gin.Context) {
	resp := gin.H{
		"message": "pong",
		"version": "v2",
	}
	c.JSON(200, resp)
}
