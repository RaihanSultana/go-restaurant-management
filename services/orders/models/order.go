package models

import "github.com/shopspring/decimal"

type Order struct {
	ID         int32
	CustomerID int32
	ItemID     int32
	Quantity   int32
	Total      decimal.Decimal
	Subtotal   decimal.Decimal
}

type OrderItem struct {
	ID      int32
	OrderID int32
	ItemID  int32
	Total   decimal.Decimal
}
