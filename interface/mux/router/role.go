package router

import (
	"github.com/i-jonathan/pharmacy-api/interface/mux/controller"
	"github.com/i-jonathan/pharmacy-api/service"
	"net/http"
)

func InitRoleRouter(svc service.RoleUseCase) {
	roleController := controller.NewRoleController(svc)
	roleRouter := router.PathPrefix("/role").Subrouter()

	roleRouter.HandleFunc("", roleController.CreateRole).Methods(http.MethodPost)
	roleRouter.HandleFunc("", roleController.FetchRoles).Methods(http.MethodGet)
	roleRouter.HandleFunc("/{slug}", roleController.FetchRoleBySlug).Methods(http.MethodGet)
	roleRouter.HandleFunc("/{slug}", roleController.UpdateRole).Methods(http.MethodPut)
	roleRouter.HandleFunc("/{slug}", roleController.DeleteRole).Methods(http.MethodDelete)
}
