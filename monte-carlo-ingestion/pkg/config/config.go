package config

type Config struct {
	KafkaBroker string
	KafkaTopic  string
}

func LoadConfig() *Config {
	return &Config{
		KafkaBroker: "kafka:9092",
		KafkaTopic:  "pii-gdpr-annotations",
	}
}
