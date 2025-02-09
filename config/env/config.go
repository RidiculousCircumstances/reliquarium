package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config структура для хранения конфигурационных параметров
type Config struct {
	DatabaseDSN                         string
	LogPath                             string
	RabbitMQURL                         string
	LogLevel                            string
	StorageEndpoint                     string
	SnapshotBucket                      string
	StorageAccessKeyId                  string
	StorageSecretAccessKey              string
	RedisAddr                           string
	RedisPassword                       string
	RedisDB                             int
	HttpPort                            string
	ReliquariumCommandExchangeName      string
	ReliquariumConfirmationExchangeName string
	ReliquariumConfirmationQueueName    string
}

// LoadConfig загружает конфигурацию из .env файла
func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}

	return &Config{
		DatabaseDSN:                         getEnv("DATABASE_DSN", ""),
		LogPath:                             getEnv("LOG_PATH", ""),
		RabbitMQURL:                         getEnv("RABBITMQ_URL", ""),
		LogLevel:                            getEnv("LOG_LEVEL", "INFO"),
		StorageEndpoint:                     getEnv("STORAGE_ENDPOINT", ""),
		StorageAccessKeyId:                  getEnv("STORAGE_ACCESS_KEY_ID", ""),
		StorageSecretAccessKey:              getEnv("STORAGE_SECRET_ACCESS_KEY", ""),
		RedisAddr:                           getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:                       getEnv("REDIS_PASSWORD", ""),
		RedisDB:                             getEnvAsInt("REDIS_DB", 0),
		HttpPort:                            getEnv("HTTP_PORT", ":80"),
		ReliquariumCommandExchangeName:      os.Getenv("RELIQUARIUM_COMMAND_EXCHANGE_NAME"),
		ReliquariumConfirmationExchangeName: os.Getenv("RELIQUARIUM_CONFIRMATION_EXCHANGE_NAME"),
		ReliquariumConfirmationQueueName:    os.Getenv("RELIQUARIUM_CONFIRMATION_QUEUE_NAME"),
	}, nil
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt возвращает значение переменной окружения как int или значение по умолчанию
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
