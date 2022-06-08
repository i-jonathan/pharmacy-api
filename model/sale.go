package model

type PaymentMethod struct {
	baseModel
	Name      string  `json:"name"`
	IsActive  bool    `json:"is_active"`
	CreatedBy Account `json:"created_by"`
	UserID    int     `json:"user_id"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	Product   Product `json:"product"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderID   int     `json:"order_id"`
	ReturnID  int     `json:"return_id"`
}

type Order struct {
	baseModel
	// TODO change price to decimal
	OrderItem       []OrderItem   `json:"order_item"`
	TotalPrice      int           `json:"total_price"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	PaymentMethodID int           `json:"payment_method_id"`
	PaymentVerified bool          `json:"payment_verified"`
	Cashier         Account       `json:"cashier"`
	CashierID       int           `json:"cashier_id"`
	AmountTendered  int           `json:"amount_tendered"`
	Change          int           `json:"change"`
}

type Return struct {
	baseModel
	Reason    string      `json:"reason"`
	Item      []OrderItem `json:"item"`
	Order     Order       `json:"order"`
	OrderID   int         `json:"order_id"`
	CreatedBy Account     `json:"created_by"`
	UserID    int         `json:"user_id"`
}
