package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func pingHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong test",
	})
}

func main() {

	// Default 返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// 输出json的结果给调用方
		// 第一个参数状态码，第二个是一个 map
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", pingHandle)

	// 参数传递
	// 1.querystring ，比如 /user/search?username=少林&address=北京
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "少林")
		//username := c.Query("username")
		address := c.Query("address")

		c.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})

	// 2.通过路径来传递， 比如 /user/search/少林/北京，这个属于精确匹配
	r.GET("/user/info/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")

		c.JSON(200, gin.H{
			"message":  "info",
			"username": username,
			"address":  address,
		})

	})

	// 3.通过表单提交 ，比如 POST /user/search
	r.POST("/user/search", func(c *gin.Context) {
		// username := c.DefaultPostForm("username","少林")
		username := c.PostForm("username")
		address := c.PostForm("address")

		c.JSON(200, gin.H{
			"message":  "info",
			"username": username,
			"address":  address,
		})
	})

	// 单文件上传
	r.POST("/upload", func(c *gin.Context) {
		// single file
		// "file" 是文件的名字
		// file 会拿到文件对象
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		log.Println(file.Filename)
		// 指定保存的路径和文件名
		//dst := fmt.Sprintf("/Users/lifeifan/go/src/gin_blog/%s", file.Filename)
		dst := fmt.Sprintf("/Users/lifeifan/go/src/gin_blog/test.doc")
		// Upload the file to specific dst
		// 保存文件
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	// 多文件上传
	r.POST("/uploads", func(c *gin.Context) {
		// 默认文件传里的内存大小是 32 MB
		// 可以通过设置 r.MaxMultipartMemory 来设置这个值
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/Users/lifeifan/go/src/gin_blog/%d_%s", index, file.Filename)
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
