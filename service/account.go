package service

import (
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/repository"
)

type accountService struct {
	repo repository.PharmacyRepository
}

func NewAccountService(r repository.PharmacyRepository) *accountService {
	return &accountService{r}
}

// Roles

func (a *accountService) FetchRoles() ([]model.Role, error) {
	panic("implement me")
}

func (a *accountService) FetchRoleBySlug(slug string) (model.Role, error) {
	panic("implement me")
}

func (a *accountService) CreateRole(role model.Role) (model.Role, error) {
	panic("implement me")
}

func (a *accountService) UpdateRole(role model.Role) error {
	panic("implement me")
}

func (a *accountService) DeleteRole(slug string) error {
	panic("implement me")
}

// Permissions

func (a *accountService) FetchPermissions() ([]model.Permission, error) {
	panic("implement me")
}

func (a *accountService) FetchPermissionBySlug(slug string) (model.Permission, error) {
	panic("implement me")
}

func (a *accountService) CreatePermission(permission model.Permission) (model.Permission, error) {
	panic("implement me")
}

func (a *accountService) UpdatePermission(permission model.Permission) (model.Permission, error) {
	panic("implement me")
}

func (a *accountService) DeletePermission(slug string) error {
	panic("implement me")
}

// Accounts

func (a *accountService) FetchAccounts() ([]model.Account, error) {
	panic("implement me")
}

func (a *accountService) FetchAccountBySlug(s string) (model.Account, error) {
	panic("implement me")
}

func (a *accountService) CreateAccount(account model.Account) (model.Account, error) {
	panic("implement me")
}

func (a *accountService) UpdateAccount(account model.Account) (model.Account, error) {
	panic("implement me")
}

func (a *accountService) DeleteAccount(s string) error {
	panic("implement me")
}