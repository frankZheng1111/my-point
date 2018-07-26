// go test -bench=. -benchmem -benchtime=3s -run=none
package main

import (
	"sync"
	"testing"
)

// 一个[]byte的对象池，每个对象为一个[]byte
// Get方法和Put方法
var bytePool = sync.Pool{
	// 当池里没有数据的时候会根据方法new一个
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func BenchmarkNewWithoutPool(b *testing.B) {
	b.ResetTimer()
	// 不使用对象池
	for i := 0; i < b.N; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
}

func BenchmarkNewWithPool(b *testing.B) {
	b.ResetTimer()
	// 使用对象池// 会快些
	for i := 0; i < b.N; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}
