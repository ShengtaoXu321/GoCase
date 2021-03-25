package main

import (
	"log"
	"net"
)

/*
Server编写步骤
1. 建立并绑定Socket
2. 使用listen监听请求
3. 使用accept接收请求
4. 进行通信
*/

// 一个简单服务器的实现
func main() {
	addr := "127.0.0.1:5000"
	s, err := net.Listen("tcp", addr)
	// 关闭连接
	defer func() {
		if err := s.Close(); err != nil {
			log.Println("关闭socket失败", err)
		}
	}()
	// 连接出错，就打印错误信息，并退出
	if err != nil {
		log.Println("连接出错", err)
	}
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Println("接收请求失败", err)
		} else {
			go connHandler(conn)
		}
	}
}

func connHandler(c net.Conn) {
	buf := make([]uint8, 1024)

	cnt, err := c.Read(buf)
	defer func() {
		c.Close()
	}()
	if err != nil {
		log.Println("读取数据失败", err)

	}
	log.Println(string(buf[:cnt]))
	c.Write(buf[:cnt])

}
