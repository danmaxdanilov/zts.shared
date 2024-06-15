package messaging

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"os"

	"github.com/danmaxdanilov/zts.shared/pkg/logger"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type IProducer interface {
}

type producer struct {
	logger      logger.Logger
	config      Config
	topicConfig TopicConfig
	w           *kafka.Writer
}

func NewProducer(
	logger logger.Logger,
	config Config,
	topicConfig TopicConfig,
) *producer {
	return &producer{
		logger:      logger,
		config:      config,
		topicConfig: topicConfig,
		w:           NewWriter(logger, config, topicConfig),
	}
}

func NewWriter(logger logger.Logger, config Config, topicConfig TopicConfig) *kafka.Writer {
	// Create a custom dial function
	dial := func(ctx context.Context, network, address string) (net.Conn, error) {
		return newDialer(logger, config).DialContext(ctx, network, address)
	}

	w := &kafka.Writer{
		Addr:         kafka.TCP(config.Brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  logger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		Async:        false,
		Topic:        topicConfig.TopicName,
		Transport:    &kafka.Transport{Dial: dial},
	}
	return w
}

func newDialer(logger logger.Logger, config Config) *kafka.Dialer {
	switch config.SecurityProtocol {
	case "PLAINTEXT":
		return &kafka.Dialer{
			Timeout: dialTimeout,
			// DualStack: true,
		}
	case "SSL":
		return &kafka.Dialer{
			Timeout: dialTimeout,
			// DualStack: true,
			TLS: newTlsConfig(logger, config),
		}
	case "SASL_SSL":
		saslMechanism, err := scram.Mechanism(scram.SHA512, config.SaslConfig.SaslUserName, config.SaslConfig.SaslPassword)
		if err != nil {
			logger.Fatalf("Failed to create SASL mechanism: %s", err)
		}
		return &kafka.Dialer{
			Timeout: dialTimeout,
			// DualStack:     true,
			TLS:           newTlsConfig(logger, config),
			SASLMechanism: saslMechanism,
		}
	default:
		return &kafka.Dialer{}
	}
}

func newTlsConfig(logger logger.Logger, config Config) *tls.Config {
	// Load client certificate
	// var cert tls.Certificate
	// var err error

	// if len(config.SSLConfig.SslPrivateKeyLocation) == 0 &&
	// 	len(config.SSLConfig.SslPublicKeyLocation) == 0 {
	// 	logger.Debug("SSL Public and Private Key are empty")
	// } else {
	// 	cert, err = tls.LoadX509KeyPair("service.cert", "service.key")
	// 	if err != nil {
	// 		logger.Fatalf("Failed to load client certificate: %s", err)
	// 	}
	// }

	// Load CA certificate
	caCert, err := os.ReadFile(config.SSLConfig.SslCaLocation)
	if err != nil {
		logger.Fatalf("Failed to read CA certificate: %s", err)
	}

	// Create a CA certificate pool and add the CA certificate
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		logger.Fatalf("Failed to append CA certificate to pool")
	}

	return &tls.Config{
		// Certificates: []tls.Certificate{cert},
		RootCAs: caCertPool,
	}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.w.Close()
}
