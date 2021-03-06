package models

import (
  "fmt"
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
