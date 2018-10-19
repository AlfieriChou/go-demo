package services

import (
  "fmt"
  "net/http"
  "encoding/json"
  model "../models"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
  users, err := model.AllUsers()
  if err != nil {
    fmt.Println(err.Error())
  }
	fmt.Println("{}", users)
	json.NewEncoder(w).Encode(users)
}