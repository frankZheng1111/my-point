package main

import "fmt"

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) == len(b):
		fallthrough
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
func main() {
	a := []byte{'a', 'b'}
	b := []byte{'a', 'b'}
	c := []byte{'b', 'a'}
	d := []byte{'b', 'a', 'a'}
	fmt.Println("a === b", compare(a, b))
	fmt.Println("a < c", compare(a, c))
	fmt.Println("d > c", compare(d, c))
}
