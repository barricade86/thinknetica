package main

import (
	"lynks/cache/pkg/api"
	"lynks/cache/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
			Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
		},
	)

	redisStorage := storage.NewRedisShortLinks(rdb)
	api := api.New(redisStorage)
	router := mux.NewRouter()
	router.HandleFunc("/cache/link/create", api.CreateShortLink)
	router.HandleFunc("/cache/link/redirect", api.RedirectToOriginal)
	http.ListenAndServe("localhost:8082", router)
}
