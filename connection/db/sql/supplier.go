package db

import (
	"database/sql"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (r *repo) CreateSupplier(supplier model.Supplier) (int, error) {
	const query = "INSERT INTO supplier (name, address, phone_number, email) VALUES ($1, $2, $3, $4) RETURNING id;"

	var id int
	err := r.Conn.QueryRow(query, supplier.Name, supplier.Address, supplier.PhoneNumber, supplier.Email).Scan(&id)

	if err != nil || id < 0 {
		return 0, err
	}
	return id, nil
}

func (r *repo) FetchSuppliers() ([]model.Supplier, error) {
	const query = "SELECT id, name, address, phone_number, email FROM supplier;"
	var result []model.Supplier

	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp model.Supplier
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Address, &temp.PhoneNumber, &temp.Email)
		if err != nil {
			log.Println(err)
			continue
		}
		temp.Slug, err = model.EncodeIDToSlug(temp.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, temp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repo) FetchSupplierByID(id int) (model.Supplier, error) {
	const query = "SELECT id, name, address, phone_number, email FROM supplier WHERE id = $1;"

	row := r.Conn.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return model.Supplier{}, err
	}
	var result model.Supplier
	err := row.Scan(&result.ID, &result.Name, &result.Address, &result.PhoneNumber, &result.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return model.Supplier{}, appError.NotFound
		default:
			return model.Supplier{}, err
		}
	}

	result.Slug, err = model.EncodeIDToSlug(result.ID)
	if err != nil {
		log.Println(err)
	}

	return result, err
}

func (r *repo) UpdateSupplier(supplier model.Supplier) error {
	const query = "UPDATE supplier SET name = $1, address = $2, phone_number = $3, email = $4 WHERE id=$5;"

	_, err := r.Conn.Exec(query, supplier.Name, supplier.Address, supplier.PhoneNumber, supplier.Email, supplier.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteSupplier(id int) error {
	const query = "DELETE FROM supplier WHERE id = $1;"
	_, err := r.Conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
