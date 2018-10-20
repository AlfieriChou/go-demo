package main

import(
  "log"
  "net/http"
  "github.com/gorilla/mux"

  model "./models"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)
  model.InitMigration()
  router.HandleFunc("/users", model.AllUsers).Methods("GET")
  log.Fatal(http.ListenAndServe(":8081", router))
}