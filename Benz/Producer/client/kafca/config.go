package kafca

import (
	"context"
	"log"

	"github.com/mithun/config"
	"github.com/segmentio/kafka-go"
)

var (
	producer *kafka.Conn

	topic     = config.Conf.Topic
	partition = config.Conf.Partition
	kafcaport = config.Conf.KafcaHost
)

func init() {

	var err error
	producer, err = kafka.DialLeader(context.Background(), "tcp", kafcaport, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
}
