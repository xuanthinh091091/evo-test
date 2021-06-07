package response

// User struct type
type User struct {
	UserID   int64  `json:"id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}
