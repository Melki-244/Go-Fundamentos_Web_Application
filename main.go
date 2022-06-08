package main

import (
	"fmt"
	"html/template"
	"net/http"
  "github.com/Melki-244/Go-Fundamentos_Web_Application/models"
)


var temp = template.Must(template.ParseGlob("templates/*.html")) 
const porta = ":8080" 


func main() {
  fmt.Println("Server Runing In localhost"+porta)
  http.HandleFunc("/", index)
  http.ListenAndServe(porta, nil) 
}

func index(rw http.ResponseWriter, rq *http.Request)  {
  produtos := models.FindProdutos()
  temp.ExecuteTemplate(rw, "Index", produtos)
}
