package db

import (
	"database/sql"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (r *repo) CreateProduct(product model.Product) (int, error) {
	const query = "INSERT INTO product (name, bar_code, description, category_id, purchase_date, production_date, expiry_date, purchase_price, selling_price, quantity_available, reorder_level, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id;"

	var id int
	err := r.Conn.QueryRow(query, product.Name, product.BarCode, product.Description, product.CategoryID, product.PurchaseDate, product.ProductionDate, product.ExpiryDate, product.PurchasePrice, product.SellingPrice, product.QuantityAvailable, product.ReorderLevel, product.UserID).Scan(&id)

	if err != nil || id < 1 {
		return 0, err
	}
	return id, nil
}

func (r *repo) FetchProducts() ([]model.Product, error) {
	const query = "SELECT id, name, bar_code, description, category_id, purchase_date, production_date, expiry_date, purchase_price, selling_price, quantity_available, reorder_level, quantity_sold, user_id FROM product;"
	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	var result []model.Product
	for rows.Next() {
		var temp model.Product
		err = rows.Scan(&temp.ID, &temp.Name, &temp.BarCode, &temp.Description, &temp.CategoryID, &temp.PurchaseDate, &temp.ProductionDate, &temp.ExpiryDate, &temp.PurchasePrice, &temp.SellingPrice, &temp.QuantityAvailable, &temp.ReorderLevel, &temp.QuantitySold, &temp.UserID)
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

func (r *repo) FetchProductByID(id int) (model.Product, error) {
	const query = "SELECT id, name, bar_code, description, category_id, purchase_date, production_date, expiry_date, purchase_price, selling_price, quantity_available, reorder_level, quantity_sold, user_id FROM product where id=$1;"

	row := r.Conn.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return model.Product{}, err
	}
	var result model.Product
	err := row.Scan(&result.ID, &result.Name, &result.BarCode, &result.Description, &result.CategoryID,
		&result.PurchaseDate, &result.ProductionDate, &result.ExpiryDate, &result.PurchasePrice, &result.SellingPrice,
		&result.QuantityAvailable, &result.ReorderLevel, &result.QuantitySold, &result.UserID)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return model.Product{}, appError.NotFound
		default:
			return model.Product{}, err
		}
	}

	result.Slug, err = model.EncodeIDToSlug(result.ID)
	if err != nil {
		log.Println(err)
	}

	return result, nil
}

func (r *repo) UpdateProduct(product model.Product) error {
	const query = "UPDATE product SET name = $1, bar_code = $2, description = $3, category_id = $4, purchase_date = $5, production_date = $6, expiry_date = $7, purchase_price = $8, selling_price = $9, quantity_available = $10, reorder_level = $11, quantity_sold = $12 WHERE id = $13;"
	_, err := r.Conn.Exec(query, product.Name, product.BarCode, product.Description, product.CategoryID, product.PurchaseDate, product.ProductionDate, product.ExpiryDate, product.PurchasePrice, product.SellingPrice, product.QuantityAvailable, product.ReorderLevel, product.QuantitySold, product.ID)

	if err != nil {
		return err
	}
	return err
}

func (r *repo) DeleteProduct(id int) error {
	const query = "DELETE FROM product WHERE id=$1;"
	_, err := r.Conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
