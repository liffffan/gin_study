package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}

func submit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit success",
	})
}

func main() {
	// 路由分组
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	router.Run()
}
