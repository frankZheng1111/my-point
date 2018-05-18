package main

// Logging libraries often provide different log levels. Unlike those logging libraries, the log package in Go does more than log if you call its Fatal*() and Panic*() functions. When your app calls those functions Go will also terminate your app :-)

import "log"

func main() {
	// log.Panicln("Panic Level: log entry") // 中断，有出错位置
	log.Fatalln("Fatal Level: log entry") //app exits here 仅中断
	log.Println("Normal Level: log entry")
}
