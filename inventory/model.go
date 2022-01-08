package inventory

import (
	"Pharmacy/account"
	"time"

	"github.com/shopspring/decimal"
)

type category struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DateCreated time.Time    `json:"date_created"`
	Creator     account.User `json:"created_by"`
	CreatorID   int          `json:"creator_id"`
}

type inventory struct {
	ItemName       string          `json:"name"`
	BarCode        string          `json:"bar_code"`
	Description    string          `json:description`
	Category       category        `json:"category"`
	CategoryID     int             `json:"category_id"`
	ExpiryDate     time.Time       `json:"expiry_date"`
	PurchaseDate   time.Time       `json"purchase_date"`
	ProductionDate time.Time       `json:"production_date"`
	Quantity       int             `json:"quantity"`
	PurchasePrice  decimal.Decimal `json:"purchase_price"`
	SellingPrice   decimal.Decimal `json:"selling_price"`
	User           account.User    `json:"user"`
	UserID         int             `json:"user_id"`
	ReorderLevel   int             `json:"reorder_level"`
	SKU            string          `json:"sku"`
	QuantitySold   int             `json:"quantity_sold"`
}
