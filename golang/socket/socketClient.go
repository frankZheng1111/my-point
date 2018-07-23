package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sender(conn net.Conn) {
	defer conn.Close()
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	words := fmt.Sprintf("%d: hello world!", rand.Intn(100))
	time.Sleep(time.Second * 3) //
	buffer := make([]byte, 2048)
	fmt.Println("My send out words: ", words)
	conn.Write([]byte(words))
	fmt.Println("send over")
	// time.Sleep(time.Second * 6) // 超过5 server 端会关闭连接
	conn.Write([]byte(words))
	fmt.Println("send over again")

	// ReadBuffer, 会在此阻塞
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	wg.Done()
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

	wg.Add(5)
	for i := 0; i < 5; i++ {
		// https://golang.org/pkg/net/#DialTCP
		// 类似于拨号建立连接
		//
		go func() {
			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
				os.Exit(1)
			}

			fmt.Println("connect success")
			sender(conn)
		}()
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
