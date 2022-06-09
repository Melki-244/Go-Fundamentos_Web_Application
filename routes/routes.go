package routes

import (
	"net/http"
  "github.com/Melki-244/Go-Fundamentos_Web_Application/controllers"
)

func AllRoutes()  {
  http.HandleFunc("/", controllers.Index)
  http.HandleFunc("/new", controllers.New)
}
