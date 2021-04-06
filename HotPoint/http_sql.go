// 这是创建api进行访问数据库的demo -- 参考网上

package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
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

	// 定义数据库的内容结构体
	type User struct {
		age  int
		name string
	}

	// 定义一个路由
	router := gin.Default() // 使用Default()方法来获取一个基本的路由变量
	// API处理程序 -- 获取用户详细信息
	router.GET("/person/:age", func(c *gin.Context) { // 使用匿名函数作为路由的处理函数，处理函数必须是func(*gin.Context)类型的函数
		var (
			user   User
			result gin.H // gin.H() 方法简化json的生成，本质就是一个map[string]interface{}
		)
		age := c.Param("age")
		fmt.Printf("输入年龄：%d", age)
		s1 := "SELECT age,name from user where age=?"
		row := db.QueryRow(s1, age)
		err = row.Scan(&user.age, &user.name)
		fmt.Printf("用户: %+v\n", user)
		if err != nil {
			result = gin.H{
				"user":  nil,
				"count": 0,
			}
		} else {
			result = gin.H{
				"age":   user.age,
				"name":  user.name,
				"count": 1,
			}
		}
		c.JSON(http.StatusOK, result)

	})
	router.Run(":18500")

}
