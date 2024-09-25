package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("dial is error")
		return
	}
	defer conn.Close()
	//var byteArray = make([]byte, 1024)
	conn.Write([]byte("are you ok"))
}
