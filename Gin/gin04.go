package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	go GetQuerryString()
	go GetForm()
	go GetJson()
	go GetPath()
	select {}
}

// 获取参数
// 1. 获取querystring参数  -- Get方式
// querystring 指的是URL中 ？ 后面携带的参数。 例如： /user/search?username=小王子&address=沙河.
// 获取querystring参数的方法如下:
func GetQuerryString() {
	r := gin.Default()
	r.GET("/querrystring/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		address := c.Query("address")
		// 输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":1026")
}

// 2. 获取form参数  -- POST方式
// 当前端请求的数据通过form表单提交时，例如向/user/search发送一个POST请求，获取的参数如下

func GetForm() {
	// Default返回一个默认的路由引擎
	r := gin.Default()
	r.POST("/form/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		username := c.DefaultPostForm("username", "小王子")
		address := c.PostForm("address")
		// 输出JSON结果给对方调用
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
		r.Run(":1027")
	})
}

// 3. 获取json参数
// 当前端请求的数据通过json提交时，例如向/json发送一个POST请求，则获取请求参数的方式如下
func GetJson() {
	r := gin.Default()
	r.POST("/json", func(c *gin.Context) {
		// 注意：这里忽略了错误处理
		res, err := c.GetRawData() // 从c.request.body读取请求数据
		if err != nil {
			log.Println("数据接收失败", err)
		}
		// 定义map或者结构体
		var m map[string]interface{}
		// 反序列化
		err1 := json.Unmarshal(res, &m)
		if err1 != nil {
			log.Println("接收数据反序列化失败", err1)
		}
		c.JSON(http.StatusOK, m)
	})
	r.Run(":1028")
}

// 4. 获取path
// 请求的参数通过URL路径传递，例如: /user/search/小王子/沙河
func GetPath() {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		// 输出的json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":1029")
}
