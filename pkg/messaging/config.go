package messaging

// Config kafka config
type Config struct {
	Brokers          []string   `mapstructure:"brokers"`
	GroupID          string     `mapstructure:"groupID"`
	InitTopics       bool       `mapstructure:"initTopics"`
	SecurityProtocol string     `mapstructure:"securityProtocol"`
	SSLConfig        SSLConfig  `mapstructure:"sSLConfig"`
	SaslConfig       SaslConfig `mapstructure:"SaslConfig"`
}

type SSLConfig struct {
	SslPrivateKeyLocation string `mapstructure:"sslPrivateKeyLocation"`
	SslPublicKeyLocation  string `mapstructure:"sslPublicKeyLocation"`
	SslCaLocation         string `mapstructure:"sslCaLocation"`
}

type SaslConfig struct {
	SaslUserName string `mapstructure:"saslUserName"`
	SaslPassword string `mapstructure:"saslPassword"`
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName         string `mapstructure:"topicName"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"`
}
