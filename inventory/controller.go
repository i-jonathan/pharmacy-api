package inventory

import (
	"Pharmacy/core"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

//Get all Items
//@Summary      Get items
//@Description  Get all items in the inventory
//@Tags         inventory
//@Produce      json
//@Success      200  {object}  core.Response{[]data=product}
//@Success      204  {object}  core.Response{[]data=product}
//@Router       /inventory/all [get]
func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

//Get Item
//@Summary Get item
//@Description Get a single item by id
//@Tags inventory
//@Produce json
//@Param id path int true "id"
//@Success 200 {object} product
//@Failure 404 {object} core.ErrorResponse
//@Router /inventory/{id} [get]
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemId := mux.Vars(r)["id"]
	_, err := strconv.Atoi(itemId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Check your ID and try again"})
		return
	}

	var item product
	db := initDatabase()
	db.Preload(clause.Associations).Find(&item, "id = ?", itemId)

	if item.ItemName == "" {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Resource not found"})
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
	}
}

//Post Item
//@Summary Post Item
//@Description Add a new item
//@Tags inventory
//@Accept json
//@Produce json
//@Success 200 {object} product
//@Failure 400 {object} core.ErrorResponse
//@Router /inventory/add [post]
func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")
	var item product
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		return
	}
	//TODO needs logged in User

	if item.ItemName == "" || item.BarCode == "" || item.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Name, Barcode and Description can not be empty"})
		return
	}

	if item.ExpiryDate.IsZero() || item.PurchaseDate.IsZero() || item.ProductionDate.IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Expiry, Purchase or Production Date can not be empty."})
		return
	}

	if item.PurchasePrice == 0 || item.SellingPrice == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Price can not be Zero (0)"})
		return
	}

	// TODO generate SKU

	// create and save item into db
	db.Create(&item)
	err = json.NewEncoder(w).Encoder(item)
	return
}

//Make Sale
//@Summary Sell Item
//@Description Sell an Item
//@Tags inventory
//@Accept json
//@Produce json
//@Success 200 {object} product
//@Failure
//@Router /inventory/sell-item [post]
func sellItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item product

}