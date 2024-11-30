package kafka

import (
	"backend/config"
	infrastructure "backend/infrastructure/log"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaConsumer(log infrastructure.LogCustom) *kafka.Consumer {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.AppConfig.Kafka.Host,
		"group.id":          config.AppConfig.Kafka.GroupId,
		"auto.offset.reset": config.AppConfig.Kafka.Reset,
	}
	consumer, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		log.Logrus.WithField("Error connection kafka consumer : ", err).Error("Failed to create consumer")
	}
	return consumer
}

func NewKafkaProducer(log infrastructure.LogCustom) *kafka.Producer {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.AppConfig.Kafka.Host,
	}

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Logrus.WithField("Error connection kafka consumer : ", err).Error("Failed to create consumer")
	}

	return producer
}
