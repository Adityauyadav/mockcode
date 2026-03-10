package queue

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	fmt.Println("Redis Connected Succesfuly")

}

func Push(submissionID string) error {

	return Client.LPush(context.Background(), "submission_queue", submissionID).Err()

}

func Pop() (string, error) {
	result, err := Client.BRPop(context.Background(), 0, "submission_queue").Result()
	if err != nil {
		return "", err
	}
	return result[1], nil

}
