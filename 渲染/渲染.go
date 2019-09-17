package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 中间件函数
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 可以设置一些公共参数
		c.Set("example", "12345")
		// 等其他中间件先执行
		c.Next()
		// 获取耗时
		latency := time.Since(t)
		log.Printf("total cost time:%d us", latency/1000)
	}
}

func main() {
	// JSON 渲染
	// 第一种方式，自己拼json ，传一个 map 进去序列化成 json
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// 第二种方式，使用结构体
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	// xml 渲染
	r.GET("/moreXML", func(c *gin.Context) {
		type MessageRecord struct {
			Name    string
			Message string
			Number  int
		}

		var msg MessageRecord
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		c.XML(http.StatusOK, msg)
	})

	// 渲染模版
	// 两种方式，一种是模版本身就是一个文件

	// 加载模版到内存中
	// 这是二级目录的，如果是三级目录就是 templates/**/**/*
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	// 静态文件的处理
	r.Static("/static", "./static")

	// 中间件
	// Gin框架允许在请求处理的过程中，加入用户自己的钩子函数，这个钩子函数就叫做中间件
	// 因此，可以使用中间件处理一些公共业务逻辑，比如耗时统计、日志打印、登陆校验
	r.Use(StatCost())
	r.GET("/test", func(c *gin.Context) {
		// 获取中间件中的公共参数，打印出来
		example := c.MustGet("example").(string)
		// it would print: "12345"
		log.Println(example)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.Run()
}
