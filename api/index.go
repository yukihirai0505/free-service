package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	mongoDBUser := os.Getenv("MONGO_USER")
	mongoDBPass := os.Getenv("MONGO_PASS")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + mongoDBUser + ":" + mongoDBPass + "@cluster0-qqdp9.mongodb.net/test"))
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
	fmt.Println(mongoDBPass)
	fmt.Fprintf(w, "Hello")
}
