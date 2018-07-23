package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//建立socket，监听端口
	// Step1: Listen
	//
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	go func() {
		for {
			Log("Waiting for new client")
			// Step2: Accept, 客户端拨号前会阻塞在此
			//
			conn, err := netListen.Accept()
			ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
			if err != nil {
				continue
			}
			go func() {
				Log(conn.RemoteAddr().String(), " tcp connect success")
				handleConnection(conn)
			}()
		}
	}()

	// 在context被结束前保持监听
	// select {
	// case <-ctx.Done():
	// 	Log("Waiting for new client too long")
	// 	os.Exit(1)
	// }
}

//处理连接
//
func handleConnection(conn net.Conn) {

	// buffer := make([]byte, 12) // 大小取决于一次能读的长度
	buffer := make([]byte, 2048)
	defer conn.Close() // 仅能关闭打开的连接

	for {
		// Step3: ReadBuffer, 会在此阻塞
		//
		Log("Wait to read content...")
		n, err := conn.Read(buffer) // 根据buffer的长度读出指定的内容，读完后阻塞

		if err != nil {
			// 若客户端断开连接(包括不限于客户端调用conn.Close(), 客户端进程停止)
			// 则err: EOF
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		conn.Write([]byte("Resp about " + string(buffer[:n]) + "I got it"))

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
