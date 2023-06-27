package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"mobilePhoneEdu/internal/controller"
	"mobilePhoneEdu/internal/repository"
	"net/http"
)

var (
	httpPort    = ":8080"
	postgresURL = "postgres://postgres:pass@127.0.0.1:5432/todo"
)

func main() {
	var err error
	controller.DB, err = repository.NewPostgres(context.Background(), postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("controller.DB = ", controller.DB)
	defer controller.DB.Close()
	r := InitHandler()
	startHttp(r)
}

func startHttp(r *chi.Mux) {
	srv := &http.Server{ // конфигурация
		Addr:    httpPort,
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func InitHandler() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/api/phones", controller.CreatePhone)
	router.Get("/api/phones", controller.GetPhone)
	router.Put("/api/phones", controller.UpdatePhone)
	router.Delete("/api/phones", controller.DeletePhone)
	return router
}
