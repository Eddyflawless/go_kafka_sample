package appkafka

import (
	"time"
)

type Kafka struct {
	topic string
	groupId string
	minBytes int //
	maxBytes int //
	batchSize int
	batchTimeout time.Duration
	brokers *[]string
}


func New(topic string, groupId string, minBytes int, maxBytes int, batchSize int, batchTimeout time.Duration, brokers *[]string) *Kafka {

	return &Kafka{

		topic: topic,
		groupId: groupId,
		minBytes: minBytes,
		maxBytes: maxBytes,
		batchSize: batchSize,
		batchTimeout: batchTimeout,
		brokers: brokers,

	}
}