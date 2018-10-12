package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  controller "./controller"
)

func main() { 
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", controller.HomePage)
  router.HandleFunc("/articles", controller.AllArticles)
  router.HandleFunc("/articles/{id}", controller.ShowArticle)
  log.Fatal(http.ListenAndServe(":10000", router))
}