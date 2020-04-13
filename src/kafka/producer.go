package kafka

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	EventProducer sarama.SyncProducer
}

func initProducer(config *Config) (error, *Producer) {

	if config.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	if config.Brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(config.Brokers, ",")
	// log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	sConfig := sarama.NewConfig()
	sConfig.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	sConfig.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	sConfig.Producer.Return.Successes = true

	sProducer, err := sarama.NewSyncProducer(brokerList, sConfig)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	log.Printf("Successfully init kafa producer")

	producer := &Producer{
		EventProducer: sProducer,
	}

	return nil, producer
}

type PublishData struct {
	EventType    string    `json:"event_type"`
	Id           string    `json:"id"`
	FullDocument string    `json:"full_document"`
	Timestamp    time.Time `json:"timestamp"`

	encoded []byte
	err     error
}

func (data *PublishData) ensureEncoded() {
	if data.encoded == nil && data.err == nil {
		data.encoded, data.err = json.Marshal(data)
	}
}

func (data *PublishData) Length() int {
	data.ensureEncoded()
	return len(data.encoded)
}

func (data *PublishData) Encode() ([]byte, error) {
	data.ensureEncoded()
	return data.encoded, data.err
}

func (k *Kafka) PublishMessage(data *PublishData) error {

	topic := "restaurants"

	data.Timestamp = time.Now()

	message := sarama.ProducerMessage{
		Topic: topic,
		Value: data,
	}

	partiotion, offset, err := k.Producer.EventProducer.SendMessage(&message)
	if err != nil {
		return err
	}

	log.Println("Message pushed at partition and offset - > ", partiotion, " ", offset)

	return nil

}
