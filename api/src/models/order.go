package models

import "time"

// Order struct type
type Order struct {
	ID              int64
	UserName        string
	BookName        string
	Quantity        int
	ShippingDetails string
	CreateAt        time.Time
}
