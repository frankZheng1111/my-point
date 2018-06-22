/*
	要点
	https://stackoverflow.com/questions/21197239/decoding-json-in-golang-using-json-unmarshal-vs-json-newdecoder-decode

	Use json.Decoder if your data is coming from an io.Reader stream, or you need to decode multiple values from a stream of data.
	参数是io.Reader(stream), 适合从文件读stream的场景
	Use json.Unmarshal if you already have the JSON data in memory.
	参数是byte[](buf)适用已经在内存里的场景
*/
package main

// http://cizixs.com/2016/12/19/golang-json-guide
import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string          `json:name`
	IsAdmin bool            `json:is_admin`
	Auth    json.RawMessage `json:auth`
}

func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	// 尽管 status 字段没有小数点，我们希望它是整数status值，解析的时候它还是会被解析成 float64，如果直接把它当做 int 访问，会出现错误：
	// 个人理解, 以.(int)为例，完成了查询是否实现了int方法并将interface类型转换为int类型
	var status = result["status"].(float64)
	fmt.Println("status value:", status)

	var userJson = []byte(`{"name": "Wang", "is_admin": false, "auth": { "token": "token1" } }`)
	var user User
	if err := json.Unmarshal(userJson, &user); err != nil {
		fmt.Println("error:", err)
		return
	}
	json.Unmarshal(user.Auth, &result)
	fmt.Println("status value:", result["token"])
}
