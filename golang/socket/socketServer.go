package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//建立socket，监听端口
	// Step1: Listen
	//
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {

		// Step2: Accept
		//
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}

//处理连接
//
func handleConnection(conn net.Conn) {

	buffer := make([]byte, 12) // 大小取决于一次能读的长度
	defer conn.Close()         // 仅能关闭打开的连接

	for {
		// Step3: ReadBuffer, 会在此阻塞
		//
		fmt.Println("Wait to read content...")
		n, err := conn.Read(buffer) // 根据buffer的长度读出指定的内容，读完后阻塞
		conn.Write([]byte("I got it"))

		if err != nil {
			// 若客户端断开连接(包括不限于客户端调用conn.Close(), 客户端进程停止)
			// 则err: EOF
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}

}
func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
