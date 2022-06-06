package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComOBancoDeDados() *sql.DB {
  conexao := "user=postgres dbname=Alura_Loja-GO password=33114665 host=localhost sslmode=disable"
  db, err := sql.Open("postgres",conexao)

  if err != nil {
    panic(err.Error())
  }
  return db
}

var temp = template.Must(template.ParseGlob("templates/*.html")) 
const porta = ":8080" 

type Produto struct {
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
  db := conectaComOBancoDeDados()
  defer db.Close()

  selectDeTodosOsProdutos, err :=  db.Query("select * from produtos")

  if err != nil {
    panic(err.Error())
  }

  p := Produto{}
  produtos := []Produto{}

  for selectDeTodosOsProdutos.Next() {
    var id, quantidade int 
    var nome, descricao  string 
    var preco float64

    err =  selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &quantidade, &preco)

    if err != nil {
      panic(err.Error())
    }

    p.Nome = nome
    p.Descricao = descricao
    p.Preco = preco
    p.Quantidade = quantidade

    produtos = append(produtos, p)
  }
  temp.ExecuteTemplate(rw, "Index", produtos)
}
