package main

import (
	"log"
	// "os"
	"github.com/joho/godotenv"
	// "time"
	"gin-rest-api/lib/kafka"
)


func main(){

	//RunCron()
	
	go appkafka.StartKafka()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Some error occrured. Err: %s", err)
	}

}