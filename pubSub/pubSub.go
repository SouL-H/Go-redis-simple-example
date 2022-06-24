package pubsub

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	
)

var ctx = context.Background()

var rdb *redis.Client


var channelName = "report"

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic("Redise bağlanılamadı => " + err.Error())
	}

	go subscriber(1)
	go subscriber(2)
	go publisher()


	c := make(chan int)
	<-c
}

func publisher() {
	for range time.Tick(time.Second * 3) {
		// Zaman değişkenini oluşturduk
		t := time.Now().Format("15:04:05")
		fmt.Println("*************************")
		fmt.Println("Kanala gönderilen => " + t)
		// Zaman değişkenini kanala publish ettik
		rdb.Publish(ctx, channelName, t)
	}
}


func subscriber(subscriberNumber int) {
	subs := rdb.Subscribe(ctx, channelName)

	for msg := range subs.Channel() {

		fmt.Println(fmt.Sprintf("Subscribe %d için kanaldan okunan => %s", subscriberNumber, msg.Payload))
	}
}