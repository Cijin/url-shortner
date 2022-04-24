package store

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)

func InitializeStore() *StorageService {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Panic(fmt.Sprintf("Error Redis Init: %s", err))
	}

	log.Print(fmt.Sprintf("Redis Ping Response: %s", pong))

	storageService.redisClient = client
	return storageService
}

/*
 * Persist shortURL -> URL mapping in Redis
 */
func SaveURL(shortURL, URL string) {

}

/*
 * Retrieve original URL given shortURL
 */
func GetURL(shortURL string) string {
	return ""
}
