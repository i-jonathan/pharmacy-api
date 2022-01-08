package inventory

import (
	"Pharmacy/core"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm/clause"
)

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var products []product
	var count int64
	db := initDatabase()
	db.Model(&product{}).Count(&count)
	db.Scopes(core.Paginate(r)).Preload(clause.Associations).Find(&products)

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	page, prev, next := core.ResponseData(int(count), r)
	response := core.Response{
		Previous: prev,
		Next:     next,
		Page:     page,
		Count:    count,
		Data:     products,
	}
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(err)
	}
	return
}
