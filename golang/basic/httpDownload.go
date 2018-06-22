package main

import (
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
	n, err := io.Copy(out, resp.Body)
	fmt.Println(n)
	if err != nil {
		return
	}
	return
}

func main() {
	err := DownloadFile("path", "url")
	fmt.Println(err)
}
