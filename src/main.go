package main

import (
	"log"
	"context"
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

	go appkafka.Producer(ctx)
	appkafka.Consumer(ctx)

}