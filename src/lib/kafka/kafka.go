package appkafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"fmt"
)

func StartKafka(){

	conf := kafka.ReaderConfig{
		Brokers: []string{ "localhost:9892"},
		Topic: "mytopic", //
		GroupID: "g1", //
		MaxBytes: 10,
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