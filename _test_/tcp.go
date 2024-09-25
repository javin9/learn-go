package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	//关闭
	defer l.Close()
	//阻塞等待用户连接
	//for {
	// Wait for a connection.
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("connect error", err)
		return
		//continue
	}
	// Handle the connection in a new goroutine.
	// The loop then returns to accepting, so that
	// multiple connections may be served concurrently.
	//接收用户请求
	var byteArray = make([]byte, 1024)
	n, connErr := conn.Read(byteArray)
	if connErr != nil {
		fmt.Println("connErr error", connErr)
		//continue
		return
	}
	fmt.Printf(string(byteArray[:n]))
	//关闭用户连接
	defer conn.Close()
	//go func(c net.Conn) {
	//	// Echo all incoming data.
	//	io.Copy(c, c)
	//	// Shut down the connection.
	//	c.Close()
	//}(conn)
	//}
}
