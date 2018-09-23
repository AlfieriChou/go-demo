package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  Read()
  Write()
}

func Read() {
  data, err := ioutil.ReadFile("test.txt")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(data))
}

func Write() {
  testData := []byte("what?")
  err := ioutil.WriteFile("test.txt", testData, 0777)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Successfully write this data.")
}