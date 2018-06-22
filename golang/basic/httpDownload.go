package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filePath string, url string) (err error) {
	out, err := os.Create(filePath)
	defer out.Close()
	if err != nil {
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	n, err := io.Copy(out, resp.Body) // n是size
	fmt.Println(n)
	if err != nil {
		return
	}
	return
}

func main() {
	var path, url string
	// 两种解析方式
	// 直接赋值
	flag.StringVar(&path, "path", "default", "file path with name")
	// 返回指针
	urlPtr := flag.String("url", "default", "download url")
	flag.Parse()
	url = *urlPtr
	err := DownloadFile(path, url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Download Success")
}
