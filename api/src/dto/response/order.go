package response

// Order struct type
type Order struct {
	ID              int64  `json:"id"`
	UserName        string `json:"user_name"`
	BookName        string `json:"book_name"`
	Quantity        int    `json:"quantity"`
	ShippingDetails string `json:"shipping_details"`
	CreateAt        string `json:"created_at"`
}
