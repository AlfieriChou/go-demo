package controller

import (
  "fmt"
  "net/http"
  "encoding/json"
  "strconv"
  "github.com/gorilla/mux"
)

type Article struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request) {
  articles := Articles{
    Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
  }
  fmt.Println("Endpoint Hit: returnAllArticles")
  json.NewEncoder(w).Encode(articles)
}

func ShowArticle(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var key string
  key = vars["id"]
  articles := Articles{
    Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
  }
  i, err := strconv.Atoi(key)
  if err != nil {
    panic(err.Error())
  }
  if i < 2 {
    fmt.Println(articles[i - 1])
  }
  fmt.Fprintf(w, "Key: " + key)
}
