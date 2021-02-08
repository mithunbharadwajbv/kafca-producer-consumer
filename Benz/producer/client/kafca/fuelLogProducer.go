package kafca

import (
	"log"
)

type Kafca interface {
	Publish(string) error
}

type kafca struct{}

func GetNewKafca() Kafca {
	return &kafca{}
}

// publish data to consumer
func (k *kafca) Publish(val string) error {
	_, err := producer.Write([]byte(val))
	if err != nil {
		log.Println("Error while writing data to the partition")
		panic(err)
	}
	log.Printf("Data sent : %s \n", val)
	return nil
}
