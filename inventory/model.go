package inventory

import (
	"Pharmacy/account"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type category struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DateCreated time.Time    `json:"date_created"`
	Creator     account.User `json:"created_by"`
	CreatorID   int          `json:"creator_id"`
}

type product struct {
	ID             uint            `json:"id"`
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

func initDatabase() *gorm.DB {
	name := os.Getenv("database_name")
	pass := os.Getenv("database_pass")
	user := os.Getenv("database_user")
	host := os.Getenv("database_host")
	port := os.Getenv("databse_port")
	ssl := os.Getenv("databse_ssl")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, name, pass, ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("err")
		return nil
	}

	err = db.AutoMigrate(&category{}, &product{})
	if err != nil {
		log.Println(err)
	}

	return db
}
