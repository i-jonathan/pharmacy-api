package router

import (
	"github.com/i-jonathan/pharmacy-api/interface/mux/controller"
	"github.com/i-jonathan/pharmacy-api/service"
	"net/http"
)

func InitPermissionRouter(svc service.PermissionUseCase) {
	permissionController := controller.NewPermissionController(svc)

	permissionRouter := router.PathPrefix("/permissions").Subrouter()

	permissionRouter.HandleFunc("", permissionController.FetchPermissions).Methods(http.MethodGet)
	permissionRouter.HandleFunc("", permissionController.CreatePermission).Methods(http.MethodPost)
	permissionRouter.HandleFunc("/{slug}", permissionController.FetchPermissionBySlug).Methods(http.MethodGet)
	permissionRouter.HandleFunc("/{slug}", permissionController.UpdatePermission).Methods(http.MethodPut)
	permissionRouter.HandleFunc("/{slug}", permissionController.DeletePermission).Methods(http.MethodDelete)
}
