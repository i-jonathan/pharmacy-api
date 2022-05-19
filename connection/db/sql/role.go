package db

import (
	"database/sql"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (r *repo) FetchRoles() ([]model.Role, error) {
	var result []model.Role
	const statement = "SELECT id, name, created_at FROM role;"

	rows, err := r.Conn.Query(statement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp model.Role
		err = rows.Scan(&temp.ID, &temp.Name, &temp.CreatedAt)
		if err != nil {
			log.Println(err)
			continue
		}
		temp.Slug, err = model.ToHashID(temp.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, temp)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return result, nil
}

func (r *repo) FetchRoleByID(id int) (model.Role, error) {
	var result model.Role
	const statement = "SELECT id, name, created_at FROM role WHERE id = $1;"
	row := r.Conn.QueryRow(statement, id)
	
	var err error
	if err = row.Err(); row != nil {
		return model.Role{}, err
	}
	
	err = row.Scan(&result.ID, &result.Name, &result.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Role{}, appError.NotFound
		}
		return model.Role{}, err
	}
	result.Slug, err = model.ToHashID(result.ID)
	if err != nil {
		log.Println(err)
	}
	return result, nil

}

func (r *repo) CreateRole(role model.Role) (int, error) {
	var id int
	const statement = "INSERT INTO role (name, description) VALUES ($1, $2) returning id;"

	err := r.Conn.QueryRow(statement, role.Name, role.Description).Scan(&id)

	if err != nil || id < 1 {
		return 0, err
	}

	return id, nil
}

func (r *repo) UpdateRole(role model.Role) error {
	const statement = "UPDATE role SET name = $1, description = $2 WHERE id = $3;"

	_, err := r.Conn.Exec(statement, role.Name, role.Description, role.ID)
	if err != nil {
		return err
	}

	return  nil
}

func (r *repo) DeleteRole(id int) error {
	const statement = "DELETE FROM role WHERE id = $1;"

	_, err := r.Conn.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}