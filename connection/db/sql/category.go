package db

import (
	"database/sql"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
	"log"
)

func (r *repo) CreateCategory(category model.Category) (int, error) {
	const query = "INSERT INTO category (name, description, user_id) VALUES ($1, $2, $3) RETURNING ID;"

	var id int
	err := r.Conn.QueryRow(query, category.Name, category.Description, category.UserID).Scan(&id)

	if err != nil || id < 1 {
		return 0, err
	}

	return id, nil
}

func (r *repo) FetchCategories() ([]model.Category, error) {
	const query = "SELECT id, name, description, user_id FROM category;"
	var result []model.Category
	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp model.Category
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.UserID)
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

	if err = rows.Err(); err != nil {
		return nil, appError.ServerError
	}

	return result, nil
}

func (r *repo) FetchCategoryByID(id int) (model.Category, error) {
	var category model.Category
	const query = "SELECT id, name, description, user_id FROM category WHERE id=$1;"

	row := r.Conn.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return model.Category{}, err
	}

	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.UserID)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return model.Category{}, appError.NotFound
		default:
			return model.Category{}, err
		}
	}

	category.Slug, err = model.ToHashID(category.ID)
	if err != nil {
		log.Println(err)
	}

	return category, nil
}

func (r *repo) UpdateCategory(category model.Category) error {
	const query = "UPDATE category SET name = $1, description = $2 WHERE id=$3;"

	_, err := r.Conn.Exec(query, category.Name, category.Description, category.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteCategory(id int) error {
	const query = "DELETE FROM category WHERE id = $1;"
	_, err := r.Conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
