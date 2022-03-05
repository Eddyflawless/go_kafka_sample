package appkafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strconv"
	"fmt"
	"time"
)

const (
	topic          = "message-log"
)

func GetBrokerAddresses() []string {

	return []string{ "localhost:9892","localhost:9893", "localhost:9894"}
}

func Producer(ctx context.Context) { 

	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: GetBrokerAddresses(),
		Topic: topic,
		// wait until we get 10 messages before writing
		BatchSize: 10,
		// no matter what happens, write all pending messages
		// every 2 seconds
		BatchTimeout: 2 * time.Second,
	})

	for { 
		// each kafka message has a key and value . The key is used 
		// to decide which partion (and consequently , which broker)
		// the message gets published on

		key := strconv.Itoa(i)

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(key),
			//create an arbitrary message payload for the value 
			Value: []byte("this is message " + key ),
		})

		if err != nil {
			panic("couldnot write message " + err.Error())
		}

		// confirmation Message
		fmt.Println("writes: ", i)

		i++
		time.Sleep(time.Second)

	}

}