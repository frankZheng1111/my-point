package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F  string `species:"gopher" color:"blue"`
		F2 string `species:"gopher2" color:"blue2"`
	}
	s := S{"f1", "f2"}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	field2 := st.Field(1)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
	fmt.Println(field2.Tag.Get("color"), field2.Tag.Get("species"))

	for m := 0; m < reflect.TypeOf(s).NumField(); m++ {
		field := reflect.TypeOf(s).Field(m)
		fmt.Println(field.Type)
		fmt.Println(field.Name)
	}

	fmt.Println(reflect.ValueOf(s).FieldByName("F"))
}
