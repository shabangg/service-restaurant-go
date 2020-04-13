package kafka

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Brokers  string
	Version  string
	GroupId  string
	Topics   string
	Assignor string
	Oldest   bool
	Verbose  bool
}

func initConfig() (error, *Config) {

	config := &Config{
		Brokers:  viper.GetString("kafka_brokers"),  // Kafka bootstrap brokers to connect to, as a comma separated list
		Version:  viper.GetString("kafka_version"),  // Kafka cluster version e.g. 2.1.1
		GroupId:  viper.GetString("kafka_group_id"), // Kafka consumer group definition
		Topics:   viper.GetString("kafka_topics"),   // Kafka topics to be consumed, as a comma seperated list
		Assignor: viper.GetString("kafka_assignor"), // Consumer group partition assignment strategy (range, roundrobin, sticky)
		Oldest:   viper.GetBool("kafka_olders"),     // Kafka consumer consume initial offset from oldest
		Verbose:  viper.GetBool("kafka_verbose"),    // Sarama logging
	}

	if config.Brokers == "" {
		return fmt.Errorf("no Kafka bootstrap brokers defined, please set the -brokers flag"), nil
	}

	if config.Topics == "" {
		return fmt.Errorf("no topics given to be consumed, please set the -topics flag"), nil
	}

	if config.GroupId == "" {
		return fmt.Errorf("no Kafka consumer group defined, please set the -group flag"), nil
	}

	return nil, config

}
