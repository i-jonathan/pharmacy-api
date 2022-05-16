package controller

import (
	"encoding/json"
	"fmt"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/interface/mux/helper"
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/service"
	"log"
	"net/http"
)

type authController struct {
	svc service.AuthUseCase
}

func NewAuthController(s service.AuthUseCase) *authController {
	return &authController{s}
}

func (controller *authController) Login(w http.ResponseWriter, r *http.Request) {
	var auth model.Auth
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.ServerError)
		return
	}

	token, err := controller.svc.SignIn(auth)

	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, map[string]string{"Token": token})
}

func (controller *authController) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	claims, err := model.ParseToken(token)
	if err != nil {
		log.Println(err)
		if err == appError.Unauthorized {
			helper.ReturnFailure(w, err)
			return
		}
		helper.ReturnFailure(w, appError.ServerError)
		return
	}

	err = controller.svc.Logout(fmt.Sprintf("%s", claims["hash"]), token)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnDelete(w)
}
