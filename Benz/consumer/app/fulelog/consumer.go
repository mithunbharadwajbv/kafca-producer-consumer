package fulelog

import (
	"consumer/config"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var (
	kafcaHost = config.Conf.KafcaHost
	topic     = config.Conf.Topic
	partition = config.Conf.Partition
)

// Works as a broker and listen to all events posted by the producer and forward for logging in a file
func FuelLogConsumer() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafcaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("failed to close connection:", err)
		}
	}()

	for {
		mes, err := conn.ReadMessage(100)
		if err != nil {
			continue
		}
		Log(string(mes.Value))
	}
}
