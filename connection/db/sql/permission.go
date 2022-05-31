package db

import (
	"database/sql"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (r *repo) FetchPermissions() ([]model.Permission, error) {
	var result []model.Permission

	const query = "SELECT id, name, description, created_at FROM permission;"

	rows, err := r.Conn.Query(query)

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	for rows.Next() {
		var perm model.Permission
		if err := rows.Scan(&perm.ID, &perm.Name, &perm.Description, &perm.CreatedAt); err != nil {
			return nil, appError.ServerError
		}
		perm.Slug, err = model.ToHashID(perm.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, perm)
	}

	if err := rows.Err(); err != nil {
		return nil, appError.ServerError
	}

	return result, nil
}

func (r *repo) FetchPermissionByID(id int) (model.Permission, error) {
	var result model.Permission

	const query = "SELECT id, name, description, created_at FROM permission WHERE id = $1;"
	row := r.Conn.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return model.Permission{}, err
	}

	err := row.Scan(&result.ID, &result.Name, &result.Description, &result.CreatedAt)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return model.Permission{}, appError.NotFound
		default:
			return model.Permission{}, err
		}
	}
	result.Slug, err = model.ToHashID(result.ID)
	if err != nil {
		log.Println(err)
	}

	return result, nil
}

func (r *repo) CreatePermission(permission model.Permission) (int, error) {
	const statement = "INSERT INTO permission (name, description) VALUES ($1, $2) returning id;"
	var id int
	err := r.Conn.QueryRow(statement, permission.Name, permission.Description).Scan(&id)

	if err != nil || id < 1{
		return 0, err
	}

	return int(id), nil
}

func (r *repo) UpdatePermission(permission model.Permission) error {
	const statement = "UPDATE permission SET name = $1, description = $2 WHERE id = $3;"

	_, err := r.Conn.Exec(statement, permission.Name, permission.Description, permission.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeletePermission(id int) error {
	const statement = "DELETE FROM permission WHERE id = $1;"

	_, err := r.Conn.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}