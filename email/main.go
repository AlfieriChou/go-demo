package main

import (
  "fmt"
  "net/smtp"
  email "github.com/jordan-wright/email"
)

func main() {
  e := email.NewEmail()
  e.From = "Jordan Wright <alfierichou@gmail.com>"
  e.To = []string{"904887302@qq.com"}
  e.Subject = "Awesome Subject"
  e.Text = []byte("Text Body is, of course, supported!")
  e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
  e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "alfierichou@gmail.com", "****", "smtp.gmail.com"))
  fmt.Println("OK")
}