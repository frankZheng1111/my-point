package main

import "fmt"

type Robot struct {
  robot_type string
}

func (robot Robot) showType() {
  fmt.Println("Robot type is", robot.robot_type)
}

// 不使用指针则无法修改对象的值
//
func (robot *Robot) setType(robot_type string) string{
  robot.robot_type = robot_type
  return robot_type
}

func main() {
  fmt.Println("Init object")
  var rx78 Robot
  rx78.setType("MS")
  rx78.showType()
}
