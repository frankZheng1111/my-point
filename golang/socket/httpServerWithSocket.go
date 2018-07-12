package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
)

var conns []*net.Conn = []*net.Conn{}

func Log(v ...interface{}) {
	log.Println(v...)
}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func holdSockets() {
	server := "127.0.0.1:8001"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)

	netListen, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)
	defer netListen.Close()

	for {

		Log("Waiting for new client")
		// Step2: Accept, 客户端拨号前会阻塞在此
		//
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go func() {
			Log(conn.RemoteAddr().String(), " tcp connect success")
			conns = append(conns, &conn)
		}()
	}
}

func publish(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(1024)
	io.WriteString(w, fmt.Sprintf("I push it! msg-id: %d", id))
	for _, conn := range conns {
		(*conn).Write([]byte(fmt.Sprintf("published msg-id: %d", id))) //Here I'm attemping to send it
	}
}

func main() {
	http.HandleFunc("/publish-msg", publish)
	go holdSockets()
	fmt.Println("HTTP SERVER Listen :8000")
	http.ListenAndServe(":8000", nil)
}
