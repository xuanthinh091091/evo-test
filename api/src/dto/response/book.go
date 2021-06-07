package response

// Book struct type
type Book struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	TotalPage     int    `json:"total_page"`
	PublishedDate string `json:"published_date"`
	Price         string `json:"price"`
	Stock         int    `json:"stock"`
}
