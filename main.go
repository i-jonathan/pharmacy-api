package main

import (
	"github.com/i-jonathan/pharmacy-api/config"
	db "github.com/i-jonathan/pharmacy-api/connection/db/postgres"
	"github.com/i-jonathan/pharmacy-api/interface/mux/router"
	"github.com/i-jonathan/pharmacy-api/service"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	mainRouter := router.InitRouter()

	repo, err := db.NewDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	accountService := service.NewAccountService(repo)
	router.InitPermissionRouter(accountService)
	router.InitAccountRouter(accountService)

	log.Println("Starting Server...")
	if err := http.ListenAndServe(":9576", mainRouter); err != nil {
		log.Panicln(err)
	}
}
