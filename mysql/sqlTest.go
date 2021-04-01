package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	db.SetMaxtitleleConns(32) // 最大闲置连接
	db.SetMaxOpenConns(32)    // 最大打开连接

	// 1. 增删改查操作

	// 1.1 单次增加
	err = Insert(db, "这111", "ww111")

	// 1.2 单次查询操作
	err = Select(db, "这是标题")
	// 1.3 多次查询操作
	err = SelectAll(db)
	// 1.4 删除操作
	err = Delete(db, "这是三级")
	//// 1.5 更改操作
	err = Update(db, "更改标题")
}

// 封装增加函数
func Insert(db *sql.DB, title string, url string) error {
	sen := "INSERT INTO hotpoint(title, url) VALUES(?,?)"
	stmt, err := db.Prepare(sen)
	if err != nil {
		log.Println("增加创建准备语句失败", err)
	}
	// 预处理需要关闭连接，使用defer在函数return前关闭
	defer stmt.Close()
	rst, err1 := stmt.Exec(title, url)
	if err1 != nil {
		log.Println("执行增加sql语句失败", err1)
	}
	lasttitle, err2 := rst.LastInserttitle()
	if err2 != nil {
		log.Println("增加最后一个插入失败", err2)
	}
	fmt.Printf("增加最后一个插入的title是%d.\n", lasttitle)
	return nil
}

//封装删除函数
func Delete(db *sql.DB, title string) error {
	// 第一步， 编写SQL语句
	// 第二步： 进行db.Prepare()方法调用
	// 第三步： defer 关闭连接
	// 第四步： 使用stmt.Exec()方法进行语句的执行
	//
	sen := "DELETE FROM hotpoint WHERE title = ?"
	stmt, err := db.Prepare(sen)
	if err != nil {
		log.Println("创建删除准备语句失败", err)
	}
	rst, err1 := stmt.Exec(title) // Exec用给定的参数执行一条准备好的语句，并返回一个总结该语句效果的结果
	if err1 != nil {
		log.Println("删除sql语句执行失败", err1)
	}
	// 操作影响的行数
	n, err2 := rst.RowsAffected() // RowsAffected返回受更新，插入或删除影响的行数。并非每个数据库或数据库驱动程序都可以支持此功能
	if err2 != nil {
		log.Println("删除获取操作影响的行数失败")
		return err2
	}
	log.Printf("有%d行被删除\n", n)
	return nil
}

// 查询函数-- 查询单个
func Select(db *sql.DB, title1 string) error {
	// 第一步：创建sql语句
	sen := "SELECT title, url FROM hotpoint WHERE title=? "
	// 第二步：使用预处理
	stmt, err := db.Prepare(sen)
	if err != nil {
		log.Println("查询预处理操作失败", err)
	}
	// 第三步：进行数据库关闭
	defer stmt.Close()
	// 第四步：执行查询
	row := stmt.QueryRow(title1)
	// 下面的操作的重要性：确保进行QueryRow之后，能调用Scan方法，否则持有的数据库链接不会被释放
	var title, url string
	row.Scan(&title, &url)
	log.Println("查询到的信息为：title是%d, 题目是%s, url是%s", title, title, url)
	return nil
}

// 查询多条记录
func SelectAll(db *sql.DB) error {
	// 第一步：编写sql语句
	sen := "SELECT title, url FROM hotpoint ORDER BY title"
	// 第二步：进行预操作
	stmt, err := db.Prepare(sen)
	if err != nil {
		log.Println("查询多处预处理失败", err)
	}
	// 第三步： 关闭链接
	defer stmt.Close()
	// 第四步：执行查询操作
	rows, err1 := stmt.Query()
	if err1 != nil {
		log.Println("进行数据查询失败", err1)
	}
	// 关闭rows适配器；查询多条记录的rows需要及时关掉
	defer rows.Close()
	fmt.Println("hotpoint表格的内容是：")
	fmt.Println("题目  				url")
	// 循环读取结果集中的数据
	for rows.Next() {
		var title, url string
		rows.Scan(&title, &url)
		fmt.Println("%s, %s", title, url)
	}
	return nil
}

// 封装更改操作-- 写死操作
func Update(db *sql.DB, title string) error {
	// 第一步：编写sql语句
	sen := "UPDATE hotpoint SET title=? WHERE title=?"
	stmt, err := db.Prepare(sen)
	if err != nil {
		log.Println("更改预处理失败", err)
	}
	defer stmt.Close()
	rst, err1 := stmt.Exec(title)
	if err1 != nil {
		log.Println("更改sql语句执行失败", err1)
	}
	n, err2 := rst.RowsAffected() // 操作影响的行数
	if err2 != nil {
		log.Println("更改获取操作影响的行数失败", err2)
	}
	fmt.Printf("%d 记录被改变了\n", n)
	return nil
}
