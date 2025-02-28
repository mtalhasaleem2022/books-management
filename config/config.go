package config

import "os"

type Config struct {
	Port         string
	DBUrl        string
	RedisUrl     string
	KafkaBrokers string
}

func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		DBUrl:        getEnv("DB_URL", "postgres://user:pass@localhost:5432/books"),
		RedisUrl:     getEnv("REDIS_URL", "redis://localhost:6379/0"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
