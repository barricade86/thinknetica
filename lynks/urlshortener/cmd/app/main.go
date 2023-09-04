package main

import (
	"context"
	"log"
	"lynks/urlshortener/pkg/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	// Подключение к БД. Функция возвращает объект БД.
	db, err := pgxpool.New(ctx, "postgres://postgres:root@0.0.0.0/lynks")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	pgStorage := storage.NewShortLinks(db)
	pgStorage.Add("http://google.com")
}
