package appkafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"fmt"
	"time"
)

func  (k *Kafka) Consumer(ctx context.Context){

	conf := kafka.ReaderConfig{
		Brokers: []string{ "localhost:9892","localhost:9893", "localhost:9894"},
		Topic: "mytopic", //
		GroupID: "my-group", //
		MinBytes: 5, //
		MaxBytes: 10,
		// the kafka library requires you to set the MaxBytes
		// in case the MinBytes are set
		MaxWait: 3 * time.Second,
		// this will start consuming messages from the earliest available
		StartOffset: kafka.FirstOffset,
	}

	reader := kafka.NewReader(conf)

	for { 

		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}

		fmt.Println("Message: ", string(message.Value))

	}

}