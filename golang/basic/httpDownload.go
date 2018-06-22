/*
 * 要点
 * 下载文件
 * 创建文件os.Create(返回*File, 有实现writer接口) 记defer关闭文件
 * 写入文件io.Copy(返回参数1需实现有实现writer接口) 记defer关闭文件
 *
 * flag 通常有两种使用方法
 * flag.String(flag名称，默认值，提示内容) 返回指针
 * flag.StringVar(需要赋值变量的地址， flag名称，默认值，提示内容)
 *

 */
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
