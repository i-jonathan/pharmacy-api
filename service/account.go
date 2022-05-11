package service

import (
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
	panic("implement me")
}

func (service *accountService) FetchAccountBySlug(s string) (model.Account, error) {
	panic("implement me")
}

func (service *accountService) CreateAccount(account model.Account) (model.Account, error) {
	panic("implement me")
}

func (service *accountService) UpdateAccount(account model.Account) (model.Account, error) {
	panic("implement me")
}

func (service *accountService) DeleteAccount(s string) error {
	panic("implement me")
}
