package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

// 任务安排
// 1. 将上次获取的到数据，保存到数据库中
// 2. 利用Gin返回数据库中内容，能够获取到热点数据

// 1.1 反序列化
type Hots1 []struct {
	Node struct {
		AvatarLarge      string        `json:"avatar_large"`
		Name             string        `json:"name"`
		AvatarNormal     string        `json:"avatar_normal"`
		Title            string        `json:"title"`
		URL              string        `json:"url"`
		Topics           int           `json:"topics"`
		Footer           string        `json:"footer"`
		Header           string        `json:"header"`
		TitleAlternative string        `json:"title_alternative"`
		AvatarMini       string        `json:"avatar_mini"`
		Stars            int           `json:"stars"`
		Aliases          []interface{} `json:"aliases"`
		Root             bool          `json:"root"`
		ID               int           `json:"id"`
		ParentNodeName   string        `json:"parent_node_name"`
	} `json:"node"`
	Member struct {
		Username     string      `json:"username"`
		Website      interface{} `json:"website"`
		Github       interface{} `json:"github"`
		Psn          interface{} `json:"psn"`
		AvatarNormal string      `json:"avatar_normal"`
		Bio          interface{} `json:"bio"`
		URL          string      `json:"url"`
		Tagline      interface{} `json:"tagline"`
		Twitter      interface{} `json:"twitter"`
		Created      int         `json:"created"`
		AvatarLarge  string      `json:"avatar_large"`
		AvatarMini   string      `json:"avatar_mini"`
		Location     interface{} `json:"location"`
		Btc          interface{} `json:"btc"`
		ID           int         `json:"id"`
	} `json:"member"`
	LastReplyBy     string `json:"last_reply_by"`
	LastTouched     int    `json:"last_touched"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	Created         int    `json:"created"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	LastModified    int    `json:"last_modified"`
	Replies         int    `json:"replies"`
	ID              int    `json:"id"`
}

type Result1 struct {
	Title string
	Url   string
}

func main() {
	var db1 *sql.DB
	// 第一步：将api的数据获取下来
	r, err := http.Get("https://www.v2ex.com/api/topics/hot.json")
	if err != nil {
		log.Println("获取api数据失败", err)
	}
	defer r.Body.Close()
	rBody, _ := ioutil.ReadAll(r.Body)
	var res Hots1
	err = json.Unmarshal(rBody, &res) // 将json字符串数据解码到相应的数据结构；Unmaeshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构
	if err != nil {
		log.Println("json字符串解析失败")
	}
	var newResTit []string
	var newResUrl []string
	for _, i := range res {
		newResTit = append(newResTit, i.Title) // 将标题取出来
		newResUrl = append(newResUrl, i.URL)   // 将url提取出来
	}
	//fmt.Println(newResTit)
	//fmt.Println(newResUrl)

	// 初始化数据库
	db1, err1 := SqlInit1()
	if err1 != nil {
		log.Println("数据库初始化失败", err)
	}

	defer db1.Close()
	// 元素赋值+SQL插入
	length := len(newResTit)
	A := make([]Result1, length)
	for i := 0; i < length; i++ {
		A[i].Title = newResTit[i]
		A[i].Url = newResUrl[i]
		s1 := A[i].Title
		s2 := A[i].Url
		sen := "INSERT INTO hotpoint(title,url) VALUES (?,?)"
		_, err1 := db1.Exec(sen, s1, s2)

		if err1 != nil {
			log.Println("数据插入失败", err1)
		}

	}

}

// 数据库初始化函数
func SqlInit1() (db1 *sql.DB, err error) {
	config := "root:123456@tcp(117.78.34.82:18100)/Hots?charset=utf8mb4"
	db1, err = sql.Open("mysql", config)
	if err != nil {
		log.Println("数据库连接失败", err)
	}
	db1.SetMaxIdleConns(100)
	db1.SetMaxOpenConns(100)
	// 判断连通性
	err = db1.Ping()
	if err != nil {
		log.Println("数据库不能连通", err)
	}
	return db1, err
}
