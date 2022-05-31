package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/i-jonathan/pharmacy-api/config"
	noSql "github.com/i-jonathan/pharmacy-api/connection/db/redis"
	db "github.com/i-jonathan/pharmacy-api/connection/db/sql"
	"github.com/i-jonathan/pharmacy-api/interface/mux/router"
	"github.com/i-jonathan/pharmacy-api/service"
)

func main() {
	config.LoadEnv()
	mainRouter := router.InitRouter()

	repo, err := db.NewDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	noSqlRepo, err := noSql.NewRedisConnection()
	if err != nil {
		log.Fatalln(err)
	}

	accountService := service.NewAccountService(repo)
	router.InitPermissionRouter(accountService)
	router.InitAccountRouter(accountService)
	router.InitRoleRouter(accountService)

	authService := service.NewAuthService(noSqlRepo)
	router.InitAuthRouter(authService)

	log.Println("Starting Server...")
	if err := http.ListenAndServe(":9576", handlers.CORS(handlers.AllowedOrigins([]string{"localhost:7060"}), handlers.AllowCredentials(), handlers.OptionStatusCode(http.StatusOK))(mainRouter)); err != nil {
		log.Panicln(err)
	}
}
