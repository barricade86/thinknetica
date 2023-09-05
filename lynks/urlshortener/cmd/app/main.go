package main

import (
	"context"
	"log"
	"lynks/urlshortener/pkg/api"
	"lynks/urlshortener/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	// Подключение к БД. Функция возвращает объект БД.
	db, err := pgxpool.New(ctx, "postgres://postgres:root@0.0.0.0/lynks")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	pgStorage := storage.NewPgShortLinks(db)
	api := api.New(pgStorage)
	router := mux.NewRouter()
	router.HandleFunc("/link/create", api.CreateShortLink)
	router.HandleFunc("/link/get", api.RedirectToOriginal)
	http.ListenAndServe("localhost:8081", router)
}
