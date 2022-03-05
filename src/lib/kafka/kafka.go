package appkafka

import (
	"github.com/segmentio/kafka-go"
	// "fmt"
)

type Kafka struct {
	topic string
	groupId string
	minBytes int //
	maxBytes int //
	batchSize int
	batchTimeout int
	brokers *[]string{}
}

func GetBrokers(){

	
}
func New(topic string, groupId string, minBytes int, maxBytes int, batchSize int, batchTimeout int, brokers *[]string){

	return &Kafka{

	}
}