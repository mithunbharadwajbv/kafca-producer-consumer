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
		mes, _ := conn.ReadMessage(100)
		Log(string(mes.Value))
	}
}
