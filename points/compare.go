package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// You can use the equality operator, ==, to compare struct variables if each structure field can be compared with the equality operator.

type data struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

type udata struct {
	num    int               //ok
	checks [10]func() bool   //not comparable
	doit   func() bool       //not comparable
	m      map[string]string //not comparable
	bytes  []byte            //not comparable
}

func main() {
	comparableStruct1 := data{}
	comparableStruct2 := data{}
	// prints: comparableStruct1 == comparableStruct2: true
	fmt.Println("comparableStruct1 == comparableStruct2: ", comparableStruct1 == comparableStruct2)

	unComparableStruct1 := udata{}
	unComparableStruct2 := udata{}
	// invalid operation: uncomparableStruct1 == unComparableStruct2 (struct containing [10]func() bool cannot be compared)
	// fmt.Println("unComparableStruct1 == unComparableStruct2:", unComparableStruct1 == unComparableStruct2)

	unComparableStruct1.m = map[string]string{"one": "a", "two": "b"}
	unComparableStruct2.m = map[string]string{"two": "b", "one": "a"}
	fmt.Println("unComparableStruct2 == unComparableStruct2(use DeepEqual): ", reflect.DeepEqual(unComparableStruct1, unComparableStruct2)) // true

	unComparableStruct1.bytes = []byte{'a', 'b'}
	unComparableStruct2.bytes = []byte{'b', 'a'}

	fmt.Println("unComparableStruct1 == unComparableStruct2(use DeepEqual with slice): ", reflect.DeepEqual(unComparableStruct1, unComparableStruct2)) // false

	var byteSlice1 []byte = nil
	byteSlice2 := []byte{}
	fmt.Println("byteSlice1 == byteSlice2(DeepEqual):", reflect.DeepEqual(byteSlice1, byteSlice2)) //prints: b1 == b2: false
	fmt.Println("byteSlice1 == byteSlice2(bytes.Equal):", bytes.Equal(byteSlice1, byteSlice2))     //prints: b1 == b2: true

	// DeepEqual() doesn't consider an empty slice to be equal to a "nil" slice.
	// This behavior is different from the behavior you get using the bytes.Equal() function.
	// bytes.Equal() considers "nil" and empty slices to be equal.

	var str string = "one"
	var interfaceStr interface{} = "one"
	fmt.Println("str == interfaceStr:", str == interfaceStr, reflect.DeepEqual(str, interfaceStr)) //prints: str == interfacestr: true true

	strSlice := []string{"one", "two"}
	interfaceStrSlice := []interface{}{"one", "two"}
	fmt.Println("strSlice == interfaceStrSlice:", reflect.DeepEqual(strSlice, interfaceStrSlice)) //prints: strSlice == interfaceStrSlice: false (not ok)

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded:", reflect.DeepEqual(data, decoded)) //prints: data == decoded: false (not ok)

}
