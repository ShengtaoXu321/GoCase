package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type HotPoint1 struct {
	Title string
	Url   string
}

func main() {
	// 1. 数据库连接
	config := "root:123456@tcp(117.78.34.82:18100)/Hots?charset=utf8mb4"
	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Println("打开数据库错误", err)
	}
	defer db.Close()
	// 确保连接是有效的
	err = db.Ping()
	if err != nil {
		log.Println("数据库不能正常连接", err)
	}

	// 2. 创建api
	router := gin.Default() // 使用Default()方法来获取一个基本的路由变量
	// API处理程序 -- 获取用户详细信息
	router.GET("/hotpoint/:title", func(c *gin.Context) { // 使用匿名函数作为路由的处理函数，处理函数必须是func(*gin.Context)类型的函数
		var (
			hp1    HotPoint1
			result gin.H // gin.H() 方法简化json的生成，本质就是一个map[string]interface{}
		)
		title := c.Param("title")
		fmt.Println("输入标题:", title)
		sen1 := "SELECT title, url FROM hotpoint WHERE title=?"
		row := db.QueryRow(sen1, title)
		err = row.Scan(&hp1.Title, &hp1.Url)
		fmt.Printf("用户: %+v\n", hp1)
		if err != nil {
			result = gin.H{
				"hp":    nil,
				"count": 0,
			}
			log.Println(err)
		} else {
			result = gin.H{
				"title": hp1.Title,
				"url":   hp1.Url,
				"count": 1,
			}
		}
		c.JSON(http.StatusOK, result)

	})
	router.Run(":18520")
}
