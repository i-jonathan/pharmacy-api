package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/i-jonathan/pharmacy-api/config"
)

type repo struct {
	Conn *sql.DB
}

func NewDBConnection() (*repo, error) {
	dbConfig := config.GetConfig()
	fmt.Println(dbConfig)
	dbRepo := new(repo)
	var err error

	//dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbProtocol, dbConfig.DbIP, dbConfig.DbPort, dbConfig.DbName)
	dsn := ""
	fmt.Println(dsn)

	dbRepo.Conn, err = sql.Open("postgres", dsn)

	if err != nil {
		return &repo{}, errors.New("an error occurred while trying to start the db up: " + err.Error())
	}
	return dbRepo, nil
}