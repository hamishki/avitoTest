package main

import (
	"avito/worker"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func getRedisClient(addr, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}

func main() {
	client := getRedisClient("localhost:6379", "pass")
	ctx := context.Background()
	db, err := sql.Open("pgx", "user=gopher password=pass host=localhost port=5432 database=test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	worker := worker.Worker{db, client, time.Duration(time.Minute * 5)}
	// запускаем обработчик
}
