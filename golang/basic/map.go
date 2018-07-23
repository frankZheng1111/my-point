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
		nevacuate  uintptr        // 搬迁进度（小于nevacuate的已经搬迁）

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
	// 4. 通过桶的索引确定在哪一个bucket中查找，bucket最后有overflow指针指向的溢出桶，形成链表结构，遍历链表(with 先遍历当前桶的tophash，若hash值的高八位相同, 判断对应位置的key是否相同，相同则返回，否则继续查询)，直到结束为止。

	// map分配值:
	// 0. flag(按位或计算)设置为当前是写状态
	// 0-again: 判断是否正在扩容(即oldbuckets中是否有内容), 若是，迁移1至2个pair
	// 1. 根据key计算hash值
	// 2. hash值通过位运算根据桶的数量取模值计算为桶的索引
	// 3. 通过桶的索引确定在哪一个bucket中查找，bucket最后有overflow指针指向的溢出桶，形成链表结构，遍历链表(
					with 先遍历当前桶的tophash，若当前桶的tophash中有空位, 先保留该空位
				)，直到结束为止。若期间找到已存在的key(先比对hash高八位, 再比对对应的key值是否精确相等), 若相等, 则更新对应值, 直接跳到最后一步
  // 4  若当前未在扩容, 若超出装载因子，直接扩容，若仅溢出桶数量大于桶的数量，设置oldbuckets为当前bucket, 若没有超出限制且未找到空位，创建一个溢出桶, 跳至again
	// 5. 键值对写入保留下的空位
	// 5. flag(按位与或计算)解除为当前是写状态

	// map删除值
	// 基本流程等同于给分配值, 改为移除对应的tophash(设为empty), 故内存并没有减少

	// map扩容
	// 每次扩容后桶的数量是当前的2倍
	// oldbuckets = buckets
	// extra: oldoverflow=overflow
	// B(+1)后, rehash的情况下 若 hash & (2^B) == 0，说明 hash < 2^B，那么它将落入与旧桶集合相同的索引位置中；
	// 否则，它将落入原先位置 + 2^B中。
	//
	// 总结
	// map是由数组+链表实现的HashTable，其大小和B息息相关
	// Golang通过hashtop快速试错加快了查找过程，利用空间换时间的思想解决了扩容的问题(增量迁移)
	//
	//   随着元素的增加，在一个bucket链中寻找特定的key会变得效率低下，所以在插入的元素个数/bucket个数达到某个阈值（当前设置为6.5，实验得来的值）时，map会进行扩容，代码中详见 hashGrow函数。首先创建bucket数组，长度为原长度的两倍，然后替换原有的bucket，原有的bucket被移动到oldbucket指针下。
	// 扩容完成后，每个hash对应两个bucket（一个新的一个旧的）。oldbucket不会立即被转移到新的bucket下，而是当访问到该bucket时，会调用growWork方法进行迁移，growWork方法会将oldbucket下的元素rehash到新的bucket中。随着访问的进行，所有oldbucket会被逐渐移动到bucket中。
  //
	// 但是这里有个问题：如果需要进行扩容的时候，上一次扩容后的迁移还没结束，怎么办？在代码中我们可以看到很多”again”标记，会不断进行迁移，知道迁移完成后才会进行下一次扩容。
	// 利用将8个key(8个value)依次放置减少了padding空间等等。
*/

import (
	"fmt"
	"reflect"
)

type CustomMap map[string]string

func (cm *CustomMap) Print() {
	fmt.Println(cm)
}

func main() {
	cm := make(CustomMap)
	cm.Print()

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

	// panic: cannot take the address of countryCapitalMap["France"]
	// 原因：map可能会随着元素的增多重新分配更大的内存空间，旧值都会拷贝到新的内存空间，因此之前的地址就会失效。
	_ = &countryCapitalMap["France"]

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}
