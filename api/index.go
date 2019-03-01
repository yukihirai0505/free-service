package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"
)

func testMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASS") + "@cluster0-qqdp9.mongodb.net/test"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("test").Collection("test")
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID
	fmt.Println(id)
}

func testRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-10054.c13.us-east-1-3.ec2.cloud.redislabs.com:10054",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	client.Set("hello", "world", 0)
	val, err := client.Get("hello").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hello", val)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//testMongo()
	testRedis()
	fmt.Fprintf(w, "Hello")
}
