package service

import (
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/repository"
	"log"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{r}
}

func (a *authService) Logout(hash, token string) error {
	done, err := a.repo.BlacklistToken(hash, token)

	if !done {
		log.Println(err)
		return appError.ServerError
	}

	return nil
}
