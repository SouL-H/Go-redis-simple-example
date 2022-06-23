package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       10,
	})
	pong, err := client.Ping().Result() //ping redis
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
	err = client.Set("username", "user1", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	//Get Value
	username, err := client.Get("username").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(username, "is the username")
}
