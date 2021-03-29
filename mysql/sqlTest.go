package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 1. 测试SQL的增删改查操作

func main() {
	// db配置信息
	user := "root"
	psw := "123456"
	host := "117.78.34.82"
	port := "18100"
	dbname := "Hots"
	charset := "utf8mb4"
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, psw, host, port, dbname, charset)
	fmt.Println(config)

	// 连接数据库，获取连接实例
	db, err := sql.Open("mysql", config)
	// 判断
	if err != nil {
		fmt.Println("打开数据库错误", err)
	}

	//设置最大闲置连接和最大的打开连接，具体数值按需指定
	db.SetMaxIdleConns(32) // 最大闲置连接
	db.SetMaxOpenConns(32) // 最大打开连接

	// 增删改查操作
	err = Insert(db, "这是标题", "www.baidu.com")
}

func Insert(db *sql.DB, title string, url string) error {
	sen := "INSERT INTO hotpoint(title, url) VALUES(?,?)"
	stmt, err := db.Prepare(sen)
	if err != nil {
		return err
	}
	// 预处理需要关闭连接，使用defer在函数return前关闭
	defer stmt.Close()
	rst, err1 := stmt.Exec(title, url)
	if err1 != nil {
		fmt.Println("执行sql语句失败")
		return err1
	}
	lastId, err2 := rst.LastInsertId()
	if err2 != nil {
		fmt.Println("最后一个插入失败")
		return err2
	}
	fmt.Printf("最后一个插入的id是%d.\n", lastId)
	return nil
}
