package messaging

import (
	"context"
	"testing"

	"github.com/danmaxdanilov/zts.shared/pkg/logger"
	"github.com/segmentio/kafka-go"
)

func TestKafkaProducer(t *testing.T) {
	t.Run("write message", func(t *testing.T) {
		cfg := &logger.Config{}
		logger := logger.NewAppLogger(cfg)
		logger.InitLogger()

		config := &Config{
			Brokers:          []string{"master.1150c21b-4e55-46bb-ab1d-32642fe61486.c.dbaas.selcloud.ru:9093"},
			GroupID:          "simple",
			InitTopics:       false,
			SecurityProtocol: "SASL_SSL",
			SSLConfig: SSLConfig{
				SslPrivateKeyLocation: "",
				SslPublicKeyLocation:  "",
				SslCaLocation:         "./root.crt",
			},
			SaslConfig: SaslConfig{
				SaslUserName: "colleen",
				SaslPassword: "hci3aSBFzDa2",
			},
		}
		topicConfig := &TopicConfig{
			TopicName: "kafka-test",
			// Partitions: 0,
		}
		p := NewProducer(logger, *config, *topicConfig)

		err := p.PublishMessage(context.Background(),
			kafka.Message{
				Key:   []byte("Key"),
				Value: []byte("Hello, Kafka!"),
				// Partition: 0,
			})

		if err != nil {
			logger.Fatalf("Failed to write message: %s", err)
		}

		logger.Info("Message written successfully")

		// Close the writer
		if err := p.Close(); err != nil {
			logger.Fatalf("Failed to close writer: %s", err)
		}
	})
}
