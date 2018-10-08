package main

import (
  "fmt"
  "log"

  "github.com/golang/protobuf/proto"
)

func main() {
  alfieri := &Person{
    Name: "Alfieri",
    Age: 100,
  }

  data, err := proto.Marshal(alfieri)
  if err != nil {
    log.Fatal("Marshal Error: ", err)
  }

  fmt.Println(data)

  newAlfieri := &Person{}
  err = proto.Unmarshal(data, newAlfieri)
  if err != nil {
    log.Fatal("Marshal Error: ", err)
  }

  fmt.Println(newAlfieri.GetAge())
  fmt.Println(newAlfieri.GetName())
}