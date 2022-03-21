package main

import (
	"log"
	"context"
	"time"
	"fmt"
	"github.com/joho/godotenv"
	"gin-rest-api/lib/kafka"
)


func main(){

	go func() {
		fmt.Println("runs in the background")
	}()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Some error occrured. Err: %s", err)
	}

	//create a new context 
	ctx := context.Background()

	// produce and consume functions are 
	// blocking
	// so lets put the messages in a go routine

	//"localhost:9892","localhost:9893", "localhost:9894" 
	
	brokers := &[]string {"localhost:29092" }
	k_instance := appkafka.New("message-log", "my-group", 5, 10, 10, time.Second * 4, brokers)

	go k_instance.Producer(ctx)
	k_instance.Consumer(ctx)


}