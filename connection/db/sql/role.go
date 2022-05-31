package db

import (
	"database/sql"
	"fmt"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
	"strings"
)

func (r *repo) FetchRoles() ([]model.Role, error) {
	var result []model.Role
	const statement = "SELECT id, name, description, created_at FROM role;"

	rows, err := r.Conn.Query(statement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp model.Role
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.CreatedAt)
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
	const statement = "SELECT id, name, description, created_at FROM role WHERE id = $1;"
	row := r.Conn.QueryRow(statement, id)

	var err error
	if err = row.Err(); err != nil {
		log.Println(err)
		return model.Role{}, err
	}

	err = row.Scan(&result.ID, &result.Name, &result.Description, &result.CreatedAt)
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

	// Fetch permissions
	const query = "SELECT id, name, description FROM role_permission LEFT JOIN permission p ON p.id = role_permission.permission_id WHERE role_id=$1;"
	rows, err := r.Conn.Query(query, id)

	if err != nil {
		log.Println(err)
		return result, nil
	}

	for rows.Next() {
		var tempPerm model.Permission
		err = rows.Scan(&tempPerm.ID, &tempPerm.Name, &tempPerm.Description)
		if err != nil {
			log.Println(err)
			continue
		}

		tempPerm.Slug, err = model.ToHashID(tempPerm.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		result.Permissions = append(result.Permissions, tempPerm)
	}

	if err = rows.Err(); err != nil {
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

	if len(role.Permissions) > 0 {
		valueString := make([]string, 0, len(role.Permissions))
		values := make([]interface{}, 0, len(role.Permissions)*2)
		i := 0
		for _, permission := range role.Permissions {
			permission.ID, err = model.DecodeID(permission.Slug)
			if err != nil {
				log.Println(err)
				continue
			}

			valueString = append(valueString, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
			values = append(values, id)
			values = append(values, permission.ID)
			i++
		}
		query := "INSERT INTO role_permission (role_id, permission_id) VALUES " + strings.Join(valueString, ",")
		_, err = r.Conn.Exec(query, values...)
		if err != nil {
			log.Println(err)
		}
	}
	return id, nil
}

func (r *repo) UpdateRole(role model.Role) error {
	const statement = "UPDATE role SET name = $1, description = $2 WHERE id = $3;"

	_, err := r.Conn.Exec(statement, role.Name, role.Description, role.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteRole(id int) error {
	const statement = "DELETE FROM role WHERE id = $1;"

	_, err := r.Conn.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}
