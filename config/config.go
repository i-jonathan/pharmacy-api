package config

import "os"

type Configuration struct {
	DBName     string
	DBPass     string
	DBHost     string
	DBUser     string
	DBProtocol string
	HashSalt   string
	HMAC       string
}

var config Configuration

func LoadEnv() {
	config = Configuration{
		DBName:     os.Getenv("database_name"),
		DBPass:     os.Getenv("database_pass"),
		DBHost:     os.Getenv("database_host"),
		DBUser:     os.Getenv("database_user"),
		DBProtocol: os.Getenv("database_protocol"),
		HashSalt:   os.Getenv("hash_salt"),
		HMAC:       os.Getenv("hmac"),
	}
}

func GetConfig() Configuration {
	return config
}
