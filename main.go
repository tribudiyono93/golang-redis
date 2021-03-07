package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	fmt.Println("Go Redis tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "2e9374e970ebba4b325aa721f9add70c7c76f8fa9fcd0e0d6e557068b6bad6d2",
		DB: 0,
		MaxRetries: 3,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	//set data to redis
	if err = client.Set(ctx, "name", "Elliot", 0).Err(); err != nil {
		fmt.Println(err)
	}

	//get data from redis
	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	client.Close()
}