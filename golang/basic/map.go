package main

// https://www.jianshu.com/p/aa0d4808cbb8
//https://juejin.im/entry/5a1e4bcd6fb9a045090942d8

/*
	map的底层实现
	map 底层实现有指针指向相应的bucket，键值对存在bucket内，也就是赋值或传参后作响应修改会影响值的原因
	指针指向如下的数据结构
	type hmap struct {
		count     int // # 元素个数
		// 状态 表示四个状态在(在遍历新的buckets(0001b,1), 遍历旧的buckets(0010b,2), 正在写数据(0100b, 4), 正在扩容(1000b, 8)), 对应位为1时, 为触发状态
		flags     uint8
		B         uint8  // 说明包含2^B个bucket
		noverflow uint16 // 溢出的bucket的个数
		hash0     uint32 // hash种子

		buckets    unsafe.Pointer // buckets的数组指针 指向一个对象是*bmap的数组
		oldbuckets unsafe.Pointer // 结构扩容的时候用于复制的buckets数组
		nevacuate  uintptr        // 搬迁进度（已经搬迁的buckets数量）

		extra *mapextra
	}

	type mapextra struct {
		overflow    *[]*bmap
		oldoverflow *[]*bmap
		nextOverflow *bmap
	}

	// tophash用于记录8个key哈希值的高8位，这样在寻找对应key的时候可以更快，不必每次都对key做全等判断。
	// 后续依次按顺序了存储了8个key值8个value(之所以没有交错存储考虑到key，value类型长度不同, 考量到对齐的原因, 节省内存)
	// 最后存了一个overflow指针, 再接下来是hash冲突发生时，下一个溢出桶的地址
	type bmap struct {
		tophash [bucketCnt]uint8
	}

	//几个比较重要的计算
	hash := alg.hash(key, uintptr(h.hash0)) // 对应键值得hash值
	top := uint8(hash >> (sys.PtrSize*8 - 8)) // 取高八位，即hash值右移(系统指针大小(字节) * 8(1byte=8bit) - 8流出的高八位)
	bucketIndex := hash & (uintptr(1)<<h.B - 1)，即 hash % 2^B(假设有8个bucket,B=3,即对111b按位取与运算(结果范围0-7))

	// map初始化
	// 0. 校验几个参数是否符合要求
	// 1. 根据具体情况设置一个合适的B, B为满足count >= bucketcnt && count/B(当前值) >= loadFactor的最大值
	//   插入的(元素个数/bucket)个数达到某个阈值（当前设置为6.5) map会进行扩容
	//   bucketCnt为一个桶可以存储多少个键值对
	//   count 为make传入参数的大小
	// 2. 申请buckets的空间
	// 3. 申请hmap的空间
	// 4. 初始化hmap(赋值buckets, B, hash0(随机数), 其他为默认初值)

	// map访问值:
	// 1. 根据key计算hash值
	// 2. hash值通过位运算根据桶的数量取模值计算为桶的索引
	// 3. 判断oldbuckets指针是否存在，存在的话表示在旧桶内查找
	// 4. 通过桶的索引确定在哪一个bucket中查找，bucket最后有overflow指针指向的溢出桶，形成链表够，遍历链表(with 先遍历当前桶的tophash，若hash值的高八位相同, 判断对应位置的key是否相同，相同则返回，否则继续查询)，直到结束为止。

*/

import (
	"fmt"
	"reflect"
)

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

	anotherMap := countryCapitalMap
	anotherMap["anotherCity"] = "anotherCaptial"
	// map 类型占8个字节，是一个指向 map 结构的指针
	// 故所有的操作会映射到原值
	fmt.Println(reflect.DeepEqual(countryCapitalMap, anotherMap)) // true

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
