package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main(){
	clinet := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 10,
	})	
	pong,err:=clinet.Ping().Result()//ping redis
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(pong)
	
}
