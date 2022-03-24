package db

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func RedisInitialize() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	dbnm := os.Getenv("REDIS_DB")
	pass := os.Getenv("REDIS_PASS")

	dbNum, err := strconv.Atoi(dbnm)
	if err != nil {
		log.Fatalln("database redis wrong")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       dbNum,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatalln("redis fail to run")
	}
	log.Printf("[INIT] Successfully connected to Redis Database\n")
	return client
}
