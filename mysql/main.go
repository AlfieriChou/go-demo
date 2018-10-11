package main

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Admin struct {
  ID   int    `json:"id"`
  Name string `json:"name"`
  Phone string `json:"phone"`
}

func main() {
  fmt.Println("Go MySql Tutorial")
  db, err := sql.Open("mysql", "koasuper:zhazhahui@tcp(106.15.230.136:3306)/koasupermall")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  results, err := db.Query("SELECT id, name, phone FROM admin")
  if err != nil {
    panic(err.Error())
  }
  fmt.Println(results)
  for results.Next() {
    var admin Admin
    err = results.Scan(&admin.ID, &admin.Name, &admin.Phone)
    if err != nil {
      panic(err.Error())
    }
    log.Printf(admin.Name + "——>" + admin.Phone)
  }

  var admin Admin
  err = db.QueryRow("SELECT id, name, phone FROM admin where id = ?", 1).Scan(&admin.ID, &admin.Name, &admin.Phone)
  if err != nil {
    panic(err.Error())
  }
  log.Printf(admin.Name + "——>" + admin.Phone)
}