package request

// BookOrder struct type
type BookOrder struct {
	UserID          int64  `json:"users_id"`
	BookID          int64  `json:"books_id"`
	Quantity        int    `json:"quantity"`
	ShippingDetails string `json:"shipping_details"`
}
