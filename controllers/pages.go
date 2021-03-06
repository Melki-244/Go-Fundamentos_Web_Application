package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Melki-244/Go-Fundamentos_Web_Application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) 

func Index(rw http.ResponseWriter, rq *http.Request)  {
  produtos := models.FindProdutos()
  temp.ExecuteTemplate(rw, "Index", produtos)
}
func New(rw http.ResponseWriter, rq *http.Request)  {
  temp.ExecuteTemplate(rw, "New", nil) 
}
func Insert(w http.ResponseWriter, r *http.Request)  {
  if r.Method == "POST" {
      nome := r.FormValue("nome")
      descricao := r.FormValue("descricao")
      preco := r.FormValue("preco")
      quantidade := r.FormValue("quantidade")
    //Convertendo Valores 
    precoToFloat64, err := strconv.ParseFloat(preco, 64)

    if err != nil {
      log.Println("Erro Na Conversão do Preço: ", err)
    }
    quantidadeToInt, err := strconv.Atoi(quantidade)

    if err != nil {
      log.Println("Erro Na Conversão de Quantidade: ", err)
    }

    models.CriarProduto(nome, descricao, precoToFloat64, quantidadeToInt)
  }  
  http.Redirect(w, r, "/", 301)
}
func Delete(w http.ResponseWriter, r *http.Request)  {
  IdProduto := r.URL.Query().Get("id")     
  models.DeletaProduto(IdProduto)
  http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request)  {
  IdProduto := r.URL.Query().Get("id")
  produto := models.EditaProduto(IdProduto)
  temp.ExecuteTemplate(w, "Edit", produto)
}
func Update(w http.ResponseWriter, r *http.Request)  {
  if r.Method == "POST" {
    id := r.FormValue("id")
    nome := r.FormValue("nome")
    descricao := r.FormValue("descricao")
    preco := r.FormValue("preco")
    quantidade := r.FormValue("quantidade")

    // Convertendo Valores
    idToInt, err := strconv.Atoi(id)
    if err != nil {
      log.Println("Erro Na Conversão do Id Para Inteiro: ", err)
    }

    quantidadeToInt, err := strconv.Atoi(quantidade) 
    if err != nil {
      log.Println("Erro Na Conversão da Quantidade Para Inteiro: ", err)
    }

    precoToFloat64, err := strconv.ParseFloat(preco, 64)
    if err != nil {
      log.Println("Erro Na Conversão do Preço Para Float64: ", err)
    }
      
    models.AtualizaProduto(idToInt, nome, descricao, precoToFloat64, quantidadeToInt)
  }

  http.Redirect(w, r, "/", 301)
}
