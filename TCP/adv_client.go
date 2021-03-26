package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "0.0.0.0:18100"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("建立连接失败", err)
	} else {
		log.Println("与服务器建立连接")
		connHandler01(conn)

	}
}

func connHandler01(c net.Conn) {
	//定义了一个slice
	var s string
	for {
		buf := make([]uint8, 1024)
		_, _ = fmt.Scan(&s)
		if s == "q" {
			log.Println("退出客户端")
			c.Close()
			return
		}
		c.Write([]uint8(s))

		n, err := c.Read(buf)

		if err != nil {
			log.Println("读写出错", err)
			continue
		} else {
			log.Println(string(buf[:n]))
		}
	}

}
