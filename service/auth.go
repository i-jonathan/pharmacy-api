package service

import (
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/repository"
	"log"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{r}
}

func (a *authService) SignIn(auth model.Auth) (string, error) {
	account, err := a.repo.FetchAccountWithPassword(auth)

	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return "", err
		}
		return "", appError.ServerError
	}

	valid, err := auth.ComparePassword(account.Password)
	if err != nil || !valid {
		log.Print(err)
		return "", appError.Unauthorized
	}

	token, err := account.CreateToken()
	if err != nil {
		log.Println(err)
		return "", appError.ServerError
	}
	return token, nil
}

func (a *authService) Logout(hash, token string) error {
	done, err := a.repo.BlacklistToken(hash, token)

	if !done {
		log.Println(err)
		return appError.ServerError
	}

	return nil
}
