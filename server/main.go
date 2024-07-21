package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jhonasalves/go-expert-fc-multithreading/docs"
	"github.com/jhonasalves/go-expert-fc-multithreading/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Go Expert - Desafio de Multithreading e APIs
// @version         1.0
// @description     Desafio de Multithreading e APIs
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jhonas Alves
// @contact.email  jhonas.alvesm@gmail.com

// @host      localhost:8000
// @BasePath  /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/cep/{cep}", handlers.GetCep)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
