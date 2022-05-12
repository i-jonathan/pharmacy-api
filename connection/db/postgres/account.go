package db

import (
	"database/sql"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (r *repo) FetchPermissions() ([]model.Permission, error) {
	var result []model.Permission

	query := "SELECT id, name, description, created_at FROM permission;"

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

	query := "SELECT id, name, description, created_at FROM permission WHERE id = ?;"
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
	return result, nil
}

func (r *repo) CreatePermission(permission model.Permission) (int, error) {
	statement := "INSERT INTO permission (name, description) VALUES ?, ?;"
	resp, err := r.Conn.Exec(statement, permission.Name, permission.Description)

	if err != nil {
		return 0, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repo) UpdatePermission(permission model.Permission) error {
	statement := "UPDATE permission SET name = ?, description = ? WHERE id = ?;"

	_, err := r.Conn.Exec(statement, permission.Name, permission.Description, permission.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeletePermission(id int) error {
	statement := "DELETE FROM permission WHERE id = ?;"

	_, err := r.Conn.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}

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

func (r *repo) FetchAccounts() ([]model.Account, error) {
	var result []model.Account
	statement := "SELECT id, first_name, last_name, email, phone_number, role_id, role.name, created_at FROM account INNER JOIN role ON account.role_id=role.id;"

	rows, err := r.Conn.Query(statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp model.Account
		err = rows.Scan(&temp.ID, &temp.FirstName, &temp.LastName, &temp.Email, &temp.PhoneNumber, &temp.RoleID, &temp.Role.Name, &temp.CreatedAt)
		if err != nil {
			log.Println(err)
			continue
		}
		temp.Slug, err = model.ToHashID(temp.ID)
		temp.Role.Slug, err = model.ToHashID(temp.RoleID)
		if err != nil {
			log.Println(err)
			continue
		}

		result = append(result, temp)
	}

	if err := rows.Err(); err != nil {
		return nil, appError.ServerError
	}

	return result, nil
}

func (r *repo) FetchAccountByID(id int) (model.Account, error) {
	var result model.Account
	query := "SELECT id, first_name, last_name, email, phone_number, role_id, role.name, created_at FROM account INNER JOIN role ON account.role_id=role.id WHERE id=?;"

	row := r.Conn.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return model.Account{}, err
	}

	err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.PhoneNumber,
		result.RoleID, &result.Role.Name, &result.CreatedAt)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return model.Account{}, appError.NotFound
		default:
			return model.Account{}, err
		}
	}

	return result, nil
}

func (r *repo) CreateAccount(account model.Account) (int, error) {
	statement := "INSERT INTO account (first_name, last_name, email, password, phone_number, role_id) VALUES ?, ?, ?, ?, ?, ?;"

	resp, err := r.Conn.Exec(statement, account.FirstName, account.LastName, account.Email, account.Password, account.PhoneNumber, account.RoleID)
	if err != nil {
		return 0, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

func (r *repo) UpdateAccount(account model.Account) error {
	statement := "UPDATE account SET first_name = ?, last_name = ?, email = ?, phone_number = ?, role_id = ?;"

	_, err := r.Conn.Exec(statement, account.FirstName, account.LastName, account.Email, account.PhoneNumber, account.Role.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteAccount(id int) error {
	statement := "DELETE FROM account WHERE id = ?"

	_, err := r.Conn.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}
