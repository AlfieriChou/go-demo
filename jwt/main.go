package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  middleware "./middleware"
  controller "./controller"
)

func main() {
  router := mux.NewRouter()
  router.Use(middleware.CommonMiddleware)
  fmt.Println("Starting the application...")
  router.HandleFunc("/", controller.HomePage)
  router.HandleFunc("/login", controller.Login).Methods("POST")
  router.HandleFunc("/protected", controller.Protected).Methods("GET")
  log.Fatal(http.ListenAndServe(":12345", router))
}