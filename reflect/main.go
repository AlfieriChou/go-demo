package main

import (
  "fmt"
  "reflect"
)

func main() {
  b := true
  s := ""
  n := 1
  f := 1.1
  a := []string{"foo", "bar", "baz"}
  fmt.Println(reflect.TypeOf(b))
  fmt.Println(reflect.TypeOf(s))
  fmt.Println(reflect.TypeOf(n))
  fmt.Println(reflect.TypeOf(f))
  fmt.Println(reflect.TypeOf(a))
}