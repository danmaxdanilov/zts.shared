package kafka

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/danmaxdanilov/zts.shared/pkg/logger"
	"github.com/segmentio/kafka-go"
)

// var cfg = &logger.Config{}
// var log = logger.NewAppLogger(cfg)

func TestKafka(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		cfg := &logger.Config{}
		logger := logger.NewAppLogger(cfg)
		logger.InitLogger()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cg := NewConsumerGroup(
			[]string{"master.1150c21b-4e55-46bb-ab1d-32642fse61486.c.dbaas.selcloud.ru:9093"},
			"kafka-group",
			logger,
		)

		cg.ConsumeTopic(ctx, []string{"kafka-test"}, 0, KafkaWorker)
	})
}

func KafkaWorker(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("Worker %d: error reading message: %v", workerID, err)
			return
		}
		log.Printf("Worker %d: message at offset %d: %s = %s\n", workerID, m.Offset, string(m.Key), string(m.Value))
	}
}
