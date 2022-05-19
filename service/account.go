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
	valid := permission.Valid()
	if !valid {
		return model.Permission{}, appError.BadRequest
	}

	result, err := service.repo.CreatePermission(permission)
	if err != nil {
		log.Println(err)
		return model.Permission{}, appError.ServerError
	}

	permission.ID = result
	permission.Slug, err = model.ToHashID(permission.ID)
	if err != nil {
		log.Println(err)
	}
	return permission, nil
}

func (service *accountService) UpdatePermission(permission model.Permission) (model.Permission, error) {
	if !permission.Valid()  {
		return model.Permission{}, appError.BadRequest
	}

	var err error
	permission.ID, err = model.DecodeID(permission.Slug)
	if err != nil {
		log.Println(err)
		return model.Permission{}, appError.BadRequest
	}

	err = service.repo.UpdatePermission(permission)

	if err != nil {
		log.Println(err)
		return model.Permission{}, appError.ServerError
	}

	return permission, nil
}

func (service *accountService) DeletePermission(slug string) error {
	id, err := model.DecodeID(slug)
	if err != nil {
		log.Println(err)
		return appError.BadRequest
	}

	_, err = service.repo.FetchPermissionByID(id)

	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return err
		}
		return appError.ServerError
	}

	err = service.repo.DeletePermission(id)
	if err != nil {
		log.Println(err)
		return appError.ServerError
	}

	return nil
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
	account.Slug, err = model.ToHashID(account.ID)
	if err != nil {
		log.Println(err)
	}
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

	token, err := account.CreateToken()
	if err != nil {
		log.Println(err)
		return "", appError.ServerError
	}
	return token, nil
}