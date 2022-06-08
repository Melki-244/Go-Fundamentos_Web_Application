package models

import "github.com/Melki-244/Go-Fundamentos_Web_Application/db"

type Produto struct {
  Nome string 
  Descricao string
  Preco float64
  Quantidade int
}

func FindProdutos()  []Produto{
  db := db.ConectaComOBancoDeDados()
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
  return produtos
}
