package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("learning redis with Go")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	setVal := client.Set(ctx, "name", "alex", 0).Err()
	if setVal != nil {
		fmt.Println(setVal)
	}

	getVal, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getVal)

	json, err := json.Marshal(Author{Name: "alex", Age: 12})
	if err != nil {
		fmt.Println(err)
	}

	setJson := client.Set(ctx, "id1", json, 0).Err()
	if setJson != nil {
		fmt.Println(setJson)
	}

	val, err := client.Get(ctx, "id1").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}
