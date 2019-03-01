package main

import (
	"fmt"
	"github.com/yukihirai0505/free-service/api/conf"
	"log"
	"net/http"
	"os"
)

// mongodb://yabaiwebyasan:<PASSWORD>@cluster0-shard-00-00-lrlto.mongodb.net:27017,cluster0-shard-00-01-lrlto.mongodb.net:27017,cluster0-shard-00-02-lrlto.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true
func Handler(w http.ResponseWriter, r *http.Request) {
	conf.LoadEnv()
	log.Println(os.Getenv("MONGO_USER"))
	fmt.Fprintf(w, "I'm hoge")
}
