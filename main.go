package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hanbarfe/rest_example/configs"
	"github.com/hanbarfe/rest_example/handlers"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/users", handlers.Create)
	router.Get("/users", handlers.List)
	router.Get("/users/{id}", handlers.Get)
	router.Delete("/users/{id}", handlers.Delete)
	router.Put("/users/{id}", handlers.Update)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)

}
