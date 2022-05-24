package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) 
var porta = ":8080" 

type Produtos struct {
  Nome string 
  Descricao string
  Preco float64
  Quantidade int
}

func main() {
  fmt.Println("Server Runing In localhost",porta)
  http.HandleFunc("/", index)
  http.ListenAndServe(porta, nil) 
}

func index(rw http.ResponseWriter, rq *http.Request)  {
  produtos := []Produtos{
    {Nome: "Camiseta", Descricao: "Azul, Bem Bonita", Preco: 200.00},
    {"Tenis", "Confortável e bonito", 250.00, 100},
    {"Fone", "Muito Bom", 70.00, 4},
  }

  temp.ExecuteTemplate(rw, "Index", produtos)
}
