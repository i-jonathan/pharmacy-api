package config

import "os"

type Configuration struct {
	DBName string
	DBPass string
	DBHost string
	DBUser string
	DBProtocol string
}

var config Configuration

func LoadEnv() {
	config = Configuration{
		DBName:     os.Getenv("database_name"),
		DBPass:     os.Getenv("database_pass"),
		DBHost:     os.Getenv("database_host"),
		DBUser:     os.Getenv("database_user"),
		DBProtocol: os.Getenv("database_protocol"),
	}
}

func GetConfig() Configuration {
	return config
}