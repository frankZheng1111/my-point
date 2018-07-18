package main

// 字符串给定数量用 string.Join
// 字符串不给定数量用 bytes.Buffer
// + 的性能相对较差

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	strs = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}
)

func TestStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(strs, "")
	}
}

func TestStringsPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < len(strs); j++ {
			s += strs[j]
		}
	}
}

func TestBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for j := 0; j < len(strs); j++ {
			b.WriteString(strs[j])
		}
	}
}

func main() {
	fmt.Println("strings.Join:")
	fmt.Println(testing.Benchmark(TestStringsJoin))
	fmt.Println("bytes.Buffer:")
	fmt.Println(testing.Benchmark(TestBytesBuffer))
	fmt.Println("+:")
	fmt.Println(testing.Benchmark(TestStringsPlus))
}
