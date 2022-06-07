package db

import (
	"database/sql"
	"log"

	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
)

func (r *repo) FetchAccounts() ([]model.Account, error) {
	var result []model.Account
	const statement = "SELECT account.id, first_name, last_name, email, phone_number, account.created_at, coalesce(role_id, 0) as role_id, coalesce(role.name, '') as role_name FROM account Left JOIN role ON account.role_id=role.id;"

	rows, err := r.Conn.Query(statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp model.Account
		err = rows.Scan(&temp.ID, &temp.FirstName, &temp.LastName, &temp.Email, &temp.PhoneNumber, &temp.CreatedAt, &temp.RoleID, &temp.Role.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		temp.Slug, _ = model.EncodeIDToSlug(temp.ID)
		temp.Role.ID = temp.RoleID
		temp.Role.Slug, err = model.EncodeIDToSlug(temp.RoleID)
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
	const query = "SELECT account.id, first_name, last_name, email, phone_number, coalesce(role_id, 0) as role_id, coalesce(role.name, '') as role_name, account.created_at FROM account LEFT JOIN role ON account.role_id=role.id WHERE account.id=$1;"

	row := r.Conn.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return model.Account{}, err
	}

	err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.PhoneNumber,
		&result.RoleID, &result.Role.Name, &result.CreatedAt)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return model.Account{}, appError.NotFound
		default:
			return model.Account{}, err
		}
	}
	result.Slug, _ = model.EncodeIDToSlug(result.ID)
	result.Role.Slug, err = model.EncodeIDToSlug(result.RoleID)
	if err != nil {
		log.Println(err)
	}
	return result, nil
}

func (r *repo) FetchAccountWithPassword(auth model.Auth) (model.Account, error) {
	var result model.Account
	const query = "SELECT id, email, password, role_id from account where email=$1;"

	err := r.Conn.QueryRow(query, auth.Email).Scan(&result.ID, &result.Email, &result.Password, &result.RoleID)

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
	var phoneExists bool
	var emailExists bool
	// TODO return apperror conflict
	row := r.Conn.QueryRow("select exists(select 1 from account where phone_number=$1);", account.PhoneNumber)
	if err := row.Scan(&phoneExists); err != nil {
		return 0, err
	}

	row = r.Conn.QueryRow("select exists(select 1 from account where email=$1);", account.Email)
	if err := row.Scan(&emailExists); err != nil {
		return 0, err
	}

	if phoneExists || emailExists {
		return 0, appError.BadRequest
	}

	const statement = "INSERT INTO account (first_name, last_name, email, password, phone_number) VALUES ($1, $2, $3, $4, $5) returning id;"
	var id int
	err := r.Conn.QueryRow(statement, account.FirstName, account.LastName, account.Email, account.Password, account.PhoneNumber).Scan(&id)

	if err != nil || id < 1 {
		return 0, err
	}

	return id, nil
}

func (r *repo) UpdateAccount(account model.Account) error {
	const statement = "UPDATE account SET first_name = $1, last_name = $2, email = $3, phone_number = $4, role_id = $5 WHERE id = $6;"
	var role interface{}

	if account.RoleID == 0 {
		role = sql.NullInt64{}
	} else {
		role = account.RoleID
	}

	_, err := r.Conn.Exec(statement, account.FirstName, account.LastName, account.Email, account.PhoneNumber, role, account.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteAccount(id int) error {
	const statement = "DELETE FROM account WHERE id = $1;"

	_, err := r.Conn.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}
