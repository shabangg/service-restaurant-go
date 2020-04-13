package kafka

type Kafka struct {
	Consumer *Consumer
	Producer *Producer
}

func New() (kafka *Kafka, err error) {

	kafka = &Kafka{}

	err, config := initConfig()
	if err != nil {
		return nil, err
	}

	err, kafka.Consumer = initConsumer(config)
	if err != nil {
		return nil, err
	}

	err, kafka.Producer = initProducer(config)
	if err != nil {
		return nil, err
	}

	return kafka, nil

}
