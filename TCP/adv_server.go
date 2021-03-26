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
	addr := "0.0.0.0:18100"
	s, err := net.Listen("tcp", addr)
	log.Println("建立连接")
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
			go connHandler02(conn)
		}
	}
}

func connHandler02(c net.Conn) {

	defer func() {
		c.Close()
	}()
	for {
		buf := make([]uint8, 1024)
		cnt, err := c.Read(buf)

		if err != nil {
			log.Println("读取数据失败", err)
			return
		}
		log.Println(string(buf[:cnt]))
		c.Write(buf[:cnt])

	}

}
