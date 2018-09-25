package main

import (
	"fmt"
)

func main() {
  fmt.Println(Add(12, 15))
  fmt.Println(Add2(13, 16))
  a, b := Swap("Hello", "World")
  fmt.Println(a, b)
  fmt.Println(Split(18))
}

func Add(x int, y int) int {
  return x + y
}

func Add2(x, y int) int {
  return x + y
}

func Swap(x, y string) (string, string) {
  return y, x
}

func Split(sum int) (x, y int) {
  x = sum * 4 / 9 
  y = sum - x
  return
}