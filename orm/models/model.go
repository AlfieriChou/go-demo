package models

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMigration() {
  db, err := gorm.Open("mysql", "lvyang:zhazhayang@tcp(106.15.230.136:3306)/test?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err.Error())
    panic("failed to connect database")
  }
  defer db.Close()
  db.AutoMigrate(&User{})
}

type User struct {
  gorm.Model
  Name string
  Email string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
  db, err := gorm.Open("mysql", "lvyang:zhazhayang@tcp(106.15.230.136:3306)/test?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err.Error())
    panic("failed to connect database")
  }
  defer db.Close()
  var users []User
  db.Find(&users)
  fmt.Println("{}", users)
  json.NewEncoder(w).Encode(users)
}
