package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	DBHost          string
	DBUsername      string
	DBPassword      string
	DBPort          string
	NatsClientPort  int64
	NatsMonitorPort int64
}

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	cfg = &Config{
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBUsername:      getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres"),
		DBPort:          getEnv("DB_PORT", "5432"),
		NatsClientPort:  getEnvAsInt("NATS_CLIENT_PORT", 4222),
		NatsMonitorPort: getEnvAsInt("NATS_MONITOR_PORT", 8222),
	}

	return nil
}

func GetConfig() *Config {
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		log.Printf("Erro ao converter %s para int, usando default: %d", key, fallback)
		return fallback
	}
	return value
}
