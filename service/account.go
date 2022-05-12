package service

import (
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/repository"
	"log"
)

type accountService struct {
	repo repository.PharmacyRepository
}

func NewAccountService(r repository.PharmacyRepository) *accountService {
	return &accountService{r}
}

// Permissions

func (service *accountService) FetchPermissions() ([]model.Permission, error) {
	result, err := service.repo.FetchPermissions()
	if err != nil {
		return nil, appError.ServerError
	}

	return result, nil
}

func (service *accountService) FetchPermissionBySlug(slug string) (model.Permission, error) {
	id, err := model.DecodeID(slug)

	if err != nil {
		return model.Permission{}, appError.BadRequest
	}

	result, err := service.repo.FetchPermissionByID(id)
	if err != nil {
		if err == appError.NotFound {
			return model.Permission{}, err
		}
		return model.Permission{}, appError.ServerError
	}

	return result, nil
}

func (service *accountService) CreatePermission(permission model.Permission) (model.Permission, error) {
	panic("implement me")
}

func (service *accountService) UpdatePermission(permission model.Permission) (model.Permission, error) {
	panic("implement me")
}

func (service *accountService) DeletePermission(slug string) error {
	panic("implement me")
}

// Roles

func (service *accountService) FetchRoles() ([]model.Role, error) {
	panic("implement me")
}

func (service *accountService) FetchRoleBySlug(slug string) (model.Role, error) {
	panic("implement me")
}

func (service *accountService) CreateRole(role model.Role) (model.Role, error) {
	panic("implement me")
}

func (service *accountService) UpdateRole(role model.Role) error {
	panic("implement me")
}

func (service *accountService) DeleteRole(slug string) error {
	panic("implement me")
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
	id, err := model.DecodeID(slug)
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

	if !valid {
		return model.Account{}, appError.BadRequest
	}

	result, err := service.repo.CreateAccount(account)
	if err != nil {
		log.Println("Error during account: ", err)
		return model.Account{}, appError.ServerError
	}
	account.ID = result
	account.Password = ""
	return account, nil
}

func (service *accountService) UpdateAccount(account model.Account) (model.Account, error) {
	id, err := model.DecodeID(account.Slug)
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
	id, err := model.DecodeID(s)
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
