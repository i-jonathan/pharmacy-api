package model

import "time"

type PaymentMethod struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedBy Account   `json:"created_by"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Slug      string    `json:"slug"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	Product   Product `json:"product"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
}

type Order struct {
	// TODO change price to decimal
	ID              int           `json:"id"`
	OrderItem       []OrderItem   `json:"order_item"`
	TotalPrice      int           `json:"total_price"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	PaymentMethodID int           `json:"payment_method_id"`
	PaymentVerified bool          `json:"payment_verified"`
	Cashier         Account       `json:"cashier"`
	CashierID       int           `json:"cashier_id"`
	AmountTendered  int           `json:"amount_tendered"`
	Change          int           `json:"change"`
	CreatedAt       time.Time     `json:"created_at"`
}
