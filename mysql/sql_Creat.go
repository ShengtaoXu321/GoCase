package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	config := "root:123456@tcp(117.78.34.82:18100)/Hots?charset=utf8mb4"
	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Println("数据库连接错误", err)
	}
	defer db.Close()
	err1 := db.Ping()
	fmt.Println(err1)
	err = CreateTable(db)
}

// 封装新建表格函数
func CreateTable(db *sql.DB) error {
	sen := `
		CREATE TABLE users(
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
);`
	stmt, err := db.Prepare(sen)
	if err != nil {
		log.Println("建表预处理失败", err)
	}
	defer stmt.Close()
	rst, err1 := stmt.Exec()
	if err1 != nil {
		log.Println("执行新建表格失败", err1)
	}
	lasttitle, err2 := rst.LastInsertId()
	if err2 != nil {
		log.Println("增加最后一个插入失败", err2)
	}
	fmt.Printf("增加最后一个插入的title是%d.\n", lasttitle)
	return nil
}
