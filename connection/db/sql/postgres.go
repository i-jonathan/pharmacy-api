package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/i-jonathan/pharmacy-api/config"
	_ "github.com/lib/pq"
)

type repo struct {
	Conn *sql.DB
}

func NewDBConnection() (*repo, error) {
	dbConfig := config.GetConfig()
	dbRepo := new(repo)
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPass)

	dbRepo.Conn, err = sql.Open("postgres", dsn)

	if err != nil {
		return &repo{}, errors.New("an error occurred while trying to start the db up: " + err.Error())
	}
	return dbRepo, nil
}
