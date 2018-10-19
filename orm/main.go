package main

import(
  "log"
  "net/http"
  "github.com/gorilla/mux"

  model "./models"
  service "./services"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)
  model.InitMigration()
  router.HandleFunc("/users", service.AllUsers).Methods("GET")
  log.Fatal(http.ListenAndServe(":8081", router))
}