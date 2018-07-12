package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn, i int) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	log.Println("Wait to read content...")
	for {
		n, err := conn.Read(buffer) // 根据buffer的长度读出指定的内容，读完后阻塞
		if err != nil {
			log.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		log.Println(conn.RemoteAddr().String(), " and ", i, " receive data string:\n", string(buffer[:n]))
	}
}

func main() {
	server := "127.0.0.1:8001"

	tcpAddr, err := net.ResolveTCPAddr("tcp", server)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	for i := 0; i <= 10; i++ {
		go func(i int) {
			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
				os.Exit(1)
			}

			fmt.Println("connect success")
			sender(conn, i)
		}(i)
	}
	time.Sleep(1000 * time.Second)
}
