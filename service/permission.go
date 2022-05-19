package service

import (
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

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
	var err error
	permission.ID, err = model.DecodeID(permission.Slug)
	if err != nil {
		log.Println(err)
		return model.Permission{}, appError.BadRequest
	}

	_, err = service.repo.FetchPermissionByID(permission.ID)
	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return model.Permission{}, appError.NotFound
		}
		return model.Permission{}, appError.BadRequest
	}

	if !permission.Valid() {
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
