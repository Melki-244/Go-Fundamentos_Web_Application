package models

import "github.com/Melki-244/Go-Fundamentos_Web_Application/db"

type Produto struct {
  Id int
  Nome string 
  Descricao string
  Preco float64
  Quantidade int
}

func FindProdutos()  []Produto{
  db := db.ConectaComOBancoDeDados()
  defer db.Close()

  selectDeTodosOsProdutos, err :=  db.Query("select * from produtos order by id asc")

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

    p.Id = id
    p.Nome = nome
    p.Descricao = descricao
    p.Preco = preco
    p.Quantidade = quantidade

    produtos = append(produtos, p)
  }
  return produtos
}
func CriarProduto(nome, descricao string, preco float64, quantidade int)  {
  db := db.ConectaComOBancoDeDados() 
  InsereOProduto, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

  if err != nil {
    panic(err.Error())
  }
  
  InsereOProduto.Exec(nome, descricao, preco, quantidade)

  defer db.Close()
}
func DeletaProduto(id string)  {
  db := db.ConectaComOBancoDeDados()
  defer db.Close()

  DeletaOProduto, err := db.Prepare("delete from produtos where id=$1")

  if err != nil {
    panic(err.Error)
  }
  DeletaOProduto.Exec(id)
}
func EditaProduto(id string) Produto {
  db := db.ConectaComOBancoDeDados() 
  defer db.Close()

  produtoDoBanco, err := db.Query("select * from produtos where id=$1",id)

  if err != nil {
    panic(err.Error())
  }

  produtoEdicao := Produto{}

  for produtoDoBanco.Next() {
    var nome, descricao string
    var id, quantidade int
    var preco float64

    produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

    produtoEdicao.Id = id 
    produtoEdicao.Nome = nome
    produtoEdicao.Descricao = descricao
    produtoEdicao.Preco = preco
    produtoEdicao.Quantidade = quantidade
  }

   return produtoEdicao
}
func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int)  {
  db := db.ConectaComOBancoDeDados() 
  defer db.Close()
  
  AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5") 
  if err != nil {
    panic(err.Error())
  }

  AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
}
