package service

import (
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (service *accountService) FetchRoles() ([]model.Role, error) {
	result, err := service.repo.FetchRoles()
	if err != nil {
		log.Println(err)
		return nil, appError.ServerError
	}

	return result, nil
}

func (service *accountService) FetchRoleBySlug(slug string) (model.Role, error) {
	id, err := model.DecodeID(slug)
	if err != nil {
		log.Println(err)
		return model.Role{}, appError.BadRequest
	}

	result, err := service.repo.FetchRoleByID(id)
	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return model.Role{}, err
		}
		return model.Role{}, appError.ServerError
	}
	return result, nil
}

func (service *accountService) CreateRole(role model.Role) (model.Role, error) {
	if !role.Valid() {
		return model.Role{}, appError.BadRequest
	}

	result, err := service.repo.CreateRole(role)
	if err != nil {
		log.Println(err)
		return model.Role{}, appError.ServerError
	}

	role.ID = result
	role.Slug, err = model.ToHashID(result)
	if err != nil {
		log.Println(err)
	}

	return role, nil
}

func (service *accountService) UpdateRole(role model.Role) error {
	var err error
	role.ID, err = model.DecodeID(role.Slug)
	if err != nil {
		log.Println(err)
		return appError.BadRequest
	}
	_, err = service.repo.FetchRoleByID(role.ID)
	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return appError.NotFound
		}
		return appError.BadRequest
	}

	if !role.Valid() {
		return appError.BadRequest
	}

	err = service.repo.UpdateRole(role)
	if err != nil {
		return appError.ServerError
	}

	return nil
}

func (service *accountService) DeleteRole(slug string) error {
	id, err := model.DecodeID(slug)
	if err != nil {
		log.Println(err)
		return appError.BadRequest
	}
	_, err = service.repo.FetchRoleByID(id)
	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return appError.NotFound
		}
		return appError.BadRequest
	}

	err = service.repo.DeleteRole(id)
	if err != nil {
		log.Println(err)
		return appError.ServerError
	}

	return nil
}
