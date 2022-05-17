package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Configuration struct {
	DBName       string
	DBPass       string
	DBHost       string
	DBUser       string
	DBPort       string
	DBProtocol   string
	HashSalt     string
	HMAC         string
	RedisAddress string
	RedisPort    string
}

var config Configuration

func LoadEnv() {

	_ = godotenv.Load()

	config = Configuration{
		DBName:       os.Getenv("database_name"),
		DBPass:       os.Getenv("database_pass"),
		DBHost:       os.Getenv("database_host"),
		DBUser:       os.Getenv("database_user"),
		DBPort:       os.Getenv("database_port"),
		DBProtocol:   os.Getenv("database_protocol"),
		HashSalt:     os.Getenv("hash_salt"),
		HMAC:         os.Getenv("hmac"),
		RedisAddress: os.Getenv("redis_address"),
		RedisPort:    os.Getenv("redis_port"),
	}
}

func GetConfig() Configuration {
	return config
}
