package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// XML渲染
func main() {
	r := gin.Default()
	// gin.H是map[string]interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{
			"msg": "Hello world!",
		})
	})

	r.GET("/moreXML", func(c *gin.Context) {
		// 方式二：使用结构体
		// 定义结构体
		type MseeageRecord struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		// 定义变量
		var msg MseeageRecord
		msg.Name = "小王子"
		msg.Message = "hello world"
		msg.Age = 18

		c.XML(http.StatusOK, msg)
	})
}
