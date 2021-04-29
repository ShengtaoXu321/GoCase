package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSON渲染
func main() {
	r := gin.Default()

	// gin.H是map[string]interface{}的缩写
	// 方式一：自己拼接JSON
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello Word!",
		})
	})
	// 方式二：使用结构体
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello World"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run("0.0.0.0:1025")
}
