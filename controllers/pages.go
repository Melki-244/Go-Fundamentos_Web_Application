package controllers

import (
  "html/template"
  "net/http"

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
