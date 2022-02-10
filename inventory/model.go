package inventory

import (
	"Pharmacy/account"
	"Pharmacy/core"
	"time"

	"github.com/shopspring/decimal"
)

type category struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DateCreated time.Time    `json:"date_created"`
	Creator     account.User `json:"created_by,-"`
	CreatorID   int          `json:"creator_id,-"`
}

type product struct {
	ID             uint            `json:"id"`
	ItemName       string          `json:"name"`
	BarCode        string          `json:"bar_code"`
	Description    string          `json:"description"`
	Category       category        `json:"category"`
	CategoryID     int             `json:"category_id"`
	ExpiryDate     string          `json:"expiry_date"`
	PurchaseDate   string          `json:"purchase_date"`
	ProductionDate string          `json:"production_date"`
	Quantity       uint            `json:"quantity"`
	PurchasePrice  decimal.Decimal `json:"purchase_price"`
	SellingPrice   decimal.Decimal `json:"selling_price"`
	User           account.User    `json:"user,-"`
	UserID         int             `json:"user_id,-"`
	ReorderLevel   int             `json:"reorder_level"`
	SKU            string          `json:"sku"`
	QuantitySold   int             `json:"quantity_sold"`
}

type saleData struct {
	ID       uint   `json:"id"`
	BarCode  string `json:"bar_code"`
	SKU      string `json:"sku"`
	Quantity uint   `json:"quantity"`
}

func init() {
	core.RegisterModel(&product{})
	core.RegisterModel(&category{})
}

//func initDatabase() *gorm.DB {
//	name := os.Getenv("database_name")
//	pass := os.Getenv("database_pass")
//	user := os.Getenv("database_user")
//	host := os.Getenv("database_host")
//	port := os.Getenv("databse_port")
//	ssl := os.Getenv("databse_ssl")
//
//	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
//		host, port, user, name, pass, ssl)
//
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Println("err")
//		return nil
//	}
//
//	err = db.AutoMigrate(&category{}, &product{})
//	if err != nil {
//		log.Println(err)
//	}
//
//	return db
//}
