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
	// 若返回俩参数， 则第二个个参数返回判断结果，否则当判断结果是否是引发panic
	//
	// var status, ok = result["status"].(int) // ok is false
	// var status = result["status"].(int) // will panic: interface conversion: interface {} is float64, not int
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
