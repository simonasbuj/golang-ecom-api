package main

import (
	"database/sql"
	repo "golang-ecom-api/internal/adapters/sqlite/sqlc"
	"golang-ecom-api/internal/products"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dbType string
	dsn    string
}

type app struct {
	config 	config
	db		*sql.DB	
}

func (app *app) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy boi"))
	})

	querier := 	repo.New(app.db)
	productService := products.NewService(querier)
	productsHandler := products.NewHandler(productService)
	r.Get("/products", productsHandler.ListProducts)

	r.Get("/products/{id}", productsHandler.GetProductByID)

	return r
}

func (app *app) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("starting server at addr %s", app.config.addr)

	return srv.ListenAndServe()
}
