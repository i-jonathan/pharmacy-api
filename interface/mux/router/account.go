package router

import (
	"github.com/i-jonathan/pharmacy-api/interface/mux/controller"
	"github.com/i-jonathan/pharmacy-api/service"
	"net/http"
)

func InitAccountRouter(svc service.AccountUseCase) {
	accountController := controller.NewAccountController(svc)
	accountRouter := router.PathPrefix("/account").Subrouter()

	accountRouter.HandleFunc("", accountController.CreateAccount).Methods(http.MethodPost)
	accountRouter.HandleFunc("", accountController.FetchAccounts).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{slug}", accountController.FetchAccountBySlug).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{slug}", accountController.UpdateAccount).Methods(http.MethodPut)
	accountRouter.HandleFunc("/{slug}", accountController.DeleteAccount).Methods(http.MethodDelete)
}
