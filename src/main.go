package main

import (
	"log"
	"context"
	"time"
	"github.com/joho/godotenv"
	"gin-rest-api/lib/kafka"
)


func main(){

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Some error occrured. Err: %s", err)
	}

	//create a new context 
	ctx := context.Background()

	// produce and consume functions are 
	// blocking
	// so lets put the messages in a go routine

	brokers := &[]string {"localhost:9892","localhost:9893", "localhost:9894" }
	k_instance := appkafka.New("message-log", "my-group", 5, 10, 10, time.Second * 4, brokers)

	go k_instance.Producer(ctx)
	k_instance.Consumer(ctx)


}