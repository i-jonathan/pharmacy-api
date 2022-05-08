package model

import "time"

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   Account   `json:"created_by"`
	UserID      int       `json:"user_id"`
	Slug        string    `json:"slug"`
}

type Supplier struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	Slug        string    `json:"slug"`
}

type Product struct {
	// TODO change price to decimal 
	ID                int        `json:"id"`
	Name              string     `json:"name"`
	BarCode           string     `json:"bar_code"`
	Description       string     `json:"description"`
	Category          Category   `json:"category"`
	CategoryID        int        `json:"category_id"`
	ExpiryDate        time.Time  `json:"expiry_date"`
	PurchaseDate      time.Time  `json:"purchase_date"`
	ProductionDate    time.Time  `json:"production_date"`
	PurchasePrice     int        `json:"purchase_price"`
	SellingPrice      int        `json:"selling_price"`
	QuantityAvailable int        `json:"quantity_available"`
	ReorderLevel      int        `json:"reorder_level"`
	SKU               string     `json:"sku"`
	QuantitySold      int        `json:"quantity_sold"`
	Supplier          []Supplier `json:"supplier"`
	CreatedBy         Account    `json:"created_by"`
	UserID            int        `json:"user_id"`
	CreatedAt         time.Time  `json:"created_at"`
	Slug              string     `json:"slug"`
}
