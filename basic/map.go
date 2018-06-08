package main

import "fmt"

func main() {
	var a map[string]int
	a = make(map[string]int)
	a["test"]++
	fmt.Println(a) // {test:1}

	var countryCapitalMap map[string]string
	fmt.Println(countryCapitalMap, countryCapitalMap == nil) // map[] true
	/* 创建集合 */
	// 分配内存，否则赋值会panic
	countryCapitalMap = make(map[string]string) // 可以
	// countryCapitalMap = map[string]string{} // 可以
	// countryCapitalMap = *new(map[string]string) //不行, new会分配初值，依旧是nil

	fmt.Println(countryCapitalMap, countryCapitalMap == nil) // map[] false

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}
