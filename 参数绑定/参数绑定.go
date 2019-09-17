package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// 如果是表单提交就是小写的 user ，如果是json的话也是小写的 user ，binding 要求参数是必须的
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	// 为什么要参数绑定，本质上是方便，提高开发效率
	// A. 通过反射机制，自动提取 querystring 、form 表单 、 json 、xml 等参数到 struct 中
	// B. 通过 http 协议中的 context type ， 识别是 json 、xml 或者表单

	router := gin.Default()
	// example for binding JSON ({"user":"manu","password":"123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err == nil {
			fmt.Printf("login info:%#v", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	})

	// example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// this will infer what binder to use depending on the content-type header.
		// 如果 content-type 是表单，它就会从表单把数据读取出来
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	// example for binding a HTML querystring (username=manu&password=123)
	router.GET("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	router.Run()
}
