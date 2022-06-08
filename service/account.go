package service

import (
	"log"

	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/repository"
)

type accountService struct {
	repo repository.PharmacyRepository
}

func NewAccountService(r repository.PharmacyRepository) *accountService {
	return &accountService{r}
}

// Accounts

func (service *accountService) FetchAccounts() ([]model.Account, error) {
	result, err := service.repo.FetchAccounts()
	if err != nil {
		return nil, appError.ServerError
	}

	return result, nil
}

func (service *accountService) FetchAccountBySlug(slug string) (model.Account, error) {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return model.Account{}, appError.BadRequest
	}

	result, err := service.repo.FetchAccountByID(id)

	if err != nil {
		if err == appError.NotFound {
			log.Println(err)
			return model.Account{}, err
		}
		return model.Account{}, appError.ServerError
	}

	return result, nil
}

func (service *accountService) CreateAccount(account model.Account) (model.Account, error) {
	valid := account.Valid()
	err := account.HashPassword()

	if err != nil {
		log.Println(err)
		return model.Account{}, appError.ServerError
	}
	if !valid {
		return model.Account{}, appError.BadRequest
	}
	result, err := service.repo.CreateAccount(account)

	if err != nil {
		log.Println("Error during account: ", err)
		if err == appError.BadRequest {
			return model.Account{}, err
		}
		return model.Account{}, appError.ServerError
	}
	account.ID = result
	account.Slug, err = model.EncodeIDToSlug(account.ID)
	if err != nil {
		log.Println(err)
	}
	account.Password = ""
	return account, nil
}

func (service *accountService) UpdateAccount(account model.Account) (model.Account, error) {
	id, err := model.DecodeSlugToID(account.Slug)
	if err != nil {
		log.Println(err)
		return model.Account{}, appError.BadRequest
	}

	account.ID = id
	_, err = service.repo.FetchAccountByID(account.ID)
	if err != nil {
		if err == appError.NotFound {
			log.Println(err)
			return model.Account{}, err
		}
		return model.Account{}, appError.ServerError
	}

	// set password to spoof valid function
	// TODO consider a dedicated password checker
	account.Password = "someText"
	if !account.Valid() {
		return model.Account{}, appError.BadRequest
	}

	err = service.repo.UpdateAccount(account)

	if err != nil {
		return model.Account{}, appError.ServerError
	}
	return account, nil
}

func (service *accountService) DeleteAccount(s string) error {
	id, err := model.DecodeSlugToID(s)
	if err != nil {
		return appError.BadRequest
	}
	_, err = service.repo.FetchAccountByID(id)
	if err != nil {
		if err == appError.NotFound {
			log.Println(err)
			return err
		}
		return appError.ServerError
	}

	err = service.repo.DeleteAccount(id)
	if err != nil {
		log.Println(err)
		return appError.ServerError
	}
	return nil
}

func (service *accountService) SignIn(auth model.Auth) (string, error) {
	account, err := service.repo.FetchAccountWithPassword(auth)

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

	if account.RoleID > 0 {
		account.Role, err = service.repo.FetchRoleByID(account.RoleID)
		if err != nil {
			log.Println(err)
			return "", appError.ServerError
		}
	}

	token, err := account.CreateToken()
	if err != nil {
		log.Println(err)
		return "", appError.ServerError
	}
	return token, nil
}
