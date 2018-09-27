package main

import (
  "fmt"
  "strings"
  "reflect"
)

func main () {
  Append()
  Map()
  Split()
}

func Append() {
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)
}

func Map() {
  slice1 := []int{1,2,3}
  slice2 := make([]int, 2)
  copy(slice2, slice1)
  fmt.Println(slice1, slice2)
}

func Split() {
  var a string = "a,b,c,d"
  b := strings.Split(a, ",")
  fmt.Println(reflect.TypeOf(b))
  fmt.Println(reflect.TypeOf(a))
  fmt.Println(b, a)
}