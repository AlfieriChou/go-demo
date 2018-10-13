package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  controller "./controller"
)

func main() {
  router := mux.NewRouter()
  fmt.Println("Starting the application...")
  router.HandleFunc("/", controller.HomePage)
  router.HandleFunc("/login", controller.CreateToken).Methods("POST")
  router.HandleFunc("/protected", controller.Protected).Methods("GET")
  log.Fatal(http.ListenAndServe(":12345", router))
}