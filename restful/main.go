package main

import (
  "log"
  "net/http"
  controller "./controller"
)

func main () {
  http.HandleFunc("/", controller.HomePage)
  http.HandleFunc("/articles", controller.ReturnAllArticles)
  log.Fatal(http.ListenAndServe(":8081", nil))
}
