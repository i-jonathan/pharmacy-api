package router

import (
	"github.com/i-jonathan/pharmacy-api/interface/mux/controller"
	"github.com/i-jonathan/pharmacy-api/service"
	"net/http"
)

func InitAuthRouter(svc service.AuthUseCase) {
	authController := controller.NewAuthController(svc)
	authRouter := router.PathPrefix("auth").Subrouter()

	authRouter.HandleFunc("/login", authController.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/logout", authController.Logout).Methods(http.MethodPost)
}
