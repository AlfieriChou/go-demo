package controller

import (
  "encoding/json"
  "fmt"
  // "log"
  "net/http"
  // "strings"
  "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  // "github.com/gorilla/mux"
  "github.com/mitchellh/mapstructure"
)

type User struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

type JwtToken struct {
  Token string `json:"token"`
}

type Exception struct {
  Message string `json:"message"`
}

func CreateToken(w http.ResponseWriter, req *http.Request) {
  var user User
  _ = json.NewDecoder(req.Body).Decode(&user)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": user.Username,
    "password": user.Password,
  })
  tokenString, error := token.SignedString([]byte("secret"))
  if error != nil {
    fmt.Println(error)
  }
  json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func Protected(w http.ResponseWriter, req *http.Request) {
  params := req.URL.Query()
  token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("There was an error")
    }
    return []byte("secret"), nil
  })
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    var user User
    mapstructure.Decode(claims, &user)
    json.NewEncoder(w).Encode(user)
  } else {
    json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
  }
}
