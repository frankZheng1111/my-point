package main

import (
  "fmt"
  "reflect"
)

func main() {
  type S struct {
    F string `species:"gopher" color:"blue"`
    F2 string `species:"gopher2" color:"blue2"`
  }
  s := S{}
  st := reflect.TypeOf(s)
  field := st.Field(0)
  field2 := st.Field(1)
  fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
  fmt.Println(field2.Tag.Get("color"), field2.Tag.Get("species"))
}
