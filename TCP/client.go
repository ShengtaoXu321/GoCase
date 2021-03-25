package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:5000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("建立连接失败", err)
	} else {
		connHandler1(conn)
	}
}

func connHandler1(c net.Conn) {
	buf := make([]uint8, 1024)
	var s string
	_, _ = fmt.Scan(&s)
	c.Write([]uint8(s))

	n, err := c.Read(buf)
	if err != nil {
		log.Println("读写出错", err)
	} else {
		log.Println(string(buf[:n]))
	}

}
