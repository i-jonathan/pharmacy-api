package db

import "github.com/i-jonathan/pharmacy-api/model"

func (r *repo) FetchRoles() ([]model.Role, error) {
	panic("implement me")
}

func (r *repo) FetchRoleByID(id int) (model.Role, error) {
	panic("implement me")
}

func (r *repo) CreateRole(role model.Role) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateRole(role model.Role) error {
	panic("implement me")
}

func (r *repo) DeleteRole(id int) error {
	panic("implement me")
}