package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	defer conn.Close()
	words := "hello world!"
	time.Sleep(time.Second * 4)
	buffer := make([]byte, 2048)
	conn.Write([]byte(words))
	fmt.Println("send over")

	// ReadBuffer, 会在此阻塞
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	time.Sleep(time.Second * 5)
}

func main() {
	server := "127.0.0.1:1024"

	// 解析一个tcp地址
	// https://golang.org/pkg/net/#ResolveTCPAddr
	//
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	// https://golang.org/pkg/net/#DialTCP
	// 类似于拨号建立连接
	//
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)
}
