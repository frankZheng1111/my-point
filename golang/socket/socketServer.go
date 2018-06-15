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

	buffer := make([]byte, 2048)

	for {

		// Step3: ReadBuffer
		//
		n, err := conn.Read(buffer)
		conn.Write([]byte("I got it"))

		if err != nil {
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
