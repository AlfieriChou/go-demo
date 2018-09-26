package main

import (
  "fmt"
  // "time"
)

func main() {
  fmt.Println(Add(12, 15))
  fmt.Println(Add2(13, 16))
  a, b := Swap("Hello", "World")
  fmt.Println(a, b)
  fmt.Println(Split(18))
  TestWhile()
  TestForContinue()
  // go f(0)
  // time.Sleep(time.Second * 1)
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

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
  }
}

func TestWhile() {
  a := 10
  for a < 20 {
    if a == 15 {
      a = a + 1
      continue
    }
    fmt.Printf("a 的值为 : %d\n", a)
    a++
  }
}

func TestForContinue() {
  for a := 10; a < 20; a++ {
    if a == 15 {
      continue
    }
    fmt.Printf("你猜猜 a 的值为 : %d\n", a)
  }
}
