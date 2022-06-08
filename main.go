package main

import (
	"net/http"
  "github.com/Melki-244/Go-Fundamentos_Web_Application/routes"
)

func main() {
  routes.AllRoutes()
  /* Rotas Devem Ser Carregadas Antes Do Servidor */
  http.ListenAndServe(":8080", nil) 
}

