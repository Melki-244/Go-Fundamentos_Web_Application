package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) 

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil) 
}
func index(rw http.ResponseWriter, rq *http.Request)  {
  temp.ExecuteTemplate(rw, "Index", nil)
}
